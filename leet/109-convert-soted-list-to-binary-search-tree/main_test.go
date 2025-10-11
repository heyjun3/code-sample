package main_test

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	prev := head
	slow := head.Next
	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	root := &TreeNode{
		Val: slow.Val,
	}

	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)
	return root
}
