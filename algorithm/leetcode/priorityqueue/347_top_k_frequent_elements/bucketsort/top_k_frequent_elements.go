package bucketsort

// Use a map to store the numbers of each frequency and then iterate the map to get top k elements
// Time: O(n)
// Space: O(x)
func topKFrequent(nums []int, k int) []int {
	counter := map[int]int{}
	for _, num := range nums {
		counter[num]++
	}
	maxFreq := 0
	freqNumsMap := map[int][]int{}
	for num, cnt := range counter {
		freqNumsMap[cnt] = append(freqNumsMap[cnt], num)
		if cnt > maxFreq {
			maxFreq = cnt
		}
	}
	results := []int{}
	for i := maxFreq; i > 0; i-- {
		if numbers, ok := freqNumsMap[i]; ok {
			for _, n := range numbers {
				results = append(results, n)
				if len(results) == k {
					return results
				}
			}
		}
	}
	return results
}
