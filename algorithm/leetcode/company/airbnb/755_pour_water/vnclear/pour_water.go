package vnclear

// A clearer version.
// Reference: https://leetcode.com/problems/pour-water/discuss/171104/Share-my-Java-straightforward-and-concise-solution
// Time: O(V*N)
// Space: O(1)
func pourWater(heights []int, V int, K int) []int {
	for i := 1; i <= V; i++ {
		// fall left
		j := K
		for j > 0 && heights[j] >= heights[j-1] {
			j--
		}
		for j < len(heights)-1 && heights[j] >= heights[j+1] {
			j++
		}
		for j > K && heights[j] >= heights[j-1] {
			j--
		}
		heights[j]++
	}
	return heights
}
