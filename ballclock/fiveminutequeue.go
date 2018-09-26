package ballclock

import (
	"unsafe"
)

type FiveMinuteQueue struct {
	array [11]int
	length int
	first unsafe.Pointer
	size uintptr
}

func FMQConstructor() *FiveMinuteQueue{
	fmq := new(FiveMinuteQueue)
	fmq.first = unsafe.Pointer(&fmq.array[0]) 
	fmq.size = unsafe.Sizeof(int(0))
	return fmq
}

func (fmq *FiveMinuteQueue) append(x int) {
	*(*int)(unsafe.Pointer(uintptr(fmq.first) + fmq.size*uintptr(fmq.length))) = x
	fmq.length++
}

func (fmq *FiveMinuteQueue) reverse() [11]int{
	return [11]int{fmq.array[10],fmq.array[9],fmq.array[8],fmq.array[7],fmq.array[6],fmq.array[5],fmq.array[4],fmq.array[3],fmq.array[2],fmq.array[1],fmq.array[0]}
}