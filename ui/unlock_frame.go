package ui

import (
	"context"
	"fmt"

	"github.com/leanovate/microtools/logging"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/untoldwind/trustless/secrets"
)

type unlockFrame struct {
	*widgets.QWidget

	logger  logging.Logger
	store   *uiStore
	secrets secrets.Secrets
}

func newUnlockFrame(store *uiStore, secrets secrets.Secrets, logger logging.Logger) *unlockFrame {
	w := &unlockFrame{
		QWidget: widgets.NewQWidget(nil, 0),

		store:   store,
		secrets: secrets,
		logger:  logger.WithField("component", "unlockframe"),
	}

	layout := widgets.NewQHBoxLayout()
	w.SetLayout(layout)

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

	popup := newPopupWidget(w)

	w.ConnectPaintEvent(func(event *gui.QPaintEvent) {
		pos := passphrase.MapFromGlobal(core.NewQPoint2(0, 0))
		posX := -pos.X()
		posY := -pos.Y()
		fmt.Println(posX)
		fmt.Println(posY)

		fmt.Println(popup.Width())
		fmt.Println(popup.Height())

		popup.SetGeometry(core.NewQRect4(posX+passphrase.Width(), posY, popup.Width(), popup.Height()))
		popup.Show()
	})

	return w
}
