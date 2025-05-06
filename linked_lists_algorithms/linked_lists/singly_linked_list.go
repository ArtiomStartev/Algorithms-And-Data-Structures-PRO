package linked_lists

import (
	"errors"
	"fmt"
)

type SinglyLinkedList[T comparable] struct {
	Head   *SLLNode[T]
	Tail   *SLLNode[T]
	Length int
}

type SLLNode[T comparable] struct {
	Value T
	Next  *SLLNode[T]
}

type FindSLLNodeResult[T comparable] struct {
	Node *SLLNode[T]
	Prev *SLLNode[T]
}

// Insert inserts a new node with the given value at the end of the linked list
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
func (sll *SinglyLinkedList[T]) Find(value T) (FindSLLNodeResult[T], error) {
	curr := sll.Head
	var prev *SLLNode[T]

	for curr != nil {
		if curr.Value == value {
			return FindSLLNodeResult[T]{Node: curr, Prev: prev}, nil
		}
		prev = curr
		curr = curr.Next
	}

	return FindSLLNodeResult[T]{}, fmt.Errorf("value %v not found in linked list", value)
}

// Remove removes the first node with the given value from the linked list
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

	curr := sll.Head
	for curr.Next != nil {
		if curr.Next.Value == value {
			curr.Next = curr.Next.Next
			sll.Length--
			if curr.Next == nil {
				sll.Tail = curr
			}
			return nil
		}
		curr = curr.Next
	}

	return fmt.Errorf("value %v not found in linked list", value)
}

// RemoveAfter removes the node after the given node from the linked list
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
func (sll *SinglyLinkedList[T]) AssertNoCycles() error {
	curr := sll.Head
	var count int

	for curr != nil {
		count++
		if count > sll.Length {
			return errors.New("cycle detected in linked list")
		}
		curr = curr.Next
	}

	return nil
}
