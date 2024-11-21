package stack_queue_tests

import (
	"algorithms/stack_queue_algorithms/stack_queue"
	"testing"
)

func TestStacks(t *testing.T) {
	t.Run("Test stack with dynamic array", func(t *testing.T) {
		stack := stack_queue.ArrayStack[int]{}
		runStackTests(t, &stack)
	})

	t.Run("Test stack with linked list", func(t *testing.T) {
		stack := stack_queue.NewLinkedListStack[int]()
		runStackTests(t, stack)
	})
}

func runStackTests(t *testing.T, stack stack_queue.Stack[int]) {
	// Test IsEmpty on an empty stack
	if !stack.IsEmpty() {
		t.Error("IsEmpty failed: expected stack to be empty")
	}

	// Test Push
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Test IsEmpty on a non-empty stack
	if stack.IsEmpty() {
		t.Error("IsEmpty failed: expected stack to not be empty")
	}

	// Test GetLatestElement
	latestElement, err := stack.GetLatestElement()
	if err != nil || *latestElement != 30 {
		t.Errorf("GetLatestElement failed: expected 30, got %v (error: %v)", latestElement, err)
	}

	// Test Pop
	err = stack.Pop()
	if err != nil {
		t.Errorf("Pop failed: unexpected error %v", err)
	}

	// Test GetLatestElement after Pop
	latestElement, err = stack.GetLatestElement()
	if err != nil || *latestElement != 20 {
		t.Errorf("GetLatestElement failed: expected 20, got %v (error: %v)", latestElement, err)
	}

	// Test Pop until empty
	err = stack.Pop()
	if err != nil {
		t.Errorf("Pop failed: unexpected error %v", err)
	}

	err = stack.Pop()
	if err != nil {
		t.Errorf("Pop failed: unexpected error %v", err)
	}

	// Test IsEmpty on an empty stack
	if !stack.IsEmpty() {
		t.Error("IsEmpty failed: expected stack to be empty")
	}

	// Test Pop on an empty stack
	err = stack.Pop()
	if err == nil {
		t.Error("Pop failed: expected error on empty stack")
	}
}
