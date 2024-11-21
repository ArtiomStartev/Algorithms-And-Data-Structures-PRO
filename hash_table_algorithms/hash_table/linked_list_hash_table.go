package hash_table

import "errors"

// BucketNode represents a node in a linked list bucket
type BucketNode struct {
	Key   string
	Value any
	Next  *BucketNode
}

// LinkedListHashTable represents a hash table that uses linked lists to handle collisions
type LinkedListHashTable struct {
	Buckets  []*BucketNode
	HashFunc HashFunc
	Capacity int
}

// NewLinkedListHashTable creates a new hash table with the given capacity and hash function
func NewLinkedListHashTable(capacity int, hashFunc HashFunc) *LinkedListHashTable {
	return &LinkedListHashTable{
		Buckets:  make([]*BucketNode, capacity),
		HashFunc: hashFunc,
		Capacity: capacity,
	}
}

// Add adds a new key-value pair to the hash table
func (ht *LinkedListHashTable) Add(key string, value any) {
	index := ht.HashFunc(key) % ht.Capacity

	// If the bucket is empty, add the new node as the head
	if ht.Buckets[index] == nil {
		ht.Buckets[index] = &BucketNode{Key: key, Value: value}
		return
	}

	// Traverse the linked list to find the last node
	current := ht.Buckets[index]
	for current.Next != nil {
		current = current.Next
	}

	// Add the new node at the end of the linked list
	current.Next = &BucketNode{Key: key, Value: value}
}

// Get retrieves the value associated with the given key from the hash table
func (ht *LinkedListHashTable) Get(key string) (any, error) {
	index := ht.HashFunc(key) % ht.Capacity

	// Traverse the linked list to find the node with the given key
	current := ht.Buckets[index]
	for current != nil {
		if current.Key == key {
			return current.Value, nil
		}
		current = current.Next
	}

	return nil, errors.New("key not found")
}

// Remove removes the key-value pair associated with the given key from the hash table
func (ht *LinkedListHashTable) Remove(key string) error {
	index := ht.HashFunc(key) % ht.Capacity

	// If the bucket is empty, return an error
	if ht.Buckets[index] == nil {
		return errors.New("key not found")
	}

	// If the head node has the given key, remove it
	if ht.Buckets[index].Key == key {
		ht.Buckets[index] = ht.Buckets[index].Next
		return nil
	}

	// Traverse the linked list to find the node with the given key
	current := ht.Buckets[index]
	for current.Next != nil {
		if current.Next.Key == key {
			current.Next = current.Next.Next
			return nil
		}
		current = current.Next
	}

	return errors.New("key not found")
}

// Free frees the memory used by the hash table
func (ht *LinkedListHashTable) Free() {
	for i := range ht.Buckets {
		ht.Buckets[i] = nil
	}
}
