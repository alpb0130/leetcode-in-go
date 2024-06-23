package twopass

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
// The problems give us a function "knows" to check whether a knows b and we are required to find
// only celebrity (every knows him but he knows nobody). A brute force problem is n^2 method but we
// can do O(n).
// First we can assume 0 is the celebrity, then check whether 0 knows other n-1 people. If knows k,
// then 0 wouldn't be celebrity and k could be. We keep go over the rest of the people. If k could
// be candidate, he must not know anybody. Then we need to recheck the candidate with other n-1
// people to see whether he is a celebrity.
// From a graph perspective, if we start any node, we can always be trapped in the potential
// celebrity if we follow the know relationship.
// Time: O(n)
// Space: O(1)
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		candidate := 0
		for i := 1; i < n; i++ {
			if knows(candidate, i) {
				candidate = i
			}
		}
		for i := 0; i < n; i++ {
			if i != candidate && (!knows(i, candidate) || knows(candidate, i)) {
				return -1
			}
		}
		return candidate
	}
}
