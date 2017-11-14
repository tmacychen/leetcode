package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (h *ListNode) Show() {
	for i := h; ; i = i.Next {
		if i.Next == nil {
			fmt.Printf("%v\n", i.Val)
			return
		}
		fmt.Printf("%v -> ", i.Val)
	}
}
func CreateList(v int) *ListNode {
	h := new(ListNode)
	l := h
	for v != 0 {
		l.Next = new(ListNode)
		l.Next.Val = v % 10
		v = v / 10
		l = l.Next
	}
	return h.Next
}
func main() {
	L1 := CreateList(89)
	L2 := CreateList(1)
	L1.Show()
	L2.Show()

	L3 := addTwoNumbers(L1, L2)
	L3.Show()

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func checkList(l1 *ListNode, l3 *ListNode, flag int) {
	if l1 != nil {
		l3.Next = l1
		tmp := flag
		flag = (l1.Val + tmp) / 10
		l1.Val = (l1.Val + tmp) % 10
		for flag != 0 {
			if l1.Next != nil {
				l1 = l1.Next
				tmp := flag
				flag = (l1.Val + tmp) / 10
				l1.Val = (l1.Val + tmp) % 10
			} else {
				l1.Next = new(ListNode)
				l1.Next.Val = flag
				break
			}
		}
	}
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	h := l3
	flag := 0

	for l1 != nil && l2 != nil {
		l3.Next = new(ListNode)
		l3 = l3.Next
		l3.Val = (l1.Val + l2.Val + flag) % 10
		flag = (l1.Val + l2.Val + flag) / 10
		l1 = l1.Next
		l2 = l2.Next
	}
	checkList(l1, l3, flag)
	checkList(l2, l3, flag)
	if l1 == nil && l2 == nil && flag != 0 {
		l3.Next = new(ListNode)
		l3.Next.Val = flag
	}
	return h.Next
}

// the optimal version

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//    ret := new(ListNode)
//    var tmp *ListNode
//    last := 0
//    for {
//        sum := last
//        if l1 == nil && l2 ==  nil {
//            break
//        } else if l1 == nil {
//            sum += l2.Val
//            l2 = l2.Next
//        } else if l2 == nil {
//            sum += l1.Val
//            l1 = l1.Next
//        } else {
//            sum += l1.Val + l2.Val
//            l1 = l1.Next
//            l2 = l2.Next
//        }
//        if tmp == nil {
//            tmp = ret
//        } else {
//            tmp.Next = new(ListNode)
//            tmp = tmp.Next
//        }
//        tmp.Val = sum % 10
//        last = sum / 10
//    }
//    if last == 1 {
//        tmp.Next = &ListNode{Val: 1, Next: nil}
//    }
//    return ret
//}
