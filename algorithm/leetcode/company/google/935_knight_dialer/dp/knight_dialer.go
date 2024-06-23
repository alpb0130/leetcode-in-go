package dp

// F(N) can be derived from F(N-1). We already know, given current position, what will the
// next hop position be. So we can easily know from last hop, how many time every number will
// appear at this hop. Then the final result is sum of count for all numbers.
// Time: O(n)
// Space: O(1)
func knightDialer(N int) int {
	if N == 0 {
		return 0
	}
	next := [][]int{
		{4, 6},
		{6, 8},
		{7, 9},
		{4, 8},
		{3, 9, 0},
		{},
		{1, 7, 0},
		{2, 6},
		{1, 3},
		{2, 4},
	}
	numCount := make([]int, 10)
	for i := range numCount {
		numCount[i] = 1
	}
	for i := 0; i < N-1; i++ {
		tmp := make([]int, 10)
		for i, cnt := range numCount {
			for _, next := range next[i] {
				tmp[next] = (tmp[next] + cnt) % 1000000007
			}
		}
		numCount = tmp
	}
	res := 0
	for _, cnt := range numCount {
		res = (res + cnt) % 1000000007
	}
	return res
}
