package int

// Time: O(log10 n) - n is the number
// Space: O(1) - always be a number regardless how big it is
func isPalindrome(x int) bool {
	return x >= 0 && reverseNum(x) == x
}

func reverseNum(x int) int {
	res := 0
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}
	return res
}
