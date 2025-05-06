package stack_queue

import "errors"

type Queue[T any] interface {
	IsEmpty() bool
	Enqueue(value T)
	GetFrontElement() (*T, error)
	Dequeue() error
}

// ----------------------- Dynamic Array Queue -----------------------

type ArrayQueue[T any] struct {
	data []T
}

func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{}
}

// IsEmpty checks if the queue is empty
func (q *ArrayQueue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Enqueue adds an element to the end of the queue
func (q *ArrayQueue[T]) Enqueue(value T) {
	q.data = append(q.data, value)
}

// GetFrontElement retrieves the front element of the queue without removing it
func (q *ArrayQueue[T]) GetFrontElement() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return &q.data[0], nil
}

// Dequeue removes the front element from the queue
func (q *ArrayQueue[T]) Dequeue() error {
	if q.IsEmpty() {
		return errors.New("queue is empty")
	}

	q.data[0] = *new(T) // Clear reference to allow GC to free memory
	q.data = q.data[1:]

	return nil
}

// ----------------------- Linked List Queue -----------------------

type QueueNode[T any] struct {
	Value T
	Next  *QueueNode[T]
}

type LinkedListQueue[T any] struct {
	Head *QueueNode[T]
	Tail *QueueNode[T]
}

func NewLinkedListQueue[T any]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{}
}

// IsEmpty checks if the queue is empty
func (q *LinkedListQueue[T]) IsEmpty() bool {
	return q.Head == nil
}

// Enqueue adds an element to the end of the queue
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

// GetFrontElement retrieves the front element of the queue without removing it
func (q *LinkedListQueue[T]) GetFrontElement() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return &q.Head.Value, nil
}

// Dequeue removes the front element from the queue
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
