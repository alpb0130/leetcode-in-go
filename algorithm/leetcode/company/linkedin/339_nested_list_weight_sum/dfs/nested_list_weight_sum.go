package dfs

import "algorithm/leetcode/company/linkedin/models"

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
// Recursive
// Time: O(n) - n is the total number of integers in the nested integer
// Space: O(d) - d is the depth of the nested integer
func depthSum(nestedList []*models.NestedInteger) int {
	return depthSumHelper(nestedList, 1)
}

func depthSumHelper(nestedList []*models.NestedInteger, depth int) int {
	if len(nestedList) == 0 {
		return 0
	}
	sum := 0
	for _, nested := range nestedList {
		if nested.IsInteger() {
			sum += depth * nested.GetInteger()
		} else {
			sum += depthSumHelper(nested.GetList(), depth+1)
		}
	}
	return sum
}
