package _map

import (
	"math/rand"
	"time"
)

type Solution struct {
	newMap map[int]int
	n      int
	r      *rand.Rand
}

// Create a new map to map blacklist element to non-blacklist element. We need to generate a random
// number less than N-len(blacklist), but some of the number might be in the blacklist. If there
// are x elements in blacklist which are less than N-len(blacklist), there must be x elements in
// [N-len(blacklist), N) which are not in blacklist. We need to create the remapping.
// For reference: https://leetcode.com/problems/random-pick-with-blacklist/discuss/144624/Java-O(B)-O(1)-HashMap
// Time: O(len(blacklist))
// Space: O(len(blacklist))
func Constructor(N int, blacklist []int) Solution {
	n := N - len(blacklist)
	// blacklistMap
	blacklistMap := map[int]bool{}
	for _, blacklistElement := range blacklist {
		blacklistMap[blacklistElement] = true
	}
	newMap := map[int]int{}
	counter := N - 1
	// construct remapping
	for _, blacklistElement := range blacklist {
		if blacklistElement < n {
			for blacklistMap[counter] {
				counter--
			}
			newMap[blacklistElement] = counter
			counter--
		}
	}
	return Solution{
		newMap: newMap,
		n:      n,
		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Time: O(1)
// Space: O(1)
func (this *Solution) Pick() int {
	num := this.r.Intn(this.n)
	if newNum, ok := this.newMap[num]; ok {
		return newNum
	}
	return num
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
