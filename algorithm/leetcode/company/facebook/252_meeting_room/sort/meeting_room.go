package sort

import "sort"

// The problem only requires us to check whether a person can attend every meeting.
// We can sort the compare the start time of current meeting with end time of prev meeting.
// Time: O(nlogn)
// Space: O(1)
func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) <= 1 {
		return true
	}
	sort.Sort(Intervals(intervals))
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

type Intervals [][]int

func (in Intervals) Len() int {
	return len(in)
}
func (in Intervals) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}
func (in Intervals) Less(i, j int) bool {
	return in[i][0] < in[j][0]
}
