package dfs

// Time: O(k*2^n) - we need to search the array k times. For each time, we need to determine whether
// 	     the element is in the subarray which is 2^n.
// Space: O(n) - n is the length of nums. We will do search to at most n height and isVisited
//        also need n space.
func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if k == 0 || sum%k != 0 {
		return false
	}
	subSum := sum / k
	isVisited := make([]bool, len(nums))
	// bfs to check
	return search(nums, isVisited, 0, k, 0, subSum)
}

func search(nums []int, isVisited []bool, start int, k int, curSum int, target int) bool {
	if k == 0 {
		return true
	}
	if curSum == target {
		return search(nums, isVisited, 0, k-1, 0, target)
	}
	for i := start; i < len(nums); i++ {
		if !isVisited[i] && curSum+nums[i] <= target {
			curSum += nums[i]
			isVisited[i] = true
			if search(nums, isVisited, i+1, k, curSum, target) {
				return true
			}
			isVisited[i] = false
			curSum -= nums[i]
		}
	}
	return false
}
