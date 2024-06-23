package quickhasnext

type Vector2D struct {
	v [][]int
	i int
	j int
}

// Quick hasNext
// Time: O(n)
// Space: O(1)
func Constructor(v [][]int) Vector2D {
	i, j := 0, 0
	for i < len(v) {
		if len(v[i]) != 0 {
			break
		}
		i++
	}

	return Vector2D{
		v: v,
		i: i,
		j: j,
	}
}

// Time: O(n)
// Space: O(1)
func (this *Vector2D) Next() int {
	x := this.v[this.i][this.j]
	this.j++
	for this.i < len(this.v) && this.j >= len(this.v[this.i]) {
		this.j = 0
		this.i++
	}
	return x
}

// Time: O(1)
// Space: O(1)
func (this *Vector2D) HasNext() bool {
	return this.i < len(this.v)
}

/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(v);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
