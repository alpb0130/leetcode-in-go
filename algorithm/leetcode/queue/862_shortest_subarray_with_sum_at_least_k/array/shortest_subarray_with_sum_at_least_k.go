package array

// Time: O(n) - amortized time is linear because each element is put and pull from queue one time
// Space: O(n)
func shortestSubarray(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}
	// get prefix sum
	B := make([]int, len(A)+1)
	for i, a := range A {
		B[i+1] = B[i] + a

	}
	result := len(A) + 1
	deque := []int{}
	// something like a double for loop but improved a lot (O(n))
	// Iterate over B (fix the upper bound of subarray) and check the lower bound
	// But when we check the lower bound, we check from deque instead of B because deque
	// has improvements and remove some unnecessary elements
	for j, b := range B {
		// iterate deque and check whether there is a validate index i which
		// B[j] - B[i] >= K
		// Because in deque, the B[deque[i]] is increasing. If we get the first i having
		// b - B[deque[i]] < K, all the index after i will have this and no need to check
		i := 0
		for {
			if i == len(deque) || b-B[deque[i]] < K {
				break
			}
			if j-deque[i] < result {
				result = j - deque[i]
			}
			deque = deque[i+1:]
			i = 0
		}
		// put B[j] into deque
		// find the first index in deque having B[deque[i]] >= b
		// remove all of them and append the new value to the deque
		// This keeps the array in ascent order and B[deque[i]]
		// is always the smallest value in deque
		i = 0
		for {
			if i == len(deque) || B[deque[i]] >= b {
				break
			}
			i++
		}
		deque = append(deque[:i], j)
	}
	if result == len(A)+1 {
		return -1
	}
	return result
}
