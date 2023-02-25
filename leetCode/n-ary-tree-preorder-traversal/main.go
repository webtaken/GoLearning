package main

type Node struct {
	Val      int
	Children []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	var pre func(n *Node)
	pre = func(n *Node) {
		ans = append(ans, n.Val)
		for _, node := range n.Children {
			pre(node)
		}
	}
	pre(root)
	return ans
}
