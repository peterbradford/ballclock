package main

import (
	"github.com/peterbradford/ballclock/ballclock"
	"time"
	"fmt"
	"os"
	"strconv"
)

func checkArgs() bool {
	argsNoPath := os.Args[1:]
	if len(argsNoPath) != 1 && len(argsNoPath) != 2 {
		fmt.Println("too many or too few args.\nUsage: 30\nOr\nUsage: 30 325\nWhere first number is number of balls\nand second number is an optional number of minutes to run")
		os.Exit(1)
	}
	return len(argsNoPath) == 2
}

func timerMode() {
	start := time.Now()
	ballclock.PushMinute()
	for !ballclock.IsOriginalPosition() {
		ballclock.PushMinute()
	}
	end := time.Since(start)
	fmt.Printf("%v balls cycle after %v days.\nCompleted in %v milliseconds (%1.3f seconds)", ballclock.Balls, ballclock.Days/2, end.Nanoseconds()/int64(time.Millisecond), float64(end.Nanoseconds())/float64(int64(time.Second)))
}

func clockMode() {
	mins, err := strconv.Atoi(os.Args[1:][1])
	if err != nil {
		fmt.Println("error with second argument, please make it an int")
		os.Exit(3)
	}

	for minsLeft := 1; minsLeft <= (mins); minsLeft++ {
		ballclock.PushMinute()
	}
	fmt.Printf("%+v\n", ballclock.GetJsonQueues())
}

func main() {
	mode := checkArgs()
	if mode {
		clockMode()
	} else {
		timerMode()
	}
}