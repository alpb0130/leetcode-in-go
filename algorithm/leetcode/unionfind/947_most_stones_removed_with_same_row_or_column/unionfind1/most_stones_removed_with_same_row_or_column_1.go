package unionfind1

// Weighted union find with path compression
// n is the number of node
// Time: O(n*alpha(n))
// Space: O(n)
func removeStones(stones [][]int) int {
	if stones == nil || len(stones) == 0 {
		return 0
	}
	parents := map[int]int{}
	size := map[int]int{}
	verticalNodeMap := map[int]int{}
	horizontalNodeMap := map[int]int{}
	for _, stone := range stones {
		index := stone[0]*10000 + stone[1]
		parents[index] = index
		size[index] = 1
		if verticalNode, ok := verticalNodeMap[stone[1]]; ok {
			union(parents, size, index, verticalNode)
		} else {
			verticalNodeMap[stone[1]] = index
		}

		if horizontalNode, ok := horizontalNodeMap[stone[0]]; ok {
			union(parents, size, index, horizontalNode)
		} else {
			horizontalNodeMap[stone[0]] = index
		}
	}

	// process set and get results
	numOfMove := 0
	for child, parent := range parents {
		if child == parent {
			numOfMove += size[child] - 1
		}
	}
	return numOfMove
}

func union(parents, size map[int]int, i, j int) {
	root1 := find(parents, i)
	root2 := find(parents, j)
	if root1 == root2 {
		return
	}
	if size[root1] > size[root2] {
		parents[root2] = root1
		size[root1] += size[root2]
	} else {
		parents[root1] = root2
		size[root2] += size[root1]
	}
}

func find(parents map[int]int, i int) int {
	if parents[i] != i {
		parents[i] = find(parents, parents[i])
	}
	return parents[i]
}
