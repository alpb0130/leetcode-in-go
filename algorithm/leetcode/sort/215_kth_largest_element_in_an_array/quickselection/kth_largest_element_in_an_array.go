package quickselection

import "math/rand"

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	// shuffle
	for i := 0; i < len(nums); i++ {
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}

	low, high := 0, len(nums)-1
	// quick selection
	for {
		if low > high {
			break
		}
		i := partition(nums, low, high)
		if i < k {
			low = i + 1
		} else if i > k {
			high = i - 1
		} else {
			break
		}
	}
	return nums[k]
}

func partition(nums []int, low, high int) int {
	i := low
	j, k := low+1, high
	for {
		for j <= high && nums[j] <= nums[i] {
			j++
		}
		for k > low && nums[k] >= nums[i] {
			k--
		}
		if j > k {
			nums[i], nums[k] = nums[k], nums[i]
			break
		}
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
	return k
}
