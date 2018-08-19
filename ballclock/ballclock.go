package ballclock

import (
	"errors"
	"strconv"
	"fmt"
	"os"
)

var mainQueue, minuteQueue, fiveMinuteQueue, hourQueue []int
var Days int
var Balls int

type JsonQueues struct {
	Min 	[]int	`json:"Min"`
	FiveMin []int	`json:"FiveMin"`
	Hour 	[]int	`json:"Hour"`
	Main 	[]int	`json:"Main"`
}

func init() {
	argsNoPath := os.Args[1:]
	balls, err := strconv.Atoi(argsNoPath[0])
	if err != nil {
		fmt.Println("error with first argument, please make it an int")
		os.Exit(2)
	}
	Balls = balls
	for i := 0; i<Balls; i++ {
		mainQueue = append(mainQueue, i+1)
	}
	minuteQueue = make([]int, 0)
	fiveMinuteQueue = make([]int, 0)
	hourQueue = make([]int, 0)
	Days = 0
}

func GetJsonQueues() JsonQueues{

	newMin := reverseQueue(&minuteQueue)
	new5Min := reverseQueue(&fiveMinuteQueue)
	newHour := reverseQueue(&hourQueue)

	return JsonQueues{Min: newMin, FiveMin:new5Min, Hour:newHour,Main:mainQueue}
}

func reverseQueue(q *[]int) []int{
	tmp := make([]int, len(*q))
	copy(tmp, *q)
	for left, right := 0, len(tmp)-1; left < right; left, right = left+1, right-1 {
		tmp[left], tmp[right] = tmp[right], tmp[left]
	}
	return tmp
}

func IsOriginalPosition() bool{
	for index, val := range mainQueue {
		if val != (index+1) {
			return false
		}
	}
	return true
}

func PushMinute() {
	x,_ := Pop(&mainQueue)

	if len(minuteQueue) == 4{
		mainQueue = append(mainQueue, minuteQueue...)
		minuteQueue = minuteQueue[:0]
		PushFiveMinute(x)
	} else {
		minuteQueue = append([]int{x}, minuteQueue...)
	}
}

func PushFiveMinute(x int) {
	if len(fiveMinuteQueue) == 11{
		mainQueue = append(mainQueue, fiveMinuteQueue...)
		fiveMinuteQueue = fiveMinuteQueue[:0]
		PushHour(x)
	} else {
		fiveMinuteQueue = append([]int{x}, fiveMinuteQueue...)
	}
}

func PushHour(x int) {
	if len(hourQueue) == 11{
		mainQueue = append(mainQueue, hourQueue...)
		hourQueue = hourQueue[:0]
		PushDay(x)
	} else {
		hourQueue = append([]int{x},hourQueue...)
	}
}

func Pop(q *[]int) (int, error){
	if len(*q) != 0{
		rtnVal := (*q)[0]
		*q = (*q)[1:]
		return rtnVal, nil
	} else {
		return 0, errors.New("queue is empty")
	}
}

func PushDay(x int) {
	Days += 1
	mainQueue = append(mainQueue, x)
}