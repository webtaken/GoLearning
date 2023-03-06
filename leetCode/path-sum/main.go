package main

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

func hasPathSum(root *TreeNode, targetSum int) bool {
	var checkSum func(n *TreeNode, sum int) bool
	checkSum = func(n *TreeNode, sum int) bool {
		if n == nil {
			return false
		}
		sum += n.Val
		if sum == targetSum && n.Left == nil && n.Right == nil {
			return true
		}
		return checkSum(n.Left, sum) || checkSum(n.Right, sum)
	}
	return checkSum(root, 0)
}
