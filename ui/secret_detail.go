package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type secretDetail struct {
	*widgets.QStackedWidget
	displayForm       *widgets.QWidget
	propertiesDisplay *propertiesDisplay

	nameLabel *widgets.QLabel

	logger logging.Logger
	store  *uiStore
}

func newSecretDetail(store *uiStore, logger logging.Logger) *secretDetail {
	w := &secretDetail{
		QStackedWidget:    widgets.NewQStackedWidget(nil),
		displayForm:       widgets.NewQWidget(nil, 0),
		nameLabel:         widgets.NewQLabel(nil, 0),
		propertiesDisplay: newPropertiesDisplay(logger),

		store:  store,
		logger: logger.WithField("component", "secretDetail"),
	}

	noSelectionLabel := widgets.NewQLabel2("No selection", nil, 0)
	noSelectionLabel.SetAlignment(core.Qt__AlignCenter)
	noSelectionLabel.SetMargin(100)
	w.AddWidget(noSelectionLabel)

	w.AddWidget(w.displayForm)
	displayFormLayout := widgets.NewQVBoxLayout2(w.displayForm)
	displayFormLayout.AddWidget(w.nameLabel, 0, core.Qt__AlignLeft)
	w.nameLabel.Font().SetBold(true)
	w.nameLabel.Font().SetPointSize(20)

	displayFormLayout.AddWidget(w.propertiesDisplay, 1, 0)

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
	w.propertiesDisplay.setProperties(next.currentSecret.Current.Properties)
}
