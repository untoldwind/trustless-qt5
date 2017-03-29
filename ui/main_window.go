package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/untoldwind/trustless/secrets"
)

type MainWindow struct {
	logger  logging.Logger
	store   *uiStore
	secrets secrets.Secrets

	window  *widgets.QMainWindow
	stacked *widgets.QStackedWidget
}

func NewMainWindow(secrets secrets.Secrets, logger logging.Logger) (*MainWindow, error) {
	initialState, err := initialUiState(secrets)
	if err != nil {
		return nil, err
	}
	mainWindow := &MainWindow{
		logger:  logger.WithField("package", "ui"),
		secrets: secrets,
		store: &uiStore{
			current: *initialState,
		},
	}
	mainWindow.init()
	return mainWindow, err
}

func (w *MainWindow) Show() {
	w.window.Show()
}

func (w *MainWindow) init() {
	w.window = widgets.NewQMainWindow(nil, 0)
	w.window.SetWindowTitle("Trustless")
	w.window.SetMinimumSize2(200, 200)

	w.stacked = widgets.NewQStackedWidget(w.window)
	w.window.SetCentralWidget(w.stacked)

	w.stacked.AddWidget(w.newUnlockFrame())

	layout := widgets.NewQVBoxLayout()
	centralWidget := widgets.NewQWidget(nil, 0)
	centralWidget.SetLayout(layout)
	w.stacked.AddWidget(centralWidget)

	button := widgets.NewQPushButton2("Click me!", nil)
	button.ConnectClicked(func(checked bool) {
		w.store.dispatch(actionLock)
	})
	layout.AddWidget(button, 0, core.Qt__AlignCenter)

	w.store.addListener(w.onStateChange)
	w.onStateChange(&w.store.current, &w.store.current)
}

func (w *MainWindow) onStateChange(prev, next *uiState) {
	if next.locked {
		w.stacked.SetCurrentIndex(0)
	} else {
		w.stacked.SetCurrentIndex(1)
	}
}
