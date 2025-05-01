/*
 * LeetCode #226: 题目226
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
226. 翻转二叉树
难度：简单
描述：给定一个二叉树的根节点 root，翻转这棵二叉树，并返回其根节点。
即交换每个节点的左右子树。
*/

// TreeNode 定义二叉树节点
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

/*
解法一：递归法
思路：递归地交换每个节点的左右子树
时间复杂度：O(n)，需要访问每个节点一次
空间复杂度：O(h)，递归栈的深度等于树的高度
*/
func invertTree(root *TreeNode) *TreeNode {
	// TODO: 实现递归翻转逻辑
	// 提示：
	// 1. 递归终止条件：节点为空
	// 2. 交换当前节点的左右子树
	// 3. 递归处理左右子树
	if root == nil {
		return nil
	}
	// 交换左右子树
	root.Left, root.Right = root.Right, root.Left
	// 递归处理左右子树
	invertTree(root.Left)
	invertTree(root.Right)
	// 返回翻转后的根节点
	return root
}

/*
解法二：迭代法（使用队列进行BFS）
思路：使用队列层序遍历二叉树，交换每个节点的左右子树
时间复杂度：O(n)，需要访问每个节点一次
空间复杂度：O(n)，最坏情况下队列需要存储所有叶子节点
*/
func invertTreeIterative(root *TreeNode) *TreeNode {
	// TODO: 实现迭代翻转逻辑
	// 提示：
	// 1. 使用队列辅助进行层序遍历
	// 2. 每次从队列取出节点时交换其左右子树
	// 3. 将非空的左右子树加入队列
	if root == nil {
		return nil
	}
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		// 取出队列的第一个节点
		node := queue[0]
		queue = queue[1:]
		// 交换左右子树
		node.Left, node.Right = node.Right, node.Left
		// 将非空的左右子树加入队列
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}
	return root
}

// 辅助函数：比较两棵二叉树是否相同
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{
			name: "测试用例1",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			name: "测试用例2",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			want: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 2},
			},
		},
	}

	for _, tt := range tests {
		// 测试递归方法
		t.Run(tt.name+" 递归", func(t *testing.T) {
			root := cloneTree(tt.root)
			if got := invertTree(root); !isSameTree(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})

		// 测试迭代方法
		t.Run(tt.name+" 迭代", func(t *testing.T) {
			root := cloneTree(tt.root)
			if got := invertTreeIterative(root); !isSameTree(got, tt.want) {
				t.Errorf("invertTreeIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 深度克隆二叉树
func cloneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	newRoot := &TreeNode{Val: root.Val}
	newRoot.Left = cloneTree(root.Left)
	newRoot.Right = cloneTree(root.Right)

	return newRoot
}
