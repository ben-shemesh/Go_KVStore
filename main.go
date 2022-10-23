package main

import (
	"fmt"
	"sync"
)

// interface has no true implementation
// it only has a function sig
type Storer[k comparable, V any] interface {
	Put(k, V) error
	Get(k) ([]byte, error)
	Update(k, V ) error
	Delete(k)  (V, error)
}
type KVStore [k comparable, V any]struct {
	// protects the variable of mu
	mu sync.RWMutex
	data map[string][]byte
}

func NewKVStore() *KVStore[string, int] {
	return &KVStore[string, int]{
		data: make(map[string][]byte),
	}
}

func StoreThings(s Storer[string, int])error {
	return  s.Put("fooo", 1)
}

func main(){


}