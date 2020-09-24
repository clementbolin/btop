package view

import (
	"fmt"
	"log"

	"github.com/rivo/tview"
	"github.com/ClementBolin/topGo/modules/time"
	"github.com/ClementBolin/topGo/modules/battery"
)

func (app *Btop) writeView(text chan string) {
	val := <-text
	// log.Fatalln(val)
	app.battery.SetText(val)
	app.app.Draw()
}

// UpdateBatteryTextView : 
func (app *Btop) UpdateBatteryTextView() {
	text := make(chan string)
	go createTextBattery(text)
	go app.writeView(text)
}

func createTextBattery(text chan string) {
	var uptime time.Uptime
	var durateList string

	uptime.UpdateUptime()
	u := uptime.GetAll()

	statBattery, err := battery.FillStatBattery()
	if err != nil {
		log.Fatalln(err)
	}

	header := fmt.Sprintf("[red]%s[white]\nLoad average : %2.2f %2.2f %2.2f\nUptime : %s", u.Now ,u.AvgOne, u.AvgFive, u.AvgFifteen, u.Up)
	durateList = fmt.Sprintf("%s\n\nPerceantage : %2.2f\nDuration : %s\n", header, statBattery.GetPercentage(), statBattery.GetDuration())
	text <- durateList
}

func queueUpdateAndDraw(app *tview.Application, f func()) {
	app.QueueUpdateDraw(f)
}
