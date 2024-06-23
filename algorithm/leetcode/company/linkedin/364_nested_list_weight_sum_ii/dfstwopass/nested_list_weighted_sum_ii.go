package dfstwopass

import (
	"math"

	"algorithm/leetcode/company/linkedin/models"
)

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
// Apply DFS twice. The first pass is to get the height of the tree. The second pass
// is to get the weighted sum.
// Be careful that not all leaves has the same weight and only the lowest leaves has
// weight 1.
// Time: O(n) - n is the total number of integers in the nested integer
// Space: O(d) - d is the depth of the nested integer
func depthSumInverse(nestedList []*models.NestedInteger) int {
	level := getHeight(nestedList)
	return getSum(nestedList, level)
}

func getHeight(nestedList []*models.NestedInteger) int {
	if len(nestedList) == 0 {
		return 0
	}
	level := 1
	for _, nested := range nestedList {
		if !nested.IsInteger() {
			subLevel := getHeight(nested.GetList())
			level = maxInt(subLevel+1, level)
		}
	}
	return level
}

func getSum(nestedList []*models.NestedInteger, level int) int {
	if len(nestedList) == 0 {
		return 0
	}
	sum := 0
	for _, nested := range nestedList {
		if nested.IsInteger() {
			sum += level * nested.GetInteger()
		} else {
			sum += getSum(nested.GetList(), level-1)
		}
	}
	return sum
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
