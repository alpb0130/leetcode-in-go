package onepassandconstspace

import "math"

// https://leetcode.com/problems/candy/discuss/42770/One-pass-constant-space-Java-solution
// Time: O(n)
// Space: O(1)
func candy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}
	res := 1
	drop := 0
	prev := 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] >= ratings[i-1] {
			// process prev drop down
			if drop > 0 {
				res += countDropDown(drop)
				if drop >= prev {
					res += (drop - prev + 1)
				}
				drop = 0
				// reset prev to 1
				prev = 1
			}
			// if current rating is equal to prev rating, reset candy (prev) to 1
			if ratings[i] == ratings[i-1] {
				prev = 1
			} else {
				// if larger, just increase the prev candy number by 1
				prev++
			}
			res += prev
		} else {
			drop++
		}
	}
	if drop > 0 {
		res += countDropDown(drop)
		if drop >= prev {
			res += (drop - prev + 1)
		}
	}
	return res
}

func countDropDown(n int) int {
	return n * (n + 1) / 2
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
