package recursive

// Solve the problem recursively. We need to support "." and "*". The hard part is to support "*".
// Three cases:
// 1. The second char of the pattern string is not "*":
// 	  look at the first char of s and first char of pattern and check whether they are matching
//    with each other. Then call isMatch(s[1:], p[1:]) recursively
// 2. The second char of the pattern string is "*":
//    2.1 If first char of s match with first char of pattern, call isMatch(s[1:], p)
//    2.2 If not, call isMatch(s, p[2:])
// Time: ???
// Space: ???
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	firstMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}
