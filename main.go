package main

import (
	"os"

	"github.com/leanovate/microtools/logging"
	"github.com/therecipe/qt/widgets"
	"github.com/untoldwind/trustless-qt5/ui"
	"github.com/untoldwind/trustless/secrets/remote"
)

func createLogger() logging.Logger {
	loggingOptions := logging.Options{
		Backend:   "logrus",
		LogFormat: "text",
		Level:     logging.Info,
	}
	return logging.NewLogrusLogger(loggingOptions).
		WithContext(map[string]interface{}{"process": "trustless-q5"})
}

func main() {
	logger := createLogger()
	secrets := remote.NewRemoteSecrets(logger)

	widgets.NewQApplication(len(os.Args), os.Args)

	mainWindow, err := ui.NewMainWindow(secrets, logger)
	if err != nil {
		logger.ErrorErr(err)
		return
	}
	mainWindow.Show()

	widgets.QApplication_Exec()
}
