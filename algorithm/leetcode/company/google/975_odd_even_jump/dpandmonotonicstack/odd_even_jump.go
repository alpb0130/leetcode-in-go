package dpandmonotonicstack

import "sort"

// Based on our DP solution, we know it take us time to find the next jump position
// for every start position. Another way (not treemap) to optimize that is using
// monotonic stack. Sort the initial array by val along with index. Then process the
// sort structure. For every element, we try to find the first smallest larger element.
// So based on the sorted array, because for any element, we know the next element must
// has a larger/equal value and we only need to check index. If index is larger, is what
// we need. Otherwise, we need to keep searching. This seems to be O(n^2) process. But it's
// actually not because when we process a element, we use a stack to store all unprocessed
// elements and we just look over the stack and if the peek element has a smaller index, we
// know the peek element's next must be current element.
// Time: O(nlogn)
// Space: O(n)
func oddEvenJumps(A []int) int {
	if len(A) == 0 {
		return 0
	}
	// get odd next
	elementsOdd := Elements{}
	for idx, a := range A {
		elementsOdd = append(elementsOdd, Element{a, idx})
	}
	sort.Sort(elementsOdd)
	oddNext := smallestLargerAfter(elementsOdd)
	// get even next
	elementsEven := Elements{}
	for idx, a := range A {
		elementsEven = append(elementsEven, Element{-a, idx})
	}
	sort.Sort(elementsEven)
	evenNext := smallestLargerAfter(elementsEven)

	// dp
	odd := make([]bool, len(A))
	odd[len(A)-1] = true
	even := make([]bool, len(A))
	even[len(A)-1] = true
	res := 1
	for i := len(A) - 2; i >= 0; i-- {
		if oddNext[i] != -1 {
			odd[i] = even[oddNext[i]]
			if odd[i] {
				res++
			}
		}
		if evenNext[i] != -1 {
			even[i] = odd[evenNext[i]]
		}
	}
	return res
}

func smallestLargerAfter(eles Elements) []int {
	s := stack{}
	res := make([]int, len(eles))
	for i := range res {
		res[i] = -1
	}
	for _, ele := range eles {
		for s.Len() > 0 && s.Peek().i < ele.i {
			res[s.Pop().i] = ele.i
		}
		s.Push(ele)
	}
	return res
}

type Element struct {
	a int
	i int
}

type Elements []Element

func (e Elements) Len() int {
	return len(e)
}
func (e Elements) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func (e Elements) Less(i, j int) bool {
	if e[i].a != e[j].a {
		return e[i].a < e[j].a
	}
	return e[i].i < e[j].i
}

type stack []Element

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) Pop() Element {
	e := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return e
}
func (s *stack) Push(e Element) {
	*s = append(*s, e)
}
func (s *stack) Peek() Element {
	return (*s)[len(*s)-1]
}
