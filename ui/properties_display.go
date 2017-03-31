package ui

import (
	"sort"

	"github.com/leanovate/microtools/logging"
	"github.com/therecipe/qt/widgets"
)

type propertyDefinition struct {
	display string
	hidden  bool
	blurred bool
}

var propertyDefinitions = map[string]propertyDefinition{
	"Tags":     {hidden: true},
	"name":     {hidden: true},
	"type":     {hidden: true},
	"username": {display: "Username"},
	"password": {display: "Password", blurred: true},
}

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
	sort.Strings(names)
	for _, name := range names {
		value := properties[name]
		if propertyDefinition, ok := propertyDefinitions[name]; ok {
			if propertyDefinition.hidden {
				continue
			}
			layout.AddRow3(propertyDefinition.display, newValueDisplay(value, propertyDefinition.blurred, w.logger))
		} else {
			layout.AddRow3(name, newValueDisplay(value, false, w.logger))
		}
	}

	w.SetWidget(form)
}
