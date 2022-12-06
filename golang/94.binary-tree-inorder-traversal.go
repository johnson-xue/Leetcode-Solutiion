/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */
package main

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// @lc code=start
/**
 * Definition for a binary tree node.
 */

func inorderTraversal(root *TreeNode) []int {
	var res []int
	inOrder(root, &res)
	return res
}
func inOrder(root *TreeNode, output *[]int) {
	if root != nil {
		inOrder(root.Left, output)
		*output = append(*output, root.Val)
		inOrder(root.Right, output)
	}
}

// @lc code=end
