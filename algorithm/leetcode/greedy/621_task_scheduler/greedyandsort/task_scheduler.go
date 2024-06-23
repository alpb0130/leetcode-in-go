package greedyandsort

import "sort"

// Greedy method with sort
// Time: O(times) - sorting takes O(1) time because the max len of the list is 26...
// Space: O(1) - map with const size
func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}
	// get num of each tasks
	numOfTasks := map[byte]int{}
	for _, task := range tasks {
		numOfTasks[task] += 1
	}
	frequency := []int{}
	for _, num := range numOfTasks {
		frequency = append(frequency, num)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(frequency)))
	time := 0
	for {
		if frequency[0] == 0 {
			break
		}
		for i := 0; i < n+1; i++ {
			if frequency[0] == 0 {
				break
			}
			if i < len(frequency) && frequency[i] > 0 {
				frequency[i] -= 1
			}
			time++
		}
		sort.Sort(sort.Reverse(sort.IntSlice(frequency)))
	}
	return time
}
