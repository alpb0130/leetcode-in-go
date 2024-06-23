package constantmemo

import "math"

// The problem requires us to find whether there is increasing triplet where
// nums[i] < nums[j] < nums[k] given i < j < k. Our thought is go over the
// num list and update the num that ends the sequence with length 1 and 2 if
// we see a smaller one. We use two variables to record the current ending
// number of sequence 1 and 2 (end1, end2). end1 means currently we have
// sequence with length 1 ended by end1. End2 means currently we have a sequence
// with length 2 ended by end2. So we know end1 will always be smaller than end2.
// When we iterate the array, we have three cases:
// 1. num <= end1, means, end1 can be update. But it not necessarily means num can
//    combine with end2 and form a 2 sequence.
// 2. num <= end2, means, end2 can be update. We get a new and better end2 which
//    can help use get more potential results later.
// 3. num > end3, means, we get a increasing triplet
// Time: O(n)
// Space: O(1)
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	end1, end2 := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num <= end1 {
			end1 = num
		} else if num <= end2 {
			end2 = num
		} else {
			return true
		}
	}
	return false
}
