package circulararray

// Use circular array. Using an array with size 3. Use a pointer to always point to
// the index we last update. Then if a new number come, we need to remove the very first
// number from sum and add current num to the sum. Where is the very first number?
// It's in the position ptr+1%size which should be on position behind the last number.
// The new number should also be put here. Remember a case where we have less than 3 number.
// Because the very first position is 0 so that would not causing too much trouble for use.
// Time: O(1)
// Space: O(1)
type MovingAverage struct {
	nums  []int
	count int
	ptr   int
	size  int
	sum   int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{make([]int, size), 0, 2, size, 0}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.count < this.size {
		this.count++
	}
	this.ptr = (this.ptr + 1) % this.size
	last := this.nums[this.ptr]
	this.sum -= last
	this.sum += val
	this.nums[this.ptr] = val
	return float64(this.sum) / float64(this.count)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
