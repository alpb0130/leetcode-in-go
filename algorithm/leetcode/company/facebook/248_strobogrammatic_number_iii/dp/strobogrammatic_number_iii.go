package dp

// Construct numbers from low length to high length and only add a number if it's within
// [low, high]. As we go through the process, we can find a lot of duplications. Like
// if we process length i and i+2, we need to process i-2 twice. We can memorize that.
// Time: O(5^n) - we actually get all strings with length 0-n which is 5^n in total
// Space: O(5^n) - the map need to store all strings and it's O((5^n)*n). The extra n
//        is string length.
func strobogrammaticInRange(low string, high string) int {
	if len(high) == 0 {
		return 0
	}
	numMap := map[byte]byte{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
	}
	lengthToStrings := map[int][]string{
		0: {""},
		1: {"0", "1", "8"},
	}

	count := 0
	// construct string from short to long
	for i := len(low); i <= len(high); i++ {
		helper(lengthToStrings, numMap, low, high, i, i, &count)
	}
	return count
}

func helper(lengthToStrings map[int][]string, numMap map[byte]byte, low, high string, curLen, target int, count *int) []string {
	if len(lengthToStrings[curLen]) > 0 && curLen < target {
		return lengthToStrings[curLen]
	}

	if len(lengthToStrings[curLen]) == 0 {
		strList := []string{}
		list := helper(lengthToStrings, numMap, low, high, curLen-2, target, count)
		for _, word := range list {
			for left, right := range numMap {
				strList = append(strList, string(left)+word+string(right))
			}
		}
		lengthToStrings[curLen] = strList
	}
	if curLen == target {
		*count += len(lengthToStrings[curLen])
		for _, word := range lengthToStrings[curLen] {
			if (len(word) > 1 && word[0] == '0') ||
				(target == len(low) && word < low) ||
				(target == len(high) && word > high) {
				*count--
			}
		}
	}

	return lengthToStrings[curLen]
}
