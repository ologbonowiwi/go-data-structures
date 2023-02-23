package datastructures

type stack[T any] struct {
	size   int
	top    int
	vector []T
}

func BuildStack[T any](size int) stack[T] {
	s := stack[T]{}

	s.size = size
	s.vector = []T{}
	s.top = -1

	return s
}

func (s *stack[T]) IsEmpty() bool {
	return s.top == -1
}

func (s *stack[T]) IsFull() bool {
	return s.top == s.size-1
}

func (s *stack[T]) Push(value T) {
	if s.IsFull() {
		return
	}

	s.top += 1
	s.vector = append(s.vector, value)
}

func (s *stack[T]) Pop() T {
	element := s.vector[s.top]
	s.vector = s.vector[:s.top]
	s.top -= 1

	return element
}

func (s *stack[T]) Peek() T {
	return s.vector[s.top]
}

func (s *stack[T]) Size() int {
	return s.size
}

func (s *stack[T]) Vector() []T {
	return s.vector
}

func (s *stack[T]) TopIndex() int {
	return s.top
}
