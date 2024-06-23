package rearrageandscan

import "strings"

// We recover the zigzag order in the row order. Actually we don't care about the space between
// two letters in the same row because we only need the order. So we just need to recover the order
// of every row and reconstruct. Like example below, we don't need to put an space between "P" and
// "A" in the first row. So we just need to know the current direction and change direction when
// hit the boundary.
// P   A   H   N
// A P L S I I G
// Y   I   R
// Time: O(n+numRows) - n is the length of s
// Space: O(n)
func convert(s string, numRows int) string {
	if len(s) == 0 || numRows == 1 {
		return s
	}
	rows := make([][]byte, numRows)
	// Start with -1 because when we start, we will change the direction.
	step := -1
	row := 0
	for i := 0; i < len(s); i++ {
		rows[row] = append(rows[row], s[i])
		if row == numRows-1 || row == 0 {
			step = -step
		}
		row += step
	}
	var res strings.Builder
	for i := 0; i < numRows; i++ {
		res.WriteString(string(rows[i]))
	}
	return res.String()
}
