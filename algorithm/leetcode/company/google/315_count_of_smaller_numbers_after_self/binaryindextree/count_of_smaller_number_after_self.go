package binaryindextree

import "sort"

// Use binary index tree. Convert every number to its rank (same number has same rank).
// When we process every number, we just need to get # of numbers that with a lower ranker
// than current number so far.
// Useful link: https://www.youtube.com/watch?v=WbafSgetDDk
//              https://www.youtube.com/watch?v=2SVLYsq5W8M
// Time: O(nlogn)
// Space: O(n)
func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	// Build rank map
	uniqueNumMap := map[int]bool{}
	for _, num := range nums {
		uniqueNumMap[num] = true
	}
	uniqueNums := make([]int, len(uniqueNumMap))
	idx := 0
	for num := range uniqueNumMap {
		uniqueNums[idx] = num
		idx++
	}
	sort.Ints(uniqueNums)
	rankMap := map[int]int{}
	for i, num := range uniqueNums {
		rankMap[num] = i + 1
	}

	// from back to front to process every number
	// and get the total # of nums with lower rank at current time (Some number may not
	// be processed so it should contribute to the results). Then this problem is converted
	// to a range sum problem
	tree := &BIT{make([]int, len(rankMap)+1)}
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = tree.query(rankMap[nums[i]] - 1)
		tree.update(rankMap[nums[i]], 1)
	}
	return res
}

type BIT struct {
	Vals []int
}

func getLowBit(a int) int {
	return a & (-a)
}

func (b *BIT) update(i int, delta int) {
	for i < len(b.Vals) {
		b.Vals[i] += delta
		i += getLowBit(i)
	}
}

func (b *BIT) query(i int) int {
	res := 0
	for i > 0 {
		res += b.Vals[i]
		i -= getLowBit(i)
	}
	return res
}
