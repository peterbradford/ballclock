package ballclock

import (
	"unsafe"
)

type MinuteQueue struct {
	array [4]int
	offset int
	first unsafe.Pointer
	size uintptr
} 

func(mq *MinuteQueue) constructor() MinuteQueue {
	return MinuteQueue{
		first: unsafe.Pointer(&mq.array[0]),
		size: unsafe.Sizeof(int(0))}
}

func(mq *MinuteQueue) append(x int) {
	mq.array[mq.offset] = x
}

func(mq *MinuteQueue) get(x int) int{
	return *(*int)(unsafe.Pointer(uintptr(mq.first) + mq.size*uintptr(x)))
}