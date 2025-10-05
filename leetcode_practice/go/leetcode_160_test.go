/*
 * LeetCode #160: 题目160
 * 难度: 未知
 * 
 * 题目描述:
 * 由大模型直接生成
 * 
 * 代码骨架完整度: 100%
 */

package leetcode

import (
	"testing"
)

// 题目：相交链表 (Intersection of Two Linked Lists)
// 难度：简单
// 题目描述：编写一个程序，找到两个单链表相交的起始节点。如果两个链表没有交点，返回 null。

// 解法1: 双指针法
// 算法思路：使用两个指针分别遍历两个链表，当到达链表末尾时切换到另一个链表头部继续遍历。
// 如果两个链表相交，指针会在交点相遇；如果不相交，最终会同时到达nil。
// 时间复杂度：O(m+n)，空间复杂度：O(1)

// ListNode 定义链表节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	
	// TODO: 初始化两个指针分别指向两个链表的头部
	// a, b := headA, headB
	prtA, prtB := headA, headB	
	// TODO: 循环遍历直到两个指针相遇
	// 当a到达链表末尾时，重定位到headB；当b到达链表末尾时，重定位到headA
	// 如果存在交点，最终a和b会在交点相遇；如果不相交，最终会同时为nil
	for prtA != prtB {
		if prtA == nil {
			prtA = headB
		} else {
			prtA = prtA.Next
		}
		if prtB == nil {
			prtB = headA
		} else {
			prtB = prtB.Next
		}
	}

	// 提示：需要处理不相交的情况，即两个指针都遍历完两个链表后同时为nil

	
	return prtA
}

// 解法2: 哈希表法
// 算法思路：遍历链表A并将所有节点地址存入哈希表，然后遍历链表B检查节点是否在哈希表中。
// 时间复杂度：O(m+n)，空间复杂度：O(m)或O(n)
func getIntersectionNodeHash(headA, headB *ListNode) *ListNode {
	// TODO: 创建哈希表存储节点地址
	// visited := make(map[*ListNode]bool)
	
	// TODO: 遍历链表A，将所有节点加入哈希表
	
	// TODO: 遍历链表B，检查每个节点是否在哈希表中存在
	
	// 提示：注意处理链表为nil的情况
	
	return nil
}

// 测试函数
func TestGetIntersectionNode(t *testing.T) {
	// 创建测试链表
	// 链表结构： 
	// listA: 4->1->8->4->5
	// listB: 5->0->1->8->4->5
	// 相交节点为值为8的节点
	
	intersectNode := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	
	listA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: intersectNode}}
	listB := &ListNode{Val: 5, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: intersectNode}}}
	
	// 测试用例1: 相交链表
	t.Run("IntersectingLists", func(t *testing.T) {
		result := getIntersectionNode(listA, listB)
		if result != intersectNode {
			t.Errorf("期望相交节点 %v, 实际得到 %v", intersectNode, result)
		}
	})
	
	// 测试用例2: 不相交链表
	t.Run("NonIntersectingLists", func(t *testing.T) {
		listC := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
		listD := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}
		
		result := getIntersectionNode(listC, listD)
		if result != nil {
			t.Errorf("期望 nil, 实际得到 %v", result)
		}
	})
	
	// 测试用例3: 其中一个链表为空
	t.Run("OneEmptyList", func(t *testing.T) {
		result := getIntersectionNode(listA, nil)
		if result != nil {
			t.Errorf("期望 nil, 实际得到 %v", result)
		}
	})
}

// 比较函数，用于验证链表节点是否相同
func compareNodes(node1, node2 *ListNode) bool {
	return node1 == node2 // 直接比较指针地址
}
