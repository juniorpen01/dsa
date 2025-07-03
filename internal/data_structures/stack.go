package datastructures

type Stack struct {
	Items []int
}

func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Size() int {
	return len(s.Items)
}

func (s *Stack) Peek() (int, bool) {
	if s.Size() == 0 {
		return 0, false
	}
	return s.Items[s.Size()-1], true
}

func (s *Stack) Pop() (int, bool) {
	if s.Size() == 0 {
		return 0, false
	}
	val := s.Items[s.Size()-1]
	s.Items = s.Items[:s.Size()-1]
	return val, true
}
