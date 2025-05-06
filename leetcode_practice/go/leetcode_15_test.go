/*
 * LeetCode #15: 题目15
 * 难度: 未知
 *
 * 题目描述:
 * 由大模型直接生成
 *
 * 代码骨架完整度: 10%
 */

package leetcode

import (
	"sort"
	"strconv"
	"testing"
)

/*
题目：#15 三数之和
难度：中等
描述：给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0？
请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

// 解法1：排序 + 双指针
// 思路：先排序，然后固定一个数，用双指针在剩余数组中寻找两数之和等于目标值
// 时间复杂度：O(n^2)，排序O(nlogn)，双指针遍历O(n^2)
// 空间复杂度：O(1) 或 O(n) 取决于排序实现
func threeSum(nums []int) [][]int {
	// 先排序以便去重和双指针操作
	sort.Ints(nums)
	var result [][]int

	// TODO: 实现双指针查找逻辑
	// 提示：
	// 1. 外层循环遍历每个元素作为固定值
	// 2. 跳过重复的固定值以避免重复解
	// 3. 内层使用双指针在固定值右侧寻找两数之和等于-target
	// 4. 注意处理指针移动和去重逻辑

	return result
}

// 解法2：哈希表法 (不推荐，因为去重较复杂)
// 思路：类似两数之和的哈希表解法，但需要处理重复问题
// 时间复杂度：O(n^2)
// 空间复杂度：O(n) 用于存储哈希表
func threeSumHash(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	// TODO: 实现哈希表解法
	// 提示：
	// 1. 外层双重循环，记录前两数的和
	// 2. 用哈希表查找是否存在需要的第三个数
	// 3. 需要特别注意去重逻辑，比双指针法更复杂

	return result
}

// 测试辅助函数：比较两个二维切片是否相同（忽略顺序）
func compareSlicesIgnoreOrder(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// 将每个子切片排序后转为字符串方便比较
	m := make(map[string]int)
	for _, triplet := range a {
		sort.Ints(triplet)
		key := ""
		for _, num := range triplet {
			key += strconv.Itoa(num) + ","
		}
		m[key]++
	}

	for _, triplet := range b {
		sort.Ints(triplet)
		key := ""
		for _, num := range triplet {
			key += strconv.Itoa(num) + ","
		}
		if _, ok := m[key]; !ok {
			return false
		}
		m[key]--
		if m[key] < 0 {
			return false
		}
	}

	return true
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "示例1",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "空数组",
			nums:     []int{},
			expected: [][]int{},
		},
		// TODO: 可以添加更多测试用例
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeSum(tt.nums)
			if !compareSlicesIgnoreOrder(result, tt.expected) {
				t.Errorf("threeSum(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestThreeSumHash(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "示例1",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "无解情况",
			nums:     []int{1, 2, 3, 4},
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeSumHash(tt.nums)
			if !compareSlicesIgnoreOrder(result, tt.expected) {
				t.Errorf("threeSumHash(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
