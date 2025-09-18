package problems

// Mooer's Voting Algorithm :  If we pair each occurrence of the majority element with a different element, the majority element will still remain because it occurs more than half the time.

// https://leetcode.com/problems/majority-element/submissions/1773605864/

func MajorityElement(nums []int) int {
	count := 0
	var curr int

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			curr = nums[i]
		}
		if len(nums) == 1 {
			return nums[0]
		}
		if curr == nums[i] {
			count++
		} else {
			count--
		}
	}
	return curr
}

// followup . Medium
//https://leetcode.com/problems/majority-element-ii/description/

// func majorityElement2(nums []int) []int {

// }
