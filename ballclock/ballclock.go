package ballclock

import (
	"fmt"
)

var MainQue MainQueue
var minuteQueue MinuteQueue
var fiveMinuteQueue FiveMinuteQueue
var hourQueue HourQueue

var Days int
var Balls int
var x int
type JsonQueues struct {
	Min 	[4]int	`json:"Min"`
	FiveMin [11]int	`json:"FiveMin"`
	Hour 	[11]int	`json:"Hour"`
	Main 	[128]int	`json:"Main"`
}

func Init2(balls int) {
	MainQue.init()
	MainQue.length = balls
	minuteQueue = MinuteQueue{}
	fiveMinuteQueue = FiveMinuteQueue{}
	hourQueue = HourQueue{}
	Days = 0
}

func GetJsonQueues() JsonQueues{
	newMin := minuteQueue.reverse()
	new5Min := fiveMinuteQueue.reverse()
	newHour := hourQueue.reverse()
	newMain := MainQue.print()

	return JsonQueues{Min: newMin, FiveMin:new5Min, Hour:newHour,Main:newMain}
}

func PushMinute() {
	
	x = MainQue.pop()

	if minuteQueue.pointer == 4{
		MainQue.append(minuteQueue.third)
		MainQue.append(minuteQueue.second)
		MainQue.append(minuteQueue.first)
		MainQue.append(minuteQueue.zero)
		
		minuteQueue.pointer = 0
		PushFiveMinute(x)
	} else {
		minuteQueue.append(x)
	}
}

func PushFiveMinute(x int) {
	if fiveMinuteQueue.pointer == 11{
		MainQue.append(fiveMinuteQueue.tenth)
		MainQue.append(fiveMinuteQueue.nineth)
		MainQue.append(fiveMinuteQueue.eighth)
		MainQue.append(fiveMinuteQueue.seventh)
		MainQue.append(fiveMinuteQueue.sixth)
		MainQue.append(fiveMinuteQueue.fifth)
		MainQue.append(fiveMinuteQueue.fourth)
		MainQue.append(fiveMinuteQueue.third)
		MainQue.append(fiveMinuteQueue.second)
		MainQue.append(fiveMinuteQueue.first)
		MainQue.append(fiveMinuteQueue.zero)
		
		fiveMinuteQueue.pointer = 0
		PushHour(x)
	} else {
		fiveMinuteQueue.append(x)
	}
}

func PushHour(x int) {
	if hourQueue.pointer == 11{
		MainQue.append(hourQueue.tenth)
		MainQue.append(hourQueue.nineth)
		MainQue.append(hourQueue.eighth)
		MainQue.append(hourQueue.seventh)
		MainQue.append(hourQueue.sixth)
		MainQue.append(hourQueue.fifth)
		MainQue.append(hourQueue.fourth)
		MainQue.append(hourQueue.third)
		MainQue.append(hourQueue.second)
		MainQue.append(hourQueue.first)
		MainQue.append(hourQueue.zero)
		hourQueue.pointer = 0
		PushDay(x)
	} else {
		hourQueue.append(x)
	}
}

func Pop(val *[]int, val2 []int){
	fmt.Printf("type1: %T\ntype2: %T\n\n", val, val2)
}

func PushDay(x int) {
	Days += 1
	MainQue.append(x)
}