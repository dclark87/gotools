package main

import "fmt"

// Problem courtesy of:
// https://leetcode.com/problems/add-two-numbers
type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(l *ListNode) {
	for {
		if l == nil {
			fmt.Printf("nil\n")
			break
		}
		fmt.Printf("[%d]-->", l.Val)
		l = l.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var v1, v2, v3, carry int
	ret := new(ListNode)
	l3 := ret
	for {
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}

		v3 = v1 + v2 + carry
		l3.Val = v3 % 10

		if v3 >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			if carry == 1 {
				l3.Next = &ListNode{Val: 1}
			}
			return ret
		}
		l3.Next = new(ListNode)
		l3 = l3.Next
	}
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l3 := addTwoNumbers(l1, l2)
	fmt.Println("inputs:")
	printListNode(l1)
	printListNode(l2)

	fmt.Println("output:")
	printListNode(l3)
}
