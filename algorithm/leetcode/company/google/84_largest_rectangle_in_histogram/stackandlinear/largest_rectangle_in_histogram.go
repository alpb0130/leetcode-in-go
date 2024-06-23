package stackandlinear

import "math"

// By look through the square time solution, we notice there are some duplicate checking
// in the algorithm, when we check two consecutive bars, some other bars might be checked
// twice. So we want to think about a way that can allow us not checking multiple times.
// The way is keep push height index into a stack in a ascending order (>) and till we see the
// first descending (>=) height. That mean some of the existing bar in queue cannot be extended
// to the right anymore. So we can calculate the area formed by peek height. We keeping popping
// height that is larger or equal to current height and compute the area (update max) till
// we find the peek is smaller then current height. Because we know at this point, the stack
// peek can still be extended to the right. And then we push the current bar index to stack
// and process the next bar.
// The last step is if after the for loop, we found there are still some bars in stack. This means
// those bar can be extended to the right most boundary. But the left boundary is the second peek.
// Time: O(n)
// Space: O(n)
func largestRectangleArea(heights []int) int {
	s := &stack{-1}
	max := 0
	for i := 0; i < len(heights); i++ {
		for s.Peek() != -1 && heights[s.Peek()] >= heights[i] {
			idx := s.Pop()
			max = maxInt(max, heights[idx]*(i-s.Peek()-1))
		}
		s.Push(i)
	}
	for s.Peek() != -1 {
		idx := s.Pop()
		max = maxInt(max, heights[idx]*(len(heights)-s.Peek()-1))
	}
	return max
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

type stack []int

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) Push(i int) {
	*s = append(*s, i)
}
func (s *stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}
