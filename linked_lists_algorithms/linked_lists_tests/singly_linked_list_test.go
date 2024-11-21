package linked_lists_tests

import (
	"algorithms/linked_lists_algorithms/linked_lists"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	list := linked_lists.SinglyLinkedList[int]{}

	// Test Insert
	node1 := list.Insert(10)
	if list.Head != node1 || list.Tail != node1 || list.Length != 1 {
		t.Errorf("Insert failed: expected Head and Tail to point to the same node with Length 1")
	}

	node2 := list.Insert(20)
	if list.Tail != node2 || list.Length != 2 {
		t.Errorf("Insert failed: expected Tail to update and Length to be 2")
	}

	// Test InsertAfter
	node3 := list.InsertAfter(node1, 15)
	if node1.Next != node3 || node3.Next != node2 || list.Length != 3 {
		t.Errorf("InsertAfter failed: node linking is incorrect")
	}

	node4 := list.InsertAfter(nil, 5)
	if list.Head != node4 || node4.Next != node1 || list.Length != 4 {
		t.Errorf("InsertAfter failed: expected Head to update and Length to be 4")
	}

	// Test AssertNoCycles
	err := list.AssertNoCycles()
	if err != nil {
		t.Errorf("AssertNoCycles failed: expected no cycles in linked list")
	}

	// Test Find
	findResult, err := list.Find(15)
	if err != nil || findResult.Node != node3 || findResult.Prev != node1 {
		t.Errorf("Find failed: expected to find node with value 15")
	}

	findResult, err = list.Find(101)
	if err == nil {
		t.Errorf("Find failed: expected error for non-existent value")
	}

	// Test Remove
	err = list.Remove(5)
	if err != nil || list.Head != node1 || list.Length != 3 {
		t.Errorf("Remove failed: expected Head to update and Length to be 3")
	}

	err = list.Remove(20)
	if err != nil || list.Tail != node3 || list.Length != 2 {
		t.Errorf("Remove failed: expected Tail to update and Length to be 2")
	}

	err = list.Remove(101)
	if err == nil {
		t.Errorf("Remove failed: expected error for non-existent value")
	}

	// Test RemoveAfter
	err = list.RemoveAfter(nil)
	if err != nil || list.Head != node3 || list.Length != 1 {
		t.Errorf("RemoveAfter failed: expected Head to update and Length to be 1")
	}

	// Test AssertNoCycles
	err = list.AssertNoCycles()
	if err != nil {
		t.Errorf("AssertNoCycles failed: expected no cycles in linked list")
	}
}
