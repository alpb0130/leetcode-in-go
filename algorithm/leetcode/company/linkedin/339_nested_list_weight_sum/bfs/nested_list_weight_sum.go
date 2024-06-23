package bfs

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
// Iterative
// Time: O(n) - n is the total number of integers in the nested integer
// Space: O(d) - max number of element of two contiguous level
func depthSum(nestedList []*models.NestedInteger) int {
	q := &queue{}
	for _, nested := range nestedList {
		q.Offer(nested)
	}
	level := 1
	sum := 0
	for q.Len() != 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			nested := q.Poll()
			if nested.IsInteger() {
				sum += level * nested.GetInteger()
			} else {
				for _, n := range nested.GetList() {
					q.Offer(n)
				}
			}
		}
		level++
	}
	return sum
}

type queue []*models.NestedInteger

func (q *queue) Poll() *models.NestedInteger {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
func (q *queue) Offer(n *models.NestedInteger) {
	*q = append(*q, n)
}
func (q *queue) Len() int {
	return len(*q)
}
