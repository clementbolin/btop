package view

import (
	"log"

	"github.com/rivo/tview"
)

// func InitView(title string) *tview.Box {
// 	box := tview.NewBox().SetBorder(true).SetTitle(title)
// 	return box
// }

func initFlexApp() *tview.Flex {
	flex := tview.NewFlex().
			AddItem(tview.NewBox().SetBorder(true).SetTitle("top"), 0, 2, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
					AddItem(tview.NewBox().SetBorder(true).SetTitle("first option"), 0, 1, false).
					AddItem(tview.NewBox().SetBorder(true).SetTitle("second option"), 0, 2, false).
					AddItem(tview.NewBox().SetBorder(true).SetTitle("bottom option"), 0, 1, false), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Right box"), 0, 2, false)
	
	return flex
}

// RunView : init and start default view
func RunView(title string) {
	app := tview.NewApplication()
	if err := app.SetRoot(initFlexApp(), true).EnableMouse(true).Run(); err != nil {
		log.Fatalln(err)
	}
}
