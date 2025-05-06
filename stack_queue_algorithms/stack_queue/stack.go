package stack_queue

import (
	"errors"
)

type Stack[T any] interface {
	IsEmpty() bool
	Push(value T)
	GetLatestElement() (*T, error)
	Pop() error
}

// ----------------------- Dynamic Array Stack -----------------------

type ArrayStack[T any] struct {
	data []T
}

func NewArrayStack[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{}
}

// IsEmpty checks if the stack is empty
func (s *ArrayStack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Push adds an element to the top of the stack
func (s *ArrayStack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// GetLatestElement retrieves the top element of the stack without removing it
func (s *ArrayStack[T]) GetLatestElement() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return &s.data[len(s.data)-1], nil
}

// Pop removes the top element from the stack
func (s *ArrayStack[T]) Pop() error {
	if s.IsEmpty() {
		return errors.New("stack is empty")
	}

	s.data[len(s.data)-1] = *new(T) // Clear reference to allow GC to free memory
	s.data = s.data[:len(s.data)-1]

	return nil
}

// ----------------------- Linked List Stack -----------------------

type StackNode[T any] struct {
	Value T
	Next  *StackNode[T]
}

type LinkedListStack[T any] struct {
	Head *StackNode[T]
}

func NewLinkedListStack[T any]() *LinkedListStack[T] {
	return &LinkedListStack[T]{}
}

// IsEmpty checks if the stack is empty
func (s *LinkedListStack[T]) IsEmpty() bool {
	return s.Head == nil
}

// Push adds an element to the top of the stack
func (s *LinkedListStack[T]) Push(value T) {
	newNode := &StackNode[T]{Value: value, Next: s.Head}
	s.Head = newNode
}

// GetLatestElement retrieves the top element of the stack without removing it
func (s *LinkedListStack[T]) GetLatestElement() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return &s.Head.Value, nil
}

// Pop removes the top element from the stack
func (s *LinkedListStack[T]) Pop() error {
	if s.IsEmpty() {
		return errors.New("stack is empty")
	}
	s.Head = s.Head.Next
	return nil
}
