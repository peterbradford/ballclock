package ballclock

type HourQueue struct {
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

func (hq *HourQueue) append(x int) {
	if hq.pointer == 0 {
		hq.zero = x
		hq.pointer++
	} else if hq.pointer == 1 {
		hq.first = x
		hq.pointer++
	} else if hq.pointer == 2 {
		hq.second = x
		hq.pointer++
	} else if hq.pointer == 3 {
		hq.third = x
		hq.pointer++
	} else if hq.pointer == 4 {
		hq.fourth = x
		hq.pointer++
	} else if hq.pointer == 5 {
		hq.fifth = x
		hq.pointer++
	} else if hq.pointer == 6 {
		hq.sixth = x
		hq.pointer++
	} else if hq.pointer == 7 {
		hq.seventh = x
		hq.pointer++
	} else if hq.pointer == 8 {
		hq.eighth = x
		hq.pointer++
	} else if hq.pointer == 9 {
		hq.nineth = x
		hq.pointer++
	} else if hq.pointer == 10 {
		hq.tenth = x
		hq.pointer++
	}
}

func (hq *HourQueue) reverse() [11]int{
	return [11]int{hq.tenth, hq.nineth, hq.eighth, hq.seventh, hq.sixth, hq.fifth, hq.fourth, hq.third,hq.second,hq.first,hq.zero}
}

