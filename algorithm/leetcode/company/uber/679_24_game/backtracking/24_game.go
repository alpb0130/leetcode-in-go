package backtracking

import "math"

// Search every number permutations and update the number list. We totally have
// 12*4*6*4*2*4 permutations commutative.
// Time: O(1)
// Space: O(1)
func judgePoint24(nums []int) bool {
	list := make([]float64, len(nums))
	for i, num := range nums {
		list[i] = float64(num)
	}
	// dfs
	return dfs(list)
}

func dfs(nums []float64) bool {
	if len(nums) == 0 {
		return false
	}
	// This is to handle some number underflow
	if len(nums) == 1 && math.Abs(nums[0]-24.0) < 0.001 {
		// if len(nums) == 1 && nums[0] == 24.0 {
		return true
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				newNums := []float64{}
				for k := 0; k < len(nums); k++ {
					if k != i && k != j {
						newNums = append(newNums, nums[k])
					}
				}
				for l := 0; l < 4; l++ {
					// if it's + or *, we can skip because the operation is
					if l < 2 && i > j {
						continue
					}
					if l == 0 {
						newNums = append(newNums, nums[i]+nums[j])
					}
					if l == 1 {
						newNums = append(newNums, nums[i]*nums[j])
					}
					if l == 2 {
						newNums = append(newNums, nums[i]-nums[j])
					}
					if l == 3 {
						if nums[j] == 0 {
							continue
						}
						newNums = append(newNums, nums[i]/nums[j])
					}
					if dfs(newNums) {
						return true
					}
					newNums = newNums[:len(newNums)-1]
				}
			}
		}
	}
	return false
}
