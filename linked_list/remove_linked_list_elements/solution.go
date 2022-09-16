package remove_linked_list_elements

// leetcode 203 esay

// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

// 执行用时： 4 ms, 在所有 Go 提交中击败了 97.12% 的用户
//内存消耗： 4.5 MB, 在所有 Go 提交中击败了 63.45% 的用户
//通过测试用例： 66 / 66

func removeElements(head *ListNode, val int) *ListNode {
	var ret = head
	for ret != nil && ret.Val == val {
		ret = ret.Next
	}
	if ret == nil {
		return nil
	}

	var pre, post = ret, ret.Next
	for post != nil {
		if post.Val != val {
			pre = post
		}

		post = post.Next
		pre.Next = post
	}

	return ret
}
