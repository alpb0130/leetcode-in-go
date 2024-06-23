package nosort

import (
	"strconv"
	"strings"
)

// Get the key for each string and a group of anagram should have the same key.
// And use a hashmap to map those string to the same group.
// Time: O(n*k) - n is # of strings. k is sting length
// Space: O(n*k)
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	strMap := map[string][]string{}
	for _, str := range strs {
		key := getKey(str)
		strMap[key] = append(strMap[key], str)
	}
	res := [][]string{}
	for _, group := range strMap {
		res = append(res, group)
	}
	return res
}

func getKey(str string) string {
	count := make([]int, 26)
	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++
	}
	var key strings.Builder
	for i, cnt := range count {
		if cnt != 0 {
			key.WriteString(string('a' + i))
			key.WriteString(strconv.Itoa(cnt))
		}
	}
	return key.String()
}
