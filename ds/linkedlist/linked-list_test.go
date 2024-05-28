package linkedlist

import "testing"

func TestAppend(t *testing.T) {
	var ll LinkedList

	ll.Append(10)
	ll.Append(20)
	ll.Append(30)

	if ll.head.data != 10 {
		t.Errorf("expected head data to be 10, got %v", ll.head.data)
	}

	if ll.head.data != 10 || ll.head.next.data != 20 || ll.head.next.next.data != 30 || ll.head.next.next.next != nil {
		t.Errorf("expected [10 -> 20 -> 30 -> nil], got [%v -> %v -> %v -> %v]", ll.head.data, ll.head.next.data, ll.head.next.next.data, ll.head.next.next.next)
	}
}

func TestDeleteWithValue(t *testing.T) {
	var ll LinkedList

	if err := ll.DeleteWithValue(1); err == nil {
		t.Errorf("expected 'LinkedList is empty', got %v", err)
	}

	ll.Append(10)
	ll.Append(20)
	ll.Append(30)

	if err := ll.DeleteWithValue(1); err == nil {
		t.Errorf("expected 'No such value in LinkedList', got %v", err)
	}

	if err := ll.DeleteWithValue(20); err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if ll.head.data != 10 || ll.head.next.data != 30 || ll.head.next.next != nil {
		t.Errorf("expected [10 -> 30 -> nil], got [%v -> %v -> %v]", ll.head.data, ll.head.next.data, ll.head.next.next)
	}
}

func TestDisplay(t *testing.T) {
	var ll LinkedList

	if err := ll.Display(); err == nil {
		t.Errorf("expected 'LinkedList is empty', got %v", err)
	}

	ll.Append(10)

	if err := ll.Display(); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestIsEmpty(t *testing.T) {
	var ll LinkedList

	if v := ll.IsEmpty(); v != true {
		t.Errorf("expected true, got %v", v)
	}

	ll.Append(10)

	if v := ll.IsEmpty(); v != false {
		t.Errorf("expected false, got %v", v)
	}

	_ = ll.DeleteWithValue(10)

	if v := ll.IsEmpty(); v != true {
		t.Errorf("expected true, got %v", v)
	}
}
