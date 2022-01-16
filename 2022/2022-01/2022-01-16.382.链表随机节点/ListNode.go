package main

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

func (s *Solution) GetRandom() int {
	var res int
	for node, i := s.head, 1; node != nil; node = node.Next {
		if rand.Intn(i) == 0 {
			res = node.Val
		}
		i++
	}
	return res
}

func main() {

}
