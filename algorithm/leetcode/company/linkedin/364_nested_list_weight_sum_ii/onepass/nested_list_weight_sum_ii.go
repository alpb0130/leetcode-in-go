package onepass

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
// One pass. Use a var to store unweighted sum and a var to store weighted sum.
// For each level, we add the unweighted sum again to the weighted sum which is same to apply
// weight at each level.
// Reference: https://leetcode.com/problems/nested-list-weight-sum-ii/discuss/83641/No-depth-variable-no-multiplication
// Time: O(n) - n is the total number of integers in the nested integer
// Space: O(n)
func depthSumInverse(nestedList []*models.NestedInteger) int {
	weighted := 0
	unweighted := 0
	for len(nestedList) > 0 {
		nextLevelList := []*models.NestedInteger{}
		for _, nested := range nestedList {
			if nested.IsInteger() {
				unweighted += nested.GetInteger()
			} else {
				nextLevelList = append(nextLevelList, nested.GetList()...)
			}
		}
		weighted += unweighted
		nestedList = nextLevelList
	}
	return weighted
}
