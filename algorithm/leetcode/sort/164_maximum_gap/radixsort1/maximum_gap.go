package radixsort1

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
		counter := make([][]int, 10)
		for _, num := range nums {
			counter[(num/base)%10] = append(counter[(num/base)%10], num)
		}
		// queue version of counting sort
		i := 0
		for _, list := range counter {
			for _, num := range list {
				nums[i] = num
				i++
			}
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
