package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use a map to store tree which are already computed. Use a clone function to copy
// a tree quickly with offset
// Time: O(1/2 * n*catalan(n)) - I am not quite sure
// Space: O()
func generateTrees(n int) []*TreeNode {
	res := []*TreeNode{}
	if n == 0 {
		return res
	}
	numTreeMap := map[int][]*TreeNode{
		0: []*TreeNode{nil},
	}
	for i := 1; i <= n; i++ {
		tempTreeNodeList := []*TreeNode{}
		for j := 1; j <= i; j++ {
			lefts := numTreeMap[j-1]
			rights := numTreeMap[i-j]
			for _, left := range lefts {
				for _, right := range rights {
					// clone left and right
					cloneLeft := left
					cloneRight := cloneTree(right, j)
					root := &TreeNode{j, cloneLeft, cloneRight}
					tempTreeNodeList = append(tempTreeNodeList, root)
				}
			}
		}
		numTreeMap[i] = tempTreeNodeList
	}
	return numTreeMap[n]
}

func cloneTree(root *TreeNode, diff int) *TreeNode {
	if root == nil {
		return nil
	}
	left := cloneTree(root.Left, diff)
	right := cloneTree(root.Right, diff)
	clonedRoot := &TreeNode{root.Val + diff, left, right}
	return clonedRoot
}
