package backtracking

// Time: O(n!)
func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	res := [][]string{}
	// dfs
	dfs(s, []string{}, &res)
	return res
}

func dfs(s string, cur []string, res *[][]string) {
	if len(s) == 0 {
		dst := make([]string, len(cur))
		copy(dst, cur)
		*res = append(*res, dst)
		return
	}
	for i := 0; i < len(s); i++ {
		if isPalindrome(s[0 : i+1]) {
			dfs(s[i+1:], append(cur, s[0:i+1]), res)
		}
	}
}

func isPalindrome(str string) bool {
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
