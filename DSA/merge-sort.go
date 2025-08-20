package DSA

import (
	"fmt"
	"strings"
)

// merge combines two sorted slices into one sorted slice
// func merge(left, right []int) []int {
// 	result := make([]int, 0, len(left)+len(right))
// 	i, j := 0, 0

// 	// Compare elements from both slices and pick the smaller
// 	for i < len(left) && j < len(right) {
// 		if left[i] <= right[j] {
// 			result = append(result, left[i])
// 			i++
// 		} else {
// 			result = append(result, right[j])
// 			j++
// 		}
// 	}

// 	// Append any remaining elements
// 	result = append(result, left[i:]...)
// 	result = append(result, right[j:]...)

// 	return result
// }

// // mergeSort recursively splits and merges
// func MergeSort(arr []int) []int {
// 	// Base case: array of length 1 is already sorted
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	// Split array into two halves
// 	mid := len(arr) / 2
// 	fmt.Println("Splitting array:", arr, "at index", mid)
// 	left := MergeSort(arr[:mid])
// 	fmt.Println("Left half after sort:", left)
// 	right := MergeSort(arr[mid:])
// 	fmt.Println("Right half after sort:", right)

// 	fmt.Println("Left half:", left, "Right half:", right)
// 	fmt.Print("------------------------------------\n")

// 	// Merge sorted halves
// 	return merge(left, right)
// }

func MergeSort(arr []int, depth int) []int {
	indent := strings.Repeat("  ", depth)
	fmt.Println(indent+"call:", arr)

	if len(arr) <= 1 {
		fmt.Println(indent+"return:", arr)
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid], depth+1)
	fmt.Println("Right:")
	right := MergeSort(arr[mid:], depth+1)
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
