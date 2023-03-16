package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var L = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}

	fmt.Println(oddEvenList(L))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	o := head
	e := head.Next
	oo := head
	ee := head.Next
	i := 1
	if e == nil {
		return o
	}
	for oo.Next != nil && ee.Next != nil {
		if (i % 2) == 1 {
			oo.Next = ee.Next
			oo = oo.Next
		} else {
			ee.Next = oo.Next
			ee = ee.Next
		}
		i++
	}
	if ee.Next != nil {
		ee.Next = oo.Next
	}
	oo.Next = e
	return o
}
