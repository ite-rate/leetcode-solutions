/*
 * LeetCode #19: 题目19
 * 难度: 未知
 *
 * 题目描述:
 * 由大模型直接生成
 *
 * 代码骨架完整度: 20%
 */

package leetcode

import (
	"testing"
)

// LeetCode #19 删除链表的倒数第N个节点
// 难度：中等
// 题目描述：给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头节点。
// 示例：给定链表 1->2->3->4->5 和 n = 2，删除倒数第二个节点后变为 1->2->3->5。

// ListNode 定义链表节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 解法1：双指针法（快慢指针）
// 算法思路：使用两个指针，快指针先走n步，然后快慢指针同时移动，当快指针到达末尾时，慢指针指向倒数第n个节点的前一个节点
// 时间复杂度：O(L)，其中L是链表的长度
// 空间复杂度：O(1)
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	// TODO: 创建哑节点，简化头节点删除的特殊情况处理
	// TODO: 初始化快慢指针都指向哑节点
	// TODO: 快指针先移动n步
	// TODO: 同时移动快慢指针，直到快指针到达链表末尾
	// TODO: 删除慢指针指向的下一个节点（即倒数第n个节点）
	// TODO: 返回哑节点的下一个节点作为新头节点

	return nil
}

// 解法2：计算链表长度法
// removeNthFromEnd2 删除链表的倒数第 n 个节点并返回新头节点。
// 思路：1) 先遍历求长度 L；2) 再走 L−n 步找到待删节点的前驱；3) 删除后返回。
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	// 1. 创建哑节点，统一处理「删除头节点」这种特殊情况
	dummy := &ListNode{Next: head}

	// 2. 第一遍遍历：统计链表长度 length
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	// 3. 第二遍遍历：让 prev 停在「倒数第 n+1 个节点」（即待删节点的前驱）
	prev := dummy
	for i := 0; i < length-n; i++ {
		prev = prev.Next
	}

	// 4. 删除 prev 的后继节点（倒数第 n 个节点）
	prev.Next = prev.Next.Next

	// 5. 返回新头节点
	return dummy.Next
}

// 辅助函数：创建链表

// 辅助函数：比较两个链表是否相等
func compareLinkedLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

// 辅助函数：链表转切片，便于调试

// 测试用例
func TestRemoveNthFromEnd1(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{
			name:     "删除倒数第二个节点",
			input:    []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "删除头节点",
			input:    []int{1, 2},
			n:        2,
			expected: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createLinkedList(tt.input)
			result := removeNthFromEnd1(head, tt.n)
			expected := createLinkedList(tt.expected)

			if !compareLinkedLists(result, expected) {
				t.Errorf("removeNthFromEnd1(%v, %d) = %v, expected %v",
					tt.input, tt.n, linkedListToSlice(result), tt.expected)
			}
		})
	}
}

func TestRemoveNthFromEnd2(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{
			name:     "删除倒数第二个节点",
			input:    []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "删除头节点",
			input:    []int{1, 2},
			n:        2,
			expected: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createLinkedList(tt.input)
			result := removeNthFromEnd2(head, tt.n)
			expected := createLinkedList(tt.expected)

			if !compareLinkedLists(result, expected) {
				t.Errorf("removeNthFromEnd2(%v, %d) = %v, expected %v",
					tt.input, tt.n, linkedListToSlice(result), tt.expected)
			}
		})
	}
}
