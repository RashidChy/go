package problems

import "fmt"

// import "fmt"

func StrStr(haystack string, needle string) int {
	occurence := []int{}

	for i := 0; i < len(haystack); i++ {
		for j := 0; j < len(needle); j++ {
			if i+j < len(haystack) && haystack[i+j] == needle[j] {
				if j == len(needle)-1 {
					occurence = append(occurence, i)
				}
			} else {
				break
			}
		}
	}
	fmt.Println("occurance  = ", occurence)
	if len(occurence) == 0 {
		return -1
	} else {
		return occurence[0]
	}
}
