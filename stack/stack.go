package stack

type StackString struct {
	elements []string
}

func (s *StackString) Push(elem string) {
	s.elements = append(s.elements, elem)
}

func (s *StackString) Empty() bool {
	return len(s.elements) == 0
}

func (s *StackString) Top() (string, bool) {
	if s.Empty() {
		return "", false
	}
	return s.elements[len(s.elements)-1], true
}

func (s *StackString) Pop() (string, bool) {
	elem, ok := s.Top()
	if !ok {
		return elem, ok
	}
	s.elements = s.elements[:len(s.elements)-1]
	return elem, ok
}
