/*
 * LeetCode #92: 题目92
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
)

// 题目：#92 反转链表 II
// 难度：中等
// 题目描述：给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right。
// 请你反转从位置 left 到位置 right 的链表节点，返回反转后的链表。

// ListNode 定义链表节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 解法1：穿针引线法
// 算法思路：
// 1. 找到left前一个节点pre和right后一个节点succ
// 2. 反转left到right之间的链表
// 3. 将pre.Next指向反转后的头节点，反转后的尾节点指向succ
// 时间复杂度：O(n)，空间复杂度：O(1)
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// TODO: 创建哑节点dummy，处理left=1的情况
	// TODO: 找到pre节点（left前一个节点）
	// TODO: 找到rightNode节点（right位置节点）和succ节点（right后一个节点）
	// TODO: 切断rightNode.Next，便于反转
	// TODO: 反转从leftNode到rightNode的链表
	// TODO: 将pre.Next指向反转后的头节点，反转后的尾节点指向succ
	// TODO: 返回dummy.Next
	return nil
}

// 解法2：头插法
// 算法思路：
// 1. 找到left前一个节点pre
// 2. 将left+1到right的节点逐个插入到pre之后
// 时间复杂度：O(n)，空间复杂度：O(1)
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	// TODO: 创建哑节点dummy，处理left=1的情况
	// TODO: 找到pre节点（left前一个节点）
	// TODO: 定义curr指向left节点
	// TODO: 循环将curr.Next节点插入到pre之后
	// TODO: 返回dummy.Next
	dummy := &ListNode{}
	dummy.Next = head
	p0 := dummy
	for i:=0;i<left-1;i++{
		p0 = p0.Next
	}
	pre, cur := p0, p0.Next
	for i:=0;i<right-left+1;i++{
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	p0.Next.Next = cur
	p0.Next = pre

	return dummy.Next
}

// compareList 比较两个链表是否相等
// func compareList(l1, l2 *ListNode) bool {
// 	for l1 != nil && l2 != nil {
// 		if l1.Val != l2.Val {
// 			return false
// 		}
// 		l1 = l1.Next
// 		l2 = l2.Next
// 	}
// 	return l1 == nil && l2 == nil
// }

// // createList 根据数组创建链表
// func createList(nums []int) *ListNode {
// 	dummy := &ListNode{}
// 	curr := dummy
// 	for _, num := range nums {
// 		curr.Next = &ListNode{Val: num}
// 		curr = curr.Next
// 	}
// 	return dummy.Next
// }

// 测试用例
func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name   string
		head   *ListNode
		left   int
		right  int
		expect *ListNode
	}{
		{
			name:   "case1",
			head:   createList([]int{1, 2, 3, 4, 5}),
			left:   2,
			right:  4,
			expect: createList([]int{1, 4, 3, 2, 5}),
		},
		{
			name:   "case2",
			head:   createList([]int{5}),
			left:   1,
			right:  1,
			expect: createList([]int{5}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.head, tt.left, tt.right); !compareList(got, tt.expect) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.expect)
			}
			// 重置head，因为链表被修改了
			tt.head = createList([]int{1, 2, 3, 4, 5})
			if tt.name == "case2" {
				tt.head = createList([]int{5})
			}
			if got := reverseBetween2(tt.head, tt.left, tt.right); !compareList(got, tt.expect) {
				t.Errorf("reverseBetween2() = %v, want %v", got, tt.expect)
			}
		})
	}
}
