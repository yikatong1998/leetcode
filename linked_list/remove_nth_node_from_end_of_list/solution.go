package remove_linked_list_elements

// leetcode 19 medium

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 2.1 MB, 在所有 Go 提交中击败了 61.71% 的用户
// 通过测试用例： 208 / 208

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre, post := head, head
	for i := 0; i < n; i++ {
		post = post.Next
	}
	if post == nil {
		return head.Next
	}

	for post.Next != nil {
		post = post.Next
		pre = pre.Next
	}

	pre.Next = pre.Next.Next

	return head
}
