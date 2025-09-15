package problems

import "fmt"

func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	curr := 0

	for i := 0; i < len(nums); i++ {
		curr = curr + nums[i]

		if curr > maxSum {
			maxSum = curr
		}

		if curr < 0 {
			curr = 0
		}
	}
	fmt.Println(maxSum)

	return maxSum
}
