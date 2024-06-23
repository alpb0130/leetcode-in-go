package recursivecomplex

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive postorder traverse. For each function call, return the value of
// found common ancestor, whether q is found and whether p is found. In callee,
// process those results from both right subtree and left subtree. If:
// 1. Left ancestor is not nil, return left ancestor, true, true
// 2. Right ancestor is not nil, return right ancestor, true, true
// 3. Get whether p/q is found in current tree (left and right ancestor are both nil):
// 	  3.1 If root.Val = p.Val, p is found
//    3.2 If root.Val = q/Val, q is found
//    3.3 If p and q are both found, return root as ancestor, true, true to callee
//    3.4 Return nil, pFound and qFound
// Time: O(n)
// Space: O(h)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res, _, _ := lowestCommonAncestorHelper(root, p, q)
	return res
}

func lowestCommonAncestorHelper(root, p, q *TreeNode) (*TreeNode, bool, bool) {
	if root == nil {
		return nil, false, false
	}
	leftAncestor, leftPFound, leftQFound := lowestCommonAncestorHelper(root.Left, p, q)
	rightAncestor, rightPFound, rightQFound := lowestCommonAncestorHelper(root.Right, p, q)
	pFound := leftPFound || rightPFound || root.Val == p.Val
	qFound := leftQFound || rightQFound || root.Val == q.Val
	var ancesstor *TreeNode
	if pFound && qFound {
		ancesstor = root
		if leftAncestor != nil {
			ancesstor = leftAncestor
		} else if rightAncestor != nil {
			ancesstor = rightAncestor
		}
	}
	return ancesstor, pFound, qFound
}
