/*
 * LeetCode #141: 题目141
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

// LeetCode #141 环形链表 (Linked List Cycle)
// 难度：简单
// 题目描述：给定一个链表，判断链表中是否有环。如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。

// ListNode 定义链表节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 解法1: 哈希表法
// 思路：遍历链表，将每个节点存入哈希表，如果遇到已经存在的节点则说明有环
// 时间复杂度：O(n)，空间复杂度：O(n)
func hasCycle1(head *ListNode) bool {
	// TODO: 实现哈希表解法
	// 步骤提示：
	// 1. 创建一个map用于存储已访问的节点
	mp := make(map[*ListNode]bool)
	// 2. 遍历链表，检查当前节点是否已在map中
	for head != nil {
		// 3. 如果节点已存在，返回true
		if mp[head] {
			return true
		}
		mp[head] = true
		head = head.Next
	}
	// 4. 如果遍历结束没有重复节点，返回false
	return false
}

// 解法2: 快慢指针法 (Floyd判圈算法)
// 思路：使用两个指针，慢指针每次走一步，快指针每次走两步，如果存在环，快慢指针最终会相遇
// 时间复杂度：O(n)，空间复杂度：O(1)
func hasCycle2(head *ListNode) bool {
	// TODO: 实现快慢指针解法
	// 步骤提示：
	// 1. 初始化慢指针和快指针都指向头节点
	// 2. 循环遍历，快指针和快指针的下一个节点不为空
	// 3. 慢指针每次移动一步，快指针每次移动两步
	// 4. 如果快慢指针相遇，说明有环
	// 5. 如果快指针到达链表尾部，说明无环
	return false
}

// 辅助函数：创建带环链表
func createCycleList(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]*ListNode, len(values))
	for i, val := range values {
		nodes[i] = &ListNode{Val: val}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	if pos >= 0 && pos < len(values) {
		nodes[len(values)-1].Next = nodes[pos]
	}

	return nodes[0]
}

// 测试函数
func TestHasCycle(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		pos      int
		expected bool
	}{
		{"空链表", []int{}, -1, false},
		{"无环链表", []int{1, 2, 3, 4}, -1, false},
		{"有环链表", []int{3, 2, 0, -4}, 1, true},
		{"单个节点无环", []int{1}, -1, false},
		{"单个节点有环", []int{1}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createCycleList(tt.values, tt.pos)

			// 测试解法1
			result1 := hasCycle1(head)
			if result1 != tt.expected {
				t.Errorf("hasCycle1() = %v, want %v", result1, tt.expected)
			}

			// 测试解法2
			result2 := hasCycle2(head)
			if result2 != tt.expected {
				t.Errorf("hasCycle2() = %v, want %v", result2, tt.expected)
			}
		})
	}
}

// 基准测试
func BenchmarkHasCycle1(b *testing.B) {
	head := createCycleList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasCycle1(head)
	}
}

func BenchmarkHasCycle2(b *testing.B) {
	head := createCycleList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasCycle2(head)
	}
}
