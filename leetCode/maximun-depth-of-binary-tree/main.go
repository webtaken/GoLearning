package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepthCount := math.MinInt
	var traverse func(n *TreeNode, count int)
	traverse = func(n *TreeNode, count int) {
		if n == nil {
			return
		}
		count++
		if count > maxDepthCount {
			maxDepthCount = count
		}
		traverse(n.Left, count)
		traverse(n.Right, count)
	}
	traverse(root, 0)
	return maxDepthCount
}
