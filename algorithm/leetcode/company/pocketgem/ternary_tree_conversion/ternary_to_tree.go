package ternary_tree_conversion

import "strings"

type TreeNode struct {
	Val       string
	TrueNode  *TreeNode
	FalseNode *TreeNode
}

// "a? b:c"
func TernaryToTree(str string) *TreeNode {
	strings.Replace(str, " ", "", -1)
	if len(str) == 0 {
		return nil
	}
	i := 0
	return toTreeHelper(str, &i)
}

// Any space? If yes, replace all " " with ""
func toTreeHelper(str string, i *int) *TreeNode {
	if *i >= len(str) {
		return nil
	}
	val := string(str[*i])
	node := &TreeNode{
		Val: val,
	}
	*i++
	if *i < len(str) && str[*i] == '?' {
		*i++
		node.TrueNode = toTreeHelper(str, i)
		*i++
		node.FalseNode = toTreeHelper(str, i)
	}
	return node
}

func TreeToTernary(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var str strings.Builder
	str.WriteString(root.Val)
	if root.TrueNode != nil && root.FalseNode != nil {
		str.WriteString("?")
		str.WriteString(TreeToTernary(root.TrueNode))
		str.WriteString(":")
		str.WriteString(TreeToTernary(root.FalseNode))
	}
	return str.String()
}

func Encode(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	var str strings.Builder
	str.WriteString("{ Val: ")
	str.WriteString(root.Val + ", ")
	str.WriteString("TrueNode: " + Encode(root.TrueNode) + ", ")
	str.WriteString("FalseNode: " + Encode(root.FalseNode) + " }")
	return str.String()
}
