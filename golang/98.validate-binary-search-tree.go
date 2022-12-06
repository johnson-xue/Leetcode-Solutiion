/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */
package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	arr := make([]int, 0)
	inOrder1(root, &arr)
	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}
	return true
}
func inOrder1(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder1(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder1(root.Right, arr)
}

// @lc code=end
