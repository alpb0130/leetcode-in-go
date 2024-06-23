package searchbackward

// The straightforward way is to search from [sx, sy] to [tx, ty] but that takes us a lot of
// time because for each node, we have two options and the call could be exponential.
// But if we solve this problem top-down, we could get a better results. Because for each
// target point, we can only generate one source point because of the limitation where two nums
// are both larger than 0. So we can only always subtract larger number by smaller number.
// The time complexity could be improved to linear.
// This will lead to timeout. We can keep improve that. Please see the optimize method.
// Time: O(max(tx, ty))
// Space: O(max(tx, ty)) - can be improved to O(1) by writing iterative function
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	if tx <= 0 || ty <= 0 || tx < sx || ty < sy {
		return false
	}
	return reachingPoints(sx, sy, tx-ty, ty) || reachingPoints(sx, sy, tx, ty-tx)
}
