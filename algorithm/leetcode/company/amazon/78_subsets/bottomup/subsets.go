package bottomup

// Bottom up method
// Time: O(2^n)
// Space O(2^n)
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	sets := subsets(nums[1:])
	res := [][]int{}
	for _, set := range sets {
		res = append(res, set)
		setCopy := make([]int, len(set))
		copy(setCopy, set)
		res = append(res, append(setCopy, nums[0]))
	}
	return res
}
