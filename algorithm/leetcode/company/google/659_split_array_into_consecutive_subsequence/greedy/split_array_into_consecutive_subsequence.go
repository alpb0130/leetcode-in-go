package greedy

// Use Greedy. Given a sorted array. We try to construct the consecutive sequence
// with minimum nums to use but make sure the len of sequence is larger than 3.
// Time: O(n)
// Space: O(n)
func isPossible(nums []int) bool {
	// Greedy
	// 1. If we can append the number to existing consecutive subsequences, do that first
	// 2. If we cannot append to any exsiting sequence, we need to start a new one but make
	//    sure you secure the next two numbers in the new sequences (need a map to do that in
	//     O(1))
	// 3. If both of those are not work, return false
	// Two place to show our greedy thoughts.
	// 1. We want to append the number to existing sequence first. If not possible, start a new
	//    sequence. We want to do that because for the case 1, 2, 3, 4, 5, when we process 4
	//    we can satisfy the problem by adding 4 only without 5 and 6.
	// 2. When we need to create a new sequence start num, we want to put num+1 and num+2 into
	//    the sequence at that time to make sure the new sequence we create meet our expectation.
	numsCount := map[int]int{}
	for _, num := range nums {
		numsCount[num]++
	}
	lastNumsCount := map[int]int{}
	for _, num := range nums {
		if numsCount[num] == 0 {
			continue
		}
		if lastNumsCount[num-1] > 0 {
			lastNumsCount[num-1]--
			lastNumsCount[num]++
			numsCount[num]--
		} else if numsCount[num+1] > 0 && numsCount[num+2] > 0 {
			lastNumsCount[num+2]++
			numsCount[num]--
			numsCount[num+1]--
			numsCount[num+2]--
		} else {
			return false
		}
	}
	return true
}
