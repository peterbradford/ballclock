package ballclock


var MainQue *MainQueue
var minuteQueue *MinuteQueue
var fiveMinuteQueue *FiveMinuteQueue
var hourQueue *HourQueue

var Days int
var Balls int
var x int
type JsonQueues struct {
	Min 	[4]int	`json:"Min"`
	FiveMin [11]int	`json:"FiveMin"`
	Hour 	[11]int	`json:"Hour"`
	Main 	[]int	`json:"Main"`
}

func Init2(balls int) {
	MainQue = MNQConstructor()
	minuteQueue = MQConstructor()
	fiveMinuteQueue = FMQConstructor()
	hourQueue = HQConstructor()
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

	if minuteQueue.length == 4{
		MainQue.append(minuteQueue.array[3])
		MainQue.append(minuteQueue.array[2])
		MainQue.append(minuteQueue.array[1])
		MainQue.append(minuteQueue.array[0])
		// MainQue.append(*(*int)(unsafe.Pointer(uintptr(minuteQueue.first) + minuteQueue.size*uintptr(3))))
		// MainQue.append(*(*int)(unsafe.Pointer(uintptr(minuteQueue.first) + minuteQueue.size*uintptr(2))))
		// MainQue.append(*(*int)(unsafe.Pointer(uintptr(minuteQueue.first) + minuteQueue.size*uintptr(1))))
		// MainQue.append(*(*int)(unsafe.Pointer(uintptr(minuteQueue.first) + minuteQueue.size*uintptr(0))))
		
		minuteQueue.length = 0
		PushFiveMinute(x)
	} else {
		minuteQueue.append(x)
	}
}

func PushFiveMinute(x int) {
	if fiveMinuteQueue.length == 11{
		MainQue.appendEleven(fiveMinuteQueue.array[10],fiveMinuteQueue.array[9],
			fiveMinuteQueue.array[8],	fiveMinuteQueue.array[7],	fiveMinuteQueue.array[6],
			fiveMinuteQueue.array[5], fiveMinuteQueue.array[4], fiveMinuteQueue.array[3],
			fiveMinuteQueue.array[2], fiveMinuteQueue.array[1],	fiveMinuteQueue.array[0])
		
		fiveMinuteQueue.length = 0
		PushHour(x)
	} else {
		fiveMinuteQueue.append(x)
	}
}

func PushHour(x int) {
	if hourQueue.length == 11{
		MainQue.appendEleven(hourQueue.array[10],hourQueue.array[9],
			hourQueue.array[8],	hourQueue.array[7],	hourQueue.array[6],
			hourQueue.array[5], hourQueue.array[4], hourQueue.array[3],
			hourQueue.array[2], hourQueue.array[1],	hourQueue.array[0])
		hourQueue.length = 0
		PushDay(x)
	} else {
		hourQueue.append(x)
	}
}

func PushDay(x int) {
	// fmt.Printf("day: %v\n", Days)
	Days += 1
	MainQue.append(x)
}