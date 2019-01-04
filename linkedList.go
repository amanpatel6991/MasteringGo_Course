package main

import (
	"fmt"
)

type Node struct {
	Value  int
	Next *Node
}

type LLIst struct {
	Head *Node
	Tail *Node
}
//
//func(node *Node) SetValue(val int) {
//	node.Value = val
//}
//
//func(node *Node) GetValue() int{
//	return node.Value
//}
//
//func GetNewNode() *Node{
//	return new(Node)
//}

func main() {

	node := LLIst{}
	node.Insert(5)
	node.Insert(9)
	node.Insert(13)
	node.Insert(22)
	node.Insert(28)

	node.Display()

}

func(list *LLIst) Insert(key int) {

	node := &Node{key , list.Head}
	list.Head = node

	l := list.Head
	for l.Next != nil {
		l = l.Next
	}

	list.Tail = l

}

func(list *LLIst) Display() {
	list1 := list.Head
	for (list1 !=nil) {
		fmt.Println(list1.Value)
		list1 = list1.Next
	}

}