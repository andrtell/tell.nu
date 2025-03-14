package data

import (
	"testing"
	"slices"
)

func FuzzStack10(f *testing.F) {
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
			t.Error("Pushed", input, "Popped", output, "want", reversed)
		}
	})
}

func TestStackPushPop(t *testing.T) {
	s := NewStack[int]()
	v := 1
	s.Push(v)
	got := s.Pop()
	if got != v {
		t.Error("Call to Pop() returned", got, "want", v)
	}
}

func TestStackPushPeek(t *testing.T) {
	s := NewStack[int]()
	v := 1
	s.Push(v)
	if s.Peek() != v {
		t.Error("Call to Peek() returned", s.Peek(), "want", v)
	}
}

func TestStackPushLen(t *testing.T) {
	s := NewStack[int]()
	v := 1
	s.Push(v)
	if s.Len() != 1 {
		t.Error("Call to Len() returned", s.Len(), "want", 1)
	}
}

func TestStackEmptyOnEmpty(t *testing.T) {
	s := NewStack[int]()
	if !s.Empty() {
		t.Error("Call to Empty() returned", s.Empty(), "want true")
	}
}

func TestStackLenOnEmpty(t *testing.T) {
	s := NewStack[int]()
	if s.Len() != 0 {
		t.Error("Call to Len() returned", s.Len(), "want 0")
	}
}

func TestStackEmptyOnNil(t *testing.T) {
	var s *Stack[int]
	if !s.Empty() {
		t.Error("Call to Empty() returned", s.Empty(), "want true")
	}
}

func TestStackLenOnNil(t *testing.T) {
	var s *Stack[int]
	if s.Len() != 0 {
		t.Error("Call to Len() returned", s.Len(), "want 0")
	}
}

func TestStackPanicPopOnEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Call to Pop() did not panic. Want panic.")
		}
	}()
	s := NewStack[int]()
	s.Pop()
}

func TestStackPanicPeekOnEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Call to Peek() did not panic. Want panic.")
		}
	}()
	s := NewStack[int]()
	s.Peek()
}

func TestStackPanicPushOnNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Call to Push() did not panic. Want panic.")
		}
	}()
	var s *Stack[int]
	s.Push(1)
}

func TestStackPanicPopOnNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Call to Pop() did not panic. Want panic.")
		}
	}()
	var s *Stack[int]
	s.Pop()
}

func TestStackPanicPeekOnNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Call to Peek() did not panic. Want panic.")
		}
	}()
	var s *Stack[int]
	s.Peek()
}
