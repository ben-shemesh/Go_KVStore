package main

import "fmt"


// interface has no true implementation 
	// it only has a function sig
type Storer interface {
	Put(string, []byte) error
	Get(string) ([]byte, error)
	Update(string, []byte ) error
	Delete(string, []byte, error)
}
type KVStore struct {
	
}

func StoreThings(s Storer) error {
	return  s.Put("fooo", []byte("barrr"))
}

func main(){


}