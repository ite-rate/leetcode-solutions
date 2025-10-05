/*
 * LeetCode #142: 题目142
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

// LeetCode #142 环形链表 II
// 难度：中等
// 题目描述：给定一个链表的头节点head，返回链表开始入环的第一个节点。如果链表无环，则返回null。

// ListNode 定义链表节点


// 解法1：哈希表法
// 思路：遍历链表，使用哈希表记录访问过的节点，第一个重复访问的节点即为环的入口
// 时间复杂度：O(n)，空间复杂度：O(n)
func detectCycleHash(head *ListNode) *ListNode {
	// TODO: 创建哈希表记录已访问节点
	mp := make(map[*ListNode]bool)
	// TODO: 遍历链表，检查当前节点是否已在哈希表中
	for head != nil {
		if mp[head] {
			return head
		}
		mp[head] = true
		head = head.Next
	}
	// TODO: 如果遇到已访问节点，返回该节点（环入口）
	// TODO: 如果遍历到链表尾部（nil），返回nil
	return nil
}

// 解法2：快慢指针法（Floyd判圈算法）
// 思路：使用快慢指针确定链表是否有环，找到相遇点后，用两个指针分别从头部和相遇点出发，相遇点即为环入口
// 时间复杂度：O(n)，空间复杂度：O(1)
func detectCycleTwoPointers(head *ListNode) *ListNode {
	// TODO: 初始化快慢指针都指向头节点
	// TODO: 快指针每次走两步，慢指针每次走一步，直到相遇或快指针到达尾部
	// TODO: 如果无环（快指针到达nil），返回nil
	// TODO: 有环情况下，将其中一个指针重置到头节点，两个指针每次走一步，再次相遇点即为环入口
	return nil
}

// 比较函数：检查两个节点是否为同一个节点（比较地址）
func isSameNode(a, b *ListNode) bool {
	return a == b
}

// 测试函数
func TestDetectCycle(t *testing.T) {
	// 测试用例1：有环链表
	t.Run("CycleExists", func(t *testing.T) {
		// 创建链表: 1->2->3->4->5->2（形成环）
		node1 := &ListNode{Val: 1}
		node2 := &ListNode{Val: 2}
		node3 := &ListNode{Val: 3}
		node4 := &ListNode{Val: 4}
		node5 := &ListNode{Val: 5}
		
		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5
		node5.Next = node2 // 形成环
		
		// 期望环入口节点是node2
		expected := node2
		
		// 测试两种解法
		result1 := detectCycleHash(node1)
		result2 := detectCycleTwoPointers(node1)
		
		if !isSameNode(result1, expected) {
			t.Errorf("哈希表法错误: 期望 %v, 得到 %v", expected, result1)
		}
		
		if !isSameNode(result2, expected) {
			t.Errorf("快慢指针法错误: 期望 %v, 得到 %v", expected, result2)
		}
	})
	
	// 测试用例2：无环链表
	t.Run("NoCycle", func(t *testing.T) {
		// 创建链表: 1->2->3->4->5
		node1 := &ListNode{Val: 1}
		node2 := &ListNode{Val: 2}
		node3 := &ListNode{Val: 3}
		node4 := &ListNode{Val: 4}
		node5 := &ListNode{Val: 5}
		
		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5
		
		// 期望结果为nil（无环）
		var expected *ListNode = nil
		
		// 测试两种解法
		result1 := detectCycleHash(node1)
		result2 := detectCycleTwoPointers(node1)
		
		if !isSameNode(result1, expected) {
			t.Errorf("哈希表法错误: 期望 nil, 得到 %v", result1)
		}
		
		if !isSameNode(result2, expected) {
			t.Errorf("快慢指针法错误: 期望 nil, 得到 %v", result2)
		}
	})
}
