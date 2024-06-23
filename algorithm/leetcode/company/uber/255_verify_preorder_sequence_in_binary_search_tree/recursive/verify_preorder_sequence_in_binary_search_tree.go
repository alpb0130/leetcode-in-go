package recursive

// Recursive and always compare the root with current min/max.
// Time: O(nlogn)
// Space: O(logn)
func verifyPreorder(preorder []int) bool {
	if len(preorder) == 0 {
		return true
	}
	return verify(preorder, nil, nil)
}

func verify(preorder []int, min, max *int) bool {
	if len(preorder) == 0 {
		return true
	}
	root := preorder[0]
	if min != nil && root <= *min {
		return false
	}
	if max != nil && root >= *max {
		return false
	}
	i := 1
	for i < len(preorder) {
		if preorder[i] > root {
			break
		}
		i++
	}
	return verify(preorder[1:i], min, &root) && verify(preorder[i:], &root, max)
}
