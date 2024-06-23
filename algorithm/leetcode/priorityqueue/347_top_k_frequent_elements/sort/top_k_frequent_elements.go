package sort

import "sort"

// Get numbers' frequency and sort the frequency map
// Time: O(nlongx) - x is the number of unique numbers
// Space: O(x)
func topKFrequent(nums []int, k int) []int {
	results := []int{}
	if k == 0 {
		return results
	}
	cnts := map[int]int{}
	for _, num := range nums {
		cnts[num]++
	}
	c := counters{}
	for num, cnt := range cnts {
		c = append(c, counter{num, cnt})
	}
	sort.Sort(c)
	for i := 0; i < k; i++ {
		results = append(results, c[i].num)
	}
	return results
}

type counter struct {
	num int
	cnt int
}
type counters []counter

func (c counters) Len() int           { return len(c) }
func (c counters) Less(i, j int) bool { return c[i].cnt > c[j].cnt }
func (c counters) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
