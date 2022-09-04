package main

import "sort"

func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)
	if nums[0] > k {
		return int64(k * (k + 1) / 2)
	}
	inner := 0
	for i, num := range nums {
		if i > 0 && nums[i-1] == num {
			continue
		} else if num <= k {
			k += 1
			inner += num
		} else {
			break
		}
	}
	return int64((k+1)*k/2 - inner)
}

func main() {
	println(minimalKSum([]int{1}, 100000000))
}
