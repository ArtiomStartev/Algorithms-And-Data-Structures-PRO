package hash_table_tests

import (
	"algorithms/hash_table_algorithms/hash_table"
	"testing"
)

func TestOpenAddressingHashTable(t *testing.T) {
	hashTable := hash_table.NewOpenAddressingHashTable(10, hash_table.HashString)

	// Test Add
	err := hashTable.Add("key1", "value1")
	if err != nil {
		t.Errorf("Expected no error when adding key 'key1', got error: %v", err)
	}

	err = hashTable.Add("key2", "value2")
	if err != nil {
		t.Errorf("Expected no error when adding key 'key2', got error: %v", err)
	}

	// Test Get
	value, err := hashTable.Get("key1")
	if err != nil {
		t.Errorf("Expected value for key 'key1', got error: %v", err)
	}
	if value != "value1" {
		t.Errorf("Expected value for key 'key1' to be 'value1', got %v", value)
	}

	value, err = hashTable.Get("key2")
	if err != nil {
		t.Errorf("Expected value for key 'key2', got error: %v", err)
	}
	if value != "value2" {
		t.Errorf("Expected value for key 'key2' to be 'value2', got %v", value)
	}

	// Test Get with a non-existing key
	_, err = hashTable.Get("non-existing-key")
	if err == nil {
		t.Errorf("Expected error for non-existing key 'non-existing-key', got nil")
	}

	// Test Remove
	err = hashTable.Remove("key1")
	if err != nil {
		t.Errorf("Expected no error when removing key 'key1', got error: %v", err)
	}

	_, err = hashTable.Get("key1")
	if err == nil {
		t.Errorf("Expected error when getting key 'key1' after removal, got nil")
	}

	// Test Remove with a non-existing key
	err = hashTable.Remove("non-existing-key")
	if err == nil {
		t.Errorf("Expected error for non-existing key 'non-existing-key', got nil")
	}

	// Test Free
	hashTable.Free()

	// Test Get after freeing
	_, err = hashTable.Get("key2")
	if err == nil {
		t.Errorf("Expected error when getting key 'key2' after freeing, got nil")
	}
}
