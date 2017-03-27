package ui

import "sync"

type uiState struct {
	locked bool
}

type uiStoreListener func(prev, next uiState)
type uiStoreAction func(prev uiState) uiState

type uiStore struct {
	lock      sync.Mutex
	current   uiState
	listeners []uiStoreListener
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
	s.current = action(prev)

	for _, listener := range s.listeners {
		listener(prev, s.current)
	}
}
