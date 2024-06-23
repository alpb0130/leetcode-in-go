package thread

import (
	"fmt"
	"sync"
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
	mutex  sync.Mutex
}

func (e *EventFire) register(c *CallBack) {
	//e.mutex.Lock()
	if e.isFire {
		// If put this unlock after call, there would be deadlock
		//e.mutex.Unlock()
		c.call()
	} else {
		e.queue.Offer(c)
		//e.mutex.Unlock()
	}
}

// The key of solving this problem is we need to make sure if there are message in queue, we need
// to process them. If we don't set lock properly, we probably will left some message in queue
// Variant 1: patch process. If we don't want to process the callback in queue 1 by 1, we
// can put all things in current queue into tmp queue and then release lock so that we don't block
// other register call.
func (e *EventFire) Fire() {
	//e.mutex.Lock()
	for e.queue.Len() > 0 {
		callBack := e.queue.Poll()
		//e.mutex.Unlock()
		callBack.call()
		//e.mutex.Lock()
	}
	// If the problem doesn't require the message in order, we can put this set before for loop
	e.isFire = true
	//e.mutex.Unlock()
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
