package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/untoldwind/trustless/secrets"
)

type secretDetail struct {
	*widgets.QStackedWidget
	displayForm       *widgets.QWidget
	displayFormLayout *widgets.QFormLayout

	nameLabel *widgets.QLabel

	logger  logging.Logger
	store   *uiStore
	secrets secrets.Secrets
}

func newSecretDetail(store *uiStore, secrets secrets.Secrets, logger logging.Logger) *secretDetail {
	w := &secretDetail{
		QStackedWidget:    widgets.NewQStackedWidget(nil),
		displayForm:       widgets.NewQWidget(nil, 0),
		displayFormLayout: widgets.NewQFormLayout(nil),
		nameLabel:         widgets.NewQLabel(nil, 0),

		store:   store,
		secrets: secrets,
		logger:  logger.WithField("component", "secretDetail"),
	}

	noSelectionLabel := widgets.NewQLabel2("No selection", nil, 0)
	noSelectionLabel.SetAlignment(core.Qt__AlignCenter)
	noSelectionLabel.SetMargin(100)
	w.AddWidget(noSelectionLabel)

	w.AddWidget(w.displayForm)
	w.displayForm.SetLayout(w.displayFormLayout)
	w.displayFormLayout.AddRow5(w.nameLabel)
	w.nameLabel.Font().SetBold(true)
	w.nameLabel.Font().SetPointSize(20)

	w.store.addListener(w.onStateChange)

	return w
}

func (w *secretDetail) onStateChange(prev, next *uiState) {
	if next.currentSecret == nil {
		w.SetCurrentIndex(0)
		return
	}
	w.SetCurrentIndex(1)
	w.nameLabel.SetText(next.currentSecret.Current.Name)
}
