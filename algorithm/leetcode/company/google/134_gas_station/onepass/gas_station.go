package onepass

// Start from the first gas station. If we stop at station k, which is not the end. It means
// the start station is not the result. And any station between start and k is not the result
// because it there is a station can reach k, the start position must be able to reach k because
// start can reach it. Once we stop at station k, we restart the simulation from this place.
// The last start point must be the result if there is a result.
// The only question left is how do we know there is a answer. We know if the total gas is larger
// than total cost, there must be a valid answer.
// Time: O(n)
// Space: O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 {
		return -1
	}
	totalDiff := 0
	curGas := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		totalDiff += gas[i] - cost[i]
		curGas += gas[i] - cost[i]
		if curGas < 0 {
			curGas = 0
			// Be careful! Should be i+1 because if we stop at position i, it means we cannot
			// reach position i+1
			start = i + 1
		}
	}
	if totalDiff < 0 {
		return -1
	}
	return start
}
