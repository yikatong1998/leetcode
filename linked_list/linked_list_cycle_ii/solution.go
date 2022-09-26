package linked_list_cycle_ii

// leetcode 142 medium

// 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
// 不允许修改 链表。

// 执行用时： 8 ms, 在所有 Go 提交中击败了 29.42% 的用户
// 内存消耗： 4.9 MB, 在所有 Go 提交中击败了 7.57% 的用户
// 通过测试用例： 17 / 17

func detectCycle(head *ListNode) *ListNode {
	m, node := make(map[*ListNode]bool), head
	for node != nil {
		if m[node] {
			return node
		}

		m[node] = true
		node = node.Next
	}

	return nil
}
