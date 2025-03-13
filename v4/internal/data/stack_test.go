package data

import (
	"testing"
	"slices"
)

func FuzzInOut(f *testing.F) {
	corpus := [][]byte{
		{}, 
		{1},
		{3, 1}, 
		{3, 2, 11}, 
		{7, 5, 13, 1}, 
		{11, 11, 11, 11}, 
		{97, 101, 13, 2, 7}, 
		{1, 43, 47, 53, 61, 67, 71, 73, 79, 83, 59},
	}
	for _, seed := range corpus {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, input []byte) {
		s := NewStack[byte]()
		for _, v := range input {
			s.Push(v)
		}
		output := make([]byte, 0)
		for !s.Empty() {
			output = append(output, s.Pop())
		}
		reversed := make([]byte, len(input))
		copy(reversed, input)
		slices.Reverse(reversed)	
		if slices.Compare(output, reversed) != 0 {
			t.Error("given", input, "got", output, "want", reversed)
		}
	})
}

func TestPushPop(t *testing.T) {
	seed := []int{1, 3, 7, 21, 33, 107}
	for _, count := range seed { 
		s := NewStack[int]()
		for i := 0; i < count; i++ {
			s.Push(i)
		}
		for i := count - 1; i >= 0; i-- {
			v := s.Pop()
			if v != i {
				t.Error("s.Pop() =", v, "want", i)
			}
		}
		if s.Len() != 0 {
			t.Error("s.Len() =", s.Len(), "want", 0)
		}
	}
}

func TestPushPeek(t *testing.T) {
	seed := []int{1, 3, 9, 23, 31, 99}
	for _, count := range seed { 
		s := NewStack[int]()
		for i := 0; i < count; i++ {
			s.Push(i)
			if s.Peek() != i {
				t.Error("s.Peek() =", s.Peek(), "want", i)
			}
		}
	}
}

func TestPushLen(t *testing.T) {
	seed := []int{1, 3, 7, 23, 79, 111}
	for _, count := range seed { 
		s := NewStack[int]()
		for v := 0; v < count; v++ {
			s.Push(v)
		}
		if s.Len() != count {
			t.Error("s.Len() =", s.Len(), "want", count)
		}
	}
}

func TestPushPopOnce(t *testing.T) {
	s := NewStack[int]()
	val := 21
	s.Push(val)
	got := s.Pop()
	if got != val {
		t.Error("s.Pop() =", got, "want", val)
	}
}

func TestPushPeekOnce(t *testing.T) {
	s := NewStack[int]()
	val := 21
	s.Push(val)
	if s.Peek() != val {
		t.Error("s.Peek() =", s.Peek(), "want", val)
	}
}

func TestEmpty(t *testing.T) {
	s := NewStack[int]()
	if s.Len() != 0 {
		t.Error("s.Len() =", s.Len(), "want 0")
	}
	if !s.Empty() {
		t.Error("s.Empty() =", s.Empty(), "want true")
	}
}

func TestEmptyPop(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("s.Pop() did not panic. Want panic.")
		}
	}()
	s := NewStack[int]()
	s.Pop()
}

func TestEmptyPeek(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("s.Peek() did not panic. Want panic.")
		}
	}()
	s := NewStack[int]()
	s.Peek()
}

func TestNilEmpty(t *testing.T) {
	var s *Stack[int]

	if s.Len() != 0 {
		t.Error("s.Len() =", s.Len(), "want 0")
	}

	if !s.Empty() {
		t.Error("s.Empty() =", s.Empty(), "want true")
	}
}

func TestNilPush(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Push() did not panic. Want panic.")
		}
	}()
	var s *Stack[int]
	s.Push(1)
}

func TestNilPop(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("s.Pop() did not panic. Want panic.")
		}
	}()
	var s *Stack[int]
	s.Pop()
}

func TestNilPeek(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("s.Peek() did not panic. Want panic.")
		}
	}()
	var s *Stack[int]
	s.Peek()
}
