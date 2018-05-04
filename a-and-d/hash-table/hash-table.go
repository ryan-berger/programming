package hash_table

import "sync"

type HashTable struct {
	items map[int]interface{}
	lock  sync.RWMutex
}

func NewHashTable() *HashTable {
	return &HashTable{
		items: make(map[int]interface{}),
		lock: sync.RWMutex{},
	}
}

func hash(key string) int {
	h := 0
	for i := 0; i < len(key); i++ {
		h = h*31 + int(key[i])
	}
	return h
}

func (hashTable *HashTable) Put(key string, value interface{}) {
	hashTable.lock.Lock()
	defer hashTable.lock.Unlock()
	i := hash(key)
	hashTable.items[i] = value
}
