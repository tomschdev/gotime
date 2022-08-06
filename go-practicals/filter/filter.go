package filter

func filter(pred func(int) bool, values []int) []int {
	filtered := make([]int, 0, len(values))
	for _, val := range values {
		if pred(val) {
			filtered = append(filtered, copy(val))
		}
	}
	return filtered
}

func isOdd(n int) bool {
	if n%2 == 1 {
		return true
	}
	return false
}
