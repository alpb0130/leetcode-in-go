package ifelse

import "strings"

// Use if else and record some already happened things
func isNumber(s string) bool {
	hasPoint := false
	hasExponent := false
	hasSign := false
	hasNumber := false
	s = strings.Trim(s, " ")
	for _, r := range s {
		if isDigit(r) {
			hasNumber = true
		} else if r == '.' {
			// There should be no "e" and point before "."
			if hasExponent || hasPoint {
				return false
			}
			hasPoint = true
		} else if r == 'e' {
			// There should be no numbers and "e" before "e"
			if !hasNumber || hasExponent {
				return false
			}
			// !!!Be careful. If an "e" happen, all the follow string could be a new INTEGER
			// without point. So we need to reset number, sign, point status to false.
			hasExponent = true
			hasNumber = false
			hasSign = false
			hasPoint = false
		} else if r == '-' || r == '+' {
			// There should be no sign, point and number before sign
			if hasSign || hasPoint || hasNumber {
				return false
			}
			hasSign = true
		} else {
			return false
		}
	}
	return hasNumber
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
