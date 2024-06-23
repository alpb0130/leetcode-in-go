package unionfind2

// Weighted union find with path compression
// n is the number of node
// Group the x and y in to groups if a node share x or y with other node.
// The results will be number of stones - number of groups
// Time: O(n*alpha(n))
// Space: O(n)
func removeStones(stones [][]int) int {
	if stones == nil || len(stones) == 0 {
		return 0
	}

	parents := map[int]int{}
	size := map[int]int{}
	for _, stone := range stones {
		if _, ok := parents[stone[0]]; !ok {
			parents[stone[0]] = stone[0]
		}
		if _, ok := parents[stone[1]+10000]; !ok {
			parents[stone[1]+10000] = stone[1] + 10000
		}
		if _, ok := size[stone[0]]; !ok {
			size[stone[0]] = 1
		}
		if _, ok := size[stone[1]+10000]; !ok {
			size[stone[1]+10000] = 1
		}
		// union
		union(parents, size, stone[0], stone[1]+10000)
	}

	numOfMove := len(stones)
	for node, parent := range parents {
		if node == parent {
			numOfMove--
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
