package main

import (
	"log"
	"fmt"
	"os"
	"flag"
	"runtime"
	"time"

	"github.com/ClementBolin/topGo/view"
	"github.com/distatus/battery"
)

func main() {
	graphicFlag := flag.Bool("ui", false, "start user interface")
	flag.Parse()

	if *graphicFlag {
		view.RunView()
		os.Exit(0)
	}
	if runtime.GOOS == "linux" {
		batteries, err := battery.GetAll()
		if err != nil {
			log.Fatalln(err)
		}
		for i, battery := range batteries {
			t := battery.State
			fmt.Printf("Bat %d :", i)
			fmt.Println("state : ", t)
			fmt.Println("BAT :", battery.Current / battery.Full * 100)
		}
	} else {
		batteries, err := battery.GetAll()
		if err != nil {
			log.Fatalln(err)
		}
		for i, battery := range batteries {
			t := byte(battery.State)
			timeNum := battery.Current / battery.ChargeRate
			duration, _ := time.ParseDuration(fmt.Sprintf("%fh", timeNum))
			fmt.Printf("Bat %d :", i)
			fmt.Println("state : ", string(t))
			fmt.Println("BAT :", battery.Current / battery.Full * 100)
			fmt.Println("DUR :", duration)
		}
	}
	os.Exit(0)
}
