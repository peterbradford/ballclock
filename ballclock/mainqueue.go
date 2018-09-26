package ballclock

import (
	"unsafe"
)

type MainQueue struct {
	array [128]int
	first unsafe.Pointer
	size uintptr
	length int
	frontPointer int
	backPointer int
	oldFront int
}

func MNQConstructor() *MainQueue{
	mnq := new(MainQueue)
	mnq.first = unsafe.Pointer(&mnq.array[0]) 
	mnq.size = unsafe.Sizeof(int(0))
	for i:=0; i<Balls; i++{
		mnq.array[i]=i+1
	}
	mnq.length = Balls
	// mnq.backPointer = 0
	return mnq
}

func (mnq *MainQueue) append(x int) {
	mnq.length++
	*(*int)(unsafe.Pointer(uintptr(mnq.first) + mnq.size*uintptr(mnq.backPointer))) = x
	if mnq.backPointer+1 >= Balls {
		mnq.backPointer = 0
	} else {
		mnq.backPointer++
	}

}

func (mnq *MainQueue) appendEleven(a int,b int,c int,d int,e int,f int,g int,h int,i int,j int,k int) {
	mnq.length++
	*(*int)(unsafe.Pointer(uintptr(mnq.first) + mnq.size*uintptr(mnq.backPointer))) = a
	if mnq.backPointer+1 >= Balls {
		mnq.backPointer = 0
	} else {
		mnq.backPointer++
	}

	mnq.length++
	*(*int)(unsafe.Pointer(uintptr(mnq.first) + mnq.size*uintptr(mnq.backPointer))) = b
	if mnq.backPointer+1 >= Balls {
		mnq.backPointer = 0
	} else {
		mnq.backPointer++
	}

	mnq.length++
	*(*int)(unsafe.Pointer(uintptr(mnq.first) + mnq.size*uintptr(mnq.backPointer))) = c
	if mnq.backPointer+1 >= Balls {
		mnq.backPointer = 0
	} else {
		mnq.backPointer++
	}

	mnq.length++
	*(*int)(unsafe.Pointer(uintptr(mnq.first) + mnq.size*uintptr(mnq.backPointer))) = d
	if mnq.backPointer+1 >= Balls {
		mnq.backPointer = 0
	} else {
		mnq.backPointer++
	}

	mnq.length++
	*(*int)(unsafe.Pointer(uintptr(mnq.first) + mnq.size*uintptr(mnq.backPointer))) = e
	if mnq.backPointer+1 >= Balls {
		mnq.backPointer = 0
	} else {
		mnq.backPointer++
	}

	mnq.length++
	*(*int)(unsafe.Pointer(uintptr(mnq.first) + mnq.size*uintptr(mnq.backPointer))) = f
	if mnq.backPointer+1 >= Balls {
		mnq.backPointer = 0
	} else {
		mnq.backPointer++
	}

	mnq.length++
	*(*int)(unsafe.Pointer(uintptr(mnq.first) + mnq.size*uintptr(mnq.backPointer))) = g
	if mnq.backPointer+1 >= Balls {
		mnq.backPointer = 0
	} else {
		mnq.backPointer++
	}

	mnq.length++
	*(*int)(unsafe.Pointer(uintptr(mnq.first) + mnq.size*uintptr(mnq.backPointer))) = h
	if mnq.backPointer+1 >= Balls {
		mnq.backPointer = 0
	} else {
		mnq.backPointer++
	}
	mnq.length++
	*(*int)(unsafe.Pointer(uintptr(mnq.first) + mnq.size*uintptr(mnq.backPointer))) = i
	if mnq.backPointer+1 >= Balls {
		mnq.backPointer = 0
	} else {
		mnq.backPointer++
	}
	mnq.length++
	*(*int)(unsafe.Pointer(uintptr(mnq.first) + mnq.size*uintptr(mnq.backPointer))) = j
	if mnq.backPointer+1 >= Balls {
		mnq.backPointer = 0
	} else {
		mnq.backPointer++
	}
	mnq.length++
	*(*int)(unsafe.Pointer(uintptr(mnq.first) + mnq.size*uintptr(mnq.backPointer))) = k
	if mnq.backPointer+1 >= Balls {
		mnq.backPointer = 0
	} else {
		mnq.backPointer++
	}
}

func (mnq *MainQueue) print() []int {
	var rtnVal []int
	for i:=mnq.frontPointer; i<Balls; i++ {
		rtnVal = append(rtnVal, mnq.array[i])
	}
	return rtnVal
}

func IsOriginalPosition(mnq *MainQueue) bool{
	if mnq.length != Balls {
		return false
	}
	// fmt.Printf("main array: %v\n", mnq.print())
	for i:=0;i<Balls;i++ {
		if mnq.array[i] != (i+1) {
			return false
		}
	}
	return true
}

func (mnq *MainQueue) get(index int) int {
	return *(*int)(unsafe.Pointer(uintptr(mnq.first) + mnq.size*uintptr(index)))
}

func (mnq *MainQueue) pop() int{
	mnq.length--
	mnq.oldFront = mnq.get(mnq.frontPointer)
	if mnq.frontPointer+1 >= Balls {
		mnq.frontPointer = 0
	} else {
		mnq.frontPointer++
	}
	return mnq.oldFront
}