package priorityqueue

import "testing"

func TestEnqueue(t *testing.T) {
	var pq PriorityQueue

	pq.Enqueue("Task 1", 2)
	pq.Enqueue("Task 2", 3)
	pq.Enqueue("Task 3", 1)

	if len(pq) != 3 {
		t.Errorf("expected length 3, got %v", len(pq))
	}

	if pq[0].value != "Task 2" && pq[1].value != "Task 1" && pq[2].value != "Task 3" {
		t.Errorf("expected [Task 2, Task 1, Task 3], got [%v, %v, %v]", pq[0].value, pq[1].value, pq[2].value)
	}
}

func TestDequeue(t *testing.T) {
	var pq PriorityQueue

	if v, err := pq.Dequeue(); v != "" || err == nil {
		t.Errorf("expected ('', PriorityQueue is empty), got (%v, %v)", v, err)
	}

	pq.Enqueue("Task 1", 2)
	pq.Enqueue("Task 2", 3)

	if v, err := pq.Dequeue(); v != "Task 2" || err != nil {
		t.Errorf("expected ('Task 2', nil), got (%v, %v)", v, err)
	}

	if len(pq) != 1 {
		t.Errorf("expected length 1, got %v", len(pq))
	}
}

func TestPeek(t *testing.T) {
	var pq PriorityQueue

	if v, err := pq.Peek(); v != "" || err == nil {
		t.Errorf("expected ('', PriorityQueue is empty), got (%v, %v)", v, err)
	}

	pq.Enqueue("Task 10", 100)
	pq.Enqueue("Task 1", 99)

	if v, err := pq.Peek(); v != "Task 10" || err != nil {
		t.Errorf("expected ('Task 10', nil), got (%v, %v)", v, err)
	}

	if len(pq) != 2 {
		t.Errorf("expected length 2, got %v", len(pq))
	}
}

func TestIsEmpty(t *testing.T) {
	var pq PriorityQueue

	if v := pq.IsEmpty(); v != true {
		t.Errorf("expected true, got %v", v)
	}

	pq.Enqueue("Task 10", 100)

	if v := pq.IsEmpty(); v != false {
		t.Errorf("expected false, got %v", v)
	}

	_, _ = pq.Dequeue()

	if v := pq.IsEmpty(); v != true {
		t.Errorf("expected true, got %v", v)
	}
}
