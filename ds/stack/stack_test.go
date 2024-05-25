package stack

import "testing"

func TestPush(t *testing.T) {
	var s Stack

	s.Push(10)
	s.Push(20)
	s.Push(30)

	if len(s) != 3 {
		t.Errorf("expected length 3, got %v", len(s))
	}

	if s[0] != 10 || s[1] != 20 || s[2] != 30 {
		t.Errorf("expected [10, 20, 30], got %v", s)
	}
}

func TestPop(t *testing.T) {
	var s Stack

	if _, err := s.Pop(); err == nil {
		t.Errorf("expected error, got %v", err)
	}

	s.Push(10)
	s.Push(20)

	val, err := s.Pop()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if val != 20 {
		t.Errorf("expected 20, got %v", val)
	}

	if len(s) != 1 {
		t.Errorf("expected length 1, got %v", len(s))
	}

	val, _ = s.Pop()
	if val != 10 {
		t.Errorf("expected 10, got %v", val)
	}

	if len(s) != 0 {
		t.Errorf("expected length 0, got %v", len(s))
	}
}

func TestPeek(t *testing.T) {
	var s Stack

	if val, err := s.Peek(); val != 0 || err == nil {
		t.Errorf("expected 0 and error, got %v and %v", val, err)
	}

	s.Push(10)
	s.Push(20)

	val, err := s.Peek()
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if val != 20 {
		t.Errorf("expected 20, got %v", val)
	}

	if len(s) != 2 {
		t.Errorf("expected length 2, got %v", len(s))
	}
}

func TestIsEmpty(t *testing.T) {
	var s Stack

	if !s.IsEmpty() {
		t.Errorf("expected true, got false")
	}

	s.Push(10)
	if s.IsEmpty() {
		t.Errorf("expected false, got true")
	}

	_, _ = s.Pop()
	if !s.IsEmpty() {
		t.Errorf("expected true, got false")
	}
}
