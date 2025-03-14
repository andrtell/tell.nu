package data

type Queue[T any] struct {
	buf    []T
	head   int
	tail   int
	size   int
	minCap int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		buf:    make([]T, 8),
		head:   0,
		tail:   0,
		size:   0,
		minCap: 8,
	}
}

func (q *Queue[T]) Len() int {
	if q == nil {
		return 0
	}
	return q.size
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}

func (q *Queue[T]) Peek() T {
	if q.Empty() {
		panic("Queue: called Peek on empty queue.")
	}
	return q.buf[q.head]
}

func (q *Queue[T]) Enqueue(item T) {
	if q == nil {
		panic("Queue: called Enqueue on nil.")
	}
	if q.full() {
		q.grow()
	}
	q.buf[q.tail] = item
	q.tail = q.next(q.tail)
	q.size++
}

func (q *Queue[T]) Dequeue() T {
	if q.Empty() {
		panic("Queue: called Dequeue on empty queue.")
	}
	if q.underused() {
		q.shrink()
	}
	ret := q.buf[q.head]
	var zero T
	q.buf[q.head] = zero
	q.head = q.next(q.head)
	q.size--
	return ret
}

func (q *Queue[T]) cap() int {
	if q == nil {
		return 0
	}
	return len(q.buf)
}

func (q *Queue[T]) full() bool {
	return q.Len() == q.cap()
}

func (q *Queue[T]) underused() bool {
	return len(q.buf) > q.minCap && (q.size<<2) == len(q.buf)
}

func (q *Queue[T]) grow() {
	q.resize(q.cap() << 1)
}

func (q *Queue[T]) shrink() {
	q.resize(q.cap() >> 1)
}

func (q *Queue[T]) resize(newSize int) {
	newBuf := make([]T, newSize)
	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}
	q.head = 0
	q.tail = q.size
	q.buf = newBuf
}

func (q *Queue[T]) prev(i int) int {
	return (i - 1) & (len(q.buf) - 1)
}

func (q *Queue[T]) next(i int) int {
	return (i + 1) & (len(q.buf) - 1)
}
