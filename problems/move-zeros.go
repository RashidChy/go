package problems

import "fmt"

/* Leetcode 283
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

0103 12
1003 12
1300 12




Example 1:

Input: nums = [0,1,0,3,12]

Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]
*/

func MoveZeroes(nums []int) {
	left := 0
	if len(nums) <= 1 {
		return
	}
	for right := 0; right < len(nums); right++ {
		if nums[left] != 0 {
			left++
		} else {
			if nums[right] != 0 && nums[left] == 0 {
				nums[left] = nums[right]
				nums[right] = 0
				left++
			}
		}
	}
	fmt.Println(nums)
}
