package dpandbinarysearch

// DP with binary search. The bottleneck of general DP is for very number at k we need to
// go over from 0 to k-1 to get the max result where k > k'. That waste us some time. If
// we want to improve that, we may need to find k' in logn time. It would be good if there
// is a sorted array storing such info.
// The way we build this data structure is use an array. But the array stores the min last
// element for all subsequence with length k. Like for array[k], it store the min last
// element for all subsequence with length k+1. If the current value is larger than array[k],
// we are confident the current element could be append to a length-k+1 list and k+2 could
// be a answer. And what we are sure is array must be sorted because the min last element
// for a short sequence must be smaller then that of a longer sequence. Otherwise, we can
// always find a smaller last element for the short sequence.
// With that data structure, we can then do binary search over the array and find the first
// element in array that is larger or equal to current value. Then we replace the found value
// with current value to update array. If all elements in array is smaller than current value,
// append current value to array and we get a new length.
// The final results would be the length of the array.
// Reference: https://leetcode.com/problems/longest-increasing-subsequence/discuss/74824/JavaPython-Binary-search-O(nlogn)-time-with-explanation
// Time: O(nlogn)
// Space: O(n)
func lengthOfLIS(nums []int) int {
	minLast := []int{}
	for _, num := range nums {
		size := len(minLast)
		i, j := 0, size-1
		for i <= j {
			mid := (i + j) / 2
			if minLast[mid] >= num && (mid == 0 || minLast[mid-1] < num) {
				i = mid
				break
			} else if minLast[mid] >= num {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}
		if i >= size {
			minLast = append(minLast, num)
		} else {
			minLast[i] = num
		}
	}
	return len(minLast)
}
