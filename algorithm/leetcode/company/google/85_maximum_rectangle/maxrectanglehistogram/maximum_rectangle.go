package maxrectanglehistogram

import "math"

// This is the best solution, we can reuse the solution to max rectangle in histogram (linear
// time).
// What we can do is get the histogram of every row base on the array and then apply the method.
// Time: O(m*n)
// Space: O(n) - if we reuse teh dp array
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res := 0
	prevHist := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		curHist := make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				curHist[j] = 0
			} else {
				curHist[j] = prevHist[j] + 1
			}
		}
		// max rectangle in histgram
		res = maxInt(res, maxRectangle(curHist))
		prevHist = curHist
	}
	return res
}

func maxRectangle(histograms []int) int {
	max := 0
	s := &stack{-1}
	for i := 0; i < len(histograms); i++ {
		for s.Peek() != -1 && histograms[s.Peek()] > histograms[i] {
			height := histograms[s.Pop()]
			max = maxInt(max, height*(i-s.Peek()-1))
		}
		s.Push(i)
	}
	for s.Peek() != -1 {
		height := histograms[s.Pop()]
		max = maxInt(max, height*(len(histograms)-s.Peek()-1))
	}
	return max
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

type stack []int

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *stack) Push(i int) {
	*s = append(*s, i)
}
func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}
