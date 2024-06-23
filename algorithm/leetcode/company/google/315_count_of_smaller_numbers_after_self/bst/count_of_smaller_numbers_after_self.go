package bst

// Use BST. The brute force solution is process every number from back to front. For every
// element, we process the number after itself and see which one is smaller than it and we
// can update the counter. The Time is O(n^2).
// We can use BST to improve that. Basically BST can help us traverse the necessary node.
// For some node that we already know which is smaller than val, we can skip it.
// As we process every number, we add it to the current BST. Also we need to record some
// important information. Like if we skip the left subtree when we add the node, we need to
// know how many numbers are in left subtree because they also contribute to the final result.
// The array also has duplicate numbers. So we need to record how many duplicate the current
// node has. So when we add a node, # of number in left subtree should contribute to the result.
// When we go right, root's count should also contribute to the results. When we go left, we also
// need to update root's left count.
// Time: O(nlogn) - O(n^2)
// Space: O(n) - BST
func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	root := &BSTNode{
		Val:       nums[len(nums)-1],
		Count:     1,
		LeftCount: 0,
	}
	res := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		// add node and return count
		res[i] = addNode(root, nums[i])
	}
	return res
}

type BSTNode struct {
	Val       int
	Count     int
	LeftCount int
	Left      *BSTNode
	Right     *BSTNode
}

func addNode(root *BSTNode, val int) int {
	res := 0
	if root.Val == val {
		root.Count++
		res += root.LeftCount
	} else if root.Val > val {
		root.LeftCount++
		if root.Left == nil {
			root.Left = &BSTNode{
				Val:       val,
				Count:     0,
				LeftCount: 0,
			}
		}
		res += addNode(root.Left, val)
	} else {
		res += root.Count
		res += root.LeftCount
		if root.Right == nil {
			root.Right = &BSTNode{
				Val:       val,
				Count:     0,
				LeftCount: 0,
			}
		}
		res += addNode(root.Right, val)
	}
	return res
}

// Add node iteratively
func addNodeIterative(root *BSTNode, val int) int {
	res := 0
	for root.Val != val {
		if root.Val < val {
			res += root.Count + root.LeftCount
			if root.Right == nil {
				root.Right = &BSTNode{
					Val:       val,
					Count:     0,
					LeftCount: 0,
				}
			}
			root = root.Right
		} else {
			root.LeftCount++
			if root.Left == nil {
				root.Left = &BSTNode{
					Val:       val,
					Count:     0,
					LeftCount: 0,
				}
			}
			root = root.Left
		}
	}
	root.Count++
	res += root.LeftCount
	return res
}
