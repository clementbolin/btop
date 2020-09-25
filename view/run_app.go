package view

import (
	"log"
)

// RunView : init and start default view
func RunView() {
	var app Btop

	app.Init()
	app.InitProcessText()
	app.CreateBatteryTextView()
	app.InitMidpView()
	app.InitNotifView()

	if err := app.app.SetRoot(app.flex, true).EnableMouse(true).Run(); err != nil {
		log.Fatalln(err)
	}
}
