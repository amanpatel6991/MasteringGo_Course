package main

import (
	"fmt"
)

type LNode interface {
	SetVal(val int)                    //Node
	GetVal() int
}

type Node struct {
	Value int
	Next *Node                                //SLLNode
}

func(node *Node) SetVal(val int) {               //The type that implements LNode interface is *Node (not Node)
	node.Value = val
}

func(node *Node) GetVal() int{
	return node.Value
}

func NewNode() *Node{                                   //similar to empty constructor
	return new(Node)     //allocates memory
}

//////

type PowerNode struct {
	Value int
	Next *PowerNode                       //PowerNode
}

func(node *PowerNode) SetVal(val int) {            //The type that implements LNode interface is *PowerNode (not PowerNode)
	node.Value = val * 10
}

func(node *PowerNode) GetVal() int{
	return node.Value
}

func NewPowerNode() *PowerNode{                      //similar to empty constructor
	return new(PowerNode)
}

//Node and PowerNode implement the LNode interface by implementing its 2 methods

func main() {

	//3.2 video

	//var node LNode
	//
	//node = NewNode()
	//node.SetVal(1)
	//node.SetVal(1)
	//fmt.Println(node.GetVal())
	//
	//node = NewPowerNode()
	//node.SetVal(1)
	//node.SetVal(1)
	//fmt.Println(node.GetVal())
	//
	////Verify if the implementor type PowerNode implements the interface LNode (ok returns true if possitive)
	//if n , ok := node.(*PowerNode); ok {
	//	fmt.Println("Power Node of val :" , n.GetVal())
	//}

	//5.1 video

	n:=createNode(2)

	fmt.Println(n.GetVal())

	switch nodeType := n.(type) {                   //Type Switch (checking which of the 2 implementors (Node or PowerNode) of interface (LNode) n is now)
	case *Node:
		fmt.Println("type is Node" , nodeType)
	case *PowerNode:
		fmt.Println("type is PowerNode" , nodeType)
	}

}


func createNode(v int) LNode {
	//newnode := Node{}                 //cannot use newnode (type Node) as type LNode in return argument: Node does not implement LNode
	//newnode := NewPowerNode()
	newnode := NewNode()
	newnode.SetVal(2)

	return newnode
	//return &Node{2,nil}
}