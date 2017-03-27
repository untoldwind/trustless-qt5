package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
	"github.com/untoldwind/trustless-qt5/ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	mainWindow := ui.NewMainWindow()
	mainWindow.Show()

	widgets.QApplication_Exec()
}
