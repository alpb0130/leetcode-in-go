package dp

import (
	"strconv"
	"strings"
)

// For a given substring, we can get it's shortest encoding by trying three different case:
// 1. If length of substring is less than or equal to 4, we don't need to encode.
// 2. Otherwise
//    2.1 Split the string into two strings and concatenate the shortest encoding of those two
//        substring.
//    2.2 If a specific pattern is showing up in the substring, we can try to encode the substring
//        by using the pattern
//    For case two, we should always maintain the current shortest encoding.
// Reference: https://leetcode.com/problems/encode-string-with-shortest-length/discuss/95599/Accepted-Solution-in-Java
// Time: O(n^3)
// Space: O(n^3)
func encode(s string) string {
	if len(s) <= 4 {
		return s
	}
	dp := make([][]string, len(s))
	for i := range dp {
		dp[i] = make([]string, len(s))
		dp[i][i] = string(s[i])
	}
	for i := 0; i < len(s); i++ {
		for j := i; j >= 0; j-- {
			dp[j][i] = s[j : i+1]
			if i-j+1 > 4 {
				// generate string by combining two substring
				for k := j + 1; k <= i; k++ {
					if len(dp[j][k-1])+len(dp[k][i]) < len(dp[j][i]) {
						dp[j][i] = dp[j][k-1] + dp[k][i]
					}
				}

				// encode the substring by itself
				for k := j; k <= (i+j)/2; k++ {
					pattern := s[j : k+1]
					if len(s[j:i+1])%len(pattern) == 0 && len(strings.Replace(s[j:i+1], pattern, "", -1)) == 0 {
						num := len(s[j:i+1]) / len(pattern)
						numStr := strconv.Itoa(num)
						if len(numStr)+len(pattern)+2 < len(dp[j][i]) {
							dp[j][i] = numStr + "[" + dp[j][k] + "]"
						}
					}
				}
			}
		}
	}
	return dp[0][len(s)-1]
}
