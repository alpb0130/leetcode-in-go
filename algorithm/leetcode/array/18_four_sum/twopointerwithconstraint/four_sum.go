package twopointerwithconstraint

import "sort"

// Solve the problem using 3sum. Take care of the duplication. But we also add some constraints
// to prune and save time.
// Reference:
// https://leetcode.com/problems/4sum/discuss/8723/There-maybe-isn't-a-n2lgn-or-better-time-complexity-solution
// https://leetcode.com/problems/4sum/discuss/8565/Lower-bound-n3
// Time: O(n^3)
// Space: O(1)
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	if len(nums) < 4 {
		return res
	}
	min := nums[0]
	max := nums[len(nums)-1]
	if 4*min > target || 4*max < target {
		return res
	}
	for i := 0; i < len(nums)-3; i++ {
		min = nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if 4*min > target {
			break
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if min+3*nums[j] > target {
				break
			}
			k := j + 1
			l := len(nums) - 1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}
	}
	return res
}
