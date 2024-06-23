package smalltolarge

import "sort"

// Start with smallest height. If there should be k people on the left side of smallest height.
// Then we know the smallest height must be on position k (start with 0). Then we process the next
// smallest height, if we know there are l people on the left side, we should put it to the (l+1)th
// empty space. What if there are two people with same height, we process people on the right first
// (with larger second number) because we know all process number must be smaller than it or after
// it.
// A O(nlogn) solution: https://leetcode.com/problems/queue-reconstruction-by-height/discuss/143403/O(nlogn)-modified-merge-sort-solution-with-detailed-explanation
// Time: O(n^2)
// Space: O(n)
func reconstructQueue(people [][]int) [][]int {
	if len(people) == 0 {
		return people
	}
	sort.Sort(People(people))
	res := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		requiredEmpty := people[i][1]
		empty := -1
		j := 0
		for j < len(res) {
			if len(res[j]) == 0 {
				empty++
			}
			if empty == requiredEmpty {
				break
			}
			j++
		}
		res[j] = people[i]
	}
	return res
}

type People [][]int

func (p People) Len() int {
	return len(p)
}
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p People) Less(i, j int) bool {
	if p[i][0] == p[j][0] {
		return p[i][1] > p[j][1]
	}
	return p[i][0] < p[j][0]
}
