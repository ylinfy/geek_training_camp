type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 保护节点, 起始节点
	dummy := &ListNode{}
	// 用于遍历链表的节点
	p, p1, p2 := dummy, l1, l2

	for p1 != nil && p2 != nil {
		// 谁小，谁排前面
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next // 指向下一个节点
	}

	// 将未完成排序的节点直接接在p上
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	// 最终返回保护节点的下一个节点，即合并链的起始节点
	return dummy.Next
}
