/*
A Queue is a linear data structure that follows the First In First Out (FIFO) principle. This means the first element added to the queue is the first
one to be removed. It has two primary operations: 'enqueue', which adds an element to the rear end, and 'dequeue' which removes and returns the front
element.
*/

package queue

import "errors"

type Queue []int

func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}

	res := (*q)[0]
	*q = (*q)[1:]
	return res, nil
}

func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}
	return (*q)[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
