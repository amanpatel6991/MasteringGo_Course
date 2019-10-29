package main

import "fmt"


// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	first1 := &ListNode{5,nil}
	sec1 := &ListNode{4,nil}
	third1 := &ListNode{3,nil}
	first1.Next = sec1
	sec1.Next = third1

	PrintLL(first1)

	first2 := &ListNode{5,nil}
	sec2 := &ListNode{6,nil}
	third2:= &ListNode{4,nil}
	first2.Next = sec2
	sec2.Next = third2

	PrintLL(first2)


	res := addTwoNumbers(first1,first2)

	PrintLL(res)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	res := &ListNode{}
	res = nil
	for l1 != nil || l2 != nil || carry != 0 {
		tempVal := 0
		if l1 != nil && l2 != nil {
			tempVal = l1.Val + l2.Val + carry
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 == nil && l2 == nil {
			tempVal = carry
		} else if l1 == nil {
			tempVal = l2.Val + carry
			l2 = l2.Next
		} else if l2 == nil {
			tempVal = l1.Val + carry
			l1 = l1.Next
		}
		//fmt.Println("temp 11 : " , tempVal)
		carry = 0

		if tempVal >= 10 {
			carry = tempVal / 10
			tempVal = tempVal - 10
		}
		//fmt.Println("temp : " , carry , " " , tempVal)
		PrintLL(res)

		tempNode := &ListNode{Val: tempVal, Next: nil}
		fmt.Println(tempNode)

		if res == nil {
			fmt.Println("first itr")
			res = tempNode
		} else {
			PrintLL(res)
			temp1 := res

			for temp1.Next != nil {
				fmt.Println("val : : " , temp1.Val)
				temp1 = temp1.Next
			}
			PrintLL(res)

			temp1.Next = tempNode

		}
		//PrintLL(res)
	}

	return res

}

func PrintLL(root *ListNode) {
	if root == nil {
		fmt.Println()
		return
	}

	fmt.Print(root.Val , "  ")
	PrintLL(root.Next)
}