package sortandarray

import "sort"

// 1. Copy to an auxiliary array
// 2. Sort the array
// 3. Copy to original array from small part and large part by turn. This step is very tricky.
//    If we start from index 0 and mid, it's not work because if the array size is small,
//    n[large] might be next to n[small+1] and they two might be equal with each other. The
//    key here is when we copy back to original array, try our best the make the neighbor element
//    separated with each other in the sorted array. So we start from mid and len-1 and then
//    decrement each index.
// Time: O(nlogn)
// Space: O(n)
func wiggleSort(nums []int)  {
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	sort.Ints(copyNums)
	i, j, k := (len(copyNums)-1)/2, len(copyNums)-1, 0
	for {
		if k == len(nums) {
			break
		}
		if k % 2 == 0 {
			nums[k] = copyNums[i]
			i--
		} else {
			nums[k] = copyNums[j]
			j--
		}
		k++
	}
}
