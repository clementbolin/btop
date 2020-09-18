package view

import (
	"log"

	"github.com/ClementBolin/topGo/modules/process"
)

// RunView : init and start default view
func RunView() {
	var app Btop

	process, err := process.ListProcessLinux()
	if err != nil {
		log.Fatalln(err)
	}
	app.Init()
	app.InitProcessText(process)
	app.InitMidpView()
	app.InitNotifView()
	if err = app.app.SetRoot(app.flex, true).EnableMouse(true).Run(); err != nil {
		log.Fatalln(err)
	}
}
