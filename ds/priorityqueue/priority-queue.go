/*
A Priority Queue is an advance version of a regular queue. Instead of following strict FIFO order, elements in a Priority Queue are assigned a priority.
Elements with a higher priority are dequeued before elements with lower priority. In the event of a tie in priority, they are dequeued in the order they
were added.
*/

package priorityqueue

import (
	"errors"
	"sort"
)

type Item struct {
	value    string
	priority int
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Enqueue(v string, p int) {
	item := &Item{
		value:    v,
		priority: p,
	}

	*pq = append(*pq, item)
	sort.SliceStable(*pq, func(i, j int) bool {
		return (*pq)[i].priority > (*pq)[j].priority
	})
}

func (pq *PriorityQueue) Dequeue() (string, error) {
	if pq.IsEmpty() {
		return "", errors.New("PriorityQueue is empty")
	}

	res := (*pq)[0].value
	*pq = (*pq)[1:]

	return res, nil
}

func (pq *PriorityQueue) Peek() (string, error) {
	if pq.IsEmpty() {
		return "", errors.New("PriorityQueue is empty")
	}

	res := (*pq)[0].value

	return res, nil
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(*pq) == 0
}
