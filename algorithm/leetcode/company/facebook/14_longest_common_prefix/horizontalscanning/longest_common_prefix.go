package horizontalscanning

import "math"

// Compare current prefix with every str and get the new prefix. A bad thing for this method
// is we need more space in worst case because string is immutable (use byte array to fix it).
// If the shortest string is in the last, we would get the worst case. Use vertical scan to fix
// that.
// Time: O(m*n) - m is number of strings and n is average length of the string
// Space: O(n)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for i := 0; i < len(strs); i++ {
		res = res[:minInt(len(res), len(strs[i]))]
		for j := 0; j < minInt(len(res), len(strs[i])); j++ {
			if res[j] != strs[i][j] {
				res = res[:j]
				break
			}
		}
	}
	return res
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
