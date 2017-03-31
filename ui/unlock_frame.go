package ui

import (
	"fmt"

	"github.com/leanovate/microtools/logging"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type unlockFrame struct {
	*widgets.QWidget

	logger logging.Logger
	store  *uiStore
}

func newUnlockFrame(store *uiStore, logger logging.Logger) *unlockFrame {
	w := &unlockFrame{
		QWidget: widgets.NewQWidget(nil, 0),

		store:  store,
		logger: logger.WithField("component", "unlockframe"),
	}

	layout := widgets.NewQHBoxLayout2(w)

	layout.AddStretch(1)
	passphraseContainer := widgets.NewQWidget(nil, 0)
	passphraseLayout := widgets.NewQVBoxLayout2(passphraseContainer)
	layout.AddWidget(passphraseContainer, 3, core.Qt__AlignVCenter)
	layout.AddStretch(1)

	passphraseLayout.AddWidget(widgets.NewQLabel2("Store is locked", nil, 0), 0, core.Qt__AlignHCenter)

	identitySelect := widgets.NewQComboBox(nil)
	passphraseLayout.AddWidget(identitySelect, 0, core.Qt__AlignVCenter)
	identities := w.store.currentState().identities
	identityItems := make([]string, len(identities))
	for i, identity := range identities {
		identityItems[i] = fmt.Sprintf("%s <%s>", identity.Name, identity.Email)
	}
	identitySelect.AddItems(identityItems)

	passphrase := widgets.NewQLineEdit(nil)
	passphrase.SetEchoMode(widgets.QLineEdit__Password)
	passphraseLayout.AddWidget(passphrase, 0, core.Qt__AlignVCenter)
	passphraseError := widgets.NewQLabel2("Invalid", nil, 0)
	passphraseLayout.AddWidget(passphraseError, 0, core.Qt__AlignHCenter)
	passphraseError.SetVisible(false)
	passphraseError.SetForegroundRole(gui.QPalette__Highlight)

	passphrase.ConnectReturnPressed(func() {
		identity := identities[identitySelect.CurrentIndex()]
		if err := w.store.actionUnlock(identity, passphrase.Text()); err != nil {
			passphraseError.SetVisible(true)
		}
	})

	return w
}
