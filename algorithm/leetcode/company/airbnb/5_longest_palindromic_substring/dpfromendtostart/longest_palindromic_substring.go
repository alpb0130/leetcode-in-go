package dpfromendtostart

// The brute-force method is enum all substrings and check whether is's palindromic or not.
// But it would introduce duplicated check. We can start from the end and short substring
// and store the palindromic status in an array.
// DP. Use an 2D array to store the palindrome status of previsited substring. For current string
// we only need to check whether the start and end letter are equal and the inner substring is
// palindrome.
// Time: O(n^2)
// Space: O(n^2)
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}
	max := 0
	str := ""
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i+1][j+1] = s[i] == s[j] && (j-i <= 1 || dp[i+2][j])
			if dp[i+1][j+1] && j-i+1 > max {
				max = j - i + 1
				str = s[i : j+1]
			}
		}
	}
	return str
}
