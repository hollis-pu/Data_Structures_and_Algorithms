package main

import (
	"fmt"
	"hash/fnv"
)

/**
* Description:
* @Author Hollis
* @Create 2023-10-15 12:10
 */

const tableSize = 100

type HashTable struct {
	data [tableSize][]KeyValue
}

type KeyValue struct {
	key   string
	value int
}

func main() {
	ht := NewHashTable()

	ht.Insert("apple", 1)
	ht.Insert("banana", 2)
	ht.Insert("cherry", 3)

	value, exists := ht.Search("banana")
	if exists {
		fmt.Println("banana:", value)
	} else {
		fmt.Println("Key not found")
	}

	ht.Delete("banana")
	_, exists = ht.Search("banana")
	if exists {
		fmt.Println("banana:", value)
	} else {
		fmt.Println("Key not found")
	}
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

func (ht *HashTable) Insert(key string, value int) {
	index := hash(key)
	kv := KeyValue{key, value}
	ht.data[index] = append(ht.data[index], kv)
}

func (ht *HashTable) Search(key string) (int, bool) {
	index := hash(key)
	for _, kv := range ht.data[index] {
		if kv.key == key {
			return kv.value, true
		}
	}
	return 0, false
}

func (ht *HashTable) Delete(key string) {
	index := hash(key)
	for i, kv := range ht.data[index] {
		if kv.key == key {
			ht.data[index] = append(ht.data[index][:i], ht.data[index][i+1:]...)
			return
		}
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32() % tableSize
}
