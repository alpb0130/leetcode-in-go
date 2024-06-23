package _map

// Use map to note down visited number and index.
// Time: O(n)
// Space: O(n)
func twoSum(nums []int, target int) []int {
	numMap := map[int]int{}
	var res []int
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			res = []int{j, i}
			break
		}
		numMap[num] = i
	}
	return res
}
