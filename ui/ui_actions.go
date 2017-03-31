package ui

import (
	"context"
	"sort"
	"strings"

	"github.com/untoldwind/trustless/api"
)

func (s *uiStore) actionUnlock(identity api.Identity, passphrase string) error {
	if !s.currentState().locked {
		return nil
	}
	if err := s.secrets.Unlock(context.Background(), identity.Name, identity.Email, passphrase); err != nil {
		s.logger.ErrorErr(err)
		return err
	}
	list, err := s.secrets.List(context.Background())
	if err != nil {
		s.logger.ErrorErr(err)
		return err
	}
	s.dispatch(func(state *uiState) *uiState {
		state.locked = false
		state.allEntries = list.Entries
		state.entryFilter = ""
		return filterSortAndVisible(state)
	})
	return nil
}

func (s *uiStore) actionLock() error {
	if s.currentState().locked {
		return nil
	}
	if err := s.secrets.Lock(context.Background()); err != nil {
		s.logger.ErrorErr(err)
		return err
	}
	s.dispatch(func(state *uiState) *uiState {
		state.locked = true
		return state
	})
	return nil
}

func (s *uiStore) actionUpdateEntryFilter(filter string) {
	s.dispatch(func(state *uiState) *uiState {
		state.entryFilter = strings.ToLower(filter)
		return filterSortAndVisible(state)
	})
}

func (s *uiStore) actionRefreshEntries() error {
	if s.currentState().locked {
		return nil
	}
	list, err := s.secrets.List(context.Background())
	if err != nil {
		s.logger.ErrorErr(err)
		return err
	}
	s.dispatch(func(state *uiState) *uiState {
		state.allEntries = list.Entries
		return filterSortAndVisible(state)
	})
	return nil
}

func (s *uiStore) actionSelectEntry(entryID string) error {
	secret, err := s.secrets.Get(context.Background(), entryID)
	if err != nil {
		s.logger.ErrorErr(err)
		return err
	}

	s.dispatch(func(state *uiState) *uiState {
		state.selectedEntry = nil
		state.currentSecret = secret
		for _, entry := range state.allEntries {
			if entry.ID == entryID {
				state.selectedEntry = entry
				return state
			}
		}
		return state
	})
	return nil
}

func filterSortAndVisible(state *uiState) *uiState {
	if state.entryFilter == "" {
		state.visibleEntries = make([]*api.SecretEntry, len(state.allEntries))
		copy(state.visibleEntries, state.allEntries)
	} else {
		state.visibleEntries = make([]*api.SecretEntry, 0, len(state.allEntries))
		for _, entry := range state.allEntries {
			if strings.HasPrefix(strings.ToLower(entry.Name), state.entryFilter) {
				state.visibleEntries = append(state.visibleEntries, entry)
			}
		}
	}
	sort.Sort(entryStoreNameAsc(state.visibleEntries))
	return state
}
