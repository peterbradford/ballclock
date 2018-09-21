package ballclock

import (
	"errors"
)

var mainQueue []int
var minuteQueue [4]int
var fiveMinuteQueue []int 
var hourQueue []int
var Days int

type JsonQueues struct {
	Min 	[4]int	`json:"Min"`
	FiveMin []int	`json:"FiveMin"`
	Hour 	[]int	`json:"Hour"`
	Main 	[]int	`json:"Main"`
}

func Init2(balls int) {
	for i := 0; i<balls; i++ {
		mainQueue = append(mainQueue, i+1)
	}
	fiveMinuteQueue = make([]int, 0)
	hourQueue = make([]int, 0)
	Days = 0
}

func GetJsonQueues() JsonQueues{
	newMin := reverseMinute(minuteQueue)
	new5Min := reverseQueue(fiveMinuteQueue)
	newHour := reverseQueue(hourQueue)

	return JsonQueues{Min: newMin, FiveMin:new5Min, Hour:newHour,Main:mainQueue}
}

func reverseQueue(tmp []int) []int{
	for left, right := 0, len(tmp)-1; left < right; left, right = left+1, right-1 {
		tmp[left], tmp[right] = tmp[right], tmp[left]
	}
	return tmp
}

func reverseMinute(tmp [4]int) [4]int{
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
	
	x,_ := PopMain()

	if len(minuteQueue) == 4{
		mainQueue = append(mainQueue, reverseMinute(minuteQueue)...)
		minuteQueue = minuteQueue[:0]
		PushFiveMinute(x)
	} else {
		minuteQueue = append(minuteQueue, x)
	}
}

func PushFiveMinute(x int) {
	if len(fiveMinuteQueue) == 11{
		mainQueue = append(mainQueue, reverseQueue(fiveMinuteQueue)...)
		fiveMinuteQueue = fiveMinuteQueue[:0]
		PushHour(x)
	} else {
		fiveMinuteQueue = append(fiveMinuteQueue, x)
	}
}

func PushHour(x int) {
	if len(hourQueue) == 11{
		mainQueue = append(mainQueue, reverseQueue(hourQueue)...)
		hourQueue = hourQueue[:0]
		PushDay(x)
	} else {
		hourQueue = append(hourQueue, x)
	}
}

func PopMain() (int, error){
	if len(mainQueue) != 0{
		rtnVal := mainQueue[0]
		mainQueue = mainQueue[1:]
		return rtnVal, nil
	} else {
		return 0, errors.New("queue is empty")
	}
}

// func Pop(val *[]int, val2 []int){
// 	// fmt.Printf("type1: %T\ntype2: %T\n\n", val, val2)
// }

func PushDay(x int) {
	Days += 1
	mainQueue = append(mainQueue, x)
}