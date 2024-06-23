package bottomup

// Recursive bottom-up method. We construct the string from middle to end
// Time: O(5^(n/2)) - from bottom to up, each level has 5^k results
// Space: O(5^(n/2))
func findStrobogrammatic(n int) []string {
	return dfs(n, n)
}

func dfs(m int, n int) []string {
	if m == 0 {
		return []string{""}
	}
	if m == 1 {
		return []string{"0", "1", "8"}
	}
	list := dfs(m-2, n)
	res := []string{}
	for _, str := range list {
		if m != n {
			res = append(res, "0"+str+"0")
		}
		res = append(res, "1"+str+"1")
		res = append(res, "6"+str+"9")
		res = append(res, "8"+str+"8")
		res = append(res, "9"+str+"6")
	}
	return res
}
