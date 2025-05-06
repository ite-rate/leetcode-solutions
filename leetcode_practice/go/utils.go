package leetcode

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


// compareList 比较两个链表是否相等
func compareList(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func copyList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{Val: head.Val}
	curr, currNew := head.Next, newHead
	for curr != nil {
		currNew.Next = &ListNode{Val: curr.Val}
		curr = curr.Next
		currNew = currNew.Next
	}
	return newHead
}
// createList 根据数组创建链表
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, num := range nums {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}
	return dummy.Next
}

// 辅助函数：将链表转换为slice便于比较
func listToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}