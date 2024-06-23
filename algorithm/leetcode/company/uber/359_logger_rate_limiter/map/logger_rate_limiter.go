package _map

type Logger struct {
	messageTimeMap map[string]int
}

// Use a map to record every string and most recent print timestamp.
// If not print out, we don't update the time.
// Cons: if manay rare string appears long time ago and never show up again,
//       we will waste space
// Time: O(1)
// Space: O(n)
/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{map[string]int{}}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	time, ok := this.messageTimeMap[message]
	if !ok || timestamp-time >= 10 {
		this.messageTimeMap[message] = timestamp
		return true
	}
	return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
