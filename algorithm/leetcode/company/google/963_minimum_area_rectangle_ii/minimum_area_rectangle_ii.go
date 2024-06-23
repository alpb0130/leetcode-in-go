package _63_minimum_area_rectangle_ii

import "math"

// Iterate triangle
// Time: O(n^3)
// Space: O(n)
func minAreaFreeRect(points [][]int) float64 {
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

	res := math.MaxFloat64
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes); j++ {
			for k := 0; k < len(nodes); k++ {
				if i == j || i == k || j == k {
					continue
				}
				vector1 := []int{nodes[j].x - nodes[i].x, nodes[j].y - nodes[i].y}
				vector2 := []int{nodes[k].x - nodes[i].x, nodes[k].y - nodes[i].y}
				if vectorProd(vector1, vector2) == 0 {
					node4 := Node{nodes[j].x + vector2[0], nodes[j].y + vector2[1]}
					if nodesMap[node4] {
						res = math.Min(res, vectorLen(vector1)*vectorLen(vector2))
					}
				}
			}
		}
	}
	if res == math.MaxFloat64 {
		return 0.0
	}
	return res
}

type Node struct {
	x int
	y int
}

func vectorProd(v1, v2 []int) int {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

func vectorLen(v []int) float64 {
	return math.Sqrt(float64(v[0]*v[0] + v[1]*v[1]))
}
