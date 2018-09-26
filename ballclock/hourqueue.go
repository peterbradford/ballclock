package ballclock

import (
	"unsafe"
)

type HourQueue struct {
	array [11]int
	length int
	first unsafe.Pointer
	size uintptr
}

func HQConstructor() *HourQueue{
	hq := new(HourQueue)
	hq.first = unsafe.Pointer(&hq.array[0]) 
	hq.size = unsafe.Sizeof(int(0))
	return hq
}

func (hq *HourQueue) append(x int) {
	*(*int)(unsafe.Pointer(uintptr(hq.first) + hq.size*uintptr(hq.length))) = x
	hq.length++
}

func (hq *HourQueue) reverse() [11]int{
	return [11]int{hq.array[10],hq.array[9],hq.array[8],hq.array[7],hq.array[6],hq.array[5],hq.array[4],hq.array[3],hq.array[2],hq.array[1],hq.array[0]}
}