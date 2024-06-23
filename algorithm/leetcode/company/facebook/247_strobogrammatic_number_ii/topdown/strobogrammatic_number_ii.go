package topdown

// Recursive top-down method. We construct the string start to mid. When we hit the
// mid, construct the whole string by flip the first half. Some edges case we need to
// take care of:
// 1. 0 cannot be the first number except the case where where n == 1 (which is
//	  half == 1 and isOdd)
// 2. If isOdd and current num is the number in the middle, it should only be 0, 1, 8
// Time: O(5^(n/2)) - from bottom to up, each level has 5^k results
// Space: O(n) - we only use a array to store current string in the recursive call
//        to store string bytes and it's always a stable memory and we don't create
//        new one. But the leaf node need to use n space to create the final string.
//        Total space should be O(3n) (call stack, current string bytes and final string bytes)
func findStrobogrammatic(n int) []string {
	if n == 0 {
		return []string{""}
	}
	// number map
	numMap := map[byte]byte{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
	}
	res := []string{}
	// dfs to generate all res
	dfs((n+1)/2, []byte{}, n%2 == 1, numMap, &res)
	return res

}

func dfs(half int, curStr []byte, isOdd bool, numMap map[byte]byte, res *[]string) {
	if len(curStr) == half {
		// add result
		*res = append(*res, constructStr(curStr, isOdd, numMap))
		return
	}
	for num1, num2 := range numMap {
		// The first element shouldn't be 0 except the case where n == 1 (which is
		// half == 1 and isOdd)
		if len(curStr) == 0 && num1 == '0' && !(half == 1 && isOdd) {
			continue
		}
		// If isOdd and current num is the number in the middle, it should only be 0, 1
		// 8
		if isOdd && len(curStr) == half-1 && num1 != num2 {
			continue
		}
		dfs(half, append(curStr, num1), isOdd, numMap, res)
	}
}

func constructStr(str []byte, isOdd bool, numMap map[byte]byte) string {
	if isOdd {
		return string(str[:len(str)-1]) + flipStr(str, numMap)
	} else {
		return string(str) + flipStr(str, numMap)
	}
}

func flipStr(str []byte, numMap map[byte]byte) string {
	res := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		res[i] = numMap[str[len(str)-i-1]]
	}
	return string(res)
}
