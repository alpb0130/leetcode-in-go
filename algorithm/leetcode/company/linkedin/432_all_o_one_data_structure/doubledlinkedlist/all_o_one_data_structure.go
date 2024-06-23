package doubledlinkedlist

import "math"

// The idea here is to group key with the the same count into the same bucket and all buckets are
// connect with each other in sorted order using double linked list to provide O(1) add and delete
// operations. Because for a give count, we can return any key. No need to sort the keys.
// Reference: https://leetcode.com/problems/all-oone-data-structure/discuss/91416/Java-AC-all-strict-O(1)-not-average-O(1)-easy-to-read
type AllOne struct {
	head           *Bucket
	tail           *Bucket
	keyCountMap    map[string]int
	countBucketMap map[int]*Bucket
}

type Bucket struct {
	prev   *Bucket
	next   *Bucket
	keyMap map[string]bool
	count  int
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	// string -> count
	// count -> bucket
	// bucket link
	head := Bucket{
		count: math.MinInt32,
	}
	tail := Bucket{
		count: math.MaxInt32,
	}
	head.next = &tail
	tail.prev = &head
	return AllOne{
		head:           &head,
		tail:           &tail,
		keyCountMap:    map[string]int{},
		countBucketMap: map[int]*Bucket{},
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	_, ok := this.keyCountMap[key]
	if !ok {
		this.keyCountMap[key] = 1
		bucket, ok := this.countBucketMap[1]
		if !ok {
			bucket = &Bucket{
				keyMap: map[string]bool{},
				count:  1,
			}
			this.addBucketAfter(bucket, this.head)
			this.countBucketMap[1] = bucket
		}
		bucket.keyMap[key] = true
	} else {
		this.changeKey(key, 1)
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	cnt := this.keyCountMap[key]
	if cnt == 0 {
		return
	}
	if cnt == 1 {
		delete(this.keyCountMap, key)
		curBucket := this.countBucketMap[cnt]
		delete(curBucket.keyMap, key)
		if len(curBucket.keyMap) == 0 {
			delete(this.countBucketMap, 1)
			this.removeBucket(curBucket)
		}
	} else {
		this.changeKey(key, -1)
	}
}

func (this *AllOne) changeKey(key string, diff int) {
	cnt := this.keyCountMap[key]
	curBucket := this.countBucketMap[cnt]
	delete(curBucket.keyMap, key)
	if len(curBucket.keyMap) == 0 {
		delete(this.countBucketMap, cnt)
		this.removeBucket(curBucket)
		curBucket = curBucket.prev
	}
	cnt += diff
	//  if the original bucket is not deleted and diff=-1, curBucket will point to
	// bucket with old cnt. If the new cnt (cnt-1) is not in any bucket. We create a
	// a new bucket but insert it after curBucket which is incorrect.
	// And example operation is [inc:a, inc:a, inc:b, inc:b, dec:b, getmin, getmax]
	if curBucket.count > cnt {
		curBucket = curBucket.prev
	}
	// update key count
	this.keyCountMap[key] = cnt
	// update bucket
	newBucket, ok := this.countBucketMap[cnt]
	if !ok {
		newBucket = &Bucket{
			keyMap: map[string]bool{},
			count:  cnt,
		}
		// update countBucketMap if we add new bucket
		this.countBucketMap[cnt] = newBucket
		this.addBucketAfter(newBucket, curBucket)
	}
	newBucket.keyMap[key] = true
}

func (this *AllOne) removeBucket(bucket *Bucket) {
	pred := bucket.prev
	succ := bucket.next
	succ.prev = pred
	pred.next = succ
}

func (this *AllOne) addBucketAfter(bucket *Bucket, pred *Bucket) {
	bucket.prev = pred
	bucket.next = pred.next
	pred.next.prev = bucket
	pred.next = bucket
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.head.next == this.tail {
		return ""
	}
	keyMap := this.tail.prev.keyMap
	var res string
	for str, _ := range keyMap {
		res = str
		break
	}
	return res
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.head.next == this.tail {
		return ""
	}
	keyMap := this.head.next.keyMap

	var res string
	for str, _ := range keyMap {
		res = str
		break
	}
	return res
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
