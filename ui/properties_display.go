package ui

import (
	"fmt"
	"sort"

	"github.com/leanovate/microtools/logging"
	"github.com/therecipe/qt/widgets"
)

type propertiesDisplay struct {
	*widgets.QScrollArea

	logger logging.Logger
}

func newPropertiesDisplay(logger logging.Logger) *propertiesDisplay {
	w := &propertiesDisplay{
		QScrollArea: widgets.NewQScrollArea(nil),
		logger:      logger.WithField("component", "propertiesDisplay"),
	}

	return w
}

func (w *propertiesDisplay) setProperties(properties map[string]string) {
	form := widgets.NewQWidget(nil, 0)
	layout := widgets.NewQFormLayout(nil)
	form.SetLayout(layout)

	names := make([]string, 0, len(properties))
	for name := range properties {
		names = append(names, name)
	}
	fmt.Println(names)
	sort.Strings(names)
	for _, name := range names {
		value := properties[name]
		layout.AddRow3(name, widgets.NewQLabel2(value, nil, 0))
	}

	w.SetWidget(form)
}
