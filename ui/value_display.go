package ui

import (
	"time"

	"github.com/leanovate/microtools/logging"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type valueDisplay struct {
	*widgets.QWidget

	logger logging.Logger
}

func newValueDisplay(value string, blurred bool, logger logging.Logger) *valueDisplay {
	w := &valueDisplay{
		QWidget: widgets.NewQWidget(nil, 0),
		logger:  logger.WithField("component", "valueDisplay"),
	}

	layout := widgets.NewQHBoxLayout2(w)
	label := widgets.NewQLabel(nil, 0)
	layout.AddWidget(label, 0, 0)
	if blurred {
		label.SetText("*************")
	} else {
		label.SetText(value)
	}

	actionButton := widgets.NewQToolButton(nil)
	layout.AddWidget(actionButton, 0, 0)
	actionButton.SetToolButtonStyle(core.Qt__ToolButtonTextOnly)
	actionButton.SetPopupMode(widgets.QToolButton__MenuButtonPopup)
	actionButton.SetText("Copy")

	actions := widgets.NewQMenu(nil)
	actionButton.SetMenu(actions)
	actions.AddAction("Copy").ConnectTriggered(func(checked bool) {
		copyToClipboard(value)
	})
	if blurred {
		actions.AddAction("Reveal").ConnectTriggered(func(checked bool) {
			label.SetText(value)
		})
	}

	actionButton.ConnectClicked(func(checked bool) {
		copyToClipboard(value)
	})

	return w
}

func copyToClipboard(value string) {
	clipboard := gui.QGuiApplication_Clipboard()
	clipboard.SetText(value, gui.QClipboard__Clipboard)
	clipboard.SetText(value, gui.QClipboard__Selection)

	go func() {
		time.Sleep(1 * time.Minute)
		if value == clipboard.Text(gui.QClipboard__Clipboard) {
			clipboard.Clear(gui.QClipboard__Clipboard)
		}
		if value == clipboard.Text(gui.QClipboard__Selection) {
			clipboard.Clear(gui.QClipboard__Selection)
		}
	}()
}
