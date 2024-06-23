package dp

import "math"

// The brute force solution is we can dfs all possible increasing subsequence and get the longest
// one. But with this method, we do a lot of repeated search. For a list of [n1,...,nk]. And for
// every element, we can choose either skip it or use it. For the choices we skip n1 and use n1,
// we could both search over the space [n2:nk] which waste a lot of time. One improvement we can
// do is use DP and store the solutions for all sub problem.
// The problem P(n) = min(G(n), G(n-1),...,G(1)) where G(k) represent the longest subsequence that
// is ended with k. For each problem G(k), we can get it by go over G(1), G(2),...,G(k-1) and check
// whether k > k'. If it is, we can update the result for G(k) by take max(G(k), G(k')+1).
// This algorithm takes O(n^2) time.
// Time: O(n^2)
// Space: O(n)
func lengthOfLIS(nums []int) int {
	length := make([]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		length[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				length[i] = maxInt(length[i], length[j]+1)
			}
		}
		res = maxInt(res, length[i])
	}
	return res
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
