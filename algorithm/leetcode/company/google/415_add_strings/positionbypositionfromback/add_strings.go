package positionbypositionfromback

import (
	"math"
	"strings"
)

// Time: O(max(m,n))
// Space: O(n)
func addStrings(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	length := maxInt(len1, len2)
	res := make([]int, length+1)
	carry := 0
	for i := 0; i < length; i++ {
		var d1, d2 int
		if len1-i-1 >= 0 {
			d1 = int(num1[len1-i-1] - '0')
		}
		if len2-i-1 >= 0 {
			d2 = int(num2[len2-i-1] - '0')
		}
		num := d1 + d2 + carry
		res[length-i] = num % 10
		carry = num / 10
	}
	res[0] = carry
	var str strings.Builder
	for i := 0; i < length+1; i++ {
		if i == 0 && res[i] == 0 {
			continue
		}
		str.WriteString(string('0' + res[i]))
	}
	return str.String()
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
