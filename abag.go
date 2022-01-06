package abag

func Solution(a []int) int {
	if isAestheticallyPleasing(a) {
		return 0
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if isAestheticallyPleasing(removeIndex(a, i)) {
			count++
		}
	}

	if count > 0 {
		return count
	}

	return -1
}

func isAestheticallyPleasing(tree []int) bool {
	for i := 1; i < len(tree)-1; i++ {
		if tree[i] >= tree[i-1] && tree[i] <= tree[i+1] ||
			tree[i] <= tree[i-1] && tree[i] >= tree[i+1] {
			return false
		}
	}

	return true
}

func removeIndex(s []int, i int) []int {
	newSlice := make([]int, 0)
	newSlice = append(newSlice, s[:i]...)
	return append(newSlice, s[i+1:]...)
}
