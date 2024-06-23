package twopass

// Two pass.
// First pass to count the number of every letter.
// Second pass to find the first letter with count 1.
// We can also do that in one pass:
// A map from letter to node. A double linked list which is in string order.
// If we find a duplicate letter, detach the node from the list but don't delete it
// from the map so that if third repeated letter comes, we know it's a repeated letter.
// In the last, we just return the head of double linked list.
// Time: O(n)
// Space: O(1)
func firstUniqChar(s string) int {
	counter := map[byte]int{}
	for i := range s {
		counter[s[i]]++
	}
	for i := range s {
		if counter[s[i]] == 1 {
			return i
		}
	}
	return -1
}
