package datastructures

type Queue struct {
	Items []int
}

func (s *Queue) Push(item int) {
	s.Items = append([]int{item}, s.Items...)
}

func (s *Queue) Size() int {
	return len(s.Items)
}

func (s *Queue) Peek() (int, bool) {
	if s.Size() == 0 {
		return 0, false
	}
	return s.Items[s.Size()-1], true
}

func (s *Queue) Pop() (int, bool) {
	if s.Size() == 0 {
		return 0, false
	}
	val := s.Items[s.Size()-1]
	s.Items = s.Items[:s.Size()-1]
	return val, true
}
