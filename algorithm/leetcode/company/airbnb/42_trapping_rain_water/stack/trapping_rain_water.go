package stack

import "math"

// One pass. Pretty tricky solution. Please refer here:
// https://leetcode.com/problems/trapping-rain-water/solution/
// Time: O(n)
// Space: O(n)
func trap(height []int) int {
	s := &stack{}
	res := 0
	for i := 0; i < len(height); i++ {
		h := height[i]
		for s.Len() > 0 && h > height[s.Peek()] {
			minIdx := s.Pop()
			if s.Len() == 0 {
				break
			}
			width := i - s.Peek() - 1
			fillingHeight := minInt(h, height[s.Peek()]) - height[minIdx]
			res += width * fillingHeight
		}
		s.Push(i)
	}
	return res
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

type stack []int

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
func (s *stack) Len() int {
	return len(*s)
}
