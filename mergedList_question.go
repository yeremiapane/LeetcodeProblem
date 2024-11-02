package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.value < list2.value {
			tail.next = list1

		} else {
			tail.next = list2
			list2 = list2.next
		}
		tail = tail.next
	}
	if list1 != nil {
		tail.next = list1
	} else {
		tail.next = list2
	}
	return dummy.next
}

func main() {
	l1 := &ListNode{value: 1}
	l1.next = &ListNode{value: 2}
	l1.next.next = &ListNode{value: 4}

	l2 := &ListNode{value: 1}
	l2.next = &ListNode{value: 3}
	l2.next.next = &ListNode{value: 4}

	merged := mergeTwoLists(l1, l2)
	fmt.Println(merged)
}
