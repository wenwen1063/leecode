/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */

// @lc code=start
func findLongestWord(s string, dictionary []string) (ans string) {
	for _, t := range dictionary {
		i := 0
		for j := range s {
			if s[j] == t[i] { //双指针慢指针向前
				i++
			}
			if i == len(t) { //匹配成功，更新结果
				if len(t) > len(ans) || len(t) == len(ans) && t < ans { //判断存在的字符串是否是最长的，如果是，更新
					ans = t
				}
				break
			}
		}
	}
	return
}

// @lc code=end

