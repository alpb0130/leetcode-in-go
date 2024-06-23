package better

// Time: O(2^n)
// Space: O(n^2)
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	res := [][]int{}
	subSetHelper(nums, 0, []int{}, &res)
	return res
}

func subSetHelper(nums []int, start int, tmp []int, res *[][]int) {
	// We don't need base case because the for loop below will take care of that
	dst := make([]int, len(tmp))
	copy(dst, tmp)
	// add result from callee
	*res = append(*res, dst)
	// Our idea is for every element, either we put it in the result, or we drop it.
	// The for loop below just do that, for every elements, we either put it in the tmp result
	// and do backtracking call (all callee will must has this element in the set). Or we
	// skip it and put the next element in the sets.
	for i := start; i < len(nums); i++ {
		subSetHelper(nums, i+1, append(tmp, nums[i]), res)
	}
}
