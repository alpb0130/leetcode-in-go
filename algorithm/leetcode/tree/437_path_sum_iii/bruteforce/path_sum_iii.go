package bruteforce

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Master theorem: T(n) = a * T(n/b) + f(n)
// T(n) = n^(log_b^a) + f (compare to n^(log_b^a))
// Here, f(n) = n (the # of map item is the # of nodes in sub tree). T(n) = nlogn.
// Time: O(nlogn) - I think it's not n^2
// Space: O((logn)^2)
func pathSum(root *TreeNode, sum int) int {
	num, _ := pathSumHelper(root, sum)
	return num
}

func pathSumHelper(root *TreeNode, sum int) (int, map[int]int) {
	if root == nil {
		return 0, map[int]int{}
	}
	leftNum, leftCnt := pathSumHelper(root.Left, sum)
	rightNum, rightCnt := pathSumHelper(root.Right, sum)
	sumNum := leftNum + rightNum
	sumNum += (leftCnt[sum-root.Val] + rightCnt[sum-root.Val])
	if root.Val == sum {
		sumNum++
	}
	newCnt := map[int]int{root.Val: 1}
	for v, c := range leftCnt {
		newCnt[root.Val+v] += c
	}
	for v, c := range rightCnt {
		newCnt[root.Val+v] += c
	}
	return sumNum, newCnt
}
