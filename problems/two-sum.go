package problems

import "fmt"

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	// for i, num := range nums {
	// 	m[num] = i
	// 	fmt.Println("Mapping:", num, "to index", i)
	// }
	for i, num := range nums {
		fmt.Println("Checking number:", num, "at index", i)
		complement := target - num
		fmt.Println("Complement for", num, "is", complement)
		if j, ok := m[complement]; ok && j != i {
			return []int{i, j}
		}

		m[num] = i
		fmt.Println("Mapping:", num, "to index", i)
	}
	// Return an empty slice if no solution is found
	return []int{}
}
