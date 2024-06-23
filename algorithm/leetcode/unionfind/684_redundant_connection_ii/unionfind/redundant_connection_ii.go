package unionfind

// There would be three cases for the invalid graph:
// 1. One node has two parents:
//     1.1 Remove one edge from parents would restore the graph
//     1.2 Remove ond edge from parents, the graph still has circle. Like A -> B <=> C, B has two
//         parents, if we remove A->B, the remaining graph would still have circle.
// 2. No node has two parents but the root has parents. There would be a circle in this case.
// To solve this problem, we need to go over all edges and note down the two edges that causes
// a node with two parents. If no such edges, use union and find to find the first edge to cause
// a circle (which is Redundant Connection I). If do have such edges, remove one of the edges
// first. If the remaining graph is valid, return the removed one. Otherwise there would be a
// circle in the remaining graph and we need to remove the edge that cause the circle.
// Reference: https://leetcode.com/problems/redundant-connection-ii/discuss/108045/C%2B%2BJava-Union-Find-with-explanation-O(n)
// Time: O(n)
// Space: O(n)
func findRedundantDirectedConnection(edges [][]int) []int {
	parents := make([]int, len(edges)+1)
	var res1, res2 []int
	for _, edge := range edges {
		if parents[edge[1]] == 0 {
			parents[edge[1]] = edge[0]
		} else {
			res2 = []int{edge[0], edge[1]}
			res1 = []int{parents[edge[1]], edge[1]}
			edge[1] = 0
		}
	}
	size := make([]int, len(edges)+1)
	for idx := range parents {
		parents[idx] = idx
		size[idx] = 1
	}
	// union find to verify the remaining graph
	for _, edge := range edges {
		if edge[1] == 0 {
			continue
		}
		if isConnected(parents, edge[0], edge[1]) {
			if res1 == nil {
				return edge
			} else {
				return res1
			}
		}
		union(parents, size, edge[0], edge[1])
	}
	return res2
}

func union(parents, size []int, i, j int) {
	rootI := find(parents, i)
	rootJ := find(parents, j)
	if size[rootI] < size[rootJ] {
		parents[rootI] = rootJ
		size[rootJ] += size[rootI]
	} else {
		parents[rootJ] = rootI
		size[rootI] += size[rootJ]
	}
}

func find(parents []int, i int) int {
	if parents[i] != i {
		parents[i] = find(parents, parents[i])
	}
	return parents[i]
}

func isConnected(parents []int, i, j int) bool {
	return find(parents, i) == find(parents, j)
}
