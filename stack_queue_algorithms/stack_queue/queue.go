package stack_queue

import "errors"

// ----------------------- Dynamic Array Queue -----------------------

// ArrayQueue is a queue implemented using a dynamic array.
type ArrayQueue[T any] struct {
	data []T
}

// NewArrayQueue creates a new ArrayQueue.
func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{}
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *ArrayQueue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Enqueue adds a new element to the end of the queue.
func (q *ArrayQueue[T]) Enqueue(value T) {
	q.data = append(q.data, value)
}

// GetFrontElement returns the element at the front of the queue without removing it.
func (q *ArrayQueue[T]) GetFrontElement() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return &q.data[0], nil
}

// Dequeue removes the element at the front of the queue.
func (q *ArrayQueue[T]) Dequeue() error {
	if q.IsEmpty() {
		return errors.New("queue is empty")
	}
	q.data = q.data[1:]
	return nil
}

// ----------------------- Linked List Queue -----------------------

// QueueNode represents a single node in the linked list queue.
type QueueNode[T any] struct {
	Value T
	Next  *QueueNode[T]
}

// LinkedListQueue is a queue implemented using a linked list.
type LinkedListQueue[T any] struct {
	Head *QueueNode[T]
	Tail *QueueNode[T]
}

// NewLinkedListQueue creates a new LinkedListQueue.
func NewLinkedListQueue[T any]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{}
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *LinkedListQueue[T]) IsEmpty() bool {
	return q.Head == nil
}

// Enqueue adds a new element to the end of the queue.
func (q *LinkedListQueue[T]) Enqueue(value T) {
	newNode := &QueueNode[T]{Value: value}

	if q.IsEmpty() {
		q.Head = newNode
		q.Tail = newNode
		return
	}

	q.Tail.Next = newNode
	q.Tail = newNode
}

// GetFrontElement returns the element at the front of the queue without removing it.
func (q *LinkedListQueue[T]) GetFrontElement() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return &q.Head.Value, nil
}

// Dequeue removes the element at the front of the queue.
func (q *LinkedListQueue[T]) Dequeue() error {
	if q.IsEmpty() {
		return errors.New("queue is empty")
	}
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
	return nil
}
