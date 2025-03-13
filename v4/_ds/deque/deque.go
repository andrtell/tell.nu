// package deque
//
// type Deque[T any] struct {
// 	buf    []T
// 	head   int
// 	tail   int
// 	size   int
// 	minCap int
// }
//
// func New[T any]() *Deque[T] {
// 	return new(Deque[T])
// }
//
// /*
//  * Cap returns the current capacity of the Deque q.
//  *
//  * Cases:
//  *
//  * A) If q (a pointer) is `nil` then q.Cap() is zero.
//  *
//  * nil => Cap() = 0
//  *
//  * B) Any other
//  *
//  * [? ?]     => Cap() = 2
//  * [? ? ? ?] => Cap() = 4
//  * ...
//  *
//  */
// func (q *Deque[T]) Cap() int {
// 	if q == nil {
// 		return 0
// 	}
// 	return len(q.buf)
// }
//
// /*
//  * Len returns the number of items currently stored in the queue.
//  *
//  * Cases:
//  *
//  * A) If q (a pointer) is `nil` then Len() is zero.
//  *
//  * nil => Len() = 0
//  *
//  * B) The tail is `after` the head.
//  *
//  * [_ 1 2 3 _ _ _ _] => Len() = 3
//  *    h     t
//  *
//  * C) The tail is `before` the head.
//  *
//  * [5 _ _ _ 1 2 3 4] => Len() = 5
//  *    t     h
//  *
//  */
// func (q *Deque[T]) Len() int {
// 	if q == nil {
// 		return 0
// 	}
// 	return q.size
// }
//
// /*
//  * PushBack appends an item to the back of the queue.
//  *
//  * Note that `head` and `tail` are asymmetric in their use.
//  *
//  * - `head` points at the position of the `first` item.
//  * - `tail` points at the position one after the `last` item.
//  *
//  * This makes it easy to `capture` all items using:
//  *
//  *	buf[q.head:q.tail]				(if q.tail > q.head)
//  *  buf[q.head:] and buf[:q.tail]	(if q.tail <= q.head)
//  *
//  * Cases:
//  *
//  * A) The queue is full and needs to grow.
//  *
//  * [2 2 1 1] => (grow) => [1 1 2 2 _ _ _ _] => (push) => [1 1 2 2 9 _ _ _]
//  *      h                  h       t                      h         t
//  *      t
//  *
//  * B0) The tail is `after` the head.
//  *
//  * [_ 1 1 1 _ _ _ _] => (push) => [_ 1 1 1 9 _ _ _]
//  *    h     t                        h       t
//  *
//  * B1) The tail is at the `end` and needs to wrap around.
//  *
//  * [_ _ _ 1 1 1 1 _] => (push) => [_ _ _ 1 1 1 1 9]
//  *        h       t                t     h
//  *
//  * C) The tail is `before` the head.
//  *
//  * [_ _ _ _ 1 1 1 1] => (push) => [9 _ _ _ 1 1 1 1]
//  *  t       h                        t     h
//  *
//  */
// func (q *Deque[T]) PushBack(item T) {
// 	if q.IsFull() {
// 		q.Grow()
// 	}
// 	q.buf[q.tail] = item
// 	q.tail = q.next(q.tail)
// 	q.size++
// }
//
// /*
//  * PushFront prepends an item to the front of the queue.
//  *
//  * Cases:
//  *
//  * A) The queue is full and needs to grow.
//  *
//  * [2 2 1 1] => (grow) => [1 1 2 2 _ _ _ _] => (push) => [1 1 2 2 _ _ _ 9]
//  *      h                  h       t                              t     h
//  *      t
//  *
//  * B0) The head is `before` the tail.
//  *
//  * [_ 1 1 1 _ _ _ _] => (push) => [9 1 1 1 _ _ _ _]
//  *    h     t                      h       t
//  *
//  * B1) The head is at the `start` and needs to wrap around.
//  *
//  * [1 1 _ _] => (push) => [1 1 _ 9]
//  *  h   t                      t h
//  *
//  * C) The head is `after` the tail.
//  *
//  * [_ _ _ _ 1 1 1 1] => (push) => [_ _ _ 9 1 1 1 1]
//  *  t       h                      t     h
//  *
//  */
// func (q *Deque[T]) PushFront(item T) {
// 	if q.IsFull() {
// 		q.Grow()
// 	}
// 	q.head = q.prev(q.head)
// 	q.buf[q.head] = item
// 	q.size++
// }
//
// /*
//  * PopBack removes and returns the last item of the queue.
//  *
//  * Cases:
//  *
//  * A) The queue is empty, panic!
//  *
//  * [_ _ _ _] => panic!
//  *  h
//  *  t
//  *
//  * B) The tail is `after` the head.
//  *
//  * [_ 1 1 7 _ _ _ _] => (pop) => [_ 1 1 _ _ _ _ _] => 7
//  *    h     t                       h   t
//  *
//  * C0) The tail is `before` the head.
//  *
//  * [7 _ _ 1 1 1 1 1] => (pop) => [_ _ _ 1 1 1 1 1] => 7
//  *    t   h                       t     h
//  *
//  * C1) The tail is at the `start` and needs to wrap around.
//  *
//  * [_ _ _ 1 1 1 1 7] => (pop) => [_ _ _ 1 1 1 1 _] => 7
//  *  t     h                             h       t
//  *
//  */
// func (q *Deque[T]) PopBack() T {
// 	if q.size <= 0 {
// 		panic("deque: PopBack() called on empty queue.")
// 	}
// 	q.tail = q.prev(q.tail)
// 	ret := q.buf[q.tail]
// 	var zero T
// 	q.buf[q.tail] = zero
// 	q.size--
// 	return ret
// }
//
// /*
//  * PopFront removes and returns the first item from the queue.
//  *
//  * Cases:
//  *
//  * A) The queue is empty, panic!
//  *
//  * [_ _ _ _] => panic!
//  *  h
//  *  t
//  *
//  * B0) The head is `before` the tail.
//  *
//  * [_ 7 1 1 _ _ _ _] => (pop) => [_ _ 1 1 _ _ _ _] => 7
//  *    h     t                         h   t
//  *
//  * B1) The head is at the `end` and needs to wrap around.
//  *
//  * [1 _ _ 7] => (pop) => [1 _ _ _] => 7
//  *    t   h               h t
//  *
//  * C) The head is `after` the tail.
//  *
//  * [_ 7 1 1] => (pop) => [_ _ 1 1] => 7
//  *  t h                   t   h
//  *
//  */
// func (q *Deque[T]) PopFront() T {
// 	if q.size <= 0 {
// 		panic("deque: PopFront() called on empty queue.")
// 	}
// 	ret := q.buf[q.head]
// 	var zero T
// 	q.buf[q.head] = zero
// 	q.head = q.next(q.head)
// 	q.size--
// 	return ret
// }
//
// /*
//  * Front returns the first item in the queue.
//  */
// func (q *Deque[T]) Front() T {
// 	if q.size <= 0 {
// 		panic("deque: Front() called on empty queue.")
// 	}
// 	return q.buf[q.head]
// }
//
// /*
//  * Back returns the last item in the queue.
//  */
// func (q *Deque[T]) Back() T {
// 	if q.size <= 0 {
// 		panic("deque: Back() called on empty queue.")
// 	}
// 	return q.buf[q.prev(q.tail)]
// }
//
// func (q *Deque[T]) At(i int) T {
// 	q.checkRange(i)
// 	return q.buf[(q.head+i)&(len(q.buf)-1)]
// }
//
// // Set
// func (q *Deque[T]) Set(i int, item T) {
// 	q.checkRange(i)
// 	q.buf[(q.head+i)&(len(q.buf)-1)] = item
// }
//
// /*
//  *
//  */
// func (q *Deque[T]) IsFull() bool {
// 	return q.size == len(q.buf)
// }
//
// func (q *Deque[T]) prev(i int) int {
// 	// 10 % 2 <=> 1010 % 10 = 1 or 0
// 	// 10 % 4 <=> 1010 % 100 = 01, 10 or 11
// 	// 0b10 - 1 = 0b01
// 	// 0b100 - 1 = 0b011
// 	// Assuming len(q.buf) is a power of 2
// 	return (i - 1) & (len(q.buf) - 1)
// }
//
// func (q *Deque[T]) next(i int) int {
// 	return (i + 1) & (len(q.buf) - 1)
// }
