package circulararray

type HitCounter struct {
	hits []int
	time []int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{
		hits: make([]int, 300),
		time: make([]int, 300),
	}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	if this.time[timestamp%300] == timestamp {
		this.hits[timestamp%300]++
	} else {
		this.time[timestamp%300] = timestamp
		this.hits[timestamp%300] = 1
	}
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	res := 0
	for i := 0; i < 300; i++ {
		// This is needed because some of element in hits array are never been updated and pretty old
		// because for some reason they are never overwrote.
		if this.time[i] > timestamp-300 {
			res += this.hits[i]
		}
	}
	return res
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
