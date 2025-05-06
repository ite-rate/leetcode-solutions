/*
 * LeetCode #3: 题目3
 * 难度: 未知
 *
 * 题目描述:
 * 由大模型直接生成
 *
 * 代码骨架完整度: 10%
 */

package leetcode

import (
	"testing"
	// "golang.org/x/net/html/charset"
)

/*
题目：#3 无重复字符的最长子串
难度：中等
描述：给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
*/

/*
解法一：滑动窗口（双指针）
时间复杂度：O(n)，n为字符串长度
空间复杂度：O(min(m, n))，m为字符集大小
关键提示：
1. 使用左右指针维护一个滑动窗口
2. 使用哈希集合记录窗口中的字符
3. 右指针不断右移，遇到重复字符时移动左指针
*/
func lengthOfLongestSubstring(s string) int {
	// TODO: 实现滑动窗口解法
	n := len(s)
	if n == 0 {
		return 0
	}
	left, right := 0, 0
	charset := make(map[byte]struct{})
	maxLength := 0
	for right < n {
		if _, exists := charset[s[right]]; !exists {
			charset[s[right]] = struct{}{}
			right++
			// 更新最大长度
			maxLength = max(maxLength, right-left)
		} else {
			// 遇到重复字符，移动左指针
			delete(charset, s[left])
			left++
		}
	}
	return maxLength
}

/*
解法二：优化的滑动窗口（直接跳转）
时间复杂度：O(n)
空间复杂度：O(min(m, n))
关键提示：
1. 使用哈希表记录字符和其索引
2. 遇到重复字符时，直接跳转到重复字符的下一个位置
*/
func lengthOfLongestSubstring2(s string) int {
	// TODO: 实现优化的滑动窗口解法
	return 0
}

// 测试辅助函数
func compareInt(a, b int) bool {
	return a == b
}

// 测试用例
func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "测试1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "测试2",
			s:    "bbbbb",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); !compareInt(got, tt.want) {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
			if got := lengthOfLongestSubstring2(tt.s); !compareInt(got, tt.want) {
				t.Errorf("lengthOfLongestSubstring2() = %v, want %v", got, tt.want)
			}
		})
	}
}
