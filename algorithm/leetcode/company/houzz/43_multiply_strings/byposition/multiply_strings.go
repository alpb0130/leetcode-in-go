package byposition

import "strings"

// We know if num1(len m)*num2(len m), we would get a number that with maximum length m+n.
// And when we process the digit in position i and j, the result should be put at position
// i+j+1 and i+j. All process should be from back to front.
// Time: O(m*n)
// Space: O(m+n)
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			num := int(num1[i]-'0') * int(num2[j]-'0')
			r := num % 10
			c := num / 10
			res[i+j+1] += r
			res[i+j] += c
		}
	}
	// Process carry and make sure number in every position is less than 10
	c := 0
	for i := len(res) - 1; i >= 0; i-- {
		r := int(res[i]+c) % 10
		c = int(res[i]+c) / 10
		res[i] = r
	}
	// Get results.
	var str strings.Builder
	for i := 0; i < len(res); i++ {
		// skip the starting 0
		if i == 0 && res[i] == 0 {
			continue
		}
		str.WriteString(string('0' + res[i]))
	}
	return str.String()
}
