package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/therecipe/qt/widgets"
	"github.com/untoldwind/trustless/secrets"
)

type MainWindow struct {
	*widgets.QMainWindow

	logger  logging.Logger
	store   *uiStore
	secrets secrets.Secrets

	stacked *widgets.QStackedWidget
}

func NewMainWindow(secrets secrets.Secrets, logger logging.Logger) (*MainWindow, error) {
	uiStore, err := newUIStore(secrets, logger)
	if err != nil {
		return nil, err
	}
	w := &MainWindow{
		QMainWindow: widgets.NewQMainWindow(nil, 0),
		logger:      logger.WithField("package", "ui").WithField("component", "mainWindow"),
		secrets:     secrets,
		store:       uiStore,
	}

	w.SetWindowTitle("Trustless")
	w.SetMinimumSize2(200, 200)
	w.stacked = widgets.NewQStackedWidget(w)
	w.SetCentralWidget(w.stacked)

	w.stacked.AddWidget(newUnlockFrame(w.store, w.secrets, w.logger))
	w.stacked.AddWidget(newSecretsFrame(w.store, w.secrets, w.logger))

	w.store.addListener(w.onStateChange)
	w.onStateChange(&w.store.current, &w.store.current)

	return w, err
}

func (w *MainWindow) onStateChange(prev, next *uiState) {
	if next.locked {
		w.stacked.SetCurrentIndex(0)
	} else {
		w.stacked.SetCurrentIndex(1)
	}
}
