package prefixsumandmap

// Use prefix sum and hashmap. For the same prefix sum, we always use the smaller index in map
// so that we can get max sub-array length. For subarray sum that equals to some value, prefix
// sum and map is general direction we should think of.
// Time: O(n)
// Space: O(n)
func maxSubArrayLen(nums []int, k int) int {
	res := 0
	sumMap := map[int]int{0: -1}
	sum := 0
	for i, num := range nums {
		sum += num
		if j, ok := sumMap[sum-k]; ok {
			if i-j > res {
				res = i - j
			}
		}
		if _, ok := sumMap[sum]; !ok {
			sumMap[sum] = i
		}
	}
	return res
}
