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
		painter.SetRenderHint(gui.QPainter__Antialiasing, true)

		rect := w.Rect()
		painter.FillRect8(rect, core.Qt__transparent)

		roundedRectDimensions := core.NewQRect4(rect.X()+5, rect.Y()+5, rect.Width()-10, rect.Height()-10)

		painter.SetBrush(gui.NewQBrush4(core.Qt__lightGray, core.Qt__SolidPattern))
		pen := gui.NewQPen3(gui.NewQColor2(core.Qt__gray))
		pen.SetWidth(3)
		painter.SetPen(pen)

		painter.DrawRoundedRect3(roundedRectDimensions, 15, 15, core.Qt__RelativeSize)

		polygon := gui.NewQPolygon3([]*core.QPoint{
			core.NewQPoint2(roundedRectDimensions.X(), roundedRectDimensions.Height()/2-5+3),
			core.NewQPoint2(roundedRectDimensions.X(), roundedRectDimensions.Height()/2+5+3),
			core.NewQPoint2(roundedRectDimensions.X()-5, roundedRectDimensions.Height()/2+3),
		})

		painter.SetPen3(core.Qt__NoPen)
		painter.SetBrush(gui.NewQBrush4(core.Qt__gray, core.Qt__SolidPattern))
		painter.DrawPolygon4(polygon, core.Qt__OddEvenFill)
	})

	return w
}

func (w *popupWidet) Show() {
	pos := w.ParentWidget().MapFromGlobal(core.NewQPoint2(0, 0))
	posX := -pos.X()
	posY := -pos.Y()

	w.SetGeometry2(posX+w.ParentWidget().Width(), posY-w.ParentWidget().Height()/2, w.Width(), w.Height())
	w.QWidget.Show()
}
