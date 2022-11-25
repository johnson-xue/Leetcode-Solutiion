/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */
package main

// @lc code=start
func minWindow(s string, t string) string {
	var sFreq, tFreq, ret = map[byte]int{}, map[byte]int{}, ""
	for _, v := range t {
		tFreq[byte(v)]++
	}
	for left, right := 0, 0; right < len(s); right++ {
		sFreq[s[right]]++
		if right-left+1 < len(t) {
			continue
		}
		for check(sFreq, tFreq) {
			if right-left+1 < len(ret) || ret == "" {
				ret = s[left : right+1]
			}
			sFreq[s[left]]--
			left++
		}
	}
	return ret
}
func check(sFreq, tFreq map[byte]int) bool {
	for k, v := range tFreq {
		if sFreq[k] < v {
			return false
		}
	}
	return true
}

// @lc code=end
