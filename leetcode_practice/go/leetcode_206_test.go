/*
 * LeetCode #206: 反转链表 (#206)
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

// 题目: 反转链表 (#206)
// 难度: 简单
// 题目描述: 给你单链表的头节点 head，请你反转链表，并返回反转后的链表。

// ListNode 定义链表节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// ================== 解法1: 迭代法 ==================
// 算法思路:
// 1. 使用三个指针: prev, curr, next
// 2. 遍历链表，逐个反转节点指向
// 3. 最后返回新的头节点(原链表的尾节点)
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func reverseList(head *ListNode) *ListNode {
	// TODO: 实现迭代法反转链表
	// 提示:
	// - 初始化prev为nil，curr为head
	// - 在循环中保存next节点，然后反转curr.Next指向prev
	// - 移动prev和curr指针向前
	// - 当curr为nil时，prev就是新的头节点
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		// 保存后续节点
		next := curr.Next
		//prev给翻转后的curr的next
		curr.Next = prev
		// 移动prev和curr指针向前
		prev = curr
		curr = next
	}
	return prev
}

// ================== 解法2: 递归法 ==================
// 算法思路:
// 1. 递归到链表末尾
// 2. 从后向前逐个反转节点指向
// 3. 返回新的头节点(原链表的尾节点)
// 时间复杂度: O(n)
// 空间复杂度: O(n) (递归栈空间)
func reverseListRecursive(head *ListNode) *ListNode {
	// TODO: 实现递归法反转链表
	// 提示:
	// - 递归终止条件: head为nil或head.Next为nil
	// - 递归调用反转head.Next之后的链表
	// - 将head.Next.Next指向head，完成反转
	// - 记得将head.Next置为nil，避免循环

	// 基本情况：空链表或只有一个节点
	if head == nil || head.Next == nil {
		return head
	}

	// 递归反转剩余部分
	newHead := reverseListRecursive(head.Next)

	// 反转当前节点与下一节点的指向
	head.Next.Next = head

	// 防止循环链表
	head.Next = nil

	// 返回新的头节点
	return newHead
}

// ================== 测试代码 ==================

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

// func copyList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	newHead := &ListNode{Val: head.Val}
// 	curr, currNew := head.Next, newHead
// 	for curr != nil {
// 		currNew.Next = &ListNode{Val: curr.Val}
// 		curr = curr.Next
// 		currNew = currNew.Next
// 	}
// 	return newHead
// }

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    *ListNode
		expected *ListNode
	}{
		{
			name: "四节点链表",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			expected: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
		// {
		// 	name:     "空链表",
		// 	input:    nil,
		// 	expected: nil,
		// },
		// {
		// 	name: "单节点链表",
		// 	input: &ListNode{
		// 		Val:  1,
		// 		Next: nil,
		// 	},
		// 	expected: &ListNode{
		// 		Val:  1,
		// 		Next: nil,
		// 	},
		// },
		// // TODO: 添加更多测试用例
		// // 例如: 多节点链表，验证反转是否正确
		// {
		// 	name: "多节点链表",
		// 	input: &ListNode{
		// 		Val: 1,
		// 		Next: &ListNode{
		// 			Val: 2,
		// 			Next: &ListNode{
		// 				Val:  3,
		// 				Next: nil,
		// 			},
		// 		},
		// 	},
		// 	expected: &ListNode{
		// 		Val: 3,
		// 		Next: &ListNode{
		// 			Val: 2,
		// 			Next: &ListNode{
		// 				Val:  1,
		// 				Next: nil,
		// 			},
		// 		},
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制链表
			input1 := copyList(tt.input)
			input2 := copyList(tt.input)
			if got := reverseList(input1); !compareList(got, tt.expected) {
				t.Errorf("reverseList() = %v, want %v", got, tt.expected)
			}
			if got := reverseListRecursive(input2); !compareList(got, tt.expected) {
				t.Errorf("reverseListRecursive() = %v, want %v", got, tt.expected)
			}
		})
	}
}

/*
以四个节点的链表为例来解释反转过程：

假设链表是：`1->2->3->4->nil`

首先，注意代码中 `prev := &ListNode{}` 是有问题的，应该是 `prev := nil`或`var prev *ListNode = nil，我会用正确的版本来讲解。

初始状态：
- `prev = nil`
- `curr = 1->2->3->4->nil`

第一次循环：
1. 保存 `next = 2->3->4->nil`
2. 反转 `1` 的指向：`1->nil`
3. 移动 `prev = 1->nil`
4. 移动 `curr = 2->3->4->nil`

第二次循环：
1. 保存 `next = 3->4->nil`
2. 反转 `2` 的指向：`2->1->nil`
3. 移动 `prev = 2->1->nil`
4. 移动 `curr = 3->4->nil`

第三次循环：
1. 保存 `next = 4->nil`
2. 反转 `3` 的指向：`3->2->1->nil`
3. 移动 `prev = 3->2->1->nil`
4. 移动 `curr = 4->nil`

第四次循环：
1. 保存 `next = nil`
2. 反转 `4` 的指向：`4->3->2->1->nil`
3. 移动 `prev = 4->3->2->1->nil`
4. 移动 `curr = nil`

循环结束，返回 `prev = 4->3->2->1->nil`，完成链表反转。
*/
