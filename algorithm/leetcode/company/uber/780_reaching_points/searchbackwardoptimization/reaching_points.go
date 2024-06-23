package searchbackwardoptimization

// We can optimize the top down mathod by keep speed up the process of reduce [tx, ty]
// We know if tx > ty, the next target we could try is [tx-ty, ty]. But consider the case where
// tx >> ty, do we still need to do tx-ty for each step which is linear? The answer is no.
// Three cases:
// 1. if ty > sy, we know at some point ty = sx+sy so sx must < ty. Then we don't need to check
//    all tx which > ty by doing tx-ty. We can use tx = tx%ty to speed up the process
// 2. if ty == sy, we know ty cannot be changed so we can only keep subtracting tx by sy till
//    tx = sx. So we can directly get result by checking (tx-sx)/sy
// 3. if ty < sy, return false
// We also need to do the same things for case where ty > tx
// Time: O(log(max(tx, ty))) - because the algorithm is based on /
// Space: O(log(max(tx, ty))) - can be improved to O(1) by writing iterative function
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	if sx > tx || sy > ty {
		return false
	}
	if sx == tx && sy == ty {
		return true
	}
	if tx > ty {
		if ty > sy {
			tx %= ty
		} else {
			return (tx-sx)%ty == 0
		}
	} else {
		if tx > sx {
			ty %= tx
		} else {
			return (ty-sy)%tx == 0
		}
	}
	return reachingPoints(sx, sy, tx, ty)
}
