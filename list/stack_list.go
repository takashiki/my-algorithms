package list

type StackList struct {
	S1 []int
	S2 []int
}

func (l *StackList) Pop() (int, bool) {
	s2Len := len(l.S2)
	if s2Len == 0 {
		for {
			s1Len := len(l.S1)
			if s1Len == 0 {
				break
			}
			l.S2 = append(l.S2, l.S1[s1Len-1])
			l.S1 = l.S1[:s1Len-1]
		}
	}

	s2Len = len(l.S2)
	if s2Len == 0 {
		return 0, false
	} else {
		res := l.S2[s2Len-1]
		l.S2 = l.S2[:s2Len-1]
		return res, true
	}
}

func (l *StackList) Push(n int) {
	l.S1 = append(l.S1, n)
}
