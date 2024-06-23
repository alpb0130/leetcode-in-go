package topdown

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// We need to print out all the root-to-leaf path of a tree. Use top-down recursive solution.
// Take care of the use of string.Builder. Remember you cannot copy a Builder if it not empty.
// So always pass into a string.
// See: https://medium.com/@thuc/8-notes-about-strings-builder-in-golang-65260daae6e9
// Time: O(n)
// Space: O(h^2) - each recursive call need h extra memory.
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	binaryTreePathsHelper(root, "", &res)
	return res
}

func binaryTreePathsHelper(root *TreeNode, str string, res *[]string) {
	if root == nil {
		return
	}
	var strBuilder strings.Builder
	strBuilder.WriteString(str)
	strBuilder.WriteString(strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		*res = append(*res, strBuilder.String())
	}
	strBuilder.WriteString("->")
	binaryTreePathsHelper(root.Left, strBuilder.String(), res)
	binaryTreePathsHelper(root.Right, strBuilder.String(), res)
}
