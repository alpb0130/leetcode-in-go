package linear

// Use prefix sum. The problem asks whether there is a subarray (length >= 2) that is sum to n*k
// where n is integer. We can brute force this problem but it's not good (O(n^2)). We can reuse the
// method in 560_subarray_sum_equals_k and use a hashmap to note down all visited sum and index.
// Because it requires us to check whether there is a subarray (length >= 2) that is equal to n*k,
// we only need to check is there a existing prefix sum s1 that current prefix sum s have
// (s-s1)%k == 0. So s%k == s1%k. All the operation can do based on modulo k.
// But some edge cases need to consider:
// 1. k is 0. We cannot do % k but in this case, we can just use sum directly. If there are at least
//    2 consecutive 0s, the answer is true.
// 2. k is negative number. Not a big trouble because the build in % operation will transfer the
//    result to non-negative number.
// Time: O(n)
// Space: O(k)
func checkSubarraySum(nums []int, k int) bool {
	sumMap := map[int]int{0: -1}
	sum := 0
	for i, num := range nums {
		sum += num
		if k != 0 {
			sum %= k
		}
		if idx, ok := sumMap[sum]; ok {
			if i-idx > 1 {
				return true
			}
		} else {
			sumMap[sum] = i
		}
	}
	return false
}
