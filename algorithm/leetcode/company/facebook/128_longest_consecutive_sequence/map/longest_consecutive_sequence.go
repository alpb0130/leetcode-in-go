package _map

// Use a map to record all nums. Then we iterate over the map.
// Every time we see a number without preceding, we start with
// this number and check the ascending sequence till a number is
// not in the map. Recording the max length as well.
// Time: O(n+n)
// Space: O(n)
func longestConsecutive(nums []int) int {
	numMap := map[int]bool{}
	for _, num := range nums {
		numMap[num] = true
	}
	res := 0
	for num := range numMap {
		if !numMap[num-1] {
			longest := 0
			for numMap[num] {
				longest++
				num++
			}
			if longest > res {
				res = longest
			}
		}
	}
	return res
}
