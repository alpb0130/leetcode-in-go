package priorityqueue

import "container/heap"

// Time: O(klogn) - n is the length of the matrix
// Space: O(n)
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	isVisited := map[int]bool{}
	h := &nums{element{0, 0, matrix[0][0]}}
	isVisited[0] = true
	heap.Init(h)
	for i := 0; i < k || i < n*n; i++ {
		min := heap.Pop(h).(element)
		if i == k-1 {
			return min.v
		}
		if min.x+1 < n && !isVisited[n*(min.x+1)+min.y] {
			isVisited[n*(min.x+1)+min.y] = true
			heap.Push(h, element{min.x + 1, min.y, matrix[min.x+1][min.y]})
		}
		if min.y+1 < n && !isVisited[n*min.x+min.y+1] {
			isVisited[n*min.x+min.y+1] = true
			heap.Push(h, element{min.x, min.y + 1, matrix[min.x][min.y+1]})
		}
	}
	return -1
}

type element struct {
	x int
	y int
	v int
}

type nums []element

func (n nums) Len() int           { return len(n) }
func (n nums) Less(i, j int) bool { return n[i].v < n[j].v }
func (n nums) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n *nums) Pop() interface{} {
	x := (*n)[len(*n)-1]
	*n = (*n)[0 : len(*n)-1]
	return x
}
func (n *nums) Push(i interface{}) {
	*n = append(*n, i.(element))
}
