/*
 * LeetCode #94: 题目94
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

// 题目：#94 二叉树的中序遍历
// 难度：简单
// 描述：给定一个二叉树的根节点root，返回它的中序遍历结果（左-根-右）。

// TreeNode 定义二叉树节点
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// 解法1：递归法
// 思路：按照左-根-右的顺序递归遍历
// 时间复杂度：O(n)，每个节点访问一次
// 空间复杂度：O(h)，h为树的高度，递归栈空间
func inorderTraversalRecursive(root *TreeNode) []int {
	var res []int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		// TODO: 按照中序遍历顺序递归调用
		// 提示：1.先递归左子树 2.访问当前节点值 3.递归右子树
		traverse(node.Left)
		res = append(res, node.Val)
		traverse(node.Right)
	}
	traverse(root)
	return res
}

// 解法2：迭代法（使用栈）
// 思路：使用栈模拟递归过程，显式维护调用栈
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func inorderTraversalIterative(root *TreeNode) []int {
	var res []int
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		// TODO: 实现迭代中序遍历
		// 提示：
		// 1. 当当前节点不为空时，将其入栈，并转向左子节点
		// 2. 先将所有左子节点入栈
		// 3. 弹出栈顶节点并访问
		// 4. 转向右子节点
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		// 转向右子节点
		curr = curr.Right
	}
	return res
}

// 测试辅助函数：比较两个切片是否相等
func compareSlices(a, b []int) bool {
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

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name: "普通二叉树",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{1, 3, 2},
		},
		{
			name:     "空树",
			root:     nil,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 测试递归解法
			// if got := inorderTraversalRecursive(tt.root); !compareSlices(got, tt.expected) {
			// 	t.Errorf("inorderTraversalRecursive() = %v, want %v", got, tt.expected)
			// }
			// 测试迭代解法
			if got := inorderTraversalIterative(tt.root); !compareSlices(got, tt.expected) {
				t.Errorf("inorderTraversalIterative() = %v, want %v", got, tt.expected)
			}
		})
	}
}
