/*
A Stack is a linear data structure that follows the Last In First Out (LIFO) principle, meaning the last element added to the stack is the first element
removed from it. Common operations associated with a stack include 'push' (to add an element) and 'pop' (to remove the top element). Another utility
operation is 'peek' or 'top', which returns the top element without popping it.
*/

package stack

import (
	"errors"
)

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("Stack is empty")
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, nil
}

func (s *Stack) Peek() (int, error) {
	if len(*s) == 0 {
		return -1, errors.New("Stack is empty")
	}
	return (*s)[len(*s)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
