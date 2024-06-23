package priorityqueue

import "container/heap"

// The straightforward solution is get all pairs fo those two array and then keep a min heap
// with size k to get a relatively good sorting cost. But this method is not computing all
// pairs before we sorting. Instead, we only compare get new pair when we need:
// We are given two sorted array. Start with nums1, get all pairs from all element in nums1 and
// combined with the first element in nums2. The smallest element must be in those pairs. Then we
// can start the process the pairs. Get the current min pair and then the next possible smallest
// pair must be pair(nums1[popped] + nums2[popped + 1]). In this way, we can always keep the length
// of heap to len(nums1).
// For reference: https://leetcode.com/problems/find-k-pairs-with-smallest-sums/discuss/84551/simple-Java-O(KlogK)-solution-with-explanation
// Time: O(m + klogm) - m is the length of nums1
// Space: O(m)
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &pairs{}
	heap.Init(h)
	results := [][]int{}
	if len(nums1) == 0 || len(nums2) == 0 {
		return results
	}
	for _, n := range nums1 {
		heap.Push(h, pair{n, nums2[0], 0})
	}
	for {
		if len(results) == k || h.Len() == 0 {
			break
		}
		min := heap.Pop(h).(pair)
		results = append(results, []int{min.num1, min.num2})
		newIndex := min.index2 + 1
		if newIndex < len(nums2) {
			heap.Push(h, pair{min.num1, nums2[newIndex], newIndex})
		}
	}
	return results
}

type pair struct {
	num1   int
	num2   int
	index2 int
}

type pairs []pair

func (p pairs) Len() int           { return len(p) }
func (p pairs) Less(i, j int) bool { return p[i].num1+p[i].num2 < p[j].num1+p[j].num2 }
func (p pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *pairs) Pop() interface{} {
	x := (*p)[len(*p)-1]
	*p = (*p)[0 : len(*p)-1]
	return x
}
func (p *pairs) Push(i interface{}) {
	*p = append(*p, i.(pair))
}
