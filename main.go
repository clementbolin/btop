package main

import (
	"fmt"
	"os"
	"flag"
	"runtime"
	// t "time"

	"github.com/ClementBolin/topGo/view"
	// "github.com/ClementBolin/topGo/modules/time"
)

// func test(up chan string, u time.Uptime) {
// 	for {
// 		u.UpdateUptime(up)
// 	}
// }

// func test2(up chan string) {
// 	for {
// 		t.Sleep(10)
// 		fmt.Println(<-up)
// 	}
// }

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
		// hour := make(chan string)
		// var uptime time.Uptime
		// go test(hour, uptime)
		// go test2(hour)
		// t.Sleep(2000000000000)
	}
	os.Exit(0)
}
