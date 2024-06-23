package twopointer

// Two pointers.
// Time: O(n)
// Space: O(1)
func isStrobogrammatic(num string) bool {
	// 1->1, 6->9, 8->8, 9->6, 0->0
	if num == "" || num[0] == '-' || (len(num) > 1 && num[len(num)-1] == '0') {
		return false
	}
	numMap := map[byte]byte{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
	}
	i, j := 0, len(num)-1
	for i <= j {
		if num[i] != numMap[num[j]] {
			return false
		}
		i++
		j--
	}
	return true
}
