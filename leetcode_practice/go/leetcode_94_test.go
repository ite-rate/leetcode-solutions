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
	"reflect"
	"testing"
)

/*
题目：94. 二叉树的中序遍历
难度：简单
描述：
给定一个二叉树的根节点 root ，返回它的中序遍历结果。
中序遍历顺序：左子树 → 根节点 → 右子树
*/

// TreeNode 定义
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// 解法一：递归法
// 思路：利用递归天然符合“左-根-右”的顺序。
// 时间复杂度：O(n)，每个节点访问一次。
func inorderTraversalRecursive(root *TreeNode) []int {
	var res []int
	// TODO: 实现递归辅助函数，按“左-根-右”顺序收集节点值
	// 提示：先递归左子树，再记录当前节点值，最后递归右子树
	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res,node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// 解法二：迭代法（显式栈）
// 思路：用栈模拟递归过程，确保“左-根-右”顺序。
// 时间复杂度：O(n)，每个节点入栈出栈各一次。
func inorderTraversalIterative(root *TreeNode) []int {
	var res []int
	// TODO: 使用栈实现中序遍历
	// 关键步骤：
	// 1. 循环将左子节点一路入栈，直到左子为空
	// 2. 弹出栈顶节点，记录值
	// 3. 转向右子节点，重复上述过程
	return res
}

// 比较函数：中序结果顺序敏感，直接比较切片即可
func equalSlice(a, b []int) bool {
	return reflect.DeepEqual(a, b)
}

// 测试用例
func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "示例1：普通二叉树",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "示例2：空树",
			root: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_recursive", func(t *testing.T) {
			got := inorderTraversalRecursive(tt.root)
			if !equalSlice(got, tt.want) {
				t.Errorf("inorderTraversalRecursive() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+"_iterative", func(t *testing.T) {
			got := inorderTraversalIterative(tt.root)
			if !equalSlice(got, tt.want) {
				t.Errorf("inorderTraversalIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}
