package problems

import "fmt"

// leetcode medium
// Algorithm: bucket sort

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	bucket := make([][]int, len(nums)+1)

	for key, val := range m {
		bucket[val] = append(bucket[val], key)
	}

	fmt.Println(bucket)
	top := []int{}
	for i := len(bucket) - 1; i > 0; i-- {
		if len(bucket[i]) > 0 {
			top = append(top, bucket[i]...)
		}
	}
	fmt.Println(top[:k])
	return top
}
