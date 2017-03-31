package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/therecipe/qt/widgets"
	"github.com/untoldwind/trustless/secrets"
)

type secretsFrame struct {
	*widgets.QWidget
	entryList    *entryList
	secretDetail *secretDetail

	logger  logging.Logger
	store   *uiStore
	secrets secrets.Secrets
}

func newSecretsFrame(store *uiStore, secrets secrets.Secrets, logger logging.Logger) *secretsFrame {
	w := &secretsFrame{
		QWidget:      widgets.NewQWidget(nil, 0),
		entryList:    newEntryList(store, secrets, logger),
		secretDetail: newSecretDetail(store, secrets, logger),

		store:   store,
		secrets: secrets,
		logger:  logger.WithField("component", "selectframe"),
	}

	layout := widgets.NewQGridLayout(w)

	splitter := widgets.NewQSplitter(nil)
	layout.AddWidget(splitter, 0, 0, 0)

	splitter.AddWidget(w.entryList)

	splitter.AddWidget(w.secretDetail)

	return w
}
