package problems

import "fmt"

// [1,2,3,4]
// [1,1,2,6]
func ProductExceptSelf(nums []int) []int {
	lenArr := len(nums)
	preffixProd := make([]int, lenArr)
	suffixVar := 1

	preffixProd[0] = 1
	for i := 1; i < lenArr; i++ {
		preffixProd[i] = preffixProd[i-1] * nums[i-1]
	}

	prod := make([]int, lenArr)
	for i := lenArr - 1; i >= 0; i-- {
		prod[i] = preffixProd[i] * suffixVar
		suffixVar = nums[i] * suffixVar
	}
	fmt.Println(prod)
	return prod
}
