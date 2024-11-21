package hash_table

import (
	"errors"
)

type HashFunc func(string) int

// OpenAddressingHashBucket represents a bucket in an open addressing hash table
type OpenAddressingHashBucket struct {
	Occupied bool
	Key      string
	Value    any
}

// OpenAddressingHashTable represents an open addressing hash table
type OpenAddressingHashTable struct {
	Buckets  []OpenAddressingHashBucket
	HashFunc HashFunc
	Capacity int
}

// NewOpenAddressingHashTable creates a new open addressing hash table with the given capacity and hash function
func NewOpenAddressingHashTable(capacity int, hashFunc HashFunc) *OpenAddressingHashTable {
	return &OpenAddressingHashTable{
		Buckets:  make([]OpenAddressingHashBucket, capacity),
		HashFunc: hashFunc,
		Capacity: capacity,
	}
}

// Add adds a new key-value pair to the hash table
func (ht *OpenAddressingHashTable) Add(key string, value any) {
	index := ht.HashFunc(key) % ht.Capacity

	for ht.Buckets[index].Occupied {
		index = (index + 1) % ht.Capacity
	}

	ht.Buckets[index] = OpenAddressingHashBucket{
		Occupied: true,
		Key:      key,
		Value:    value,
	}
}

// Get retrieves the value associated with the given key from the hash table
func (ht *OpenAddressingHashTable) Get(key string) (any, error) {
	index := ht.HashFunc(key) % ht.Capacity

	for ht.Buckets[index].Occupied {
		if ht.Buckets[index].Key == key {
			return ht.Buckets[index].Value, nil
		}

		index = (index + 1) % ht.Capacity
	}

	return nil, errors.New("key not found")
}

// Remove removes the key-value pair associated with the given key from the hash table
func (ht *OpenAddressingHashTable) Remove(key string) error {
	index := ht.HashFunc(key) % ht.Capacity

	for ht.Buckets[index].Occupied {
		if ht.Buckets[index].Key == key {
			ht.Buckets[index] = OpenAddressingHashBucket{}
			return nil
		}

		index = (index + 1) % ht.Capacity

		// if we have traversed the whole hash table and still not found the key
		if index == ht.HashFunc(key)%ht.Capacity {
			break
		}
	}

	return errors.New("key not found")
}

// Free frees the memory used by the hash table
func (ht *OpenAddressingHashTable) Free() {
	for i := range ht.Buckets {
		ht.Buckets[i] = OpenAddressingHashBucket{}
	}
}
