// Package abag provides the solution function for the alternating trees problem.
package abag

// Solution returns the number of ways of cutting out one tree if given an array a consisting of n
// integers, where A[K] denotes the height of the K-th tree, so that
// the remaining trees are aesthetically pleasing. If it is not possible to
// achieve the desired result, your function should return −1. If the trees are
// already aesthetically pleasing without any removal,
func Solution(a []int) int {
	// check if trees are already pleasing before cutting
	if isAestheticallyPleasing(a) {
		return 0
	}

	count := 0 // init count
	for i := 0; i < len(a); i++ { // iterate through trees
		// Check if the condition passes if the tree at i is removed
		if isAestheticallyPleasing(removeIndex(a, i)) {
			count++
		}
	}

	if count > 0 {
		return count
	}

	// no condition passed after cutting down exactly one tree
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
	newSlice := make([]int, len(s)-1)
	copy(newSlice[:i], s[:i])
	copy(newSlice[i:], s[i+1:])

	return newSlice
}
