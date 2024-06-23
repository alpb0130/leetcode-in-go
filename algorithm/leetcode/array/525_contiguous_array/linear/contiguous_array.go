package linear

// The problem required us to find the max length of subarray who contains the same number of 0 and
// 1 in a binary array. We using some method like prefix sum but not exactly same. Use a var
// "count", if current number is 0, count--, else count++. When the count is 0, we know the current
// subarray (start from position 0) has the same number of 0 and 1. But another case is if there is
// a prefix array having the same count number with current count. In this case the subarray should
// also be eligible.
// Reference: https://leetcode.com/problems/contiguous-array/solution/
//            https://leetcode.com/problems/contiguous-array/discuss/99655/Python-O(n)-Solution-with-Visual-Explanation
// Time: O(n)
// Space: O(n)
func findMaxLength(nums []int) int {
	count := 0
	// Do forget the initialization for prefix sum problem!!!
	countMap := map[int]int{0: -1}
	res := 0
	for i, n := range nums {
		if n == 0 {
			count--
		} else {
			count++
		}
		if idx, ok := countMap[count]; ok {
			if i-idx > res {
				res = i - idx
			}
		} else {
			countMap[count] = i
		}
	}
	return res
}
