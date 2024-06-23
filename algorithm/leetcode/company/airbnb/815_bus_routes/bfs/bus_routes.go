package bfs

// BFS. The problem requires us to find the min bus need to take from S to T. We can regard every
// stop as a vertex but the edge between them indicated whether they reachable by taking a the bus
// If two stop are belong to one bus, then they have an edge between them.
// If we want to find the min bus to take, we can just use the bfs to find the shortest path from
// S to T.
// One optimization we could do is record whether a bus has been visited instead of a stop. Because
// if a bus is visited, we don't need to visit all its stops.
// Time: O(m*n) + O(m*n*k*n) - m is the # bus, n is the average # stop for each bus and k is the
//       average # bus that would pass the stop. The complexity here is 1. process the routes
//       2. bfs. We need to visit all stops. All node we need to visit all buss. For all buses, we
//       need to visit all stops the bus pass.
// Space: O(m*n) - the map.
func numBusesToDestination(routes [][]int, S int, T int) int {
	stopToBusMap := map[int][]int{}
	for bus, route := range routes {
		for _, stop := range route {
			if _, ok := stopToBusMap[stop]; ok {
				stopToBusMap[stop] = append(stopToBusMap[stop], bus)
			} else {
				stopToBusMap[stop] = []int{bus}
			}
		}
	}

	q := &queue{S}
	// We should check whether a bus is visited yet. If it is, we don't need to visit
	// all stops belong to this bus which would save us some time.
	isBusVisited := map[int]bool{}
	isVisited := map[int]bool{S: true}
	res := 0
	if S == T {
		return res
	}
	for q.Len() > 0 {
		size := q.Len()
		res++
		for i := 0; i < size; i++ {
			stop := q.Poll()
			for _, bus := range stopToBusMap[stop] {
				// bus check
				if isBusVisited[bus] {
					continue
				}
				isBusVisited[bus] = true
				for _, nextStop := range routes[bus] {
					if nextStop == T {
						return res
					}
					// stop check
					if !isVisited[nextStop] {
						isVisited[nextStop] = true
						q.Offer(nextStop)
					}

				}
			}
		}
	}
	return -1
}

type queue []int

func (q *queue) Len() int {
	return len(*q)
}
func (q *queue) Offer(x int) {
	*q = append(*q, x)
}
func (q *queue) Poll() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
