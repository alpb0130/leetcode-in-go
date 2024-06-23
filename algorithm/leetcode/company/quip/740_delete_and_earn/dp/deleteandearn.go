package dp

import "math"

// DP. We are given a list of positive numbers with repeat. We need to find the largest earn by the
// method that if we earn a number, we need to delete the adjacent number and cannot earn it any
// more. Think about paint house and they are similar. Because we earn a number, we can earn all
// other numbers that are equal to it. So first we convert the input array to a map which is
// num -> sum (num*cnt). Then we can start from the first number and keep two values. One is
// for skip the current number and they other is keep the current number.
//
// Time: O(n)
// Space: O(n)
func deleteAndEarn(nums []int) int {
	sum := make([]int, 10002)
	for _, num := range nums {
		sum[num] += num
	}
	skip := 0
	keep := 0
	for i := 1; i < len(sum); i++ {
		newSkip := maxInt(skip, keep)
		newKeep := skip + sum[i]
		skip = newSkip
		keep = newKeep
	}
	return maxInt(skip, keep)
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
