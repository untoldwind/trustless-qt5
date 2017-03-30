package ui

import (
	"sort"
	"strings"

	"github.com/untoldwind/trustless/api"
)

func actionUnlock(state *uiState) *uiState {
	if !state.locked {
		return nil
	}
	state.locked = false
	return state
}

func actionLock(state *uiState) *uiState {
	if state.locked {
		return nil
	}
	state.locked = true
	return state
}

func actionUpdateEntryFilter(filter string) uiStoreAction {
	return func(state *uiState) *uiState {
		state.entryFilter = strings.ToLower(filter)
		return actionFilterAndSortVisible(state)
	}
}

func actionUpdateEntries(list *api.SecretList) uiStoreAction {
	return func(state *uiState) *uiState {
		state.allEntries = list.Entries
		return actionFilterAndSortVisible(state)
	}
}

func actionFilterAndSortVisible(state *uiState) *uiState {
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

func actionSelectEntry(entryID string) uiStoreAction {
	return func(state *uiState) *uiState {
		state.selectedEntry = nil
		for _, entry := range state.allEntries {
			if entry.ID == entryID {
				state.selectedEntry = entry
				return state
			}
		}
		return state
	}
}
