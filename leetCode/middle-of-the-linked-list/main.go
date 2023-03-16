package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	size := 0
	ptr := head
	for ptr != nil {
		size++
		ptr = ptr.Next
	}
	ptr = head
	for i := 0; i < size/2; i++ {
		ptr = ptr.Next
	}
	return ptr
}
