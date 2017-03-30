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
	entryList      *widgets.QListView
	entryListModel *gui.QStandardItemModel

	logger  logging.Logger
	store   *uiStore
	secrets secrets.Secrets
}

func newEntryList(store *uiStore, secrets secrets.Secrets, logger logging.Logger) *entryList {
	w := &entryList{
		QWidget:        widgets.NewQWidget(nil, 0),
		entryList:      widgets.NewQListView(nil),
		entryListModel: gui.NewQStandardItemModel(nil),

		store:   store,
		secrets: secrets,
		logger:  logger.WithField("component", "entryList"),
	}

	layout := widgets.NewQVBoxLayout2(w)
	layout.AddWidget(w.entryList, 0, 0)

	w.entryList.SetModel(w.entryListModel)
	w.entryListModel.SetSortRole(int(core.Qt__DisplayRole))

	w.store.addListener(w.onStateChange)
	if !w.store.current.locked {
		w.refresh()
	}

	return w
}

func (w *entryList) onStateChange(prev, next *uiState) {
	if !next.locked && prev.locked {
		w.refresh()
	}
	for i, entry := range next.entries {
		var listItem *gui.QStandardItem
		if i < w.entryListModel.RowCount(core.NewQModelIndex()) {
			listItem = w.entryListModel.Item(i, 0)
		} else {
			listItem = gui.NewQStandardItem()
			w.entryListModel.AppendRow([]*gui.QStandardItem{listItem})
		}
		listItem.SetText(entry.Name)
		listItem.SetData(core.NewQVariant14(entry.Name), int(core.Qt__DisplayRole))
	}
	if w.entryListModel.RowCount(core.NewQModelIndex()) > len(next.entries) {
		w.entryListModel.RemoveRows(len(next.entries), w.entryListModel.RowCount(core.NewQModelIndex())-len(next.entries), nil)
	}
	w.entryListModel.Sort(0, core.Qt__AscendingOrder)
}

func (w *entryList) refresh() {
	entries, err := w.secrets.List(context.Background())
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	w.store.dispatch(actionUpdateEntries(entries))
}
