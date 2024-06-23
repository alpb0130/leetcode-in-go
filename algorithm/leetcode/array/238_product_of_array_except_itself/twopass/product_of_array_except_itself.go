package twopass

// Calculate prefix prod first from start to end and put into result value.
// Then calculate suffix prod from end to start and put into result value.
// Time: O(n)
// Space: O(1)
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	prefixProd := 1
	for i := 0; i < len(nums); i++ {
		res[i] = prefixProd
		prefixProd *= nums[i]
	}
	suffixProd := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= suffixProd
		suffixProd *= nums[i]
	}
	return res
}
