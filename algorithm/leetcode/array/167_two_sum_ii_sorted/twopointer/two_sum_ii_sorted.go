package twopointer

// Two pointer.
// Time: O(n)
// Space: O(1)
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	var res []int
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			res = []int{i + 1, j + 1}
			break
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return res
}
