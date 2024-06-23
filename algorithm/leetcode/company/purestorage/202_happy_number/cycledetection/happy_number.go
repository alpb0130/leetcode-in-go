package cycledetection

// No extra space. We use Floyd cycle detection to detect whether we could see cycle in the process
func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	slow := getNext(n)
	fast := getNext(getNext(n))
	for slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	if slow == 1 {
		return true
	}
	return false

}

func getNext(n int) int {
	m := 0
	for n != 0 {
		r := n % 10
		m += r * r
		n /= 10
	}
	return m
}
