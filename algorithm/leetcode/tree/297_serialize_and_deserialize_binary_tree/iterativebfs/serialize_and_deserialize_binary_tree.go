package iterativebfs

import (
	"strconv"
	"strings"

	"algorithm/leetcode/tree/models"
)

func Encode(root *models.TreeNode) string {
	var str strings.Builder
	if root == nil {
		return "#"
	}
	q := &models.TreeNodeQueue{root}
	str.WriteString(strconv.Itoa(root.Val))
	str.WriteString(" ")
	for q.Len() != 0 {
		n := q.Poll()
		if n.Left != nil {
			q.Offer(n.Left)
			str.WriteString(strconv.Itoa(n.Left.Val))
			str.WriteString(" ")
		} else {
			str.WriteString("#")
			str.WriteString(" ")
		}
		if n.Right != nil {
			q.Offer(n.Right)
			str.WriteString(strconv.Itoa(n.Right.Val))
			str.WriteString(" ")
		} else {
			str.WriteString("#")
			if q.Len() > 0 {
				str.WriteString(" ")
			}
		}
	}
	return str.String()
}

func Decode(str string) *models.TreeNode {
	valStrs := strings.Split(str, " ")
	var root *models.TreeNode
	if valStrs[0] == "#" {
		return nil
	} else {
		val, _ := strconv.Atoi(valStrs[0])
		root = &models.TreeNode{Val: val}
	}
	q := &models.TreeNodeQueue{root}
	i := 1
	for i < len(valStrs) && q.Len() > 0 {
		node := q.Poll()
		if valStrs[i] != "#" {
			val, _ := strconv.Atoi(valStrs[i])
			node.Left = &models.TreeNode{Val: val}
			q.Offer(node.Left)
		}
		i++
		if i < len(valStrs) && valStrs[i] != "#" {
			val, _ := strconv.Atoi(valStrs[i])
			node.Right = &models.TreeNode{Val: val}
			q.Offer(node.Right)
		}
		i++
	}
	return root
}
