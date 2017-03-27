package ui

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func (w *MainWindow) newUnlockFrame(store *uiStore) widgets.QWidget_ITF {
	layout := widgets.NewQHBoxLayout()
	central := widgets.NewQWidget(nil, 0)
	central.SetLayout(layout)

	layout.AddStretch(1)
	passphrase := widgets.NewQLineEdit(nil)
	passphrase.SetEchoMode(widgets.QLineEdit__Password)
	passphrase.ConnectActionEvent(func(event *gui.QActionEvent) {
		fmt.Println("Action")
	})
	passphrase.ConnectReturnPressed(func() {
		store.dispatch(actionUnlock)
	})
	layout.AddWidget(passphrase, 3, core.Qt__AlignVCenter)
	layout.AddStretch(1)

	return central
}
