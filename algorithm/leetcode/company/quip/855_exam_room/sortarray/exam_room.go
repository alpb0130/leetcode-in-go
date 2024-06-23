package sortarray

// This is the
type ExamRoom struct {
	seats []int
	size  int
}

// 1. Most naive idea: maintain an array with size N and mark it to 1 if the seat is occupied.
//    If we need to seat at some place, we just iterate over all seats.
// 2. The above method is slow if we have too many seats but only a few students. We can only
//    maintain an array for seats that are occupied. That would improve the time a lot. We we
//    need to seat a student. Just 1. Process the start. 2.Iterate over the middle seats. 3.
//    Process the end. Finally we can decide where to seat the student
// 3. Method 2 spend a lot of time on array operation, we can improve that by using double linked
//    list and maintain the order. (In java, we can also use treeset).
// 4. A method based on priority queue (O(logn))
//    https://leetcode.com/problems/exam-room/discuss/148595/Java-PriorityQueue-with-customized-object.-seat%3A-O(logn)-leave-O(n)-with-explanation
func Constructor(N int) ExamRoom {
	return ExamRoom{[]int{}, N}
}

// Time: O(s) - s is the # of seated students (occupied seats)
// Space: O(s) - insert into array is O(s) space
func (this *ExamRoom) Seat() int {
	if len(this.seats) == 0 {
		this.seats = append(this.seats, 0)
		return 0
	}
	seat := 0
	seatIndex := 0
	dist := 0
	// start. We can probably seat at 0. So we only compare with the first occupied seat.
	if this.seats[0] > dist {
		dist = this.seats[0]
		seat = 0
		seatIndex = 0
	}
	// mid
	prev := this.seats[0]
	for i := 1; i < len(this.seats); i++ {
		if (this.seats[i]-prev)/2 > dist {
			dist = (this.seats[i] - prev) / 2
			seat = prev + dist
			seatIndex = i
		}
		prev = this.seats[i]
	}
	// end. We can probably seat at the end. So we only compare with the last occupied seat.
	if this.size-1-this.seats[len(this.seats)-1] > dist {
		this.seats = append(this.seats, this.size-1)
		seat = this.size - 1
	} else {
		// insert mid
		right := make([]int, len(this.seats[seatIndex:]))
		copy(right, this.seats[seatIndex:])
		this.seats = append(this.seats[:seatIndex], seat)
		this.seats = append(this.seats, right...)
	}
	return seat
}

// Time: O(s) - s is the # of seated students (occupied seats)
// Space: O(s) - insert into array is O(n) space
func (this *ExamRoom) Leave(p int) {
	for i := 0; i < len(this.seats); i++ {
		if this.seats[i] == p {
			right := make([]int, len(this.seats[i+1:]))
			copy(right, this.seats[i+1:])
			this.seats = append(this.seats[:i], right...)
		}
	}
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(N);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
