package internal

type Stack[T any] struct {
	data []*T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, &v)
}

func (s *Stack[T]) Pop() *T {
	if len(s.data) == 0 {
		return nil
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Top() *T {
	if len(s.data) == 0 {
		return nil
	}

	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Clear() {
	s.data = make([]*T, 0)
}
