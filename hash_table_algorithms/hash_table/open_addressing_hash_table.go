package hash_table

import (
	"errors"
)

type OpenAddressingHashTable struct {
	Buckets  []OpenAddressingHashBucket
	HashFunc HashFunc
	Capacity int
}

type OpenAddressingHashBucket struct {
	Occupied bool
	Key      string
	Value    any
}

func NewOpenAddressingHashTable(capacity int, hashFunc HashFunc) *OpenAddressingHashTable {
	return &OpenAddressingHashTable{
		Buckets:  make([]OpenAddressingHashBucket, capacity),
		HashFunc: hashFunc,
		Capacity: capacity,
	}
}

// Add adds a new key-value pair to the hash table
func (ht *OpenAddressingHashTable) Add(key string, val any) error {
	index := ht.HashFunc(key) % ht.Capacity
	origIndex := index

	for {
		if !ht.Buckets[index].Occupied {
			ht.Buckets[index] = OpenAddressingHashBucket{
				Occupied: true,
				Key:      key,
				Value:    val,
			}
			return nil
		}

		index = (index + 1) % ht.Capacity

		// If we've looped back to the original index, the hash table is full
		if index == origIndex {
			break
		}
	}

	return errors.New("hash table is full")
}

// Get retrieves the value associated with the given key from the hash table
func (ht *OpenAddressingHashTable) Get(key string) (any, error) {
	index := ht.HashFunc(key) % ht.Capacity
	origIndex := index

	for ht.Buckets[index].Occupied {
		if ht.Buckets[index].Key == key {
			return ht.Buckets[index].Value, nil
		}

		index = (index + 1) % ht.Capacity

		// If we've looped back to the original index, the key is not in the hash table
		if index == origIndex {
			break
		}
	}

	return nil, errors.New("key not found")
}

// Remove removes the key-value pair associated with the given key from the hash table
func (ht *OpenAddressingHashTable) Remove(key string) error {
	index := ht.HashFunc(key) % ht.Capacity
	origIndex := index

	for {
		if ht.Buckets[index].Occupied && ht.Buckets[index].Key == key {
			ht.Buckets[index] = OpenAddressingHashBucket{}
			return nil
		}

		index = (index + 1) % ht.Capacity

		// If we've looped back to the original index, the key is not in the hash table
		if index == origIndex {
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
