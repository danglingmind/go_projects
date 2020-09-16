package practice

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

func CreateList(items []int) *ListNode {
	head := NewListNode(-1)
	iter := head
	for i := 0; i < len(items); i++ {
		iter.Next = NewListNode(items[i])
		iter = iter.Next
	}
	return head
}

func PrintList(head *ListNode) {
	cur := head.Next
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func Reverse(head *ListNode) *ListNode {
	cur := head.Next
	var prv *ListNode
	var third *ListNode
	for cur != nil {
		third = cur.Next
		cur.Next = prv
		prv = cur
		cur = third
	}
	dummyHead := NewListNode(-1)
	dummyHead.Next = prv
	return dummyHead
}

func ReverseMN(head *ListNode, m, n int) *ListNode {
	var prv *ListNode
	cur := head.Next

	for m > 1 {
		prv = cur
		cur = cur.Next
		m--
		n--
	}

	con, tail := prv, cur

	// reverse
	var third *ListNode
	for n > 0 {
		third = cur.Next
		cur.Next = prv
		prv = cur
		cur = third
		n--
	}

	con.Next = prv
	tail.Next = cur

	return head

}
