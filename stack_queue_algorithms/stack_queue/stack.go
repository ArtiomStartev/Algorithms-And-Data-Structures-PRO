package stack_queue

import (
	"errors"
)

// Stack is an interface that both ArrayStack and LinkedListStack should implement
type Stack[T any] interface {
	IsEmpty() bool
	Push(value T)
	GetLatestElement() (*T, error)
	Pop() error
}

// ----------------------- Dynamic Array Stack -----------------------

// ArrayStack is a stack implemented using a dynamic array.
type ArrayStack[T any] struct {
	data []T
}

// NewArrayStack creates a new ArrayStack.
func NewArrayStack[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{}
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *ArrayStack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Push adds a new element to the top of the stack.
func (s *ArrayStack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// GetLatestElement returns the latest element added to the stack without removing it.
func (s *ArrayStack[T]) GetLatestElement() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return &s.data[len(s.data)-1], nil
}

// Pop removes the latest element added to the stack.
func (s *ArrayStack[T]) Pop() error {
	if s.IsEmpty() {
		return errors.New("stack is empty")
	}
	s.data = s.data[:len(s.data)-1]
	return nil
}

// ----------------------- Linked List Stack -----------------------

// StackNode represents a single node in the linked list stack.
type StackNode[T any] struct {
	Value T
	Next  *StackNode[T]
}

// LinkedListStack is a stack implemented using a linked list.
type LinkedListStack[T any] struct {
	Head *StackNode[T]
}

// NewLinkedListStack creates a new LinkedListStack.
func NewLinkedListStack[T any]() *LinkedListStack[T] {
	return &LinkedListStack[T]{}
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *LinkedListStack[T]) IsEmpty() bool {
	return s.Head == nil
}

// Push adds a new element to the top of the stack.
func (s *LinkedListStack[T]) Push(value T) {
	newNode := &StackNode[T]{Value: value, Next: s.Head}
	s.Head = newNode
}

// GetLatestElement returns the latest element added to the stack without removing it.
func (s *LinkedListStack[T]) GetLatestElement() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return &s.Head.Value, nil
}

// Pop removes the latest element added to the stack.
func (s *LinkedListStack[T]) Pop() error {
	if s.IsEmpty() {
		return errors.New("stack is empty")
	}
	s.Head = s.Head.Next
	return nil
}
