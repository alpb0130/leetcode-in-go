package smartinit

type MyCircularQueue struct {
	first  int
	last   int
	length int
	queue  []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
// All operation with O(1) time complexity, O(1) space complexity
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		// Initialization of first and last is a little bit tricky
		// In enqueue, we will only update last, so initialize it k-1 and first to 0
		// In dequeue, we will only update first, it will always point to the next first element
		// We can also initialize them to both -1, but need to do some special check
		first:  0,
		last:   k - 1,
		length: 0,
		queue:  make([]int, k),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	// Update the last index, which is initialized to k-1
	this.last = (this.last + 1) % cap(this.queue)
	this.queue[this.last] = value
	this.length++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	// Update the first index, which will point to the future first position
	this.first = (this.first + 1) % cap(this.queue)
	this.length--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	result := -1
	if this.length > 0 {
		result = this.queue[this.first]
	}
	return result
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	result := -1
	if this.length > 0 {
		result = this.queue[this.last]
	}
	return result
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.length == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.length == cap(this.queue)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
