package main

import "fmt"

type Node struct {
	Data  int
	Left   *Node
	Right  *Node
}

func(node *Node) GetNode(val int , left , right *Node) *Node {
	return &Node{Data: val,Left: left,Right: right}
}


func main() {
	node := Node{}

	c1 := node.GetNode(1 , nil , nil)

	c3 := node.GetNode(4 , nil , nil)
	c2 := node.GetNode(3 , nil , c3)

	root := node.GetNode(2 , c1 , c2)
	fmt.Println("root" , root.Data)

	//PrintTree(root)

	newRoot := InsertInTree(root , 0)
	fmt.Print("after insertion : ")
	PrintTree(newRoot)
}

func PrintTree(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.Data , "  ")
	PrintTree(root.Left)
	PrintTree(root.Right)
}

func InsertInTree(root *Node , val int) *Node {
	if root == nil {
		//PrintTree(root)
		return root
	}

	//if root.Left == nil && root.Right == nil {
	//
	//}

	if val < root.Data {
		if root.Left != nil {
			InsertInTree(root.Left , val)
		} else {
			node := Node{}
			newNode := node.GetNode(val , nil , nil)
			root.Left = newNode
		}
	} else if val >= root.Data {
		if root.Right != nil {
			InsertInTree(root.Right , val)
		} else {
			node := Node{}
			newNode := node.GetNode(val , nil , nil)
			root.Right = newNode
		}
	}

	return root
}