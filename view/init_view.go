package view

import (
	"log"

	"github.com/rivo/tview"
)

// func InitView(title string) *tview.Box {
// 	box := tview.NewBox().SetBorder(true).SetTitle(title)
// 	return box
// }

// RunView : init and start default view
func RunView(title string) {
	box := tview.NewBox().SetBorder(true).SetTitle(title)
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		log.Fatalln(err);
	}	
}
