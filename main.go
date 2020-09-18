package main

import (
	"fmt"
	"os"
	"flag"
	"runtime"

	"github.com/ClementBolin/topGo/modules/process"
	"github.com/ClementBolin/topGo/view"
)

func main() {
	graphicFlag := flag.Bool("ui", false, "start user interface")
	flag.Parse()

	if *graphicFlag {
		view.RunView()
		os.Exit(0)
	}
	if runtime.GOOS == "linux" {
		p, _ := process.ListProcessLinux()
		if p != nil {
			fmt.Println("p not null :", p)
		}
	} else {
		process.ListProcessUnix()
	}
	os.Exit(0)
}
