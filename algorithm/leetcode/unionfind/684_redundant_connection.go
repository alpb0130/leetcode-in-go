package unionfind

// Union find with path compression (using map)
// Time: O(nlogn)
// Space: O(n)
func findRedundantConnection(edges [][]int) []int {
	// handle extreme case
	if edges == nil || len(edges) == 0 {
		return nil
	}

	parents := map[int]int{}
	for _, edge := range edges {
		if isConnected_684(parents, edge[0], edge[1]) {
			return edge
		}
		union_684(parents, edge[0], edge[1])
	}
	return nil
}

func union_684(parents map[int]int, i, j int) {
	rootI := find_684(parents, i)
	rootJ := find_684(parents, j)
	if rootI == rootJ {
		return
	}
	parents[rootI] = rootJ

}

func find_684(parents map[int]int, i int) int {
	for {
		parent, ok := parents[i]
		if ok && parent != i {
			parents[i] = parents[parent]
			i = parent
		} else {
			parents[i] = i
			return i
		}
	}
}

func isConnected_684(parents map[int]int, i, j int) bool {
	if find_684(parents, i) == find_684(parents, j) {
		return true
	}
	return false
}

// Union find with path compression (using array)
// Time: O(nlogn)
// Space: O(n)
func findRedundantConnection1(edges [][]int) []int {
	// handle extreme case
	if edges == nil || len(edges) == 0 {
		return nil
	}

	parents := make([]int, len(edges)+1)
	for i := 0; i < len(edges)+1; i++ {
		parents[i] = i
	}
	for _, edge := range edges {
		if isConnected_684_1(parents, edge[0], edge[1]) {
			return edge
		}
		union_684_1(parents, edge[0], edge[1])
	}
	return nil
}

func union_684_1(parents []int, i, j int) {
	rootI := find_684_1(parents, i)
	rootJ := find_684_1(parents, j)
	if rootI == rootJ {
		return
	}
	parents[rootI] = rootJ
}

func find_684_1(parents []int, i int) int {
	for parents[i] != i {
		parent := parents[i]
		parents[i] = parents[parent]
		i = parent
	}
	return i
}

func isConnected_684_1(parents []int, i, j int) bool {
	if find_684_1(parents, i) == find_684_1(parents, j) {
		return true
	}
	return false
}
