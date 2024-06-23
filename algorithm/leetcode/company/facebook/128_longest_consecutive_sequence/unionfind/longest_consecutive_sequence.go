package unionfind

// Weighted union find with path compression
// Time: O(n*log*n)
// Space: O(n)
func longestConsecutive(nums []int) int {
	root := map[int]int{}
	size := map[int]int{}
	for _, num := range nums {
		// This check is every import.
		// If we don't have this check, we will reset some already process number to be the init
		// status
		if _, ok := root[num]; !ok {
			root[num] = num
			size[num] = 1
		}
		if _, ok := root[num-1]; ok {
			// union
			union(root, size, num, num-1)
		}
		if _, ok := root[num+1]; ok {
			// union
			union(root, size, num, num+1)
		}
	}
	res := 0
	for _, s := range size {
		if s > res {
			res = s
		}
	}
	return res
}

func union(root map[int]int, size map[int]int, i, j int) {
	rootI := find(root, i)
	rootJ := find(root, j)
	if rootI == rootJ {
		return
	}
	if size[rootI] < size[rootJ] {
		root[rootI] = rootJ
		size[rootJ] += size[rootI]
	} else {
		root[rootJ] = rootI
		size[rootI] += size[rootJ]
	}
}

func find(root map[int]int, i int) int {
	if root[i] != i {
		root[i] = find(root, root[i])
	}
	return root[i]
}
