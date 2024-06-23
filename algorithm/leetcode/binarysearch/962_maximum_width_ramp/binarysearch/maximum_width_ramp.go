package binarysearch

// binary search
// Time: O(nlogn)
// Space: O(n)
func maxWidthRamp(A []int) int {
	width := 0
	// keep elements in desc order because any large value with larger index is unnecessary to put
	// into the queue and take into consideration.
	// If i < j and A[i] <= A[j], then for any k, we must have A[k] <= A[i], A[k] <= A[j]. But if
	// A[k] > A[i], k-i must be a better answer than k-j
	queue := []int{}
	for i, a := range A {
		// improve by binary search
		j := binarySearch(A, queue, a)
		if j == len(queue) {
			queue = append(queue, i)
		} else if width < i-queue[j] {
			width = i - queue[j]
		}
	}
	return width
}

// find the index of first element which is smaller than or equal to value
func binarySearch(A, queue []int, value int) int {
	first, last := 0, len(queue)-1
	for {
		if first > last {
			break
		}
		mid := first + (last-first)/2
		if A[queue[mid]] > value {
			first = mid + 1
		} else if mid-1 >= first && A[queue[mid-1]] <= value {
			last = mid - 1
		} else {
			return mid
		}
	}
	return len(queue)
}
