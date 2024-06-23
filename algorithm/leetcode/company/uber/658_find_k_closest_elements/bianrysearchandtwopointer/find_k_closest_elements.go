package bianrysearchandtwopointer

import "math"

// Use binary search to find the one/two index where the target number is in the mid. Be careful
// about the case where the input array has duplications. We always want to find the smallest case
// Time: O(logn+k)
// Space: O(1)
func findClosestElements(arr []int, k int, x int) []int {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}
	// Two edge case
	if x <= arr[0] {
		return arr[:k]
	}
	if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}
	//binary search to find the pivotal index
	// if x is in the array, i==j and i, j point to the value with min index
	// if x is not in the array, i = j-1 and nums[i] < x < nums[j]
	i, j := binarySearch(arr, x)
	// final subarray boundary which should be inclusive
	left, right := i+1, j-1
	// two pointer to update the final subarray boundary
	// Every time when we decide to use i or j, we update the
	// boundary.
	for right-left+1 < k && (i >= 0 || j < len(arr)) {
		var idx int
		if i < 0 {
			idx = j
			j++
		} else if j > len(arr)-1 {
			idx = i
			i--
		} else {
			if x-arr[i] <= arr[j]-x {
				idx = i
				i--
			} else {
				idx = j
				j++
			}
		}
		left = minInt(left, idx)
		right = maxInt(right, idx)
	}
	return arr[left : right+1]
}

func binarySearch(arr []int, x int) (int, int) {
	i, j := 0, len(arr)-1
	for i <= j {
		mid := (i + j) / 2
		if arr[mid] == x && (mid == 0 || arr[mid-1] < x) {
			// we find the target the value but we only want to return the one
			// with smallest index if there are duplications.
			return mid, mid
		} else if arr[mid] >= x {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	// in this case, i > j but we want to return the correct order
	return j, i
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
