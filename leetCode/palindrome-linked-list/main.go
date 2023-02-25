package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var D = &ListNode{Val: 1, Next: nil}
	var C = &ListNode{Val: 2, Next: D}
	var B = &ListNode{Val: 2, Next: C}
	var A = &ListNode{Val: 1, Next: B}
	fmt.Println(isPalindrome(A))
	var E = &ListNode{Val: 3, Next: nil}
	D = &ListNode{Val: 2, Next: E}
	C = &ListNode{Val: 4, Next: D}
	B = &ListNode{Val: 2, Next: C}
	A = &ListNode{Val: 1, Next: B}
	fmt.Println(isPalindrome(A))
	var Z = &ListNode{Val: 1, Next: nil}
	fmt.Println(isPalindrome(Z))

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	size := 0
	tmp := head
	for tmp != nil {
		size++
		tmp = tmp.Next
	}
	if size == 1 {
		return true
	}
	reverse := func(head *ListNode) *ListNode {
		var prev *ListNode
		for head != nil {
			next := head.Next
			head.Next = prev
			prev = head
			head = next
		}
		return prev
	}

	tmp = head
	var l, r, prevL *ListNode
	t := size / 2
	if size%2 == 0 {
		t -= 1
	}
	for i := 0; i < t; i++ {
		prevL = tmp
		tmp = tmp.Next
	}
	l = tmp
	r = tmp.Next
	// fmt.Println(l.Val, r.Val)
	l.Next = nil
	if size%2 == 1 {
		prevL.Next = nil
	}
	tmp1 := reverse(head)
	tmp2 := r
	// fmt.Println(tmp1.Val, tmp2.Val)

	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val != tmp2.Val {
			return false
		}
		tmp1 = tmp1.Next
		tmp2 = tmp2.Next
	}
	return true
}
