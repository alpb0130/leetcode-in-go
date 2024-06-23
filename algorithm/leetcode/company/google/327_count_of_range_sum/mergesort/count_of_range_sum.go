package mergesort

import "math"

// We still use prefix sum to help us with this problem. If we have prefix sum, what we want to
// find is sum pairs that meet lower <= prefix[j] - prefix[i] <= upper. Think about problem 493
// reverse pairs. We know we can counting pairs from left part and right part which meet those
// requirements. BTW, we need two index to track the range. Because the relationship is monotonic,
// we can finish the counting process in O(n). Finally, don't forget merge.
// Reference: https://leetcode.com/problems/count-of-range-sum/discuss/77990/Share-my-solution
// Time: O(nlogn)
// Space: O(nlogn)
func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}

	prefix := make([]int, len(nums)+1)
	for i, num := range nums {
		prefix[i+1] = prefix[i] + num
	}
	return mergeSort(prefix, 0, len(nums), lower, upper)
}

func mergeSort(prefix []int, left, right, lower, upper int) int {
	if left >= right {
		return 0
	}
	mid := (left + right) / 2
	count := mergeSort(prefix, left, mid, lower, upper) + mergeSort(prefix, mid+1, right, lower, upper)

	// counting
	i, j, k := left, mid+1, mid+1
	for i <= mid {
		for j <= right && prefix[j]-prefix[i] < lower {
			j++
		}
		for k <= right && prefix[k]-prefix[i] <= upper {
			k++
		}
		count += k - j
		i++
	}

	// copy
	tmp := make([]int, right-left+1)
	i, j, k = left, mid+1, 0
	for i <= mid || j <= right {
		if i <= mid && j <= right {
			if prefix[j] < prefix[i] {
				tmp[k] = prefix[j]
				j++
			} else {
				tmp[k] = prefix[i]
				i++
			}
		} else if i <= mid {
			tmp[k] = prefix[i]
			i++
		} else {
			tmp[k] = prefix[j]
			j++
		}
		k++
	}
	for i := 0; i < len(tmp); i++ {
		prefix[left+i] = tmp[i]
	}

	return count
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
