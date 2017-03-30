package ui

import (
	"context"
	"sync"

	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/trustless/api"
	"github.com/untoldwind/trustless/secrets"
)

type uiState struct {
	locked  bool
	entries []*api.SecretEntry
}

type uiStoreListener func(prev, next *uiState)
type uiStoreAction func(prev *uiState) *uiState

type uiStore struct {
	lock      sync.Mutex
	logger    logging.Logger
	current   uiState
	listeners []uiStoreListener
	actions   chan uiStoreAction
}

func newUIStore(secrets secrets.Secrets, logger logging.Logger) (*uiStore, error) {
	status, err := secrets.Status(context.Background())
	if err != nil {
		return nil, err
	}
	store := &uiStore{
		logger: logger.WithField("package", "ui").WithField("component", "uiStore"),
		current: uiState{
			locked: status.Locked,
		},
		actions: make(chan uiStoreAction, 100),
	}

	go store.loop()

	return store, nil
}

func (s *uiStore) addListener(listener uiStoreListener) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.listeners = append(s.listeners, listener)
}

func (s *uiStore) dispatch(action uiStoreAction) {
	select {
	case s.actions <- action:
	default:
		s.logger.Error("Action queue exhausted")
	}
}

func (s *uiStore) runAction(action uiStoreAction) {
	s.lock.Lock()
	defer s.lock.Unlock()

	prev := s.current
	if next := action(&s.current); next != nil {
		s.current = *next

		for _, listener := range s.listeners {
			listener(&prev, &s.current)
		}
	}
}

func (s *uiStore) loop() {
	for action := range s.actions {
		s.runAction(action)
	}
}
