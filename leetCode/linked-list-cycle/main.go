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
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	ptr := head
	ptrsMap := make(map[*ListNode]bool, 0)
	for ptr != nil {
		if ok := ptrsMap[ptr]; ok {
			return true
		}
		ptrsMap[ptr] = true
		ptr = ptr.Next
	}
	return false
}
