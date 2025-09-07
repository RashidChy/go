package problems

import "fmt"

// my sol complexity: Time: O(n) , Space: O(n)
func ContainsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for i, num := range nums {
		if _, exists := m[num]; exists {
			fmt.Println("true")
			return true
		} else {
			m[num] = i
		}
	}
	fmt.Println("true")
	return false
}
