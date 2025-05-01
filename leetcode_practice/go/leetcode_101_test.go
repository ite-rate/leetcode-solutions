/*
 * LeetCode #101: 题目101
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

/*
题目：#101 对称二叉树（Symmetric Tree）
难度：简单
描述：给定一个二叉树的根节点root，检查它是否轴对称。
*/

// TreeNode 定义二叉树节点
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

/*
解法一：递归法
算法思路：
1. 对称二叉树的条件是左子树和右子树互为镜像
2. 两个树互为镜像的条件：
   - 它们的根节点值相同
   - 每个树的左子树与另一个树的右子树互为镜像
时间复杂度：O(n)，需要遍历所有节点
空间复杂度：O(h)，h为树的高度，递归调用栈的深度
*/
func isSymmetric(root *TreeNode) bool {
	// TODO: 实现递归判断逻辑
	// 提示：可以定义一个辅助函数来比较两个子树是否对称
	if root == nil {
		return true
	}
	var isMirror func(t1, t2 *TreeNode) bool
	isMirror = func(t1, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil {
			return false
		}
		return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
	}
	return isMirror(root.Left, root.Right)
}

/*
解法二：迭代法（使用队列）
算法思路：
1. 使用队列来按层比较节点
2. 每次从队列中取出两个节点比较它们的值
3. 将左子树的左节点和右子树的右节点入队，左子树的右节点和右子树的左节点入队
时间复杂度：O(n)，需要遍历所有节点
空间复杂度：O(n)，最坏情况下队列需要存储所有叶子节点
*/
func isSymmetricIterative(root *TreeNode) bool {
	// TODO: 实现迭代判断逻辑
	// 提示：使用队列或栈来存储需要比较的节点对
	return false
}

// 测试用例
func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "对称二叉树",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 3},
				},
			},
			want: true,
		},
		{
			name: "非对称二叉树",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
			// if got := isSymmetricIterative(tt.root); got != tt.want {
			// 	t.Errorf("isSymmetricIterative() = %v, want %v", got, tt.want)
			// }
		})
	}
}
