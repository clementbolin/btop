package view

import (
	"log"
	"runtime"

	"github.com/ClementBolin/topGo/modules/process"
)

// RunView : init and start default view
func RunView() {
	var app Btop
	var processList []process.UnixProcess
	var err error

	if runtime.GOOS == "linux" {
		processList, err = process.ListProcessLinux()
	} else {
		processList = process.ListProcessUnix()
	}
	if err != nil {
		log.Fatalln(err)
	}
	app.Init()
	app.InitProcessText(processList)
	app.InitMidpView()
	app.InitNotifView()
	if err = app.app.SetRoot(app.flex, true).EnableMouse(true).Run(); err != nil {
		log.Fatalln(err)
	}
}
