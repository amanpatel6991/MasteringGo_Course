package main

import "fmt"

type MagicStore struct {
	Value interface{}
	Name string
}

func(store *MagicStore) SetVal(val interface{}) {
	store.Value = val
}

func(store *MagicStore) GetVal() interface{}{
	return store.Value
}

func NewMagicStore(name string) *MagicStore{                                   //similar to parameterized constructor
	return &MagicStore{Name: name}
}

func main() {

	store1 := NewMagicStore("IntStore")
	store1.SetVal(2)

	if _ , ok := store1.GetVal().(int); ok {
		fmt.Println("Int store setup : " , *store1)
	}

	store2 := NewMagicStore("StringStore")
	store2.SetVal("str")

	if _ , ok := store2.GetVal().(string); ok {
		fmt.Println("String store setup : " , *store2)
	}

}
