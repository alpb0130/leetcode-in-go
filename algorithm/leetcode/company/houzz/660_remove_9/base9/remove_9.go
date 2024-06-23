package base9

// The problem requires us to find the nth number in the list of numbers without digit 9.
// We can convert this problem to find the nth number in a base-9 number series. So our
// task is given n, convert it to base-9 expression.
// In every for loop, we need to use modulo to decide current digit and then multiply with
// base which should be 1, 10, 100, ...
// Time: O(logn)
// Space: O(1)
func newInteger(n int) int {
	res := 0
	base := 1
	for n != 0 {
		res += base * (n % 9)
		n /= 9
		base *= 10
	}
	return res
}
