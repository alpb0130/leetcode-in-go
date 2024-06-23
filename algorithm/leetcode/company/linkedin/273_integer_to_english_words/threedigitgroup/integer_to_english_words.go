package threedigitgroup

import "strings"

// The problem limit the number to 0-2^32-1. ASK the number range before you write code. So
// the max number would be 4,294,967,295 which is billion something. We we don't need rotation
// magnitude notation like "Thousand Billion Million Thousand"
// Divide the number to group of 3 digit and process the group. Shouldn't
// be too hard. We process the number from low significant digit to high significant so remember
// to process the res string list reversely.
// Some clear version: https://leetcode.com/problems/integer-to-english-words/solution/
// Time: O(n) - n is the length of the number
// Space: O(n) - string is immutable is we need some temporary space
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	belowTwentyMap := map[int]string{
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
	}
	belowHundredMap := map[int]string{
		20: "Twenty",
		30: "Thirty",
		40: "Forty",
		50: "Fifty",
		60: "Sixty",
		70: "Seventy",
		80: "Eighty",
		90: "Ninety",
	}
	magMap := map[int]string{
		1: "Thousand",
		2: "Million",
		3: "Billion",
	}
	mag := 0
	res := []string{}
	for num > 0 {
		remainder := num % 1000
		if mag != 0 && remainder != 0 {
			res = append(res, magMap[mag])
		}
		num /= 1000
		if num > 0 {
			mag++
		}
		// process remainder
		// less than 100
		a := remainder % 100
		if a < 20 {
			if str, ok := belowTwentyMap[a]; ok {
				res = append(res, str)
			}
		} else {
			b := (a / 10) * 10
			c := a % 10
			if str, ok := belowTwentyMap[c]; ok {
				res = append(res, str)
			}
			res = append(res, belowHundredMap[b])
		}
		// hundred
		remainder /= 100
		if str, ok := belowTwentyMap[remainder]; ok {
			res = append(res, "Hundred")
			res = append(res, str)
		}
	}
	var s strings.Builder
	for i := len(res) - 1; i >= 0; i-- {
		s.WriteString(res[i])
		if i != 0 {
			s.WriteString(" ")
		}
	}
	return s.String()
}
