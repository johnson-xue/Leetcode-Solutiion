/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */
package main

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// if list1 == nil {
	// 	return list2
	// }
	// if list2 == nil {
	// 	return list1
	// }
	// if list1.Val < list2.Val {
	// 	list1.Next = mergeTwoLists(list1.Next, list2)
	// 	return list1
	// }
	// list2.Next = mergeTwoLists(list1, list2.Next)
	// return list2
	head := &ListNode{Next: list1}
	cur := head
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}
	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}
	return head.Next
}

// @lc code=end
