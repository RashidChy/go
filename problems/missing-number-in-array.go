package problems

func MissingNumber(nums []int) int {
	var flag int
	for i := 0; i <= len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == nums[j] {
				break
			}
			if j == len(nums)-1 {
				flag = i
			}
		}
	}
	return flag
}
