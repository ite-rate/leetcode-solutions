/*
 * LeetCode #438: 题目438
 * 难度: 未知
 *
 * 题目描述:
 * 由大模型直接生成
 *
 * 代码骨架完整度: 10%
 */

package leetcode

import (
	// "reflect"
	"testing"
)

/*
题目：#438 找到字符串中所有字母异位词
难度：中等
描述：给定一个字符串s和一个非空字符串p，找到s中所有是p的字母异位词的子串，返回这些子串的起始索引。
     字母异位词指字母相同，但排列不同的字符串。
*/

// 解法1：滑动窗口 + 数组哈希
// 思路：使用固定长度的滑动窗口，比较窗口内字符频率与p的字符频率
// 时间复杂度：O(n), 空间复杂度：O(1)
func findAnagrams(s string, p string) []int {
	var cntP [26]int // 使用固定大小数组代替切片
	for _, c := range p {
		cntP[c-'a']++
	}

	res := []int{}
	n, m := len(s), len(p)
	if n < m {
		return res
	}

	var cntS [26]int // 使用固定大小数组代替切片
	for i := 0; i < n; i++ {
		cntS[s[i]-'a']++
		if i >= m {
			cntS[s[i-m]-'a']--
		}
		if cntS == cntP { // 直接比较数组，无需遍历和检查i>=m-1
			res = append(res, i-m+1)
		}
	}
	return res
}

// 解法2：滑动窗口优化版
// 思路：在解法1基础上优化比较过程，使用一个计数器减少不必要的全量比较
// 时间复杂度：O(n), 空间复杂度：O(1)
func findAnagrams2(s string, p string) []int {
	// TODO: 初始化结果切片和频率数组

	// TODO: 统计p的字符频率并初始化计数器

	// TODO: 使用滑动窗口，维护计数器

	// TODO: 当计数器为零时，表示找到异位词

	return nil
}

// 辅助函数：比较两个切片是否相等（忽略顺序）
// func compareSlices(a, b []int) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	for i := range a {
// 		if a[i] != b[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want []int
	}{
		{
			name: "test1",
			s:    "cbaebabacd",
			p:    "abc",
			want: []int{0, 6},
		},
		{
			name: "test2",
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.s, tt.p); !compareSlices(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
			// if got := findAnagrams2(tt.s, tt.p); !compareSlices(got, tt.want) {
			// 	t.Errorf("findAnagrams2() = %v, want %v", got, tt.want)
			// }
		})
	}
}

// 提示：
// 1. 解法1关键步骤：
//    - 先统计p的字符频率
//    - 使用双指针维护一个固定长度的窗口
//    - 比较窗口内字符频率是否与p相同
// 2. 解法2关键步骤：
//    - 使用一个计数器跟踪匹配的字符数
//    - 滑动窗口时动态更新计数器
//    - 当计数器等于p的长度时，表示找到异位词
