package main

import (
	"fmt"

	"github.com/RashidChy/go/problems"
)

func main() {
	fmt.Println("Hello, Go=")

	// problems

	//fmt.Println(problems.MissingNumber([]int{0, 2}))
	//fmt.Println(problems.TwoSum([]int{2, 7, 11, 15, 13}, 15))
	//problems.Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3) // merge two sorted array
	//problems.MaxProfit2([]int{7, 1, 5, 3, 6, 4}) // best time to buy and sell stocks
	//problems.ContainsDuplicate([]int{1, 2, 3, 1})
	//problems.MaxProduct([]int{2, 3, -2, 4})
	//problems.ProductExceptSelf([]int{1, 2, 3, 4})
	problems.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) // kadane's Algorithm

	// DSA

	//fmt.Println("sorted array:", DSA.MergeSort([]int{38, 27, 43, 3, 10}, 0))
}
