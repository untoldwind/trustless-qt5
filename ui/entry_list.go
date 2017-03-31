package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/untoldwind/trustless/secrets"
)

type entryList struct {
	*widgets.QWidget
	filter         *widgets.QLineEdit
	entryList      *widgets.QListView
	entryListModel *entryListModel
	contextMenu    *widgets.QMenu

	logger  logging.Logger
	store   *uiStore
	secrets secrets.Secrets
}

func newEntryList(store *uiStore, logger logging.Logger) *entryList {
	w := &entryList{
		QWidget:        widgets.NewQWidget(nil, 0),
		filter:         widgets.NewQLineEdit(nil),
		entryList:      widgets.NewQListView(nil),
		entryListModel: newEntryListModel(),
		contextMenu:    widgets.NewQMenu(nil),

		store:  store,
		logger: logger.WithField("component", "entryList"),
	}

	layout := widgets.NewQVBoxLayout2(w)
	layout.AddWidget(w.filter, 0, 0)
	layout.AddWidget(w.entryList, 1, 0)

	w.entryList.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
	w.entryList.SetModel(w.entryListModel)
	w.entryList.ConnectCurrentChanged(w.onCurrentChanged)
	w.entryList.ConnectContextMenuEvent(w.onListContextMenu)
	w.entryList.ConnectKeyReleaseEvent(w.onListKeyRelease)

	w.contextMenu.AddAction("Copy username").ConnectTriggered(w.onCopyUsername)
	w.contextMenu.AddAction("Copy password").ConnectTriggered(w.onCopyPassword)

	w.filter.ConnectKeyReleaseEvent(w.onFilterChange)
	w.filter.ConnectReturnPressed(w.onFilterReturn)

	w.store.addListener(w.onStateChange)
	if !w.store.current.locked {
		w.store.actionRefreshEntries()
	}

	return w
}

func (w *entryList) onListContextMenu(event *gui.QContextMenuEvent) {
	w.contextMenu.Exec2(gui.QCursor_Pos(), nil)
}

func (w *entryList) onListKeyRelease(event *gui.QKeyEvent) {
	if event.Key() == int(core.Qt__Key_Right) {
		index := w.entryList.CurrentIndex()
		if !index.IsValid() {
			return
		}
		rect := w.entryList.VisualRect(index)
		w.contextMenu.Exec2(w.entryList.MapToGlobal(rect.Center()), nil)
	}
}

func (w *entryList) onCopyUsername(checked bool) {
	secret := w.store.currentState().currentSecret
	if secret == nil {
		return
	}
	if username, ok := secret.Current.Properties["username"]; ok {
		safeCopyToClipboard(username)
	}
}

func (w *entryList) onCopyPassword(checked bool) {
	secret := w.store.currentState().currentSecret
	if secret == nil {
		return
	}
	if password, ok := secret.Current.Properties["password"]; ok {
		safeCopyToClipboard(password)
	}
}

func (w *entryList) onFilterChange(event *gui.QKeyEvent) {
	w.store.actionUpdateEntryFilter(w.filter.Text())
}

func (w *entryList) onFilterReturn() {
	w.entryList.SetFocus(core.Qt__OtherFocusReason)
}

func (w *entryList) onCurrentChanged(current *core.QModelIndex, previous *core.QModelIndex) {
	w.store.actionSelectEntry(current.Data(int(entityIDRole)).ToString())
}

func (w *entryList) onStateChange(prev, next *uiState) {
	w.entryListModel.updateEntries(next.visibleEntries)
}
