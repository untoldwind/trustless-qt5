package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/therecipe/qt/widgets"
)

type secretsFrame struct {
	*widgets.QWidget
	entryList    *entryList
	secretDetail *secretDetail

	logger logging.Logger
	store  *uiStore
}

func newSecretsFrame(store *uiStore, logger logging.Logger) *secretsFrame {
	w := &secretsFrame{
		QWidget:      widgets.NewQWidget(nil, 0),
		entryList:    newEntryList(store, logger),
		secretDetail: newSecretDetail(store, logger),

		store:  store,
		logger: logger.WithField("component", "selectframe"),
	}

	layout := widgets.NewQGridLayout(w)

	splitter := widgets.NewQSplitter(nil)
	layout.AddWidget(splitter, 0, 0, 0)

	splitter.AddWidget(w.entryList)

	splitter.AddWidget(w.secretDetail)

	return w
}
