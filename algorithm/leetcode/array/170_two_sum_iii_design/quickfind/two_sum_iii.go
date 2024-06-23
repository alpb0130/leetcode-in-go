package quickfind

type TwoSum struct {
	numMap map[int]bool
	sumMap map[int]bool
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{
		numMap: map[int]bool{},
		sumMap: map[int]bool{},
	}
}

// Time: O(n)
/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	for num, _ := range this.numMap {
		this.sumMap[num+number] = true
	}
	this.numMap[number] = true
}

// Time: O(1)
/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	return this.sumMap[value]
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
