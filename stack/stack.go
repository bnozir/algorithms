package stack

type StackInt struct {
	elements []int
}

func (s *StackInt) Push(elem int) {
	s.elements = append(s.elements, elem)
}

func (s *StackInt) Empty() bool {
	return len(s.elements) == 0
}

func (s *StackInt) Top() (int, bool) {
	if s.Empty() {
		return 0, false
	}
	return s.elements[len(s.elements)-1], true
}

func (s *StackInt) Pop() (int, bool) {
	elem, ok := s.Top()
	if !ok {
		return elem, ok
	}
	s.elements = s.elements[:len(s.elements)-1]
	return elem, ok
}
