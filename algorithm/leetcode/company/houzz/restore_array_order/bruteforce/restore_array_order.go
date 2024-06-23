package bruteforce

func RestoreArray(input []int) []int{
	if len(input) == 0 {
		return []int{}
	}
	res := []int{}
	used := make([]bool, len(input))
	for _, i := range input {
		cnt := 0
		var tmp int
		for j := 0; j < len(used); j++ {
			if used[j] {
				continue
			}
			if cnt == i {
				tmp = j
				break
			}
			cnt++

		}
		used[tmp] = true
		res = append(res, tmp+1)
	}
	return res
}
