package ui

import (
	"context"

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

	logger  logging.Logger
	store   *uiStore
	secrets secrets.Secrets
}

func newEntryList(store *uiStore, secrets secrets.Secrets, logger logging.Logger) *entryList {
	w := &entryList{
		QWidget:        widgets.NewQWidget(nil, 0),
		filter:         widgets.NewQLineEdit(nil),
		entryList:      widgets.NewQListView(nil),
		entryListModel: newEntryListModel(),

		store:   store,
		secrets: secrets,
		logger:  logger.WithField("component", "entryList"),
	}

	layout := widgets.NewQVBoxLayout2(w)
	layout.AddWidget(w.filter, 0, 0)
	layout.AddWidget(w.entryList, 1, 0)

	w.entryList.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
	w.entryList.SetModel(w.entryListModel)
	w.entryList.ConnectCurrentChanged(w.onCurrentChanged)

	w.filter.ConnectKeyReleaseEvent(w.onFilterChange)

	w.store.addListener(w.onStateChange)
	if !w.store.current.locked {
		w.refresh()
	}

	return w
}

func (w *entryList) onFilterChange(event *gui.QKeyEvent) {
	w.store.dispatch(actionUpdateEntryFilter(w.filter.Text()))
}

func (w *entryList) onCurrentChanged(current *core.QModelIndex, previous *core.QModelIndex) {
	w.store.dispatch(actionSelectEntry(current.Data(int(entityIDRole)).ToString()))
}

func (w *entryList) onStateChange(prev, next *uiState) {
	if !next.locked && prev.locked {
		w.refresh()
	}
	w.entryListModel.updateEntries(next.visibleEntries)
	w.entryList.SetCurrentIndex(w.entryListModel.indexOf(next.selectedEntry))
}

func (w *entryList) refresh() {
	entries, err := w.secrets.List(context.Background())
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	w.store.dispatch(actionUpdateEntries(entries))
}
