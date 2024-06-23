package quickadd

type TwoSum struct {
	numMap map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{
		numMap: map[int]int{},
	}
}

// O(1)
/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.numMap[number]++
}

// O(n)
/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for n, _ := range this.numMap {
		cnt, ok := this.numMap[value-n]
		if (n == value-n && cnt > 1) || (n != value-n && ok) {
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
