/*
 * LeetCode #234: 题目234
 * 难度: 未知
 *
 * 题目描述:
 * 由大模型直接生成
 *
 * 代码骨架完整度: 30%
 */

package leetcode

import (
	"testing"
)

// LeetCode #234 回文链表
// 难度：简单
// 题目描述：给定一个单链表的头节点 head，判断该链表是否为回文链表。
// 回文链表是指正序和倒序读取结果一致的链表。

// 解法1：使用数组存储+双指针
// 思路：将链表值复制到数组中，然后使用双指针判断数组是否为回文
// 时间复杂度：O(n)，空间复杂度：O(n)
func isPalindromeArray(head *ListNode) bool {
	// TODO: 实现数组解法
	// 步骤提示：
	// 1. 遍历链表，将每个节点的值存入数组
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	// 2. 使用双指针，一个从数组开头，一个从数组末尾，向中间移动比较值
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 解法2：快慢指针+链表反转
// 思路：使用快慢指针找到链表中点，反转后半部分链表，然后比较前后两部分
// 时间复杂度：O(n)，空间复杂度：O(1)
func isPalindromeReverse(head *ListNode) bool {
	// TODO: 实现快慢指针+链表反转解法
	// 步骤提示：
	// 1. 使用快慢指针找到链表中点（慢指针每次走一步，快指针每次走两步）
	// 2. 反转后半部分链表
	// 3. 比较前半部分和反转后的后半部分是否相同
	// 4. 恢复链表（可选）
	return false
}

// 链表节点定义


// 测试函数
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    *ListNode
		expected bool
	}{
		{
			name:     "空链表",
			input:    nil,
			expected: true,
		},
		{
			name:     "单个节点",
			input:    &ListNode{Val: 1},
			expected: true,
		},
		{
			name: "回文链表-奇数长度",
			input: func() *ListNode {
				// 构建链表：1->2->3->2->1
				head := &ListNode{Val: 1}
				head.Next = &ListNode{Val: 2}
				head.Next.Next = &ListNode{Val: 3}
				head.Next.Next.Next = &ListNode{Val: 2}
				head.Next.Next.Next.Next = &ListNode{Val: 1}
				return head
			}(),
			expected: true,
		},
		{
			name: "非回文链表",
			input: func() *ListNode {
				// 构建链表：1->2->3->4
				head := &ListNode{Val: 1}
				head.Next = &ListNode{Val: 2}
				head.Next.Next = &ListNode{Val: 3}
				head.Next.Next.Next = &ListNode{Val: 4}
				return head
			}(),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name+"_数组解法", func(t *testing.T) {
			result := isPalindromeArray(test.input)
			if result != test.expected {
				t.Errorf("输入: %v, 期望: %v, 结果: %v", listToSlice(test.input), test.expected, result)
			}
		})

		t.Run(test.name+"_反转解法", func(t *testing.T) {
			// 复制链表用于测试（避免修改影响其他测试）
			copiedList := copyList(test.input)
			result := isPalindromeReverse(copiedList)
			if result != test.expected {
				t.Errorf("输入: %v, 期望: %v, 结果: %v", listToSlice(test.input), test.expected, result)
			}
		})
	}
}

// 辅助函数：链表转切片（用于测试输出）

// 辅助函数：复制链表

// 运行测试: go test -v
