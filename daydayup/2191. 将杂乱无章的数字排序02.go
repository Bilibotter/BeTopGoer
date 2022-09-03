package main

import "sort"

// Go的字典真是弱爆了，比结构体慢
func getRealNum2(mapping []int, num int) int {
	if num < 10 {
		return mapping[num]
	}
	pos := 1
	conv := 0
	var c int
	for num > 0 {
		c = num % 10
		conv += mapping[c] * pos
		num /= 10
		pos *= 10
	}
	return conv
}

func sortJumbled2(mapping []int, nums []int) []int {
	maps := make(map[int]int, len(nums))
	for _, num := range nums {
		maps[num] = getRealNum2(mapping, num)
	}
	sort.Slice(nums, func(i, j int) bool {
		return maps[nums[i]] < maps[nums[j]]
	})
	return nums
}
