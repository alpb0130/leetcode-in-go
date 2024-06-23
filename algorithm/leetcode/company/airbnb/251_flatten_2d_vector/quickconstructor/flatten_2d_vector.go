package quickconstructor

type Vector2D struct {
	v [][]int
	i int
	j int
}

// Quick constructor
func Constructor(v [][]int) Vector2D {
	return Vector2D{
		v: v,
		i: 0,
		j: 0,
	}
}

// Next call is always valid but we need to do some preprocess before we actually
// access the array because it's possible the current index is not valid. We can
// do the index process separately in Constructor, Next but it would be nice if
// we can consolidate the function into HasNext.
// Time: O(n)
// Space: O(1)
func (this *Vector2D) Next() int {
	this.HasNext()
	x := this.v[this.i][this.j]
	this.j++
	return x
}

// Time: O(n)
// Space: O(1)
func (this *Vector2D) HasNext() bool {
	for this.i < len(this.v) && this.j >= len(this.v[this.i]) {
		this.i++
		this.j = 0
	}
	return this.i < len(this.v)
}

/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(v);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
