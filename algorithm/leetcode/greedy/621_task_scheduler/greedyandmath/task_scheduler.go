package greedyandmath

import "math"


// Greedy method and math - the best solution is always arranging most frequent task first and make
// sure they have n intervals, then fill intervals with other less frequent tasks. Two special
// cases: 1. more than 1 most frequent tasks. 2. The num of available idle intervals is less than
// that of non-most-frequent tasks. For more detailed explanation, please refer to
// https://leetcode.com/problems/task-scheduler/discuss/104500/Java-O(n)-time-O(1)-space-1-pass-no-sorting-solution-with-detailed-explanation
// or https://leetcode.com/problems/task-scheduler/discuss/104496/concise-Java-Solution-O(N)-time-O(26)-space
// Time: O(len(tasks)) - the time cost is iterate over the task list
// Space: O(1)
func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}
	// task frequency
	taskMap := map[byte]int{}
	maxFreq := 0
	numMax := 0

	for _, task := range tasks {
		taskMap[task] += 1
		if maxFreq < taskMap[task] {
			maxFreq = taskMap[task]
			numMax = 1
		} else if maxFreq == taskMap[task] {
			numMax++
		}
	}
	return int(math.Max(float64(len(tasks)), float64((maxFreq-1)*(n+1)+numMax)))
}
