package countingsort

// Counting sort but we need to note down the largest citation which might hurt the time complexity
// One improvement we can do is to limit the citation to len(citations) because we know that the h
// index cannot be larger than len(citations). So for all citations which are larger than n, we can
// count all those citations as n.
// Time: O(k + n) - k is the largest citation. n is len(citations)
// Space: O(n)
func hIndex(citations []int) int {
	max := 0
	counter := map[int]int{}
	for _, citation := range citations {
		counter[citation]++
		if citation > max {
			max = citation
		}
	}
	num := 0
	i := max
	for i >= 0 {
		num += counter[i]
		if num >= i {
			break
		}
		i--
	}
	return i
}
