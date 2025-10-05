/*
 * LeetCode #2: 题目2
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

// 题目：两数相加 (Add Two Numbers)
// 难度：中等
// 题目描述：给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字0之外，这两个数都不会以0开头。



// 解法一：迭代法
// 算法思路：同时遍历两个链表，逐位相加并处理进位
// 时间复杂度：O(max(m,n))，其中m和n分别是两个链表的长度
// 空间复杂度：O(1)（不考虑结果链表）
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// TODO: 创建哑节点作为结果链表的头节点
	dummy := &ListNode{}
	// TODO: 初始化当前指针和进位值
	current := dummy
	carry := 0
	// TODO: 遍历两个链表，直到都为空且进位为0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	// TODO: 计算当前位的和（包括进位）
	// TODO: 更新进位值
	// TODO: 创建新节点存储当前位的值
	// TODO: 移动指针到下一个节点
	// TODO: 返回哑节点的下一个节点
	return dummy.Next
}

// 解法二：递归法
// 算法思路：递归处理每一位的相加，将进位传递给下一层递归
// 时间复杂度：O(max(m,n))
// 空间复杂度：O(max(m,n))（递归栈空间）
func addTwoNumbersRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersHelper(l1, l2, 0)
}

func addTwoNumbersHelper(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	// TODO: 递归终止条件：两个链表都为空且进位为0
	// TODO: 计算当前位的和（包括进位）
	// TODO: 创建新节点
	// TODO: 递归处理下一个节点
	// TODO: 返回新节点
	return nil
}

// 比较两个链表是否相等
func isEqual(l1 *ListNode, l2 *ListNode) bool {
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
func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			name: "基本测试用例",
			l1:   createLinkedList([]int{2, 4, 3}),
			l2:   createLinkedList([]int{5, 6, 4}),
			want: createLinkedList([]int{7, 0, 8}),
		},
		{
			name: "不同长度链表",
			l1:   createLinkedList([]int{9, 9, 9, 9}),
			l2:   createLinkedList([]int{1}),
			want: createLinkedList([]int{0, 0, 0, 0, 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 测试迭代法
			if got := addTwoNumbers(tt.l1, tt.l2); !isEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}

			// 测试递归法
			if got := addTwoNumbersRecursive(tt.l1, tt.l2); !isEqual(got, tt.want) {
				t.Errorf("addTwoNumbersRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
