package main

import (
	"os"
	"flag"

	"github.com/ClementBolin/topGo/modules/process"
	"github.com/ClementBolin/topGo/view"
)

func main() {
	graphicFlag := flag.Bool("ui", false, "start user interface")
	flag.Parse()

	if *graphicFlag {
		view.RunView("top-go")
		os.Exit(0)
	}
	process.ListAllProcess()
	os.Exit(0)
}
