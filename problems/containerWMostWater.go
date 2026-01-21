package problems

func MaxArea(height []int) int {
	left := 0
	maxArea := 0
	for right := len(height) - 1; right >= 0; right-- {
		minHeight := height[right]
		if height[left] < minHeight {
			minHeight = height[left]
		}
		area := (right - left) * minHeight

		if area > maxArea {
			maxArea = area
		}
		if left == right {
			return maxArea
		}
		if height[left] < height[right] {
			left++
			right++
		}
	}
	return maxArea
}
