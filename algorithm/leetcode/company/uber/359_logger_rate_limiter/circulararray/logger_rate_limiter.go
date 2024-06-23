package circulararray

type Logger struct {
	time     [10]int
	messages [10]map[string]bool
}

/** Initialize your data structure here. */
func Constructor() Logger {
	messages := [10]map[string]bool{}
	for i := 0; i < 10; i++ {
		messages[i] = map[string]bool{}
	}
	return Logger{
		time:     [10]int{},
		messages: messages,
	}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	for i := 0; i < 10; i++ {
		if this.time[i] > timestamp-10 && this.messages[i][message] {
			return false
		}
	}
	if this.time[timestamp%10] != timestamp {
		this.time[timestamp%10] = timestamp
		this.messages[timestamp%10] = map[string]bool{message: true}
	} else {
		this.messages[timestamp%10][message] = true
	}
	return true

}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
