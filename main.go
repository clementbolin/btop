package main

import (
	"fmt"
	"os"
	"flag"
	"runtime"

	"github.com/ClementBolin/topGo/view"
	"github.com/ClementBolin/topGo/modules/docker"
)

func main() {
	graphicFlag := flag.Bool("ui", false, "start user interface")
	flag.Parse()

	if *graphicFlag {
		view.RunView()
		os.Exit(0)
	}
	if runtime.GOOS == "linux" {
		fmt.Println("Linux")
	} else {
		fmt.Println("Mac")
		var docker docker.DockerWidget
		docker.Init()
		docker.GetSystemInfo()
	}
	os.Exit(0)
}
