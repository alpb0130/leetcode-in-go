package sort

import (
	"sort"
)

// Sort the array by value but keep the index after sorting. Iterate the sorted array
// and note down the visited smallest index. Every time we compute the diff between current
// orignal index and visited smallest index and update the max result and visited smallest index.
// Because the array has been sorted by value, we can guarantee that current value is larger than
// visited smallest index.
// Actually this problem is transferred to stock problem
// Time: O(nlogn) - sort
// Space: O(n)
func maxWidthRamp(A []int) int {
	if len(A) == 0 {
		return 0
	}
	max := 0
	// construct array
	array := make(Array, len(A))
	for i, a := range A {
		array[i] = ArrayValue{
			index: i,
			value: a,
		}
	}
	sort.Sort(array)
	arrayValueWithMinIndex := array[0]
	for _, element := range array {
		if element.index-arrayValueWithMinIndex.index > max {
			max = element.index - arrayValueWithMinIndex.index
		} else if element.index-arrayValueWithMinIndex.index < 0 {
			arrayValueWithMinIndex = element
		}
	}
	return max
}

type ArrayValue struct {
	index int
	value int
}

type Array []ArrayValue

func (a Array) Len() int { return len(a) }
func (a Array) Less(i, j int) bool {
	if a[i].value != a[j].value {
		return a[i].value < a[j].value
	}
	return a[i].index < a[j].index
}
func (a Array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
