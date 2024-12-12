/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {
	//双指针
	//俩个指针从两端向中间移动，相加判断是否满足条件,满足条件就交换，知道俩个指针相遇
	left := 0
	right := len(s) - 1
	t := []byte(s)
	for left < right {
		if isVowel(string(t[left])) && isVowel(string(t[right])) {
			t[left], t[right] = t[right], t[left]
			left++
			right--
		}
		if !isVowel(string(t[left])) {
			left++
		}
		if !isVowel(string(t[right])) {
			right--
		}
	}
	return string(t)
}

func isVowel(s string) bool {
	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" || s == "A" || s == "E" || s == "I" || s == "O" || s == "U" {
		return true
	}
	return false
}

// @lc code=end

