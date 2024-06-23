package mapandarray

import "math/rand"

type RandomizedSet struct {
	set    map[int]int
	values []int
}

// Keep a map and an array.
// Map is used to insert and review data in O(1) time (value to array idx)
// Array is used to get random
/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		set:    map[int]int{},
		values: []int{},
	}
}

// append the value to array
/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}
	this.values = append(this.values, val)
	this.set[val] = len(this.values) - 1

	return true
}

// 1. Remove it in map
// 2. Get the index in array and swap with the last element and remove it (I was struggling
// with it before)
/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.set[val]
	if !ok {
		return false
	}
	if idx != len(this.values)-1 {
		// Be careful here. If the value is the last one, don't do those.
		this.values[idx], this.values[len(this.values)-1] = this.values[len(this.values)-1], this.values[idx]
		this.set[this.values[idx]] = idx
	}
	this.values = this.values[:len(this.values)-1]
	delete(this.set, val)
	return true
}

// Use random to get a random value. Because of the method we take in Remove, all number should be
// with index less than length.
/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.values))
	return this.values[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
