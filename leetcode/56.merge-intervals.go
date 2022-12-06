/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */
package main

// @lc code=start

func merge1(intervals [][]int) [][]int {

	if len(intervals) == 0 {
		return intervals
	}
	quickSort(intervals, 0, len(intervals)-1)
	currIndex := 0
	res := [][]int{}
	res = append(res, intervals[currIndex])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[currIndex][1] {
			currIndex++
			res = append(res, intervals[i])
		} else {
			res[currIndex][1] = max2(res[currIndex][1], intervals[i][1])
		}
	}
	return res
}
func max2(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min2(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
func partitionSort(a [][]int, low, high int) int {
	pivot := high
	i := low - 1
	for j := low; j < high; j++ {
		if (a[j][0] < a[pivot][0]) || (a[j][0] == a[pivot][0]) && (a[j][1] <= a[pivot][0]) {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}
func quickSort(a [][]int, low, high int) {
	if low >= high {
		return
	}
	pivot := partitionSort(a, low, high)
	quickSort(a, low, pivot-1)
	quickSort(a, pivot+1, high)
}

// @lc code=end
