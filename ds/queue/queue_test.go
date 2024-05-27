package queue

import "testing"

func TestEnqueue(t *testing.T) {
	var q Queue

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	if len(q) != 3 {
		t.Errorf("expected length 3, got %v", len(q))
	}

	if q[0] != 10 || q[1] != 20 || q[2] != 30 {
		t.Errorf("expected [10, 20, 30], got %v", q)
	}
}

func TestDequeue(t *testing.T) {
	var q Queue

	if v, err := q.Dequeue(); err == nil || v != 0 {
		t.Errorf("expected (0, Queue is empty), got (%v, %v)", v, err)
	}

	q.Enqueue(10)
	q.Enqueue(20)

	if v, err := q.Dequeue(); err != nil || v != 10 {
		t.Errorf("expected (10, nil), got (%v, %v)", v, err)
	}

	if len(q) != 1 {
		t.Errorf("expected length 1, got %v", len(q))
	}
}

func TestFront(t *testing.T) {
	var q Queue

	if v, err := q.Front(); v != 0 || err == nil {
		t.Errorf("expected (0, Queue is empty), got (%v, %v)", v, err)
	}

	q.Enqueue(20)
	q.Enqueue(10)

	if v, err := q.Front(); v != 20 || err != nil {
		t.Errorf("expected (20, nil), got (%v, %v)", v, err)
	}

	if len(q) != 2 {
		t.Errorf("expected length 2, got %v", len(q))
	}
}

func TestIsEmpty(t *testing.T) {
	var q Queue

	if v := q.IsEmpty(); v != true {
		t.Errorf("expected true, got %v", v)
	}

	q.Enqueue(10)

	if v := q.IsEmpty(); v != false {
		t.Errorf("expected false, got %v", v)
	}

	_, _ = q.Dequeue()

	if v := q.IsEmpty(); v != true {
		t.Errorf("expected true, got %v", v)
	}
}
