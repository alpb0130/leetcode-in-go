package nolock

import (
	"fmt"
)

type CallBack struct {
	callback string
}

func (c *CallBack) call() {
	fmt.Println(c.callback)
}

type EventFire struct {
	isFire bool
	queue  Queue
}

func (e *EventFire) register(c *CallBack) {
	if e.isFire {
		c.call()
	} else {
		e.queue.Offer(c)
		if e.isFire {
			callback := e.queue.Poll()
			callback.call()
		}
	}
}

// The key of solving this problem is we need to make sure if there are message in queue, we need
// to process them. If we don't set lock properly, we probably will left some message in queue
func (e *EventFire) Fire() {
	e.isFire = true
	for e.queue.Len() > 0 {
		callBack := e.queue.Poll()
		callBack.call()
	}

}

type Queue []*CallBack

func (q *Queue) Poll() *CallBack {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
func (q *Queue) Offer(c *CallBack) {
	*q = append(*q, c)
}
func (q *Queue) Len() int {
	return len(*q)
}
