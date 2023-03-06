package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root1
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var merge func(root1 *TreeNode, root2 *TreeNode)
	merge = func(root1 *TreeNode, root2 *TreeNode) {
		if root1 != nil && root2 != nil {
			root1.Val += root2.Val
		} else if root1 == nil && root2 != nil {
			root1 = root2
		} else if root1 != nil && root2 == nil {
			return
		} else if root1 == nil && root2 == nil {
			return
		}
		merge(root1.Left, root2.Left)
		merge(root1.Right, root2.Right)
	}
	merge(root1, root2)
	return root1
}
