package binarysearch

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		pings: []int{},
	}
}

// Time: O(1) - if no call can happen at the same time /  O(logn) - if
// calls can happen at the same time. n is the # of calls
// Space: O(1) if no call can happen at the same time
func (this *RecentCounter) Ping(t int) int {
	this.pings = append(this.pings, t)
	index := binarySearch(this.pings, t-3000)
	this.pings = this.pings[index:]
	return len(this.pings)
}

// find the first element which is larger than or equal to the given value. If no such value, return len()
func binarySearch(array []int, value int) int {
	start, end := 0, len(array)-1
	for {
		if start > end {
			break
		}
		mid := start + (end-start)/2
		if array[mid] < value {
			start = mid + 1
		} else if mid > start && array[mid-1] >= value {
			end = mid - 1
		} else {
			return mid
		}

	}
	return len(array)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
