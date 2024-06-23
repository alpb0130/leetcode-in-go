package extraarray

func SortArray(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	numsCopy := make([]bool, 200)
	for _, num := range nums {
		numsCopy[num] = true
	}
	res := []int{}
	for i, isExisted := range numsCopy {
		if isExisted {
			res = append(res, i)
		}
	}
	return res
}
