package view

import (
	"fmt"
	// "strconv"

	"github.com/rivo/tview"
	"github.com/ClementBolin/topGo/modules/process"
)

// Btop : struct who 
type Btop struct {
	app		*tview.Application;
	flex	*tview.Flex;
	process *tview.TextView;
}


/*------------- Export Function ----------------*/

// Init : init app
func (app *Btop) Init() {
	app.app = tview.NewApplication()
	app.flex = tview.NewFlex()
	app.process = nil
}

// InitProcessText : init text view with process
func (app *Btop) InitProcessText(process []process.UnixProcess) {
	var processList string

	for _, item := range process {
		processList = processList + fmt.Sprintf("name : %s, pid : %d, ppid : %d\n", item.GetName(), item.GetPid(), item.GetPpid())
	}
	app.process = tview.NewTextView()
	app.process.SetBorder(false)
	app.process.SetText(processList)
	app.process.SetTextAlign(tview.AlignLeft)
	app.process.SetBorder(true)
	app.process.SetTitle("top")
	app.flex.AddItem(app.process, 0, 1, false)
}

// InitMidpView : init mid view with 3 empty box
func (app *Btop) InitMidpView() {
	app.flex.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
					AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
					AddItem(tview.NewBox().SetBorder(true).SetTitle("Mid"), 0, 1, false).
					AddItem(tview.NewBox().SetBorder(true).SetTitle("Bootom"), 0, 1, false), 0, 1, false)
}

// InitNotifView : Init notification view
func (app *Btop) InitNotifView() {
	app.flex.AddItem(tview.NewBox().SetBorder(true).SetTitle("Notification"), 0, 1, false)
}

/*------------- Export Function ----------------*/

/*------------- No export Func -----------------*/

// func initFlexApp() *tview.Flex {
// 	flex := tview.NewFlex().
// 			AddItem(textViewTest().SetBorder(true).SetTitle("top"), 0, 2, false).
// 			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
// 					AddItem(tview.NewBox().SetBorder(true).SetTitle("first option"), 0, 1, false).
// 					AddItem(tview.NewBox().SetBorder(true).SetTitle("second option"), 0, 2, false).
// 					AddItem(tview.NewBox().SetBorder(true).SetTitle("bottom option"), 0, 1, false), 0, 3, false).
// 			AddItem(tview.NewBox().SetBorder(true).SetTitle("Right box"), 0, 2, false)
// 	return flex
// }

/*------------- No export Func -----------------*/
