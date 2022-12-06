/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	l, j, curDir := 0, 0, false
	for len(queue) > 0 {
		l = len(queue)
		j = l - 1
		tmp := make([]int, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if !curDir {
				tmp[i] = queue[i].Val
			} else {
				tmp[j] = queue[i].Val
				j--
			}
		}
		queue = queue[l:]
		res = append(res, tmp)
		curDir = !curDir
	}
	return res
}

func zigzagLevelOrder2(root *TreeNode) [][]int {
	var res [][]int
	search1(root, 0, &res)
	return res
}
func search1(node *TreeNode, depth int, res *[][]int) {
	if node == nil {
		return
	}
	for len(*res) < depth+1 {
		*res = append(*res, []int{})
	}
	if depth%2 == 0 {
		(*res)[depth] = append((*res)[depth], node.Val)
	} else {
		(*res)[depth] = append([]int{node.Val}, (*res)[depth]...)
	}
	search1(node.Left, depth+1, res)
	search1(node.Right, depth+1, res)
}

func zigzagLevelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := make([][]int, 0)
	curDir := false
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i].Val)
		}
		queue = queue[l:]
		if curDir {
			for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
				tmp[i], tmp[j] = tmp[j], tmp[i]
			}
			curDir = false
		} else {
			curDir = true
		}
		res = append(res, tmp)
	}
	return res
}

// @lc code=end
