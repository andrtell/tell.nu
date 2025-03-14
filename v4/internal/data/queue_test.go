package data

import (
	"testing"
	"slices"
)

func FuzzQueue10(f *testing.F) {
	corpus := [][]byte{
		{}, 
		{1},
		{3, 1}, 
		{3, 2, 11}, 
		{7, 5, 13, 1}, 
		{11, 11, 11, 11}, 
		{97, 101, 13, 2, 7}, 
		{1, 43, 47, 53, 61, 67, 71, 73, 79, 83, 59}, }
	for _, seed := range corpus {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, input []byte) {
		q := NewQueue[byte]()
		for _, v := range input {
			q.Enqueue(v)
		}
		output := make([]byte, 0)
		for !q.Empty() {
			output = append(output, q.Dequeue())
		}
		if slices.Compare(input, output) != 0 {
			t.Error("Enqueued", input, "Dequeued", output, "want", input)
		}
	})
}

func TestQueueEnqueueDequeue(t *testing.T) {
	q := NewQueue[int]()
	v := 1
	q.Enqueue(v)
	got := q.Dequeue()
	if got != v {
		t.Error("Call to Dequeue() returned", got, "want", v)
	}
	if q.Len() != 0 {
		t.Error("Call to Len() returned", q.Len(), "want", 1)
	}
}

func TestQueueEnqueuePeek(t *testing.T) {
	s := NewQueue[int]()
	v := 1
	s.Enqueue(v)
	if s.Peek() != v {
		t.Error("Call to Peek() returned", s.Peek(), "want", v)
	}
}

func TestQueueEnqueueLen(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	if q.Len() != 1 {
		t.Error("Call to Len() returned", q.Len(), "want", 1)
	}
}

func TestQueueEmptyOnEmpty(t *testing.T) {
	q := NewQueue[int]()
	if !q.Empty() {
		t.Error("Call to Empty() returned", q.Empty(), "want true")
	}
}

func TestQueueLenOnEmpty(t *testing.T) {
	q := NewQueue[int]()
	if q.Len() != 0 {
		t.Error("Call to Len() returned", q.Len(), "want 0")
	}
}

func TestQueueEmptyOnNil(t *testing.T) {
	var q *Queue[int]
	if !q.Empty() {
		t.Error("Call to Empty() returned", q.Empty(), "want true")
	}
}

func TestQueueLenOnNil(t *testing.T) {
	var q *Queue[int]
	if q.Len() != 0 {
		t.Error("Call to Len() returned", q.Len(), "want 0")
	}
}

func TestQueuePanicDequeueOnEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Call to Dequeue() did not panic. Want panic.")
		}
	}()
	q := NewQueue[int]()
	q.Dequeue()
}

func TestQueuePanicPeekOnEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Call to Peek() did not panic. Want panic.")
		}
	}()
	q := NewQueue[int]()
	q.Peek()
}

func TestQueuePanicEnqueueOnNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Call to Enqueue() did not panic. Want panic.")
		}
	}()
	var q *Queue[int]
	q.Enqueue(1)
}

func TestQueuePanicDequeueOnNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Call to Dequeue() did not panic. Want panic.")
		}
	}()
	var q *Queue[int]
	q.Dequeue()
}

func TestQueuePanicPeekOnNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Call to Peek() did not panic. Want panic.")
		}
	}()
	var q *Queue[int]
	q.Peek()
}

// Internal

func TestQueueMonotoneCap(t *testing.T) {
	q := NewQueue[int]()
	prevCap := q.cap()
	for i := 0; i < 131; i++ {
		q.Enqueue(i)
		if q.cap() < prevCap {
			t.Error("Cap did not grow (curr cap)", q.cap(), "< (prev cap)", prevCap)
		}
		prevCap = q.cap()
	}
	for i := 0; i < 131; i++ {
		q.Dequeue()
		if q.cap() > prevCap {
			t.Error("Cap did not shrink (curr cap)", q.cap(), "< (prev cap)", prevCap)
		}
		prevCap = q.cap()
	}
}
