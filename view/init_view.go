package view

import (
	"fmt"
	"log"
	"strconv"
	"os"
	"runtime"

	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
	"github.com/ClementBolin/topGo/modules/process"
	"github.com/ClementBolin/topGo/modules/battery"
	"github.com/ClementBolin/topGo/modules/time"
	"github.com/ClementBolin/topGo/modules/gitstat"
)

// Btop : struct who 
type Btop struct {
	app		*tview.Application;
	flex	*tview.Flex;
	process *tview.TextView;
	battery *tview.TextView;
	gitStat *tview.TextView;
}

/*------------- Export Function ----------------*/

// Init : init app
func (app *Btop) Init() {
	app.app = tview.NewApplication()
	app.flex = tview.NewFlex()
	app.process = nil
	app.battery = nil
	app.bindsInit()
}

// InitGitStatText : init git stat text view
func (app *Btop) InitGitStatText() {
	gitStat := gitstat.GitStat()
	
	app.gitStat = tview.NewTextView()
	app.gitStat.SetBorder(true)
	app.gitStat.SetText(gitStat)
	app.gitStat.SetTextAlign(tview.AlignCenter)
	app.gitStat.SetTitle("Git Stat " + os.Getenv("PWD"))
	app.gitStat.SetBorderColor(tcell.ColorBlueViolet)
	app.gitStat.SetWrap(true)
	app.gitStat.SetDynamicColors(true)
}

// InitProcessText : init text view with process
func (app *Btop) InitProcessText() {
	var processList string = fmt.Sprintf("[green]Name[white]%s[red]Pid[white]%s[yellow]Ppid[white]\n\n", calculSpaceProcessList("Name"), calculSpaceProcessList("pid  "))
	var processL []process.UnixProcess
	var err error

	if runtime.GOOS == "linux" {
		processL, err = process.ListProcessLinux()
	} else {
		processL = process.ListProcessUnix()
	}

	if err != nil {
		processList = fmt.Sprintf("[green]Name[white]%s[red]Pid[white]%s[yellow]Ppid[white]\n\n", calculSpaceProcessList("Name"), calculSpaceProcessList("pid  "))
	} else {
		for _, item := range processL {
			processList = processList + fmt.Sprintf("[green]%s[white]%s[red]%d[white]%s[yellow]%d[white]\n", item.GetName(), calculSpaceProcessList(item.GetName()) ,item.GetPid(), calculSpaceProcessList(strconv.Itoa(item.GetPid())), item.GetPpid())
		}
	}

	app.process = tview.NewTextView()
	app.process.SetDynamicColors(true)
	app.process.SetText(processList)
	app.process.SetTextAlign(tview.AlignLeft)
	app.process.SetBorder(true)
	app.process.SetTitle("top")
	app.process.SetBorderColor(tcell.ColorBlue)
	app.flex.AddItem(app.process, 0, 1, false)
	UpdateTextView(refreshTextProcess, app.process, app.app)
}

// CreateBatteryTextView : create, update Battery TextView
func (app *Btop) CreateBatteryTextView() {
	var durateList string = ""
	var uptime time.Uptime

	uptime.InitUptime()
	u := uptime.GetAll()
	header := fmt.Sprintf("[red]%s[white]\nLoad average : %2.2f %2.2f %2.2f\nUptime : %s", u.Now ,u.AvgOne, u.AvgFive, u.AvgFifteen, u.Up)

	statBattery, err := battery.FillStatBattery()
	if err != nil {
		log.Fatalln(err)
	}

	durateList = fmt.Sprintf("%s\n\nPerceantage : %2.2f\nDuration : %s\n", header, statBattery.GetPercentage(), statBattery.GetDuration())
	app.battery = tview.NewTextView()
	app.battery.SetBorder(true)
	app.battery.SetText(durateList)
	app.battery.SetTextAlign(tview.AlignCenter)
	app.battery.SetTitle("Power System")
	app.battery.SetBorderColor(tcell.ColorBlueViolet)
	app.battery.SetDynamicColors(true)
	UpdateTextView(refreshTextBattery, app.battery, app.app)

}

// InitMidpView : init mid view with 3 empty box
func (app *Btop) InitMidpView() {
	app.flex.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
					AddItem(app.battery, 0, 1, false).
					AddItem(app.gitStat, 0, 2, false).
					AddItem(tview.NewBox().SetBorder(true).SetTitle("Docker Stat"), 0, 2, false), 0, 1, false)
}

// InitNotifView : Init notification view
func (app *Btop) InitNotifView() {
	app.flex.AddItem(tview.NewBox().SetBorder(true).SetTitle("Notification/Security"), 0, 1, false)
}

/*------------- Export Function ----------------*/

func calculSpaceProcessList(str string) string {
	var totalSpace = 30;
	for i := 0; i != len(str) - 1; i++ { totalSpace-- }
	var spaceStr = ""
	for i := 0; i != totalSpace; i++ { spaceStr = spaceStr + " " }
	return spaceStr
}
