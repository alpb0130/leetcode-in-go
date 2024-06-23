package sort

import (
	"bytes"
	"sort"
	"strconv"
)

// Time: O(nlog) - quick sort
// Space: O(logn) - quick sort need O(logn) at best
func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	sort.Sort(Nums(nums))
	if nums[0] == 0 {
		return "0"
	}
	// bytes.Buffer is a more efficient way to concat string
	var result bytes.Buffer
	for _, num := range nums {
		result.WriteString(strconv.Itoa(num))
	}
	return result.String()
}

type Nums []int

func (n Nums) Len() int {
	return len(n)
}

func (n Nums) Less(i, j int) bool {
	strI := strconv.Itoa(n[i])
	strJ := strconv.Itoa(n[j])
	return strI+strJ > strJ+strI
}

func (n Nums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
