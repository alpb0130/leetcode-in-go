package scaninroworder

import "strings"

// Instead of recovering the original matrix, we do the scan in memory and append letters in
// correct order. For the example below, our order should be
// 0 (0*2(numRows-1)+0), 6 (1*2(numRows-1)+0), 12 (2*2(rows-1)+0)
// This is for the first row. For the second row, it's
// 1 (0*2(numRows-1)+1), 5 ((0+1)*2(numRows-1)-1), 7 (1*2(numRows-1)+1), 11 ((1+1)*2(numRows-1)-1), 13 (2*2(numRows-1)+1)
// For the third row, it's
// 2 (0*2(numRows-1)+2), 4 ((0+1)*2(numRows-1)-2), 8 (1*2(numRows-1)+2), 10 ((1+1)*2(numRows-1)-2), 14 (2*2(numRows-1)+2)
// For the last row, it's
// 3 (0*2(numRows-1)+3), 9 (1*2(numRows-1)+3)
//
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
//
// Time: O(n)
// Space: O(1)
func convert(s string, numRows int) string {
	if len(s) == 0 || numRows == 1 {
		return s
	}
	var res strings.Builder
	step := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; i+j < len(s); j += step {
			res.WriteString(string(s[i+j]))
			if i != 0 && i != numRows-1 && j+step-i < len(s) {
				res.WriteString(string(s[j+step-i]))
			}
		}
	}
	return res.String()
}
