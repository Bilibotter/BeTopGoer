package main

import "sort"

// 击败100%
// Go的字典真是弱爆了，比结构体慢

type Node struct {
	num1, num2 int
}

func getReal(mapping []int, num int) int {
	if num < 10 {
		return mapping[num]
	}
	pos := 1
	num2 := 0
	for num > 0 {
		num2 += mapping[num%10] * pos
		num /= 10
		pos *= 10
	}
	return num2
}

func sortJumbled(mapping []int, nums []int) []int {
	nodes := make([]*Node, len(nums))
	for i, num := range nums {
		nodes[i] = &Node{num, getReal(mapping, num)}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].num2 < nodes[j].num2
	})
	for i, node := range nodes {
		nums[i] = node.num1
	}
	return nums
}
