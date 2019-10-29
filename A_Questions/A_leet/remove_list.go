package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	first1 := &ListNode{1,nil}
	sec1 := &ListNode{2,nil}
	third1 := &ListNode{3,nil}
	fourth1 := &ListNode{4,nil}
	fivth1 := &ListNode{5,nil}
	first1.Next = sec1
	sec1.Next = third1
	third1.Next = fourth1
	fourth1.Next = fivth1


	PrintLL(first1)

	removeNthFromEnd(first1 , 2 , 0)

}

func removeNthFromEnd(head *ListNode, n int , len int) *ListNode {

	if head == nil {
		return
	} else {
		removeNthFromEnd(head.Next , n , len++)
	}

}

func PrintLL(root *ListNode) {
	if root == nil {
		fmt.Println()
		return
	}

	fmt.Print(root.Val, "  ")
	PrintLL(root.Next)
}