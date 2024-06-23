package backtracking

// Backtracking and dfs. Start from bottom and generate all possible string in upper level.
// If all pairs in bottom have been processed, replace bottom with upper level string.x
// Time: O(k^(n^2)) - n is length of bottom, k is the size of allowed. This should be the worst
//       case complexity.
// Space: O(n^2)
func pyramidTransition(bottom string, allowed []string) bool {
	bottomToUpMap := map[string][]string{}
	for _, allow := range allowed {
		bottomToUpMap[allow[:2]] = append(bottomToUpMap[allow[:2]], allow[2:])
	}
	return dfs(bottom, "", bottomToUpMap)
}

func dfs(bottom string, curUp string, bottomToUpMap map[string][]string) bool {
	if len(bottom) == 1 && len(curUp) == 0 {
		return true
	}
	if len(bottom) == 1 {
		return dfs(curUp, "", bottomToUpMap)
	}
	for _, up := range bottomToUpMap[bottom[:2]] {
		if dfs(bottom[1:], curUp+up, bottomToUpMap) {
			return true
		}
	}
	return false
}
