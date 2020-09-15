package main

import (
	"log"
	"os"

	"github.com/rivo/tview"
)

func main() {
	box := tview.NewBox().SetBorder(true).SetTitle("Top-Go")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		log.Fatalln(err);
	}
	os.Exit(0)
}
	