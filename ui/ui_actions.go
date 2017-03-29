package ui

import "github.com/untoldwind/trustless/api"

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

func actionUpdateEntries(list *api.SecretList) uiStoreAction {
	return func(state *uiState) *uiState {
		state.entries = list.Entries
		return state
	}
}
