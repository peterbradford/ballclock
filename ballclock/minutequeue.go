package ballclock

type MinuteQueue struct {
	pointer int
	zero int
	first int
	second int
	third int
}

func (mq *MinuteQueue) append(x int) {
	if mq.pointer == 0 {
		mq.zero = x
		mq.pointer++
	} else if mq.pointer == 1 {
		mq.first = x
		mq.pointer++
	} else if mq.pointer == 2 {
		mq.second = x
		mq.pointer++
	} else if mq.pointer == 3 {
		mq.third = x
		mq.pointer++
	}
}

func (mq *MinuteQueue) reverse() [4]int{
	return [4]int{mq.third,mq.second,mq.first,mq.zero}
}

