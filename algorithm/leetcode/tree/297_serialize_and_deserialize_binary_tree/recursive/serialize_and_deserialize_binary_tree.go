package recursive

import (
	"strconv"
	"strings"

	"algorithm/leetcode/tree/models"
)

// Time: O(n)
// Space: O(n) - n + n/2 + n/4 + n/8... = n
func Encode(root *models.TreeNode) string {
	var str strings.Builder
	if root == nil {
		return "#"
	}
	str.WriteString(strconv.Itoa(root.Val))
	str.WriteString(" ")
	left := Encode(root.Left)
	str.WriteString(left)
	str.WriteString(" ")
	right := Encode(root.Right)
	str.WriteString(right)
	return str.String()
}

// Time: O(n)
// Space: O(h) - for each recursive call, constant memo use
func Decode(str string) *models.TreeNode {
	nodeStrs := strings.Split(str, " ")
	i := 0
	return decodeHelper(nodeStrs, &i)
}

func decodeHelper(strs []string, i *int) *models.TreeNode {
	if strs[*i] == "#" {
		*i++
		return nil
	}
	val, _ := strconv.Atoi(strs[*i])
	root := models.TreeNode{Val: val}
	*i++
	root.Left = decodeHelper(strs, i)
	root.Right = decodeHelper(strs, i)
	return &root
}
