package radixsort2

// Radix sort. Sort number by the digit of each number. When largest number in the array is very
// small, the algorithm is almost linear.
// Reference1: https://en.wikipedia.org/wiki/Radix_sort
// Reference1: https://leetcode.com/problems/maximum-gap/solution/
// Time: O((n+b)*logb_k) - n is len(nums). b is the base of number (10 in this case). k
// is the largest number in the least. So logb_k is the max len of the number. If k is very
// large, this solution is not efficient.
// Space: O(n)
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	// maximum
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	// LSD radix sort
	base := 1
	for {
		if max/base == 0 {
			break
		}
		counter := make([]int, 10)
		for _, num := range nums {
			counter[(num/base)%10]++
		}
		for i := 1; i < len(counter); i++ {
			counter[i] = counter[i] + counter[i-1]
		}
		copyNums := make([]int, len(nums))
		for i := len(nums) - 1; i >= 0; i-- {
			counter[(nums[i]/base)%10]--
			copyNums[counter[(nums[i]/base)%10]] = nums[i]
		}
		for i, copyNum := range copyNums {
			nums[i] = copyNum
		}
		base *= 10
	}
	result := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > result {
			result = nums[i] - nums[i-1]
		}
	}
	return result
}
