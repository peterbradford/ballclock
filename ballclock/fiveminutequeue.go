package ballclock

type FiveMinuteQueue struct {
	pointer int
	zero int
	first int
	second int
	third int
	fourth int
	fifth int
	sixth int
	seventh int
	eighth int
	nineth int
	tenth int
}

func (fmq *FiveMinuteQueue) append(x int) {
	if fmq.pointer == 0 {
		fmq.zero = x
		fmq.pointer++
	} else if fmq.pointer == 1 {
		fmq.first = x
		fmq.pointer++
	} else if fmq.pointer == 2 {
		fmq.second = x
		fmq.pointer++
	} else if fmq.pointer == 3 {
		fmq.third = x
		fmq.pointer++
	} else if fmq.pointer == 4 {
		fmq.fourth = x
		fmq.pointer++
	} else if fmq.pointer == 5 {
		fmq.fifth = x
		fmq.pointer++
	} else if fmq.pointer == 6 {
		fmq.sixth = x
		fmq.pointer++
	} else if fmq.pointer == 7 {
		fmq.seventh = x
		fmq.pointer++
	} else if fmq.pointer == 8 {
		fmq.eighth = x
		fmq.pointer++
	} else if fmq.pointer == 9 {
		fmq.nineth = x
		fmq.pointer++
	} else if fmq.pointer == 10 {
		fmq.tenth = x
		fmq.pointer++
	}
}

func (fmq *FiveMinuteQueue) reverse() [11]int{
	return [11]int{fmq.tenth, fmq.nineth, fmq.eighth, fmq.seventh, fmq.sixth, fmq.fifth, fmq.fourth, fmq.third,fmq.second,fmq.first,fmq.zero}
}

