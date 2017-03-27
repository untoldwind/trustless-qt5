package ui

func actionUnlock(state uiState) uiState {
	state.locked = false
	return state
}

func actionLock(state uiState) uiState {
	state.locked = true
	return state
}
