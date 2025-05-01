/*
 * LeetCode #104: 二叉树的最大深度
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

// TreeNode 定义
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

/*
LeetCode 104: 二叉树的最大深度
难度: 简单
描述: 给定一个二叉树，找出其最大深度。最大深度是从根节点到最远叶子节点的最长路径上的节点数。
*/

/*
方法一: 递归DFS
思路: 递归计算左右子树深度，取较大值+1
时间复杂度: O(n)，每个节点访问一次
空间复杂度: O(h)，递归栈空间，h为树高度
*/
func maxDepthRecursive(root *TreeNode) int {
	// TODO: 实现递归解法
	// 提示: 处理空节点，递归左右子树，返回较大值+1
    if root == nil {
        return 0
    }
    left := maxDepthRecursive(root.Left)
    right := maxDepthRecursive(root.Right)
    if left > right {
        return left + 1
    }
    return right + 1
}

/*
方法二: 迭代BFS
思路: 使用队列层序遍历，记录遍历层数
时间复杂度: O(n)，每个节点访问一次
空间复杂度: O(n)，队列最大存储节点数
*/
func maxDepthBFS(root *TreeNode) int {
	// TODO: 实现BFS解法
	// 提示: 使用队列保存当前层节点，逐层遍历并计数
	return 0
}

// 测试函数
func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name: "测试用例1",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: 3,
		},
		{
			name:     "测试用例2",
			root:     &TreeNode{Val: 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthRecursive(tt.root); got != tt.expected {
				t.Errorf("递归法错误: 预期 %d, 实际 %d", tt.expected, got)
			}
			// if got := maxDepthBFS(tt.root); got != tt.expected {
			// 	t.Errorf("BFS法错误: 预期 %d, 实际 %d", tt.expected, got)
			// }
		})
	}
}
