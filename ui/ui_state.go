package ui

import (
	"context"
	"sync"

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
	current   uiState
	listeners []uiStoreListener
}

func initialUiState(secrets secrets.Secrets) (*uiState, error) {
	status, err := secrets.Status(context.Background())
	if err != nil {
		return nil, err
	}
	return &uiState{
		locked: status.Locked,
	}, nil
}

func (s *uiStore) addListener(listener uiStoreListener) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.listeners = append(s.listeners, listener)
}

func (s *uiStore) dispatch(action uiStoreAction) {
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
