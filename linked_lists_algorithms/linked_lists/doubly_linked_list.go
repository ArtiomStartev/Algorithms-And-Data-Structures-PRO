package linked_lists

import (
	"fmt"
)

type DLLNode[T comparable] struct {
	Value T
	Next  *DLLNode[T]
	Prev  *DLLNode[T]
}

type DoublyLinkedList[T comparable] struct {
	Head   *DLLNode[T]
	Tail   *DLLNode[T]
	Length int
}

// Insert inserts a new node with the given value at the end of the linked list
// It returns the newly inserted node
func (dll *DoublyLinkedList[T]) Insert(value T) *DLLNode[T] {
	newNode := &DLLNode[T]{Value: value}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		dll.Tail.Next = newNode
		newNode.Prev = dll.Tail
		dll.Tail = newNode
	}

	dll.Length++
	return newNode
}

// InsertBefore inserts a new node with the given value before the given node
// It returns the newly inserted node
func (dll *DoublyLinkedList[T]) InsertBefore(node *DLLNode[T], value T) *DLLNode[T] {
	newNode := &DLLNode[T]{Value: value}

	if node == nil {
		if dll.Head == nil {
			dll.Head = newNode
			dll.Tail = newNode
		} else {
			newNode.Next = dll.Head
			dll.Head.Prev = newNode
			dll.Head = newNode
		}
	} else {
		newNode.Prev = node.Prev
		newNode.Next = node
		node.Prev = newNode
		if newNode.Prev != nil {
			newNode.Prev.Next = newNode
		} else {
			dll.Head = newNode
		}
	}

	dll.Length++
	return newNode
}

// InsertAfter inserts a new node with the given value after the given node
// It returns the newly inserted node
func (dll *DoublyLinkedList[T]) InsertAfter(node *DLLNode[T], value T) *DLLNode[T] {
	newNode := &DLLNode[T]{Value: value}

	if node == nil {
		if dll.Head == nil {
			dll.Head = newNode
			dll.Tail = newNode
		} else {
			newNode.Next = dll.Head
			dll.Head.Prev = newNode
			dll.Head = newNode
		}
	} else {
		newNode.Prev = node
		newNode.Next = node.Next
		node.Next = newNode
		if newNode.Next != nil {
			newNode.Next.Prev = newNode
		} else {
			dll.Tail = newNode
		}

	}

	dll.Length++
	return newNode
}

// Find searches for a node with the given value in the linked list
// It returns the node if found, otherwise an error
func (dll *DoublyLinkedList[T]) Find(value T) (*DLLNode[T], error) {
	current := dll.Head

	for current != nil {
		if current.Value == value {
			return current, nil
		}
		current = current.Next
	}

	return nil, fmt.Errorf("value %v not found in the linked list", value)
}

// Remove removes the given node from the linked list
// It returns an error if the node is nil
func (dll *DoublyLinkedList[T]) Remove(node *DLLNode[T]) error {
	if node == nil {
		return fmt.Errorf("node is nil")
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		dll.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		dll.Tail = node.Prev
	}

	dll.Length--
	return nil
}

// AssertNoCycles checks for cycles in the doubly linked list
// It returns an error if a cycle is detected
func (dll *DoublyLinkedList[T]) AssertNoCycles() error {
	visited := make(map[*DLLNode[T]]bool)
	current := dll.Head

	for current != nil {
		if visited[current] {
			return fmt.Errorf("cycle detected in linked list")
		}
		visited[current] = true
		current = current.Next
	}

	return nil
}
