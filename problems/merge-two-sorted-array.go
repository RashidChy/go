package problems

import "fmt"

// nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
func Merge(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, 0, m+n)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] && nums1[i] != 0 {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	nums1 = result
	fmt.Println(nums1)
}
