package halfint

// Only flip half of the number (abcd and abc is two different case so we need the one more check
// at the end). Also we cannot use the logic to half-flip number ending with 0 (try some case) so
// we need to extra checks at the beginning. And is a number is ending with 0, it cannot be
// palindromic except it 0.
// Time: O(log10 n) - n is the number
// Space: O(1) - always be a number regardless how big it is
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	res := 0
	for x > res {
		res = res*10 + x%10
		x /= 10
	}
	return res == x || res/10 == x
}
