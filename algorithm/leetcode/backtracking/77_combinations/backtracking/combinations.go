package backtracking

// Solve this problem recursively. F(n, k) = F(n-1, k-1) + F(n-2, k-1) + ... + F(1, k-1)
// Reference:https://leetcode.com/problems/combinations/discuss/27002/Backtracking-Solution-Java
// Time: O(n*C(n, k)) = O(n* n!/(k!*(n-k)!))
// Space: O(n*k)
func combine(n int, k int) [][]int {
	if k == 0 || k > n {
		return [][]int{}
	}
	res := [][]int{}
	combineHelper(n, k, 1, []int{}, &res)
	return res
}

func combineHelper(n, k, start int, list []int, res *[][]int) {
	if len(list) == k {
		des := make([]int, len(list))
		copy(des, list)
		*res = append(*res, des)
		return
	}
	if start > n {
		return
	}
	for i := start; i <= n; i++ {
		combineHelper(n, k, i+1, append(list, i), res)
	}
}
