package queue

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		pings: []int{},
	}
}

// Time: O(1) - if no call can happen at the same time /  O(n) - if
// calls can happen at the same time. n is the # of calls
// Space: O(1) if no call can happen at the same time
func (this *RecentCounter) Ping(t int) int {
	this.pings = append(this.pings, t)
	for i, ping := range this.pings {
		if t-ping <= 3000 {
			this.pings = this.pings[i:]
			break
		}
	}
	return len(this.pings)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
