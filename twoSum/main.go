package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		current = current.Next
	}

	return dummy.Next
}

func createLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}

	return dummy.Next
}

func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	list1 := createLinkedList([]int{2, 4, 3})
	list2 := createLinkedList([]int{5, 4, 5})

	fmt.Println("input: ")
	printLinkedList(list1)
	printLinkedList(list2)

	result := addTwoNumbers(list1, list2)

	fmt.Println("output: ")
	printLinkedList(result)
}
