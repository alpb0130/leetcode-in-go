package valid_pick_up_and_delivery

import "strconv"

func IsValidPickupAndDelivery(orders []string) bool {
	if len(orders) == 0 {
		return true
	}
	if len(orders)%2 == 1 {
		return false
	}

	// Is duplicate id possible?
	orderMap := map[string]int{}
	for _, order := range orders {
		id := getID(order)
		if isPickup(order) {
			orderMap[id]++
		} else {
			if orderMap[id] == 0 {
				return false
			}
			orderMap[id]--
			if orderMap[id] == 0 {
				delete(orderMap, id)
			}
		}
	}
	return len(orderMap) == 0
}

func isPickup(str string) bool {
	return str[0] == 'P'
}

func getID(str string) string {
	return str[1:]
}

func AllValidPermutations(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}

	dp := make([][][]string, n+1)
	//dp[1] = [][]string{{""}}
	dp[1] = [][]string{{"P1", "D1"}}
	for i := 1; i <= n; i++ {
		for _, strs := range dp[i-1] {
			for j := 0; j <= len(strs); j++ {
				for k := j; k <= len(strs); k++ {
					tmp := []string{}
					tmp = append(tmp, strs[:j]...)
					tmp = append(tmp, "P"+strconv.Itoa(i))
					tmp = append(tmp, strs[j:k]...)
					tmp = append(tmp, "D"+strconv.Itoa(i))
					tmp = append(tmp, strs[k:]...)
					dp[i] = append(dp[i], tmp)
				}
			}
		}
	}
	return dp[n]
}
