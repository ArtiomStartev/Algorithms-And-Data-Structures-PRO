package stack_queue_tests

import (
	"algorithms/stack_queue_algorithms/stack_queue"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	queue := stack_queue.ArrayQueue[int]{}

	// Test IsEmpty on an empty queue
	if !queue.IsEmpty() {
		t.Error("IsEmpty failed: expected queue to be empty")
	}

	// Test Enqueue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	// Test IsEmpty on a non-empty queue
	if queue.IsEmpty() {
		t.Error("IsEmpty failed: expected queue to not be empty")
	}

	// Test GetFrontElement
	frontElement, err := queue.GetFrontElement()
	if err != nil || *frontElement != 10 {
		t.Errorf("GetFrontElement failed: expected 10, got %v (error: %v)", frontElement, err)
	}

	// Test Dequeue
	err = queue.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: unexpected error %v", err)
	}

	// Test GetFrontElement after Dequeue
	frontElement, err = queue.GetFrontElement()
	if err != nil || *frontElement != 20 {
		t.Errorf("GetFrontElement failed: expected 20, got %v (error: %v)", frontElement, err)
	}

	// Test Dequeue until empty
	err = queue.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: unexpected error %v", err)
	}

	err = queue.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: unexpected error %v", err)
	}

	// Test IsEmpty on an empty queue
	if !queue.IsEmpty() {
		t.Error("IsEmpty failed: expected queue to be empty")
	}

	// Test Dequeue on an empty queue
	err = queue.Dequeue()
	if err == nil {
		t.Error("Dequeue failed: expected error on empty queue")
	}
}

func TestLinkedListQueue(t *testing.T) {
	queue := stack_queue.LinkedListQueue[int]{}

	// Test IsEmpty on an empty queue
	if !queue.IsEmpty() {
		t.Error("IsEmpty failed: expected queue to be empty")
	}

	// Test Enqueue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	// Test IsEmpty on a non-empty queue
	if queue.IsEmpty() {
		t.Error("IsEmpty failed: expected queue to not be empty")
	}

	// Test GetFrontElement
	frontElement, err := queue.GetFrontElement()
	if err != nil || *frontElement != 10 {
		t.Errorf("GetFrontElement failed: expected 10, got %v (error: %v)", frontElement, err)
	}

	// Test Dequeue
	err = queue.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: unexpected error %v", err)
	}

	// Test GetFrontElement after Dequeue
	frontElement, err = queue.GetFrontElement()
	if err != nil || *frontElement != 20 {
		t.Errorf("GetFrontElement failed: expected 20, got %v (error: %v)", frontElement, err)
	}

	// Test Dequeue until empty
	err = queue.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: unexpected error %v", err)
	}

	err = queue.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: unexpected error %v", err)
	}

	// Test IsEmpty on an empty queue
	if !queue.IsEmpty() {
		t.Error("IsEmpty failed: expected queue to be empty")
	}

	// Test Dequeue on an empty queue
	err = queue.Dequeue()
	if err == nil {
		t.Error("Dequeue failed: expected error on empty queue")
	}
}
