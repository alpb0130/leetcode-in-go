package priorityqueue

import "container/heap"

type Interval struct {
	Start int
	End   int
}

// The idea is convert intervals to time point. Because the schedule for every employee have been
// sorted, after the conversion, we can get k sorted lists. Then using the k sorted lists, we can
// get free intervals for all employees.
// At which time, we know there is free interval? The answer is at current time, if all intervals
// have end and no interval start yet, it free. We use a counter to counting how many intervals is
// still alive. If we see a start, counter++. If we see an end, counter--. If counter == 0, it
// means not live interval at that time. So we note down this time as prev. Next time a new
// interval start, we compute the free interval which is [prev, newI.Start]
// We also have O(nlogn) method.
// 1. Extract every timestamp and sort. Then based on method above, we can get the results
// 2. Sort all intervals by start (then by end) and iterate. Every time if current end is
//    smaller than next start, we get a result. Otherwise, update current end.
// Time: O(nlogk) - n is the total number of intervals. k is the number of employees.
// Space: O(k)
func employeeFreeTime(schedule [][]Interval) []Interval {
	res := []Interval{}
	if len(schedule) == 0 {
		return res
	}
	// construct time struct
	processedSchedule := make([][]time, len(schedule))
	for i, s := range schedule {
		for j, interval := range s {
			processedSchedule[i] = append(processedSchedule[i], time{interval.Start, true, i, 2 * j})
			processedSchedule[i] = append(processedSchedule[i], time{interval.End, false, i, 2*j + 1})
		}
	}
	pq := &priorityQueue{}
	for _, s := range processedSchedule {
		if len(s) > 0 {
			heap.Push(pq, s[0])
		}
	}
	prev := -1
	counter := 0
	for pq.Len() > 0 {
		t := heap.Pop(pq).(time)
		if t.isStart {
			counter++
			if counter == 1 && prev != -1 {
				res = append(res, Interval{prev, t.timestamp})
			}
		} else {
			counter--
			if counter == 0 {
				prev = t.timestamp
			}
		}
		if t.j+1 < len(processedSchedule[t.i]) {
			heap.Push(pq, processedSchedule[t.i][t.j+1])
		}
	}
	return res
}

type time struct {
	timestamp int
	isStart   bool
	i         int
	j         int
}

type priorityQueue []time

func (q priorityQueue) Len() int {
	return len(q)
}
func (q priorityQueue) Less(i, j int) bool {
	if q[i].timestamp != q[j].timestamp {
		return q[i].timestamp < q[j].timestamp
	}
	return q[i].isStart
}
func (q priorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q *priorityQueue) Pop() interface{} {
	x := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return x
}
func (q *priorityQueue) Push(i interface{}) {
	*q = append(*q, i.(time))
}
