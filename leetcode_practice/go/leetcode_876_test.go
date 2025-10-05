/*
 * LeetCode #876: 题目876
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

// 题目：链表的中间结点 (#876) 难度：简单
// 题目描述：给定一个头结点为 head 的非空单链表，返回链表的中间结点。
// 如果有两个中间结点，则返回第二个中间结点。

// ListNode 定义链表节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 解法一：两次遍历法
// 算法思路：第一次遍历统计链表长度，第二次遍历到中间位置
// 时间复杂度：O(n)
func middleNodeTwoPass(head *ListNode) *ListNode {
	// TODO: 第一次遍历计算链表长度
	// TODO: 第二次遍历到长度/2的位置
	len := 0
	for cur := head; cur != nil; cur = cur.Next {
		len++
	}
	mid := len / 2
	for i := 0; i < mid; i++ {
		head = head.Next
	}
	return head
}

// 解法二：快慢指针法
// 算法思路：使用快慢指针，快指针每次走两步，慢指针每次走一步
// 时间复杂度：O(n)
func middleNodeFastSlow(head *ListNode) *ListNode {
	slow, fast := head, head
	// TODO: 循环条件判断
	// TODO: 移动快慢指针
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 测试辅助函数：创建链表
func createLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// 测试辅助函数：链表转切片
func linkedListToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

// 比较函数：验证结果是否正确
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

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"奇数长度", []int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
		{"偶数长度", []int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_TwoPass", func(t *testing.T) {
			head := createLinkedList(tt.input)
			middle := middleNodeTwoPass(head)
			result := linkedListToSlice(middle)
			if !compareSlices(result, tt.expected) {
				t.Errorf("两次遍历法结果错误，期望%v，得到%v", tt.expected, result)
			}
		})

		t.Run(tt.name+"_FastSlow", func(t *testing.T) {
			head := createLinkedList(tt.input)
			middle := middleNodeFastSlow(head)
			result := linkedListToSlice(middle)
			if !compareSlices(result, tt.expected) {
				t.Errorf("快慢指针法结果错误，期望%v，得到%v", tt.expected, result)
			}
		})
	}
}
