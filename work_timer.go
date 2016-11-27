package main

import (
	"github.com/sam-killgallon/timer/lib"
	"os"
	"time"
)

func main() {
	if len(os.Args[:]) == 1 {
		panic("Need to supply a command")
	}

	command := os.Args[1]
	current_time := time.Now()
	switch command {
	case "start":
		timer.Save_start_time(current_time)
	case "end":
		timer.Save_end_time(current_time)
	case "overtime":
		timer.Overtime()
	default:
		panic("Unrecognised command")
	}
}
