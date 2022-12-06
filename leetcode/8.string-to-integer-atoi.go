/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
package main

// @lc code=start
func myAtoi(s string) int {
	signAllowed, whitespaceAllowed, sign, digits := true, true, 1, []int{}
	for _, c := range s {
		if c == ' ' && whitespaceAllowed {
			continue
		}
		if signAllowed {
			if c == '+' {
				signAllowed = false
				whitespaceAllowed = false
				continue
			} else if c == '-' {
				sign = -1
				signAllowed = false
				whitespaceAllowed = false
				continue
			}
		}
		if c < '0' || c > '9' {
			break
		}
		signAllowed, whitespaceAllowed = false, false
		digits = append(digits, int(c-48))
	}
	maxInt := int64(2 << 30)
	if sign > 0 {
		maxInt = maxInt - 1
	}
	lastLeading0Index := -1
	for i, d := range digits {
		if d == 0 {
			lastLeading0Index = i
		} else {
			break
		}
	}
	if lastLeading0Index > -1 {
		digits = digits[lastLeading0Index+1:]
	}
	digitsCount, num, place := len(digits), int64(0), int64(1)
	for i := digitsCount - 1; i >= 0; i-- {
		num += int64(digits[i]) * place
		place *= 10
		if digitsCount-i > 10 || num > maxInt {
			return int(int64(sign) * maxInt)
		}
	}
	num *= int64(sign)
	return int(num)
}

// @lc code=end
