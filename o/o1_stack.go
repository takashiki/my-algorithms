package o

import "math"

type O1Stack struct {
	data []int
	min  []int
}

func (s *O1Stack) GetMin() int {
	length := len(s.min)
	if length == 0 {
		return math.MaxInt
	}

	return s.min[length-1]
}

func (s *O1Stack) Push(n int) {
	s.data = append(s.data, n)
	if n <= s.GetMin() {
		s.min = append(s.min, n)
	}
}

func (s *O1Stack) Pop() int {
	length := len(s.data)
	n := s.data[length-1]
	s.data = s.data[:length-1]
	if n == s.GetMin() {
		s.min = s.min[:len(s.min)-1]
	}

	return n
}
