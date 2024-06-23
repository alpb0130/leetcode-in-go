package twopointer

import "sort"

// Sort the input array first. Fix the first number, use two pointer to solve the remaining
// array. For a given second number, find the first number from the end that meet the sum
// requirement. Then we don't need to move the second point forward because we know all nums
// ahead are smaller and can meeting the requirements. The we move the first pointer one more
// step further the check again. After first pointer move, we don't need to be afraid about
// number after the current second pointer because they must not work.
// Time: O(n^2)
// Space: O(logn) - sorting need O(logn) space
func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < target {
				res += (k - j)
				j++
			} else {
				k--
			}
		}
	}
	return res
}
