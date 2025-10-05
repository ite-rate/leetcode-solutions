/*
 * LeetCode #24: 题目24
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

// LeetCode #24 - 两两交换链表中的节点
// 难度：中等
// 题目描述：给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// 链表节点定义

// 解法1：迭代法
// 思路：使用虚拟头节点，**每次处理两个相邻节点**，调整指针指向
// 时间复杂度：O(n)，空间复杂度：O(1)
func swapPairs(head *ListNode) *ListNode {
	// TODO: 创建虚拟头节点dummy，指向head
	dummy := &ListNode{}
	dummy.Next = head
	// TODO: 设置当前指针cur指向dummy
	cur := dummy
	// TODO: 循环条件：cur.Next != nil && cur.Next.Next != nil
	for cur.Next != nil && cur.Next.Next != nil {
		// TODO: 在循环中：
		//       1. 获取第一个节点node1 = cur.Next
		//       2. 获取第二个节点node2 = cur.Next.Next
		//       3. 调整指针：cur.Next = node2, node1.Next = node2.Next, node2.Next = node1
		//       4. 移动cur指针到node1
		node1 := cur.Next      // 待交换的第一个节点
		node2 := cur.Next.Next // 待交换的第二个节点

		cur.Next = node2        // 让前驱指向第二个节点
		node1.Next = node2.Next // 第一个节点接到第二个节点后面的链
		node2.Next = node1      // 第二个节点指向第一个节点，完成交换

		cur = node1 // cur 前进到本对交换后的尾节点（原 node1）
	}

	// TODO: 返回dummy.Next
	return dummy.Next
}

// 解法2：递归法
// 思路：递归地交换每两个节点，将问题分解为子问题
// 时间复杂度：O(n)，空间复杂度：O(n) - 递归栈空间
func swapPairsRecursive(head *ListNode) *ListNode {
	// TODO: 递归终止条件：head == nil || head.Next == nil
	// TODO: 获取第二个节点newHead = head.Next
	// TODO: 递归交换后续节点：head.Next = swapPairsRecursive(newHead.Next)
	// TODO: 交换当前两个节点：newHead.Next = head
	// TODO: 返回新的头节点newHead
	return nil
}

// 辅助函数：比较两个链表是否相等
func compareLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

// 测试函数
func TestSwapPairs(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "空链表",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "四个节点",
			input:    []int{1, 2, 3, 4},
			expected: []int{2, 1, 4, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 测试迭代法
			inputList := createList(tt.input)
			expectedList := createList(tt.expected)
			result := swapPairs(inputList)
			if !compareLists(result, expectedList) {
				t.Errorf("swapPairs() = %v, want %v", listToSlice(result), listToSlice(expectedList))
			}

			// 测试递归法
			inputList2 := createList(tt.input)
			result2 := swapPairsRecursive(inputList2)
			if !compareLists(result2, expectedList) {
				t.Errorf("swapPairsRecursive() = %v, want %v", listToSlice(result2), listToSlice(expectedList))
			}
		})
	}
}

// 辅助函数：链表转切片（用于错误信息显示）
