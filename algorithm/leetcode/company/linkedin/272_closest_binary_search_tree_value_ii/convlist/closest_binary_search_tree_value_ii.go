package convlist

import (
	"algorithm/leetcode/company/linkedin/models"
	"math"
)

// Convert BST to sorted list and use binary search to find the pivotal idx (where value smaller
// than target but the next value larger or equal than target). And then use two pointer to find
// the top K value.
// Time: O(n)+O(logn)+O(k)
// Space: O(n)
func closestKValues(root *models.TreeNode, target float64, k int) []int {
	list := []int{}
	for root != nil {
		if root.Left != nil {
			pred := root.Left
			for pred.Right != nil && pred.Right != root {
				pred = pred.Right
			}
			if pred.Right == root {
				pred.Right = nil
			} else {
				pred.Right = root
				root = root.Left
				continue
			}
		}
		list = append(list, root.Val)
		root = root.Right
	}
	start, end := 0, len(list)-1
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		if float64(list[mid]) < target && (mid+1 >= len(list) || float64(list[mid+1]) >= target) {
			break
		} else if float64(list[mid]) < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	left := mid
	right := mid + 1
	res := []int{}
	for len(res) < k {
		if isValid(len(list), left) && isValid(len(list), right) {
			leftDiff := math.Abs(float64(list[left]) - target)
			rightDiff := math.Abs(float64(list[right]) - target)
			if leftDiff <= rightDiff {
				res = append(res, list[left])
				left--
			} else {
				res = append(res, list[right])
				right++
			}
		} else if isValid(len(list), left) {
			res = append(res, list[left])
			left--
		} else {
			res = append(res, list[right])
			right++
		}
	}
	return res
}

func isValid(length, idx int) bool {
	return idx >= 0 && idx < length
}
