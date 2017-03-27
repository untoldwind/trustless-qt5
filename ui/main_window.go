package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type MainWindow struct {
	store *uiStore

	window  *widgets.QMainWindow
	stacked *widgets.QStackedWidget
}

func NewMainWindow() *MainWindow {
	mainWindow := &MainWindow{
		store: &uiStore{
			current: uiState{
				locked: true,
			},
		},
	}
	mainWindow.init()
	return mainWindow
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

	w.stacked.AddWidget(w.newUnlockFrame(w.store))

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
}

func (w *MainWindow) onStateChange(prev, next uiState) {
	if next.locked {
		w.stacked.SetCurrentIndex(0)
	} else {
		w.stacked.SetCurrentIndex(1)
	}
}
