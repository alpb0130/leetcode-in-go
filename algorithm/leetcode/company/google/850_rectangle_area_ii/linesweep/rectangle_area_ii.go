package linesweep

import (
	"math"
	"sort"
)

// Line sweep. Imaging there is a virtual horizontal line from bottom to the top. We keep moving the
// line up and see the area that overlapping with rectangle. What we need to do is for every rectangle,
// we want to add the bottom line and top line in a list and also mark it's y value and whether it's
// start or end. Then we sort all interval lines by its y value (if tie, any order). Start from the bottom
// line, we keep record of current intervals and keep them in ascending order by line start value.
// Every time we process the line, we know there are some intervals to be added to/removed from current
// interval list and we calculate the rectangle size so far. Then we update the intervals. If's start line,
// we add interval. If's end line, we remove the interval from the interval lists
// Reference: https://leetcode.com/problems/rectangle-area-ii/solution/
//            https://leetcode.com/problems/rectangle-area-ii/discuss/137914/C%2B%2BPython-Discretization-and-O(NlogN)
// Time: O(n^2)
// Space: O(n)
func rectangleArea(rectangles [][]int) int {
	if len(rectangles) == 0 {
		return 0
	}
	lines := []Line{}
	for _, rectangle := range rectangles {
		lines = append(lines, Line{rectangle[1], rectangle[0], rectangle[2], true})
		lines = append(lines, Line{rectangle[3], rectangle[0], rectangle[2], false})
	}
	sort.Slice(lines, func(i, j int) bool {
		return lines[i].y < lines[j].y
	})
	ranges := [][]int{}
	res := 0
	prev := lines[0].y
	for _, line := range lines {
		// get area
		height := line.y - prev
		res += height * getLen(ranges)
		prev = line.y
		if line.isStart {
			ranges = addRange(ranges, []int{line.i, line.j})
		} else {
			ranges = removeRange(ranges, []int{line.i, line.j})
		}
	}
	return res % 1000000007
}

func getLen(ranges [][]int) int {
	if len(ranges) == 0 {
		return 0
	}
	res := 0
	left := ranges[0][0]
	for _, ran := range ranges {
		left = maxInt(left, ran[0])
		rangeLen := maxInt(0, ran[1]-left)
		left = maxInt(left, ran[1])
		res += rangeLen
	}
	return res
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

// We can also simply sort the range list but it takes longer time.
func addRange(ranges [][]int, r []int) [][]int {
	if len(ranges) == 0 {
		return append(ranges, r)
	}
	idx := len(ranges)
	for i, ran := range ranges {
		if ran[0] >= r[0] {
			idx = i
			break
		}
	}
	secondPart := make([][]int, len(ranges[idx:]))
	copy(secondPart, ranges[idx:])
	ranges = append(ranges[:idx], r)
	ranges = append(ranges, secondPart...)
	return ranges
}

func removeRange(ranges [][]int, r []int) [][]int {
	var idx int
	for i, ran := range ranges {
		if ran[0] == r[0] && ran[1] == r[1] {
			idx = i
			break
		}
	}
	ranges = append(ranges[:idx], ranges[idx+1:]...)
	return ranges
}

type Line struct {
	y       int
	i       int
	j       int
	isStart bool
}
