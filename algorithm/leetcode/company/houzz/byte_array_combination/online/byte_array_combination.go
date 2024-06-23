package online

import "strings"

/*
	[
		['a', 'b', 'c'],
		['d'],
		['e', 'f'],
	]
*/

type CombinationIterator struct {
	matrix  [][]byte
	index   []int
	hasNext bool
}

func NewCombinationIterator(matrix [][]byte) CombinationIterator {
	return CombinationIterator{
		matrix: matrix,
		index:  make([]int, len(matrix)),
	}
}

func (c CombinationIterator) Next() string {
	var res strings.Builder
	for i := 0; i < len(c.index); i++ {
		res.WriteString(string(c.matrix[i][c.index[i]]))
	}
	for i := len(c.index) - 1; i >= 0; i-- {
		c.index[i]++
		if c.index[i] == len(c.matrix[i]) {
			c.index[i] = 0
			if i == 0 {
				c.hasNext = false
			}
		} else {
			break
		}
	}

	return res.String()
}

func (c CombinationIterator) HasNext() bool {
	return c.hasNext
}
