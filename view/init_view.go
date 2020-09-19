package view

import (
	"fmt"
	"strconv"

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
	var processList string = fmt.Sprintf("Name%sPid%sPpid\n\n", calculSpaceProcessList("Name"), calculSpaceProcessList("pid  "))

	for _, item := range process {
		processList = processList + fmt.Sprintf("%s%s%d%s%d\n", item.GetName(), calculSpaceProcessList(item.GetName()) ,item.GetPid(), calculSpaceProcessList(strconv.Itoa(item.GetPid())), item.GetPpid())
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

func calculSpaceProcessList(str string) string {
	var totalSpace = 30;
	for i := 0; i != len(str) - 1; i++ { totalSpace-- }
	var spaceStr = ""
	for i := 0; i != totalSpace; i++ { spaceStr = spaceStr + " " }
	return spaceStr
}
