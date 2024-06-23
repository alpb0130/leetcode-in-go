package stringprocess

import (
	"math"
	"strconv"
)

// Some special case need to be take care of:
// 1. 1-9
// 2. 10-19
// 3. 98 and 999
// 4. 100 and 1000
// Time: O(len(n))
// Space: O(len(n))
func nearestPalindromic(n string) string {
	prefix, isOdd := getPrefixAddIsOdd(n)

	palindromes := []string{
		getSmaller(prefix, isOdd),
		getPalindrome(prefix, isOdd),
		getLarger(prefix, isOdd),
	}
	minDiff := math.MaxInt32
	var res string
	for _, p := range palindromes {
		diff := getDiff(p, n)
		if diff == 0 {
			continue
		}
		if diff < minDiff {
			minDiff = diff
			res = p
		}
	}
	return res
}

func getDiff(n1, n2 string) int {
	num1, _ := strconv.Atoi(n1)
	num2, _ := strconv.Atoi(n2)
	return int(math.Abs(float64(num1 - num2)))
}

func reverseString(str string) string {
	bytes := []byte(str)
	i, j := 0, len(str)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}

func getPrefixAddIsOdd(n string) (string, bool) {
	return n[:(len(n)+1)/2], len(n)%2 == 1
}

func getPalindrome(prefix string, isOdd bool) string {
	if isOdd {
		return prefix + reverseString(prefix[:len(prefix)-1])
	} else {
		return prefix + reverseString(prefix)
	}
}

func getLarger(prefix string, isOdd bool) string {
	larger := stringAdd(prefix, 1)
	if len(larger) > len(prefix) {
		if isOdd {
			larger = larger[:len(larger)-1]
		}
		isOdd = !isOdd
	}
	return getPalindrome(larger, isOdd)
}

func getSmaller(prefix string, isOdd bool) string {
	smaller := stringAdd(prefix, -1)
	if smaller == "0" && !isOdd {
		return "9"
	}
	if len(smaller) < len(prefix) {
		if !isOdd {
			smaller = smaller + "9"
		}
		isOdd = !isOdd
	}
	return getPalindrome(smaller, isOdd)
}

func stringAdd(str string, delta int) string {
	n, _ := strconv.Atoi(str)
	return strconv.Itoa(n + delta)
}
