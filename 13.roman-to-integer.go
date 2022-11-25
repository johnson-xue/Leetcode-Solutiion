/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
package main

// @lc code=start
func romanToInt(s string) int {
	var roman = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	if s == "" {
		return 0
	}
	num, lastInt, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman[char]
		if num < lastInt {
			total = total - num
		} else {
			total = total + num
		}
		lastInt = num
	}
	return total
}

// @lc code=end
