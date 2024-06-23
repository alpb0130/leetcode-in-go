package strings_only_in_one_list

func StringsInOneList(strs [][]string) []string {
	if len(strs) == 0 {
		return []string{}
	}

	strCounter := map[string]int{}
	for _, list := range strs {
		strMap := map[string]bool{}
		for _, str := range list {
			if !strMap[str] {
				strCounter[str]++
				strMap[str] = true
			}
		}
	}
	res := []string{}
	for str, cnt := range strCounter {
		if cnt == 1 {
			res = append(res, str)
		}
	}
	return res
}
