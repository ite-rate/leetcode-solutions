/*
 * LeetCode #21: 题目21
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

// LeetCode #21 合并两个有序链表
// 难度：简单
// 题目描述：将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 解法1：迭代法
// 算法思路：使用一个哑节点作为新链表的起始点，比较两个链表的当前节点值，将较小值的节点连接到新链表
// 时间复杂度：O(n+m)，其中n和m分别是两个链表的长度
// 空间复杂度：O(1)
func mergeTwoListsIterative(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建哑节点作为新链表的起始点
	dummy := &ListNode{}
	current := dummy

	// TODO: 遍历两个链表，比较节点值，将较小值的节点连接到新链表
	// 关键步骤提示：
	// 1. 当两个链表都还有节点时，比较当前节点值
	// 2. 将较小值的节点连接到current.Next
	// 3. 移动被连接节点的指针到下一个节点
	// 4. 移动current指针到新连接的节点
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// TODO: 处理剩余节点（如果某个链表还有剩余节点）
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

// 解法2：递归法
// 算法思路：递归地比较两个链表的头节点，将较小值的节点作为当前节点，然后递归处理剩余部分
// 时间复杂度：O(n+m)
// 空间复杂度：O(n+m)，递归调用栈的深度
func mergeTwoListsRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	// TODO: 实现递归终止条件
	// 关键步骤提示：
	// 1. 如果l1为空，返回l2
	// 2. 如果l2为空，返回l1

	// TODO: 比较两个链表的头节点值
	// 关键步骤提示：
	// 1. 如果l1的值小于等于l2的值，将l1作为当前节点
	// 2. l1.Next指向递归调用合并l1.Next和l2的结果
	// 3. 返回l1
	// 4. 否则，将l2作为当前节点，类似处理

	return nil
}

// 测试用例
func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			name:     "两个空链表",
			l1:       createList([]int{}),
			l2:       createList([]int{}),
			expected: createList([]int{}),
		},
		{
			name:     "一个空链表和一个非空链表",
			l1:       createList([]int{}),
			l2:       createList([]int{1, 3, 5}),
			expected: createList([]int{1, 3, 5}),
		},
		{
			name:     "两个非空链表",
			l1:       createList([]int{1, 2, 4}),
			l2:       createList([]int{1, 3, 4}),
			expected: createList([]int{1, 1, 2, 3, 4, 4}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_Iterative", func(t *testing.T) {
			result := mergeTwoListsIterative(tt.l1, tt.l2)
			if !compareList(result, tt.expected) {
				t.Errorf("迭代法测试失败 %s", tt.name)
			}
		})

		t.Run(tt.name+"_Recursive", func(t *testing.T) {
			// 重新创建链表，因为迭代法已经修改了原链表
			l1 := createList(getListValues(tt.l1))
			l2 := createList(getListValues(tt.l2))
			expected := createList(getListValues(tt.expected))

			result := mergeTwoListsRecursive(l1, l2)
			if !compareList(result, expected) {
				t.Errorf("递归法测试失败 %s", tt.name)
			}
		})
	}
}

// 辅助函数：获取链表的值切片
func getListValues(head *ListNode) []int {
	values := []int{}
	current := head
	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}
	return values
}

// 运行测试: go test -v
