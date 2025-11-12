/*
 * LeetCode #23: 题目23
 * 难度: 未知
 * 
 * 题目描述:
 * 由大模型直接生成
 * 
 * 代码骨架完整度: 30%
 */

package main

import (
	"container/heap"
	"testing"
)

/*
题目：23. 合并 K 个升序链表
难度：困难
描述：
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/

// ListNode 单链表节点定义
type ListNode struct {
	Val  int
	Next *Node
}

// 解法一：优先队列（最小堆）合并
// 时间复杂度：O(N·log k)，其中 N 为所有链表节点总数，k 为链表条数
// 思路：
// 1. 维护一个大小为 k 的最小堆，堆顶是当前所有链表头中的最小值
// 2. 每次取出堆顶节点，将其 next 节点（如果存在）重新放入堆
// 3. 重复直到堆为空
func mergeKListsHeap(lists []*ListNode) *ListNode {
	// TODO: 实现一个支持 *ListNode 的最小堆
	// 提示：需要实现 heap.Interface 的 Len/Less/Swap/Push/Pop 五个方法
	// 注意：Less 方法应基于 ListNode.Val 进行比较

	// TODO: 初始化堆，将所有非空链表头节点加入堆

	// TODO: 创建 dummy 头节点，用于简化链表拼接

	// TODO: 循环弹出堆顶节点，拼接到结果链表，并将该节点的 next 重新入堆（如果非空）

	return nil // TODO: 返回 dummy.Next
}

// 解法二：分治合并（两两合并）
// 时间复杂度：O(N·log k)，每次合并两个链表代价为 O(n)，共 log k 层
// 思路：
// 1. 将 k 条链表两两配对，合并成 ⌈k/2⌉ 条
// 2. 递归或迭代进行，直到只剩一条链表
func mergeKListsDivideConquer(lists []*ListNode) *ListNode {
	// TODO: 实现辅助函数 mergeTwoLists(a, b *ListNode) *ListNode
	// 提示：使用双指针归并两个已排序链表

	// TODO: 处理边界：lists 为空或长度为 0

	// TODO: 采用分治策略，递归地对半合并

	return nil // TODO: 返回最终合并后的头节点
}

// 比较函数：判断两个链表是否相等（值序列相同）
func listEqual(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

// 辅助函数：将切片转换为链表
func sliceToList(s []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range s {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// 测试用例
func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{
			name:  "示例1",
			lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			want:  []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:  "示例2",
			lists: [][]int{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 构造输入链表数组
			lists := make([]*ListNode, len(tt.lists))
			for i, s := range tt.lists {
				lists[i] = sliceToList(s)
			}
			wantList := sliceToList(tt.want)

			// 测试堆解法
			gotHeap := mergeKListsHeap(lists)
			if !listEqual(gotHeap, wantList) {
				t.Errorf("heap 方法失败")
			}

			// 测试分治解法
			gotDC := mergeKListsDivideConquer(lists)
			if !listEqual(gotDC, wantList) {
				t.Errorf("分治 方法失败")
			}
		})
	}
}
