package linked_lists

import (
	"errors"
	"fmt"
)

// SLLNode represents a node in a singly linked list
type SLLNode[T comparable] struct {
	Value T
	Next  *SLLNode[T]
}

// FindSLLNodeResult represents the result of a find operation in a singly linked list
type FindSLLNodeResult[T comparable] struct {
	Node *SLLNode[T]
	Prev *SLLNode[T]
}

// SinglyLinkedList represents a singly linked list
type SinglyLinkedList[T comparable] struct {
	Head   *SLLNode[T]
	Tail   *SLLNode[T]
	Length int
}

// Insert inserts a new node with the given value at the end of the linked list
// It returns the newly inserted node
func (sll *SinglyLinkedList[T]) Insert(value T) *SLLNode[T] {
	newNode := &SLLNode[T]{Value: value}

	if sll.Head == nil {
		sll.Head = newNode
		sll.Tail = newNode
	} else {
		sll.Tail.Next = newNode
		sll.Tail = newNode
	}

	sll.Length++
	return newNode
}

// InsertAfter inserts a new node with the given value after the given node
// It returns the newly inserted node
func (sll *SinglyLinkedList[T]) InsertAfter(node *SLLNode[T], value T) *SLLNode[T] {
	newNode := &SLLNode[T]{Value: value}

	if node == nil {
		newNode.Next = sll.Head
		sll.Head = newNode
		if sll.Tail == nil {
			sll.Tail = newNode
		}
	} else {
		newNode.Next = node.Next
		node.Next = newNode
		if node == sll.Tail {
			sll.Tail = newNode
		}
	}

	sll.Length++
	return newNode
}

// Find searches for a node with the given value in the linked list
// It returns the node and the previous node if the value is found, otherwise an error
func (sll *SinglyLinkedList[T]) Find(value T) (*FindSLLNodeResult[T], error) {
	current := sll.Head
	var prev *SLLNode[T]

	for current != nil {
		if current.Value == value {
			return &FindSLLNodeResult[T]{Node: current, Prev: prev}, nil
		}
		prev = current
		current = current.Next
	}

	return nil, fmt.Errorf("value %v not found in linked list", value)
}

// Remove removes the node with the given value from the linked list
// It returns an error if the value is not found
func (sll *SinglyLinkedList[T]) Remove(value T) error {
	if sll.Head == nil {
		return errors.New("linked list is empty")
	}

	if sll.Head.Value == value {
		sll.Head = sll.Head.Next
		sll.Length--
		if sll.Head == nil {
			sll.Tail = nil
		}
		return nil
	}

	current := sll.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			sll.Length--
			if current.Next == nil {
				sll.Tail = current
			}
			return nil
		}
		current = current.Next
	}

	return fmt.Errorf("value %v not found in linked list", value)
}

// RemoveAfter removes the node after the given node
// It returns an error if the node has no next node to remove
func (sll *SinglyLinkedList[T]) RemoveAfter(node *SLLNode[T]) error {
	if node == nil {
		if sll.Head != nil {
			sll.Head = sll.Head.Next
			sll.Length--
			if sll.Head == nil {
				sll.Tail = nil
			}
			return nil
		} else {
			return errors.New("linked list is empty")
		}
	}

	if node.Next != nil {
		node.Next = node.Next.Next
		sll.Length--
		if node.Next == nil {
			sll.Tail = node
		}
		return nil
	}

	return errors.New("node has no next node to remove")
}

// AssertNoCycles checks if the linked list has any cycles
// It returns an error if a cycle is detected
func (sll *SinglyLinkedList[T]) AssertNoCycles() error {
	current := sll.Head
	var count int

	for current != nil {
		count++
		if count > sll.Length {
			return errors.New("cycle detected in linked list")
		}
		current = current.Next
	}

	return nil
}
