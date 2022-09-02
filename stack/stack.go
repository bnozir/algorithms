package stack

type Stack struct {
	elements []int
}

func (s *Stack) Push(elem int) {
	s.elements = append(s.elements, elem)
}

func (s *Stack) Empty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Top() (int, bool) {
	if s.Empty() {
		return 0, false
	}
	return s.elements[len(s.elements)-1], true
}

func (s *Stack) Pop() (int, bool) {
	elem, ok := s.Top()
	if !ok {
		return elem, ok
	}
	s.elements = s.elements[:len(s.elements)-1]
	return elem, ok
}
