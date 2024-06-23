package sortedrange

import "math"

// General idea is to maintain a sorted list of range without overlapping. Split a range or merge
// two ranges if needed.
type RangeModule struct {
	ranges []Range
}

type Range struct {
	left  int
	right int
}

func Constructor() RangeModule {
	return RangeModule{
		ranges: []Range{},
	}
}

// Time: O(logn+n) - logn for search and n for copy array
// Space: O(n)
func (this *RangeModule) AddRange(left int, right int) {
	if len(this.ranges) == 0 {
		this.ranges = append(this.ranges, Range{left, right})
		return
	}
	i, j := this.findInterval(left, right)
	if i == -1 {
		this.ranges = append(this.ranges, Range{left, right})
	} else if j == -1 {
		this.ranges = append([]Range{Range{left, right}}, this.ranges...)
	} else if j < i {
		dst := make([]Range, len(this.ranges[i:]))
		copy(dst, this.ranges[i:])
		this.ranges = append(this.ranges[:j+1], Range{left, right})
		this.ranges = append(this.ranges, dst...)
	} else {
		l := minInt(left, this.ranges[i].left)
		r := maxInt(right, this.ranges[j].right)
		dst := make([]Range, len(this.ranges[j+1:]))
		copy(dst, this.ranges[j+1:])
		this.ranges = append(this.ranges[:i], Range{l, r})
		this.ranges = append(this.ranges, dst...)
	}
}

// Time: O(logn)
// Space: O(1)
func (this *RangeModule) QueryRange(left int, right int) bool {
	i, j := this.findInterval(left, right)
	return i != -1 && i == j && this.ranges[i].left <= left && this.ranges[i].right >= right
}

// Time: O(logn+n) - logn for search and n for copy array
// Space: O(n)
func (this *RangeModule) RemoveRange(left int, right int) {
	i, j := this.findInterval(left, right)
	if i != -1 && j != -1 && j >= i {
		dst := make([]Range, len(this.ranges[j+1:]))
		copy(dst, this.ranges[j+1:])
		res := make([]Range, len(this.ranges[:i]))
		copy(res, this.ranges[:i])
		if left > this.ranges[i].left {
			res = append(res, Range{this.ranges[i].left, left})
		}
		if right < this.ranges[j].right {
			res = append(res, Range{right, this.ranges[j].right})
		}
		res = append(res, dst...)
		this.ranges = res
	}
}

func (this *RangeModule) findInterval(left, right int) (int, int) {
	start, end := 0, len(this.ranges)-1
	i := -1
	for start <= end {
		mid := (start + end) / 2
		if this.ranges[mid].right >= left && (mid == 0 || this.ranges[mid-1].right < left) {
			i = mid
			break
		} else if this.ranges[mid].right < left {
			start++
		} else {
			end--
		}
	}

	start, end = 0, len(this.ranges)-1
	j := -1
	for start <= end {
		mid := (start + end) / 2
		if this.ranges[mid].left <= right && (mid == len(this.ranges)-1 || this.ranges[mid+1].left > right) {
			j = mid
			break
		} else if this.ranges[mid].left > right {
			end--
		} else {
			start++
		}
	}
	return i, j
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
