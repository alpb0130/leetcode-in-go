package knuthshuffle

import (
	"math/rand"
	"time"
)

// Use knuth shuffle.
// Time: O(n)
// Space: O(n)
type Solution struct {
	nums []int
	r    *rand.Rand
}

func Constructor(nums []int) Solution {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return Solution{nums, r}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	clone := make([]int, len(this.nums))
	copy(clone, this.nums)
	for i := 0; i < len(clone); i++ {
		idx := this.r.Intn(i + 1)
		clone[i], clone[idx] = clone[idx], clone[i]
	}
	return clone
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
