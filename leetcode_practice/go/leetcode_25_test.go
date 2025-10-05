/*
 * LeetCode #25: 题目25
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

// ListNode 定义链表节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

/*
题目25. K 个一组翻转链表（困难）
描述：给你链表的头节点 head ，每 k 个节点一组进行翻转，返回修改后的链表。
如果节点总数不是 k 的整数倍，最后剩余的节点保持原有顺序。
*/

/*
解法一：递归法
算法思路：
1. 先检查后续是否还有k个节点
2. 翻转当前k个节点
3. 递归处理后续链表
4. 将翻转后的子链表连接起来
时间复杂度：O(n)，每个节点被处理一次
空间复杂度：O(n/k) 递归栈深度
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	// TODO: 检查剩余节点是否足够k个
	// TODO: 翻转当前k个节点
	// TODO: 递归处理后续链表并连接
	return head
}

/*
解法二：迭代法
算法思路：
1. 使用dummy节点简化头节点处理
2. 维护pre和end指针标记当前组的前驱和末尾
3. 翻转当前组并重新连接链表
时间复杂度：O(n)
空间复杂度：O(1)
*/
func reverseKGroup2(head *ListNode, k int) *ListNode {
	len := 0
	for cur := head; cur != nil; cur = cur.Next {
		len++
	}

	dummy := &ListNode{}
	dummy.Next = head
	p0 := dummy
	pre, cur := p0, p0.Next

	for n := len; n >= k; n -= k {
		for i := 0; i < k; i++ {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		currentGroupOriginalHead := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = currentGroupOriginalHead
	}
	return dummy.Next
}

// 辅助函数：将链表转换为slice便于比较
// func listToSlice(head *ListNode) []int {
// 	var res []int
// 	for head != nil {
// 		res = append(res, head.Val)
// 		head = head.Next
// 	}
// 	return res
// }

// 测试用例
func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		k    int
		want []int
	}{
		{
			name: "case1",
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:    2,
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "case2",
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:    3,
			want: []int{3, 2, 1, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(tt.head, tt.k)
			if !compareSlice(listToSlice(got), tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", listToSlice(got), tt.want)
			}
		})
	}
}

// 比较slice是否相等
func compareSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
