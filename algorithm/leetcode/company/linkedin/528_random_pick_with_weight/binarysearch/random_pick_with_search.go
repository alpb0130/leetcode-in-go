package binarysearch

import (
	"math/rand"
	"time"
)

// For each index, compute the cumulative weight so that we know the range of every idx.
// Then conduct binary search to find the idx.
// Time: O(logn) - n is the length of numbers
// Space: O(n)
type Solution struct {
	weights []int
	r       *rand.Rand
	sum     int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return Solution{w, r, w[len(w)-1]}
}

func (this *Solution) PickIndex() int {
	num := this.r.Intn(this.sum)
	l := len(this.weights)
	left, right := 0, l-1
	for left < right {
		mid := (left + right) / 2
		if this.weights[mid] > num && (mid == 0 || this.weights[mid-1] <= num) {
			return mid
		} else if this.weights[mid] > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
