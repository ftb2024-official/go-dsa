/*
A Linked List is a foundational data structure, comprising a collection of nodes, where each node contains a value and a reference to the next node in
the sequence. This design enables efficient insertion or removal of elements from any position within the sequence.
*/

package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Append(v int) {
	newNode := &Node{data: v}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) DeleteWithValue(v int) {
	if ll.head == nil {
		return
	}

	if ll.head.data == v {
		ll.head = ll.head.next
		return
	}

	current := ll.head
	for current.next != nil && current.next.data != v {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}
