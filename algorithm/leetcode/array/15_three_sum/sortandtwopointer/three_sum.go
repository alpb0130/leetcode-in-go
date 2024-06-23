package sortandtwopointer

import "sort"

// Sort and use two pointer to solve the problem.
// The problem requires there is no duplication in the result so we need to do dedupe.
// First of all, if two consecutive pivotal numbers are the same, skip the later one.
// Second, during the two pointer process, if we find a solution, we need to move both
// j and k but here we should be super careful because there would be case like [a,a,b,b].
// If j, k = 0, 3, after moving, nums at j, k are still the same. So we should move j, k
// till they are the last duplication and then additional moving.
// Time: O(n^2)
// Space: O(logn) - sorting need O(logn) space
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}
