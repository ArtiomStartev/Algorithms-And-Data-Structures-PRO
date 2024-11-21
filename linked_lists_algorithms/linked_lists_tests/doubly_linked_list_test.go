package linked_lists_tests

import (
	"algorithms/linked_lists_algorithms/linked_lists"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	list := linked_lists.DoublyLinkedList[int]{}

	// Test Insert
	node1 := list.Insert(10)
	if list.Head != node1 || list.Tail != node1 || list.Length != 1 {
		t.Errorf("Insert failed: expected Head and Tail to point to the same node with Length 1")
	}

	node2 := list.Insert(20)
	if list.Tail != node2 || node1.Next != node2 || node2.Prev != node1 || list.Length != 2 {
		t.Errorf("Insert failed: expected Tail to update and node linking to be correct")
	}

	// Test InsertBefore
	node3 := list.InsertBefore(node2, 15)
	if node1.Next != node3 || node3.Prev != node1 || node3.Next != node2 || node2.Prev != node3 || list.Length != 3 {
		t.Errorf("InsertBefore failed: expected node linking to be correct")
	}

	node4 := list.InsertBefore(nil, 5)
	if list.Head != node4 || node4.Next != node1 || node1.Prev != node4 || list.Length != 4 {
		t.Errorf("InsertBefore failed: expected Head to update and node linking to be correct")
	}

	// Test InsertAfter
	node5 := list.InsertAfter(node1, 12)
	if node1.Next != node5 || node5.Prev != node1 || node5.Next != node3 || node3.Prev != node5 || list.Length != 5 {
		t.Errorf("InsertAfter failed: expected node linking to be correct")
	}

	node6 := list.InsertAfter(nil, 1)
	if list.Head != node6 || node6.Next != node4 || node4.Prev != node6 || list.Length != 6 {
		t.Errorf("InsertAfter failed: expected Head to update and node linking to be correct")
	}

	// Test AssertNoCycles
	err := list.AssertNoCycles()
	if err != nil {
		t.Errorf("AssertNoCycles failed: expected no cycles in linked list")
	}

	// Test Find
	findResult, err := list.Find(15)
	if err != nil || findResult != *node3 || findResult.Prev != node5 || findResult.Next != node2 {
		t.Errorf("Find failed: expected to find node with value 15")
	}

	findResult, err = list.Find(101)
	if err == nil {
		t.Errorf("Find failed: expected error for non-existent value")
	}

	// Test Remove
	err = list.Remove(node6)
	if err != nil || list.Head != node4 || node4.Next != node1 || node1.Prev != node4 || list.Length != 5 {
		t.Errorf("Remove failed: expected Head to update and node linking to be correct")
	}

	err = list.Remove(node2)
	if err != nil || list.Tail != node3 || node5.Next != node3 || node3.Prev != node5 || list.Length != 4 {
		t.Errorf("Remove failed: expected Tail to update and node linking to be correct")
	}

	err = list.Remove(nil)
	if err == nil {
		t.Errorf("Remove failed: expected error for nil node")
	}

	// Test AssertNoCycles
	err = list.AssertNoCycles()
	if err != nil {
		t.Errorf("AssertNoCycles failed: expected no cycles in linked list")
	}
}
