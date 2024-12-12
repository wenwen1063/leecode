/*
 * @lc app=leetcode.cn id=744 lang=golang
 *
 * [744] 寻找比目标字母大的最小字母
 */

// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
	// 二分查找
	l := 0
	h := len(letters) - 1
	for l <= h {
		mid := l + (h-l)/2
		if letters[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	if l == len(letters) {
		return letters[0]
	}

	return letters[l]
}

// @lc code=end

