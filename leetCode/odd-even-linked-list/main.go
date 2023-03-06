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
	if head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		return head
	}

	evenStart := head.Next
	even_i := head.Next
	odd_i := head

	for ; odd_i != nil; odd_i = odd_i.Next {
		odd_i.Next = odd_i.Next.Next
	}
	for ; even_i != nil; even_i = even_i.Next {
		even_i.Next = even_i.Next.Next
	}
	odd_i.Next = evenStart
	return head
}
