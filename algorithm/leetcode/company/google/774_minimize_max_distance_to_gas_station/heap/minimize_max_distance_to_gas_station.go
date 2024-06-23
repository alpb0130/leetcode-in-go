package heap

import "container/heap"

// We keep adding 1 station into current max interval. Current max interval is
// max original interval length/# of station added. This is a greedy method and
// we know the minimum interval happens when we evenly distributed stations in
// the interval.
// Time: O(k*logn + n*logn) - n is # of stations, k is K
// Space: O(n)
func minmaxGasDist(stations []int, K int) float64 {
	pq := &PQ{}
	for i := 1; i < len(stations); i++ {
		dist := stations[i] - stations[i-1]
		heap.Push(pq, StationInterval{float64(dist), 0.0})
	}
	for i := 0; i < K; i++ {
		maxInterval := heap.Pop(pq).(StationInterval)
		maxInterval.Station++
		heap.Push(pq, maxInterval)
	}
	return heap.Pop(pq).(StationInterval).getInterval()
}

type StationInterval struct {
	Length  float64
	Station float64
}

func (i StationInterval) getInterval() float64 {
	return i.Length / (i.Station + 1)
}

type PQ []StationInterval

func (p PQ) Len() int {
	return len(p)
}
func (p PQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PQ) Less(i, j int) bool {
	return p[i].getInterval() > p[j].getInterval()
}
func (p *PQ) Pop() interface{} {
	i := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return i
}
func (p *PQ) Push(i interface{}) {
	*p = append(*p, i.(StationInterval))
}
