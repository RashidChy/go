package problems

func IsAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)

	for _, ch := range s {
		if _, ok := m[ch]; ok {
			m[ch]++
		} else {
			m[ch] = 1
		}
	}

	for _, ch := range t {
		if _, ok := m[ch]; ok {
			m[ch]--
		}
	}

	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}
