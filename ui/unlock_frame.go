package ui

import (
	"context"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func (w *MainWindow) newUnlockFrame() widgets.QWidget_ITF {
	layout := widgets.NewQHBoxLayout()
	central := widgets.NewQWidget(nil, 0)
	central.SetLayout(layout)

	layout.AddStretch(1)
	passphraseLayout := widgets.NewQVBoxLayout()
	passphraseContainer := widgets.NewQWidget(nil, 0)
	passphraseContainer.SetLayout(passphraseLayout)
	layout.AddWidget(passphraseContainer, 3, core.Qt__AlignVCenter)
	layout.AddStretch(1)

	passphraseLayout.AddWidget(widgets.NewQLabel2("Store is locked", nil, 0), 0, core.Qt__AlignHCenter)
	passphrase := widgets.NewQLineEdit(nil)
	passphrase.SetEchoMode(widgets.QLineEdit__Password)
	passphraseLayout.AddWidget(passphrase, 0, core.Qt__AlignVCenter)
	passphraseError := widgets.NewQLabel2("Invalid", nil, 0)
	passphraseLayout.AddWidget(passphraseError, 0, core.Qt__AlignHCenter)
	passphraseError.SetVisible(false)
	passphraseError.SetForegroundRole(gui.QPalette__Highlight)

	passphrase.ConnectReturnPressed(func() {
		if err := w.secrets.Unlock(context.Background(), "Bodo Junglas", "junglas@objectcode.de", passphrase.Text()); err != nil {
			w.logger.ErrorErr(err)
			passphraseError.SetVisible(true)
		} else {
			w.store.dispatch(actionUnlock)
		}
	})

	return central
}
