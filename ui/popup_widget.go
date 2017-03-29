package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type popupWidet struct {
	*widgets.QWidget
	label *widgets.QLabel
}

func newPopupWidget(parent widgets.QWidget_ITF) *popupWidet {
	w := &popupWidet{
		QWidget: widgets.NewQWidget(parent, core.Qt__FramelessWindowHint|core.Qt__Tool),
		label:   widgets.NewQLabel(nil, 0),
	}

	layout := widgets.NewQGridLayout2()
	w.SetLayout(layout)

	w.label.SetAlignment(core.Qt__AlignCenter)
	w.label.SetText("hurra")
	layout.AddWidget(w.label, 0, 0, 0)

	w.ConnectPaintEvent(func(event *gui.QPaintEvent) {
		painter := gui.NewQPainter2(w)
		defer painter.DestroyQPainter()
		painter.SetRenderHint(gui.QPainter__Antialiasing, true)

		rect := w.Rect()
		painter.FillRect8(rect, core.Qt__transparent)

		roundedRectDimensions := core.NewQRect4(rect.X()+5, rect.Y()+5, rect.Width()-10, rect.Height()-10)

		painter.SetBrush(gui.NewQBrush4(core.Qt__lightGray, core.Qt__SolidPattern))
		pen := gui.NewQPen3(gui.NewQColor2(core.Qt__gray))
		pen.SetWidth(3)
		painter.SetPen(pen)

		painter.DrawRoundedRect3(roundedRectDimensions, 15, 15, core.Qt__RelativeSize)
	})

	return w
}
