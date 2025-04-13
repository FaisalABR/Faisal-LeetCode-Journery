package main

import "fmt"

func main() {
	list1Node3 := &ListNode{Val: 6, Next: nil}
	list1Node2 := &ListNode{Val: 5, Next: list1Node3}
	list1Node1 := &ListNode{Val: 2, Next: list1Node2}

	list2Node3 := &ListNode{Val: 1, Next: nil}
	list2Node2 := &ListNode{Val: 8, Next: list2Node3}
	list2Node1 := &ListNode{Val: 5, Next: list2Node2}

	result := addTwoNumbers(list1Node1, list2Node1)

	for result != nil {
		fmt.Print(result.Val)
		if result.Next != nil {
			fmt.Print(" -> ")
		}
		result = result.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// buat dummy head node sebagai awal list hasil
	dummy := &ListNode{}

	// buat pointer curr dan variabel carry = 0
	curr := dummy
	carry := 0

	// loop saat l1 != nil | l2 != nil | carry > 0
	for l1 != nil || l2 != nil || carry > 0 {
		// a. ambil val1 dari l1, ambil val2 dari l1, jika nil pakai 0
		val1 := 0
		val2 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}

		// b. total = val1 + val2+ carry
		total := val1 + val2 + carry

		// c. carry = total / 10
		carry = total / 10

		// d. curr.Next = total % 10
		curr.Next = &ListNode{Val: total % 10}

		// e. curr = curr.Next
		curr = curr.Next

		// f. l1 = l1.Next, l2 = l2.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	// return dummy.Next
	return dummy.Next
}
