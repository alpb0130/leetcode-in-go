package backtracking

import "sort"

// Sort and backtracking
// Time: O(2^n)
// Space: O(n^2)
func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	res := [][]int{}
	sort.Ints(nums)
	subSetHelper(nums, 0, []int{}, &res)
	return res
}

func subSetHelper(nums []int, start int, tmp []int, res *[][]int) {
	dst := make([]int, len(tmp))
	copy(dst, tmp)
	*res = append(*res, dst)
	for i := start; i < len(nums); i++ {
		// deduplication
		if i != start && nums[i] == nums[i-1] {
			continue
		}
		subSetHelper(nums, i+1, append(tmp, nums[i]), res)
	}
}
