package customizeddeque

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
	d := deque{}
	// something like a double for loop but improved a lot (O(n))
	// Iterate over B (fix the upper bound of subarray) and check the lower bound
	// But when we check the lower bound, we check from deque instead of B because deque
	// has improvements and remove some unnecessary elements
	for j, b := range B {
		// iterate deque and check whether there is a validate index i which
		// B[j] - B[i] >= K
		// Because in deque, the B[deque[i]] is increasing. If we get the first i having
		// b - B[deque[i]] < K, all the index after i will have this and no need to check
		var first int
		for {
			if d.isEmpty() || b-B[d[0]] < K {
				break
			}
			d, first = d.pollFirst()
			if j-first < result {
				result = j - first
			}
		}
		// put B[j] into deque
		// find the first index in deque having B[deque[i]] >= b
		// remove all of them and append the new value to the deque
		// This keeps the array in ascent order and B[deque[i]]
		// is always the smallest value in deque
		for {
			if d.isEmpty() || b > B[d[len(d)-1]] {
				break
			}
			d, _ = d.pollLast()
		}
		d = d.offerLast(j)
	}
	if result == len(A)+1 {
		return -1
	}
	return result
}

type deque []int

func (d deque) pollFirst() (deque, int) {
	return d[1:], d[0]
}

func (d deque) offerFirst(value int) deque {
	return append(deque{value}, d...)
}

func (d deque) pollLast() (deque, int) {
	l := len(d)
	return d[:l-1], d[l-1]
}

func (d deque) offerLast(value int) deque {
	return append(d, value)
}

func (d deque) isEmpty() bool {
	return len(d) == 0
}
