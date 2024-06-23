package dp

import "math"

// The brute force method is for each start position, we simulate the jump process. For odd jump,
// we try to find the element which is the smallest numbers that are larger thank start num. Then
// we can jump to this position. Keep this process until we can jump the last position. For evey
// start position, the jump process is O(n^2) - O(n) to find next position to jump and O(n) for all
// jumps. But as we process this problem, we find many problem are processed multiple times. Like
// when we decide whether we can jump from i, j to last. We probably consider k to last twice where
// k > i and j. And whether we can jump from k to last is not related to i or j. We can process the
// array from back to front and record the odd/even jump status for every process position.
// So we got a DP solution: process from back to front and for every start position, we only use
// O(n) time to decide the next jump position. Whether we can jump from the next position to end
// is already recode in array and we can know that in O(1) time. So the total time is O(n^2).
// Keep thinking about the process, find the next jump position eat us a lot of time. We can optimize
// that by using treemap(red-black tree) in java and get the next element in O(logn)
// Time: O(n^2)
// Space: O(n)
func oddEvenJumps(A []int) int {
	if len(A) == 0 {
		return 0
	}
	res := 1
	odd := make([]bool, len(A))
	even := make([]bool, len(A))
	odd[len(A)-1] = true
	even[len(A)-1] = true
	for i := len(A) - 2; i >= 0; i-- {
		// odd jump and contribute to final results
		min := math.MaxInt32
		max := math.MinInt32
		indexOdd := i
		indexEven := i
		for j := i + 1; j < len(A); j++ {
			if A[j] >= A[i] && A[j] < min {
				min = A[j]
				indexOdd = j
			}
			if A[j] <= A[i] && A[j] > max {
				max = A[j]
				indexEven = j
			}
		}
		if indexOdd != i && even[indexOdd] {
			odd[i] = true
			res++
		}
		if indexEven != i {
			even[i] = odd[indexEven]
		}
	}
	return res
}
