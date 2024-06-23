package bitmanipulation

// Time: O(n) - n is the length of the string
// Space: O(n)
func findRepeatedDnaSequences(s string) []string {
	res := []string{}
	seqMap := map[int]int{}
	encode := map[byte]int{}
	encode['C'-'A'] = 1
	encode['G'-'A'] = 2
	encode['T'-'A'] = 3
	for i := 0; i < len(s)-9; i++ {
		v := 0
		for j := i; j < i+10; j++ {
			v |= encode[s[j]-'A']
			v <<= 2
		}
		if seqMap[v] == 1 {
			res = append(res, s[i:i+10])
		}
		seqMap[v]++
	}
	return res
}
