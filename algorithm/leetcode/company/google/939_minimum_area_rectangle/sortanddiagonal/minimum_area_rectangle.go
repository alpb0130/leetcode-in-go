package sortanddiagonal

import (
	"math"
)

// For all pair of points which is not in the same horizontal and vertical line, we
// try to check whether we have the other two points. If yes, try to update the area.
// Sort the points to reduce the time a little bit.
// Time: O(n^2)
// Space: O(n)
func minAreaRect(points [][]int) int {
	if len(points) < 4 {
		return 0
	}
	nodes := []Node{}
	nodesMap := map[Node]bool{}
	for _, pointPair := range points {
		point := Node{pointPair[0], pointPair[1]}
		nodes = append(nodes, point)
		nodesMap[point] = true
	}
	res := math.MaxInt32
	minX, minY := math.MaxInt32, math.MaxInt32
	for i := 0; i < len(nodes)-1; i++ {
		for j := i + 1; j < len(nodes); j++ {
			x := intAbs(nodes[j].x - nodes[i].x)
			y := intAbs(nodes[j].y - nodes[i].y)
			if nodes[j].x == nodes[i].x || nodes[j].y == nodes[i].y || (x >= minX && y >= minY) {
				continue
			}
			n1 := Node{nodes[i].x, nodes[j].y}
			n2 := Node{nodes[j].x, nodes[i].y}
			if nodesMap[n1] && nodesMap[n2] && x*y < res {
				res = x * y
				minX = x
				minY = y
			}
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func intAbs(a int) int {
	return int(math.Abs(float64(a)))
}

type Node struct {
	x int
	y int
}
