package remove_linked_list_elements

// leetcode 24 easy

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗： 1.9 MB, 在所有 Go 提交中击败了 99.84% 的用户
//通过测试用例： 55 / 55

func swapPairs(head *ListNode) *ListNode {
	var p = head
	for p != nil && p.Next != nil {
		var tmp = p.Val
		p.Val = p.Next.Val
		p.Next.Val = tmp
		p = p.Next.Next
	}

	return head
}
