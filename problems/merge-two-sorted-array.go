package problems

import "fmt"

//	j                         k k k
//	  i i i
//
// nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
func Merge(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	j, k := m-1, n-1
	for j >= 0 && k >= 0 {
		fmt.Println("   j=", j, "  k=", k, "  i=", i)
		if nums1[j] > nums2[k] {
			fmt.Print("j>k  :")
			nums1[i] = nums1[j]
			fmt.Println(nums1)
			j--
		} else {
			fmt.Print("j<k  :")
			nums1[i] = nums2[k]
			fmt.Println(nums1)
			k--
		}
		i--

		println("----------")
	}
	if k > 0 {
		copy(nums1[:k+1], nums2[:k+1])
	}

	fmt.Println(nums1) //

}
