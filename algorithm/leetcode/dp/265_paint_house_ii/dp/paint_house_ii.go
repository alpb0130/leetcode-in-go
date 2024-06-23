package dp

import "math"

// Similar to Paint House 1 but 2 requires that there are k different colors. The straightforward
// way is like 1. Compute the min cost for house k using every color and when we compute the same
// thing for house k+1, for color l, we get the min cost using l for house k+1 + cost of previous
// k house and using non-l color for house k. We only need to keep the top 2 cheapest cost color
// and we can finish that.
// Time: O(nk)
// Space: O(1)
func minCostII(costs [][]int) int {
	min1 := colorCost{-1, 0}
	min2 := colorCost{-1, 0}
	for _, cost := range costs {
		curMin1 := createColorCost()
		curMin2 := createColorCost()
		for i, c := range cost {
			curCost := 0
			if i != min1.color {
				curCost = c + min1.cost
			} else {
				curCost = c + min2.cost
			}
			if curCost < curMin2.cost {
				curMin2.color = i
				curMin2.cost = curCost
			}
			if curCost < curMin1.cost {
				curMin2.color = curMin1.color
				curMin2.cost = curMin1.cost
				curMin1.color = i
				curMin1.cost = curCost
			}
		}
		min1, min2 = curMin1, curMin2
	}
	return min1.cost
}

type colorCost struct {
	color int
	cost  int
}

func createColorCost() colorCost {
	return colorCost{-1, math.MaxInt32}
}
