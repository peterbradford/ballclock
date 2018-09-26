package ballclock

import (
	"unsafe"
)

type MinuteQueue struct {
	array [4]int
	length int
	first unsafe.Pointer
	size uintptr
}

func MQConstructor() *MinuteQueue{
	mq := new(MinuteQueue)
	mq.first = unsafe.Pointer(&mq.array[0]) 
	mq.size = unsafe.Sizeof(int(0))
	return mq
}


func (mq *MinuteQueue) append(x int) {
	*(*int)(unsafe.Pointer(uintptr(mq.first) + mq.size*uintptr(mq.length))) = x
	mq.length++
}

func (mq *MinuteQueue) reverse() [4]int{
	return [4]int{mq.array[3],mq.array[2],mq.array[1],mq.array[0]}
}

