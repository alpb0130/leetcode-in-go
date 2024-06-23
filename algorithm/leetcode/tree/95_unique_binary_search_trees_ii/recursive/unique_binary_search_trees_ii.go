package recursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n*Catalan(n)) - there are Catalan(n) trees in the results and each results has n nodes
// Space: O(1) - except the func stack, we don't use any extra space
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateHelper(1, n)
}

func generateHelper(start int, end int) []*TreeNode {
	res := []*TreeNode{}
	if end-start < 0 {
		res := []*TreeNode{nil}
		return res
	}
	for i := start; i <= end; i++ {
		left := generateHelper(start, i-1)
		right := generateHelper(i+1, end)
		for _, leftChild := range left {
			for _, rightChild := range right {
				res = append(res, &TreeNode{i, leftChild, rightChild})
			}
		}
	}
	return res
}
