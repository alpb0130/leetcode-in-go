package queue

type HitCounter struct {
	hits  *queue
	count int
}

// Use a queue to always maintain the hits in 5 mins. and its size should be 300
// We can also use circular array
/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{
		hits:  &queue{},
		count: 0,
	}
}

// When we call Hit, we firstly update the queue and remove out-of-date hits.
// Then we add the new hits to existing hit/create a new one
/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	// remove all out of date hits (we can also move this part to GetHits function)
	this.update(timestamp)
	// update hits or add new hits
	if this.hits.Len() > 0 && this.hits.Last().timestamp == timestamp {
		this.hits.Last().count++
	} else {
		this.hits.Offer(&HitsInSec{timestamp, 1})
	}
	this.count++
}

// Every time when we need to get hits, we update the queue first and remove out-of-date hits
// Then return counts.
/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	// Don't forget this!!!!
	this.update(timestamp)
	return this.count
}

func (this *HitCounter) update(timestamp int) {
	for this.hits.Len() > 0 && this.hits.First().timestamp <= timestamp-300 {
		first := this.hits.Poll()
		this.count -= first.count
	}
}

type HitsInSec struct {
	timestamp int
	count     int
}

type queue []*HitsInSec

func (q *queue) Len() int {
	return len(*q)
}
func (q *queue) Poll() *HitsInSec {
	h := (*q)[0]
	*q = (*q)[1:]
	return h
}
func (q *queue) Offer(h *HitsInSec) {
	*q = append(*q, h)
}
func (q *queue) First() *HitsInSec {
	return (*q)[0]
}
func (q *queue) Last() *HitsInSec {
	return (*q)[len(*q)-1]
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
