package priorityqueue

import (
	"container/heap"
	"math"
	"sort"
)

// Some key points:
// 1. If we hire a group of workers, we need to at lease pay the minimum wage to 1 worker and
//    extend the wage to other workers based on the wage.
// 2. If we pay a worker w1, and its quality is q1, then the total cost is c = w1*(q/q1). q is
//    total quality. So to minimize c, we need to minimize q
// 3. Another constraint is all money we pay should be larger than min wage of every worker.
//    If we pay w1 to a worker with q1, then based on ratio, for another worker with q2, we need
//    to pay w1/q1 = w2/q1 => w2 = w1*(q2/q1). w2 is > than w2_min, so w1*(q2/q1) > w2_min.
//    Then we get w1/q1 > w2_min/q2.
// Here we got our solution:
// 1. Sort the workers based on ratio (w/q) in ascending order. Then we know when we choose group,
//    we should always pay money based on the last worker in the group.
// 2. Maintain a group of K workers and every time we pop workers with max quality because we try
//    to minimize the total quality.
// Reference: https://leetcode.com/problems/minimum-cost-to-hire-k-workers/discuss/141768/Detailed-explanation-O(NlogN)
//            https://leetcode.com/problems/minimum-cost-to-hire-k-workers/discuss/177269/N-log-N-explanation-no-code
// Time: O(nlogn+nlogk)
// Space: O(n)
func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	if len(quality) == 0 || K == 0 || len(quality) < K {
		return 0
	}
	workers := []Worker{}
	for i := 0; i < len(quality); i++ {
		workers = append(workers, Worker{float64(wage[i]) / float64(quality[i]), float64(quality[i])})
	}
	sort.Slice(workers, func(i, j int) bool {
		return workers[i].Ratio < workers[j].Ratio
	})
	pq := &PQ{}
	sumQuality := 0.0
	res := math.MaxFloat64
	for _, worker := range workers {
		if pq.Len() == K {
			sumQuality -= heap.Pop(pq).(Worker).Quality
		}
		heap.Push(pq, worker)
		sumQuality += worker.Quality
		if pq.Len() == K {
			res = math.Min(res, sumQuality*worker.Ratio)
		}
	}
	return res
}

type Worker struct {
	Ratio   float64
	Quality float64
}

type PQ []Worker

func (p PQ) Len() int {
	return len(p)
}
func (p PQ) Less(i, j int) bool {
	return p[i].Quality > p[j].Quality
}
func (p PQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p *PQ) Pop() interface{} {
	x := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return x
}
func (p *PQ) Push(i interface{}) {
	*p = append(*p, i.(Worker))
}
