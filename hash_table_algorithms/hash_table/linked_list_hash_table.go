package hash_table

import "errors"

type BucketNode struct {
	Key   string
	Value any
	Next  *BucketNode
}

type LinkedListHashTable struct {
	Buckets  []*BucketNode
	HashFunc HashFunc
	Capacity int
}

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

	if ht.Buckets[index] == nil {
		ht.Buckets[index] = &BucketNode{Key: key, Value: value}
		return
	}

	curr := ht.Buckets[index]
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = &BucketNode{Key: key, Value: value}
}

// Get retrieves the value associated with the given key from the hash table
func (ht *LinkedListHashTable) Get(key string) (any, error) {
	index := ht.HashFunc(key) % ht.Capacity

	curr := ht.Buckets[index]
	for curr != nil {
		if curr.Key == key {
			return curr.Value, nil
		}
		curr = curr.Next
	}

	return nil, errors.New("key not found")
}

// Remove removes the key-value pair associated with the given key from the hash table
func (ht *LinkedListHashTable) Remove(key string) error {
	index := ht.HashFunc(key) % ht.Capacity

	if ht.Buckets[index] == nil {
		return errors.New("key not found")
	}

	if ht.Buckets[index].Key == key {
		ht.Buckets[index] = ht.Buckets[index].Next
		return nil
	}

	curr := ht.Buckets[index]
	for curr.Next != nil {
		if curr.Next.Key == key {
			curr.Next = curr.Next.Next
			return nil
		}
		curr = curr.Next
	}

	return errors.New("key not found")
}

// Free frees the memory used by the hash table
func (ht *LinkedListHashTable) Free() {
	for i := range ht.Buckets {
		ht.Buckets[i] = nil
	}
}
