package data

type Stack[T any] struct {
	buf    []T
	tail   int
	size   int
	minCap int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		buf:    make([]T, 8),
		tail:   0,
		size:   0,
		minCap: 8,
	}
}

func (s *Stack[T]) Len() int {
	if s == nil {
		return 0
	}
	return s.size
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Push(item T) {
	if s == nil {
		panic("Stack: called Push on nil stack.")
	}
	if s.full() {
		s.grow()
	}
	s.buf[s.tail] = item
	s.tail++
	s.size++
}

func (s *Stack[T]) Pop() T {
	if s.Empty() {
		panic("Stack: called Pop on empty stack.")
	}
	if s.underused() {
		s.shrink()
	}
	s.tail--
	ret := s.buf[s.tail]
	var zero T
	s.buf[s.tail] = zero
	s.size--
	return ret
}

func (s *Stack[T]) Peek() T {
	if s.Empty() {
		panic("Stack: called Peek on empty stack.")
	}
	return s.buf[s.tail-1]
}

func (s *Stack[T]) cap() int {
	if s == nil {
		return 0
	}
	return len(s.buf)
}

func (s *Stack[T]) full() bool {
	return s.Len() == s.cap()
}

func (s *Stack[T]) underused() bool {
	return len(s.buf) > s.minCap && (s.size<<2) == len(s.buf)
}

func (s *Stack[T]) grow() {
	s.resize(s.cap() << 1)
}

func (s *Stack[T]) shrink() {
	s.resize(s.cap() >> 1)
}

func (s *Stack[T]) resize(newSize int) {
	newBuf := make([]T, newSize)
	copy(newBuf, s.buf)
	s.buf = newBuf
}
