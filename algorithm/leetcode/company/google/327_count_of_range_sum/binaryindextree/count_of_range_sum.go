package binaryindextree

import (
	"math"
	"sort"
)

// Use binary index tree (similar to problem 315).
// Firstly, we get the ranking of all needed sums (including all prefix sum and
// prefix - lower/upper). The reason we need prefix - lower/upper is we need to identify
// the index when we query the BIT. We can also find the closest prefix sum in the range of
// [prefix-upper, prefix-lower] but it's very complicate. Actually when we process the prefix
// sum array, the ranking index of [prefix-upper, prefix-lower] will mostly don't have a count.
// Then we process the prefix sum array. For every prefix sum, we try to query the total # of
// prefix sum in the range of [prefix-upper, prefix-lower]. Because we have ranking index, we
// will use ranking index to do the query.
// Last, add every prefix sum into the BIT
// Time: O(nlogn)
// Space: O(n)
func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}

	prefix := make([]int, len(nums)+1)
	sumMap := map[int]bool{}
	for i, num := range nums {
		prefix[i+1] = prefix[i] + num
		sumMap[prefix[i+1]] = true
		sumMap[prefix[i+1]-lower] = true
		sumMap[prefix[i+1]-upper] = true
	}
	sums := []int{0, 0 - lower, 0 - upper}
	for sum, _ := range sumMap {
		sums = append(sums, sum)
	}
	sort.Ints(sums)
	rankMap := map[int]int{}
	for i, sum := range sums {
		rankMap[sum] = i + 1
	}

	tree := &BIT{make([]int, len(sums)+1)}
	res := 0
	// for prefix[0], it's not contribute to the final result but we need to put it
	// into tree first.
	for i := 0; i < len(prefix); i++ {
		s := prefix[i]
		sRank := rankMap[s]
		h := prefix[i] - lower
		hRank := rankMap[h]
		l := prefix[i] - upper
		lRank := rankMap[l]
		// It's possible lRank-1 and BIT's 0 index will give us 0 although it's
		// a invalid index
		res += tree.query(hRank) - tree.query(lRank-1)
		tree.update(sRank, 1)
	}
	return res
}

type BIT struct {
	Vals []int
}

func (b *BIT) query(i int) int {
	res := 0
	for i > 0 {
		res += b.Vals[i]
		i -= (i & -i)
	}
	return res
}

func (b *BIT) update(i, delta int) {
	for i < len(b.Vals) {
		b.Vals[i] += delta
		i += (i & -i)
	}
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
