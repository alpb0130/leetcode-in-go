package _map

// The problem is not hard. Just use a map to record whether a sequence has shown up or not.
// But a couple of things needs to be taken care of:
// 1. Process the current 10-char sequence right away and don't wait till next new seq
//    because you may left out the last seq if you always process the previous one.
// 2. The map should record how many times the seq showing up already because we shouldn't
//    add it to the results if it has shown up for third time.
// Time: O(n) - n is the length of the string
// Space: O(n)
func findRepeatedDnaSequences(s string) []string {
	res := []string{}
	seqMap := map[string]int{}
	for i := 9; i < len(s); i++ {
		subStr := s[i-9 : i+1]
		if seqMap[subStr] == 1 {
			res = append(res, subStr)
		}
		seqMap[subStr]++
	}
	return res
}
