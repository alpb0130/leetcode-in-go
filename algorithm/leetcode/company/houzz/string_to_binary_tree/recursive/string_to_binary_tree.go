package recursive

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

func StringToTree(str string) *TreeNode {
	if len(str) == 0 {
		return nil
	}
	if str[0] == '(' && str[len(str)-1] == ')' {
		str = str[1 : len(str)-1]
	}
	if len(str) == 1 {
		return &TreeNode{str, nil, nil}
	}
	underCnt := 0
	pCnt := 0
	rootIndex := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			pCnt++
		} else if str[i] == ')' {
			pCnt--
		} else if str[i] == '_' && pCnt == 0 {
			underCnt++
		}
		if underCnt == 1 {
			rootIndex = i + 1
			break
		}
	}
	root := &TreeNode{string(str[rootIndex]), nil, nil}
	root.Left = StringToTree(str[:rootIndex-1])
	root.Right = StringToTree(str[rootIndex+2:])
	return root
}
