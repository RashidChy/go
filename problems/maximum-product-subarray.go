package problems

func MaxProduct(nums []int) int {
	product := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			product = product * nums[i]
		}
	}
	return product
}
