package main

import (
	"github.com/pkg/errors"
	"fmt"
)

var errInvalidNode = errors.New("Invalid Node !!")

type LNode interface {
	setVal(val int) error                   //Node
	getVal() int
}

type Node struct {
	Value int
	Next  *Node //SLLNode
}

func (node *Node) setVal(val int) error {
	//The type that implements LNode interface is *Node (not Node)
	if node == nil {
		return errInvalidNode
	}
	node.Value = val
	return nil
}

func (node *Node) getVal() int {
	return node.Value
}

func (*Node) nilRecieverTest() string {
	return "nil Receiver called for Node type"
}

func newNode() *Node {
	//similar to empty constructor
	return new(Node)
}

//////

type PowerNode struct {
	Value int
	Next  *PowerNode //PowerNode
}

func (node *PowerNode) setVal(val int) error {
	//The type that implements LNode interface is *PowerNode (not PowerNode)
	if node == nil {
		return errInvalidNode
	}
	node.Value = val * 10
	return nil
}

func (node *PowerNode) getVal() int {
	return node.Value
}

func newPowerNode() *PowerNode {
	//similar to empty constructor
	return new(PowerNode)
}

//Node and PowerNode implement the LNode interface by implementing its 2 methods

func main() {

	//var node *Node                                 //does not assign memory
	node := newNode()                                //assigns memory
	fmt.Println(node.setVal(2))

	fmt.Println(node.getVal())
	//fmt.Println(node.nilRecieverTest())

}
