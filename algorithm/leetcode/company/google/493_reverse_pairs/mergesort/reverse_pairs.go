package mergesort

// For question like the reverse order, we can think about solving those problems using merge sort.
// The brute force solution is from back to front, use two for loop to get the result in O(n^2).
// We actually use O(n) time when we process every element to find valid pairs.
// This process can be reduce to nlogn time using merge sort. Refer to:
// https://www.youtube.com/watch?v=9LjxnRH-C6I
// https://leetcode.com/problems/reverse-pairs/solution/
// We can also solve this problem using BST or BIT
// Time: O(nlogn)
// Space: O(nlogn)
func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) int {
	if left == right {
		return 0
	}
	mid := (left + right) / 2
	leftRes := mergeSort(nums, left, mid)
	rightRes := mergeSort(nums, mid+1, right)
	// get larger
	res := 0
	i, j := left, mid+1
	for i <= mid {
		for j <= right && 2*nums[j] < nums[i] {
			j++
		}
		res += j - (mid + 1)
		i++
	}

	// merge
	buff := make([]int, right-left+1)
	idx := 0
	i, j = left, mid+1
	for i <= mid || j <= right {
		if j > right || (i <= mid && nums[i] <= nums[j]) {
			buff[idx] = nums[i]
			i++
		} else {
			buff[idx] = nums[j]
			j++
		}
		idx++
	}
	for i := 0; i < len(buff); i++ {
		nums[left+i] = buff[i]
	}
	return res + leftRes + rightRes
}
