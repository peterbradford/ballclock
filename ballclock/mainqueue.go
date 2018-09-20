package ballclock

type MainQueue struct {
	frontPointer int
	backPointer  int
	length		 int
	oldFront	 int

	num0th int
	num1th int
	num2th int
	num3th int
	num4th int
	num5th int
	num6th int
	num7th int
	num8th int
	num9th int
	
	num10th int
	num11th int
	num12th int
	num13th int
	num14th int
	num15th int
	num16th int
	num17th int
	num18th int
	num19th int

	num20th int
	num21th int
	num22th int
	num23th int
	num24th int
	num25th int
	num26th int
	num27th int
	num28th int
	num29th int

	num30th int
	num31th int
	num32th int
	num33th int
	num34th int
	num35th int
	num36th int
	num37th int
	num38th int
	num39th int

	num40th int
	num41th int
	num42th int
	num43th int
	num44th int
	num45th int
	num46th int
	num47th int
	num48th int
	num49th int

	num50th int
	num51th int
	num52th int
	num53th int
	num54th int
	num55th int
	num56th int
	num57th int
	num58th int
	num59th int

	num60th int
	num61th int
	num62th int
	num63th int
	num64th int
	num65th int
	num66th int
	num67th int
	num68th int
	num69th int

	num70th int
	num71th int
	num72th int
	num73th int
	num74th int
	num75th int
	num76th int
	num77th int
	num78th int
	num79th int

	num80th int
	num81th int
	num82th int
	num83th int
	num84th int
	num85th int
	num86th int
	num87th int
	num88th int
	num89th int

	num90th int
	num91th int
	num92th int
	num93th int
	num94th int
	num95th int
	num96th int
	num97th int
	num98th int
	num99th int

	num100th int
	num101th int
	num102th int
	num103th int
	num104th int
	num105th int
	num106th int
	num107th int
	num108th int
	num109th int
	
	num110th int
	num111th int
	num112th int
	num113th int
	num114th int
	num115th int
	num116th int
	num117th int
	num118th int
	num119th int

	num120th int
	num121th int
	num122th int
	num123th int
	num124th int
	num125th int
	num126th int
	num127th int
}

func (mnq *MainQueue) init() {
	mnq.num0th = 0
	mnq.num1th = 1
	mnq.num2th = 2
	mnq.num3th = 3
	mnq.num4th = 4
	mnq.num5th = 5
	mnq.num6th = 6
	mnq.num7th = 7
	mnq.num8th = 8	
	mnq.num9th = 9
	
	mnq.num10th = 10
	mnq.num11th = 11	
	mnq.num12th = 12	
	mnq.num13th = 13	
	mnq.num14th = 14	
	mnq.num15th = 15	
	mnq.num16th = 16	
	mnq.num17th = 17	
	mnq.num18th = 18	
	mnq.num19th = 19

	mnq.num20th = 20
	mnq.num21th = 21
	mnq.num22th = 22
	mnq.num23th = 23
	mnq.num24th = 24
	mnq.num25th = 25
	mnq.num26th = 26
	mnq.num27th = 27
	mnq.num28th = 28
	mnq.num29th = 29

	mnq.num30th = 30
	mnq.num31th = 31
	mnq.num32th = 32
	mnq.num33th = 33
	mnq.num34th = 34
	mnq.num35th = 35
	mnq.num36th = 36
	mnq.num37th = 37
	mnq.num38th = 38
	mnq.num39th = 39

	mnq.num40th = 40
	mnq.num41th = 41
	mnq.num42th = 42
	mnq.num43th = 43
	mnq.num44th = 44
	mnq.num45th = 45
	mnq.num46th = 46
	mnq.num47th = 47
	mnq.num48th = 48
	mnq.num49th = 49

	mnq.num50th = 50
	mnq.num51th = 51
	mnq.num52th = 52
	mnq.num53th = 53
	mnq.num54th = 54
	mnq.num55th = 55
	mnq.num56th = 56
	mnq.num57th = 57
	mnq.num58th = 58
	mnq.num59th = 59

	mnq.num60th = 60
	mnq.num61th = 61
	mnq.num62th = 62
	mnq.num63th = 63
	mnq.num64th = 64
	mnq.num65th = 65
	mnq.num66th = 66
	mnq.num67th = 67
	mnq.num68th = 68
	mnq.num69th = 69

	mnq.num70th = 70
	mnq.num71th = 71
	mnq.num72th = 72
	mnq.num73th = 73
	mnq.num74th = 74
	mnq.num75th = 75
	mnq.num76th = 76
	mnq.num77th = 77
	mnq.num78th = 78
	mnq.num79th = 79

	mnq.num80th = 80
	mnq.num81th = 81
	mnq.num82th = 82
	mnq.num83th = 83
	mnq.num84th = 84
	mnq.num85th = 85
	mnq.num86th = 86
	mnq.num87th = 87
	mnq.num88th = 88
	mnq.num89th = 89

	mnq.num90th = 90
	mnq.num91th = 91
	mnq.num92th = 92
	mnq.num93th = 93
	mnq.num94th = 94
	mnq.num95th = 95
	mnq.num96th = 96
	mnq.num97th = 97
	mnq.num98th = 98
	mnq.num99th = 99

	mnq.num100th = 100
	mnq.num101th = 101
	mnq.num102th = 102
	mnq.num103th = 103
	mnq.num104th = 104
	mnq.num105th = 105
	mnq.num106th = 106
	mnq.num107th = 107
	mnq.num108th = 108
	mnq.num109th = 109

	mnq.num110th = 110
	mnq.num111th = 111
	mnq.num112th = 112
	mnq.num113th = 113
	mnq.num114th = 114
	mnq.num115th = 115
	mnq.num116th = 116
	mnq.num117th = 117
	mnq.num118th = 118
	mnq.num119th = 119

	mnq.num120th = 120
	mnq.num121th = 121
	mnq.num122th = 122
	mnq.num123th = 123
	mnq.num124th = 124
	mnq.num125th = 125
	mnq.num126th = 126
	mnq.num127th = 127
} 
 
func (mnq *MainQueue) append(x int) {
	mnq.length++
	if mnq.backPointer+1 > Balls {
		mnq.backPointer = 0
	}

	if mnq.backPointer == 0 {					mnq.num0th = x
	} else if mnq.backPointer == 1 {			mnq.num1th = x
	} else if mnq.backPointer == 2 {			mnq.num2th = x
	} else if mnq.backPointer == 3 {			mnq.num3th = x
	} else if mnq.backPointer == 4 {			mnq.num4th = x
	} else if mnq.backPointer == 5 {			mnq.num5th = x
	} else if mnq.backPointer == 6 {			mnq.num6th = x
	} else if mnq.backPointer == 7 {			mnq.num7th = x
	} else if mnq.backPointer == 8 {			mnq.num8th = x
	} else if mnq.backPointer == 9 {			mnq.num9th = x
	
	} else if mnq.backPointer == 10 {			mnq.num10th = x
	} else if mnq.backPointer == 11 {			mnq.num11th = x
	} else if mnq.backPointer == 12 {			mnq.num12th = x
	} else if mnq.backPointer == 13 {			mnq.num13th = x
	} else if mnq.backPointer == 14 {			mnq.num14th = x
	} else if mnq.backPointer == 15 {			mnq.num15th = x
	} else if mnq.backPointer == 16 {			mnq.num16th = x
	} else if mnq.backPointer == 17 {			mnq.num17th = x
	} else if mnq.backPointer == 18 {			mnq.num18th = x
	} else if mnq.backPointer == 19 {			mnq.num19th = x

	} else if mnq.backPointer == 20 {			mnq.num20th = x
	} else if mnq.backPointer == 21 {			mnq.num21th = x
	} else if mnq.backPointer == 22 {			mnq.num22th = x
	} else if mnq.backPointer == 23 {			mnq.num23th = x
	} else if mnq.backPointer == 24 {			mnq.num24th = x
	} else if mnq.backPointer == 25 {			mnq.num25th = x
	} else if mnq.backPointer == 26 {			mnq.num26th = x
	} else if mnq.backPointer == 27 {			mnq.num27th = x
	} else if mnq.backPointer == 28 {			mnq.num28th = x
	} else if mnq.backPointer == 29 {			mnq.num29th = x	

	} else if mnq.backPointer == 30 {			mnq.num30th = x
	} else if mnq.backPointer == 31 {			mnq.num31th = x
	} else if mnq.backPointer == 32 {			mnq.num32th = x
	} else if mnq.backPointer == 33 {			mnq.num33th = x
	} else if mnq.backPointer == 34 {			mnq.num34th = x
	} else if mnq.backPointer == 35 {			mnq.num35th = x
	} else if mnq.backPointer == 36 {			mnq.num36th = x
	} else if mnq.backPointer == 37 {			mnq.num37th = x
	} else if mnq.backPointer == 38 {			mnq.num38th = x
	} else if mnq.backPointer == 39 {			mnq.num39th = x

	} else if mnq.backPointer == 40 {			mnq.num40th = x
	} else if mnq.backPointer == 41 {			mnq.num41th = x
	} else if mnq.backPointer == 42 {			mnq.num42th = x
	} else if mnq.backPointer == 43 {			mnq.num43th = x
	} else if mnq.backPointer == 44 {			mnq.num44th = x
	} else if mnq.backPointer == 45 {			mnq.num45th = x
	} else if mnq.backPointer == 46 {			mnq.num46th = x
	} else if mnq.backPointer == 47 {			mnq.num47th = x
	} else if mnq.backPointer == 48 {			mnq.num48th = x
	} else if mnq.backPointer == 49 {			mnq.num49th = x

	} else if mnq.backPointer == 50 {			mnq.num50th = x
	} else if mnq.backPointer == 51 {			mnq.num51th = x
	} else if mnq.backPointer == 52 {			mnq.num52th = x
	} else if mnq.backPointer == 53 {			mnq.num53th = x
	} else if mnq.backPointer == 54 {			mnq.num54th = x
	} else if mnq.backPointer == 55 {			mnq.num55th = x
	} else if mnq.backPointer == 56 {			mnq.num56th = x
	} else if mnq.backPointer == 57 {			mnq.num57th = x
	} else if mnq.backPointer == 58 {			mnq.num58th = x
	} else if mnq.backPointer == 59 {			mnq.num59th = x

	} else if mnq.backPointer == 60 {			mnq.num60th = x
	} else if mnq.backPointer == 61 {			mnq.num61th = x
	} else if mnq.backPointer == 62 {			mnq.num62th = x
	} else if mnq.backPointer == 63 {			mnq.num63th = x
	} else if mnq.backPointer == 64 {			mnq.num64th = x
	} else if mnq.backPointer == 65 {			mnq.num65th = x
	} else if mnq.backPointer == 66 {			mnq.num66th = x
	} else if mnq.backPointer == 67 {			mnq.num67th = x
	} else if mnq.backPointer == 68 {			mnq.num68th = x
	} else if mnq.backPointer == 69 {			mnq.num69th = x

	} else if mnq.backPointer == 70 {			mnq.num70th = x
	} else if mnq.backPointer == 71 {			mnq.num71th = x
	} else if mnq.backPointer == 72 {			mnq.num72th = x
	} else if mnq.backPointer == 73 {			mnq.num73th = x
	} else if mnq.backPointer == 74 {			mnq.num74th = x
	} else if mnq.backPointer == 75 {			mnq.num75th = x
	} else if mnq.backPointer == 76 {			mnq.num76th = x
	} else if mnq.backPointer == 77 {			mnq.num77th = x
	} else if mnq.backPointer == 78 {			mnq.num78th = x
	} else if mnq.backPointer == 79 {			mnq.num79th = x

	} else if mnq.backPointer == 80 {			mnq.num80th = x
	} else if mnq.backPointer == 81 {			mnq.num81th = x
	} else if mnq.backPointer == 82 {			mnq.num82th = x
	} else if mnq.backPointer == 83 {			mnq.num83th = x
	} else if mnq.backPointer == 84 {			mnq.num84th = x
	} else if mnq.backPointer == 85 {			mnq.num85th = x
	} else if mnq.backPointer == 86 {			mnq.num86th = x
	} else if mnq.backPointer == 87 {			mnq.num87th = x
	} else if mnq.backPointer == 88 {			mnq.num88th = x
	} else if mnq.backPointer == 89 {			mnq.num89th = x
		
	} else if mnq.backPointer == 90 {			mnq.num90th = x
	} else if mnq.backPointer == 91 {			mnq.num91th = x
	} else if mnq.backPointer == 92 {			mnq.num92th = x
	} else if mnq.backPointer == 93 {			mnq.num93th = x
	} else if mnq.backPointer == 94 {			mnq.num94th = x
	} else if mnq.backPointer == 95 {			mnq.num95th = x
	} else if mnq.backPointer == 96 {			mnq.num96th = x
	} else if mnq.backPointer == 97 {			mnq.num97th = x
	} else if mnq.backPointer == 98 {			mnq.num98th = x
	} else if mnq.backPointer == 99 {			mnq.num99th = x

	} else if mnq.backPointer == 100 {			mnq.num100th = x
	} else if mnq.backPointer == 101 {			mnq.num101th = x
	} else if mnq.backPointer == 102 {			mnq.num102th = x
	} else if mnq.backPointer == 103 {			mnq.num103th = x
	} else if mnq.backPointer == 104 {			mnq.num104th = x
	} else if mnq.backPointer == 105 {			mnq.num105th = x
	} else if mnq.backPointer == 106 {			mnq.num106th = x
	} else if mnq.backPointer == 107 {			mnq.num107th = x
	} else if mnq.backPointer == 108 {			mnq.num108th = x
	} else if mnq.backPointer == 109 {			mnq.num109th = x

	} else if mnq.backPointer == 110 {			mnq.num110th = x
	} else if mnq.backPointer == 111 {			mnq.num111th = x
	} else if mnq.backPointer == 112 {			mnq.num112th = x
	} else if mnq.backPointer == 113 {			mnq.num113th = x
	} else if mnq.backPointer == 114 {			mnq.num114th = x
	} else if mnq.backPointer == 115 {			mnq.num115th = x
	} else if mnq.backPointer == 116 {			mnq.num116th = x
	} else if mnq.backPointer == 117 {			mnq.num117th = x
	} else if mnq.backPointer == 118 {			mnq.num118th = x
	} else if mnq.backPointer == 119 {			mnq.num119th = x

	} else if mnq.backPointer == 120 {			mnq.num120th = x
	} else if mnq.backPointer == 121 {			mnq.num121th = x
	} else if mnq.backPointer == 122 {			mnq.num122th = x
	} else if mnq.backPointer == 123 {			mnq.num123th = x
	} else if mnq.backPointer == 124 {			mnq.num124th = x
	} else if mnq.backPointer == 125 {			mnq.num125th = x
	} else if mnq.backPointer == 126 {			mnq.num126th = x
	} else if mnq.backPointer == 127 {			mnq.num127th = x
	}

	

	//cycle to back to zero if it goes over the Balls amount,
	
	mnq.backPointer++
	
}

func (mnq *MainQueue) print() [128]int {
	return [128]int{
		mnq.num0th,
		mnq.num1th,
		mnq.num2th,  
		mnq.num3th,
		mnq.num4th,
		mnq.num5th,
		mnq.num6th,
		mnq.num7th,
		mnq.num8th,
		mnq.num9th,
		mnq.num10th, 
		mnq.num11th, 
		mnq.num12th, 
		mnq.num13th, 
		mnq.num14th, 
		mnq.num15th, 
		mnq.num16th, 
		mnq.num17th, 
		mnq.num18th, 
		mnq.num19th, 
		mnq.num20th, 
		mnq.num21th, 
		mnq.num22th, 
		mnq.num23th, 
		mnq.num24th, 
		mnq.num25th, 
		mnq.num26th, 
		mnq.num27th, 
		mnq.num28th, 
		mnq.num29th, 
		mnq.num30th, 
		mnq.num31th, 
		mnq.num32th, 
		mnq.num33th, 
		mnq.num34th, 
		mnq.num35th, 
		mnq.num36th, 
		mnq.num37th, 
		mnq.num38th, 
		mnq.num39th, 
		mnq.num40th, 
		mnq.num41th, 
		mnq.num42th, 
		mnq.num43th, 
		mnq.num44th, 
		mnq.num45th, 
		mnq.num46th, 
		mnq.num47th, 
		mnq.num48th, 
		mnq.num49th, 
		mnq.num50th, 
		mnq.num51th, 
		mnq.num52th, 
		mnq.num53th, 
		mnq.num54th, 
		mnq.num55th, 
		mnq.num56th, 
		mnq.num57th, 
		mnq.num58th, 
		mnq.num59th, 
		mnq.num60th, 
		mnq.num61th, 
		mnq.num62th, 
		mnq.num63th, 
		mnq.num64th, 
		mnq.num65th, 
		mnq.num66th, 
		mnq.num67th, 
		mnq.num68th, 
		mnq.num69th, 
		mnq.num70th, 
		mnq.num71th, 
		mnq.num72th, 
		mnq.num73th, 
		mnq.num74th, 
		mnq.num75th, 
		mnq.num76th, 
		mnq.num77th, 
		mnq.num78th, 
		mnq.num79th, 
		mnq.num80th, 
		mnq.num81th, 
		mnq.num82th, 
		mnq.num83th, 
		mnq.num84th, 
		mnq.num85th, 
		mnq.num86th, 
		mnq.num87th, 
		mnq.num88th, 
		mnq.num89th, 
		mnq.num90th, 
		mnq.num91th, 
		mnq.num92th, 
		mnq.num93th, 
		mnq.num94th, 
		mnq.num95th, 
		mnq.num96th, 
		mnq.num97th, 
		mnq.num98th, 
		mnq.num99th, 
		mnq.num100th,
		mnq.num101th,
		mnq.num102th,
		mnq.num103th,
		mnq.num104th,
		mnq.num105th,
		mnq.num106th,
		mnq.num107th,
		mnq.num108th,
		mnq.num109th,
		mnq.num110th,
		mnq.num111th,
		mnq.num112th,
		mnq.num113th,
		mnq.num114th,
		mnq.num115th,
		mnq.num116th,
		mnq.num117th,
		mnq.num118th,
		mnq.num119th,
		mnq.num120th,
		mnq.num121th,
		mnq.num122th,
		mnq.num123th,
		mnq.num124th,
		mnq.num125th,
		mnq.num126th,
		mnq.num127th,
	}
}

func IsOriginalPosition(mnq MainQueue) bool{
	if mnq.length != Balls {
		return false
	}
	if 	mnq.num0th != 1 { return false 
	} else if mnq.num1th != 2 { return false
	} else if mnq.num2th != 3 { return false
	} else if mnq.num3th != 4 { return false
	} else if mnq.num4th != 5 { return false
	} else if mnq.num5th != 6 { return false
	} else if mnq.num6th != 7 { return false
	} else if mnq.num7th != 8 { return false
	} else if mnq.num8th != 9 { return false
	
	} else if mnq.num9th != 10 { return false
	} else if mnq.num10th != 11 { return false 
	} else if mnq.num11th != 12 { return false
	} else if mnq.num12th != 13 { return false
	} else if mnq.num13th != 14 { return false
	} else if mnq.num14th != 15 { return false
	} else if mnq.num15th != 16 { return false
	} else if mnq.num16th != 17 { return false
	} else if mnq.num17th != 18 { return false
	} else if mnq.num18th != 19 { return false
	
	} else if mnq.num19th != 20 { return false
	} else if mnq.num20th != 21 { return false 
	} else if mnq.num21th != 22 { return false
	} else if mnq.num22th != 23 { return false
	} else if mnq.num23th != 24 { return false
	} else if mnq.num24th != 25 { return false
	} else if mnq.num25th != 26 { return false
	} else if mnq.num26th != 27 { return false
	} else if mnq.num27th != 28 { return false
	} else if mnq.num28th != 29 { return false

	} else if mnq.num29th != 30 { return false
	} else if mnq.num30th != 31 { return false 
	} else if mnq.num31th != 32 { return false
	} else if mnq.num32th != 33 { return false
	} else if mnq.num33th != 34 { return false
	} else if mnq.num34th != 35 { return false
	} else if mnq.num35th != 36 { return false
	} else if mnq.num36th != 37 { return false
	} else if mnq.num37th != 38 { return false
	} else if mnq.num38th != 39 { return false

	} else if mnq.num39th != 40 { return false
	} else if mnq.num40th != 41 { return false 
	} else if mnq.num41th != 42 { return false
	} else if mnq.num42th != 43 { return false
	} else if mnq.num43th != 44 { return false
	} else if mnq.num44th != 45 { return false
	} else if mnq.num45th != 46 { return false
	} else if mnq.num46th != 47 { return false
	} else if mnq.num47th != 48 { return false
	} else if mnq.num48th != 49 { return false
	
	} else if mnq.num49th != 50 { return false
	} else if mnq.num50th != 51 { return false 
	} else if mnq.num51th != 52 { return false
	} else if mnq.num52th != 53 { return false
	} else if mnq.num53th != 54 { return false
	} else if mnq.num54th != 55 { return false
	} else if mnq.num55th != 56 { return false
	} else if mnq.num56th != 57 { return false
	} else if mnq.num57th != 58 { return false
	} else if mnq.num58th != 59 { return false
	
	} else if mnq.num59th != 60 { return false
	} else if mnq.num60th != 61 { return false 
	} else if mnq.num61th != 62 { return false
	} else if mnq.num62th != 63 { return false
	} else if mnq.num63th != 64 { return false
	} else if mnq.num64th != 65 { return false
	} else if mnq.num65th != 66 { return false
	} else if mnq.num66th != 67 { return false
	} else if mnq.num67th != 68 { return false
	} else if mnq.num68th != 69 { return false

	} else if mnq.num69th != 70 { return false
	} else if mnq.num70th != 71 { return false 
	} else if mnq.num71th != 72 { return false
	} else if mnq.num72th != 73 { return false
	} else if mnq.num73th != 74 { return false
	} else if mnq.num74th != 75 { return false
	} else if mnq.num75th != 76 { return false
	} else if mnq.num76th != 77 { return false
	} else if mnq.num77th != 78 { return false
	} else if mnq.num78th != 79 { return false

	} else if mnq.num79th != 80 { return false
	} else if mnq.num80th != 81 { return false 
	} else if mnq.num81th != 82 { return false
	} else if mnq.num82th != 83 { return false
	} else if mnq.num83th != 84 { return false
	} else if mnq.num84th != 85 { return false
	} else if mnq.num85th != 86 { return false
	} else if mnq.num86th != 87 { return false
	} else if mnq.num87th != 88 { return false
	} else if mnq.num88th != 89 { return false

	} else if mnq.num89th != 90 { return false
	} else if mnq.num90th != 91 { return false 
	} else if mnq.num91th != 92 { return false
	} else if mnq.num92th != 93 { return false
	} else if mnq.num93th != 94 { return false
	} else if mnq.num94th != 95 { return false
	} else if mnq.num95th != 96 { return false
	} else if mnq.num96th != 97 { return false
	} else if mnq.num97th != 98 { return false
	} else if mnq.num98th != 99 { return false

	} else if mnq.num99th  != 100 { return false
	} else if mnq.num100th != 101 { return false 
	} else if mnq.num101th != 102 { return false
	} else if mnq.num102th != 103 { return false
	} else if mnq.num103th != 104 { return false
	} else if mnq.num104th != 105 { return false
	} else if mnq.num105th != 106 { return false
	} else if mnq.num106th != 107 { return false
	} else if mnq.num107th != 108 { return false
	} else if mnq.num108th != 109 { return false

	} else if mnq.num109th != 110 { return false
	} else if mnq.num110th != 111 { return false 
	} else if mnq.num111th != 112 { return false
	} else if mnq.num112th != 113 { return false
	} else if mnq.num113th != 114 { return false
	} else if mnq.num114th != 115 { return false
	} else if mnq.num115th != 116 { return false
	} else if mnq.num116th != 117 { return false
	} else if mnq.num117th != 118 { return false
	} else if mnq.num118th != 119 { return false
	
	} else if mnq.num119th != 120 { return false
	} else if mnq.num120th != 121 { return false 
	} else if mnq.num121th != 122 { return false
	} else if mnq.num122th != 123 { return false
	} else if mnq.num123th != 124 { return false
	} else if mnq.num124th != 125 { return false
	} else if mnq.num125th != 126 { return false
	} else if mnq.num126th != 127 { return false
	} else if mnq.num127th != 128 { return false}
	
	return true
}

func (mnq *MainQueue) get(index int) int {
	if index == 0 { return mnq.num0th
	}else if index == 1 { return mnq.num1th
	}else if index == 2 { return mnq.num2th
	}else if index == 3 { return mnq.num3th
	}else if index == 4 { return mnq.num4th
	}else if index == 5 { return mnq.num5th
	}else if index == 6 { return mnq.num6th
	}else if index == 7 { return mnq.num7th
	}else if index == 8 { return mnq.num8th
	}else if index == 9 { return mnq.num9th

	}else if index == 10 { return mnq.num10th
	}else if index == 11 { return mnq.num11th
	}else if index == 12 { return mnq.num12th
	}else if index == 13 { return mnq.num13th
	}else if index == 14 { return mnq.num14th
	}else if index == 15 { return mnq.num15th
	}else if index == 16 { return mnq.num16th
	}else if index == 17 { return mnq.num17th
	}else if index == 18 { return mnq.num18th
	}else if index == 19 { return mnq.num19th

	}else if index == 20 { return mnq.num20th
	}else if index == 21 { return mnq.num21th
	}else if index == 22 { return mnq.num22th
	}else if index == 23 { return mnq.num23th
	}else if index == 24 { return mnq.num24th
	}else if index == 25 { return mnq.num25th
	}else if index == 26 { return mnq.num26th
	}else if index == 27 { return mnq.num27th
	}else if index == 28 { return mnq.num28th
	}else if index == 29 { return mnq.num29th

	}else if index == 30 { return mnq.num30th
	}else if index == 31 { return mnq.num31th
	}else if index == 32 { return mnq.num32th
	}else if index == 33 { return mnq.num33th
	}else if index == 34 { return mnq.num34th
	}else if index == 35 { return mnq.num35th
	}else if index == 36 { return mnq.num36th
	}else if index == 37 { return mnq.num37th
	}else if index == 38 { return mnq.num38th
	}else if index == 39 { return mnq.num39th

	}else if index == 40 { return mnq.num40th
	}else if index == 41 { return mnq.num41th
	}else if index == 42 { return mnq.num42th
	}else if index == 43 { return mnq.num43th
	}else if index == 44 { return mnq.num44th
	}else if index == 45 { return mnq.num45th
	}else if index == 46 { return mnq.num46th
	}else if index == 47 { return mnq.num47th
	}else if index == 48 { return mnq.num48th
	}else if index == 49 { return mnq.num49th

	}else if index == 50 { return mnq.num50th
	}else if index == 51 { return mnq.num51th
	}else if index == 52 { return mnq.num52th
	}else if index == 53 { return mnq.num53th
	}else if index == 54 { return mnq.num54th
	}else if index == 55 { return mnq.num55th
	}else if index == 56 { return mnq.num56th
	}else if index == 57 { return mnq.num57th
	}else if index == 58 { return mnq.num58th
	}else if index == 59 { return mnq.num59th

	}else if index == 60 { return mnq.num60th
	}else if index == 61 { return mnq.num61th
	}else if index == 62 { return mnq.num62th
	}else if index == 63 { return mnq.num63th
	}else if index == 64 { return mnq.num64th
	}else if index == 65 { return mnq.num65th
	}else if index == 66 { return mnq.num66th
	}else if index == 67 { return mnq.num67th
	}else if index == 68 { return mnq.num68th
	}else if index == 69 { return mnq.num69th

	}else if index == 70 { return mnq.num70th
	}else if index == 71 { return mnq.num71th
	}else if index == 72 { return mnq.num72th
	}else if index == 73 { return mnq.num73th
	}else if index == 74 { return mnq.num74th
	}else if index == 75 { return mnq.num75th
	}else if index == 76 { return mnq.num76th
	}else if index == 77 { return mnq.num77th
	}else if index == 78 { return mnq.num78th
	}else if index == 79 { return mnq.num79th

	}else if index == 80 { return mnq.num80th
	}else if index == 81 { return mnq.num81th
	}else if index == 82 { return mnq.num82th
	}else if index == 83 { return mnq.num83th
	}else if index == 84 { return mnq.num84th
	}else if index == 85 { return mnq.num85th
	}else if index == 86 { return mnq.num86th
	}else if index == 87 { return mnq.num87th
	}else if index == 88 { return mnq.num88th
	}else if index == 89 { return mnq.num89th

	}else if index == 90 { return mnq.num90th
	}else if index == 91 { return mnq.num91th
	}else if index == 92 { return mnq.num92th
	}else if index == 93 { return mnq.num93th
	}else if index == 94 { return mnq.num94th
	}else if index == 95 { return mnq.num95th
	}else if index == 96 { return mnq.num96th
	}else if index == 97 { return mnq.num97th
	}else if index == 98 { return mnq.num98th
	}else if index == 99 { return mnq.num99th

	}else if index == 100 { return mnq.num100th
	}else if index == 101 { return mnq.num101th
	}else if index == 102 { return mnq.num102th
	}else if index == 103 { return mnq.num103th
	}else if index == 104 { return mnq.num104th
	}else if index == 105 { return mnq.num105th
	}else if index == 106 { return mnq.num106th
	}else if index == 107 { return mnq.num107th
	}else if index == 108 { return mnq.num108th
	}else if index == 109 { return mnq.num109th

	}else if index == 110 { return mnq.num110th
	}else if index == 111 { return mnq.num111th
	}else if index == 112 { return mnq.num112th
	}else if index == 113 { return mnq.num113th
	}else if index == 114 { return mnq.num114th
	}else if index == 115 { return mnq.num115th
	}else if index == 116 { return mnq.num116th
	}else if index == 117 { return mnq.num117th
	}else if index == 118 { return mnq.num118th
	}else if index == 119 { return mnq.num119th

	}else if index == 120 { return mnq.num120th
	}else if index == 121 { return mnq.num121th
	}else if index == 122 { return mnq.num122th
	}else if index == 123 { return mnq.num123th
	}else if index == 124 { return mnq.num124th
	}else if index == 125 { return mnq.num125th
	}else if index == 126 { return mnq.num126th
	}else if index == 127 { return mnq.num127th}
	return -1
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