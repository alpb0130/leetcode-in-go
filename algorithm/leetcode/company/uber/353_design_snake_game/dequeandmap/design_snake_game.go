package dequeandmap

type SnakeGame struct {
	body         *deque
	bodyMap      map[int]bool
	food         [][]int
	foodPtr      int
	width        int
	height       int
	directionMap map[string][]int
}

/** Initialize your data structure here.
  @param width - screen width
  @param height - screen height
  @param food - A list of food positions
  E.g food = [[1,1], [1,0]] means the first food is positioned at [1,1], the second is at [1,0]. */
func Constructor(width int, height int, food [][]int) SnakeGame {
	return SnakeGame{
		body:    &deque{Pos{0, 0}},
		bodyMap: map[int]bool{0: true},
		food:    food,
		foodPtr: 0,
		width:   width,
		height:  height,
		directionMap: map[string][]int{
			"R": {0, 1},
			"L": {0, -1},
			"U": {-1, 0},
			"D": {1, 0},
		},
	}
}

// 1. Move head
// 2. Remove tail (lazy remove is better - only remove from map but not queue)
// 3. Valid - boundary and body check
// 4. Add new head
// 5. check food. If food, add back tail. Else remove tail.
// Time: O(n) - because of offer first. We can improve that by not poll tail immediately
// Space: O(n)
/** Moves the snake.
  @param direction - 'U' = Up, 'L' = Left, 'R' = Right, 'D' = Down
  @return The game's score after the move. Return -1 if game over.
  Game over when snake crosses the screen boundary or bites its body. */
func (this *SnakeGame) Move(direction string) int {
	dir := this.directionMap[direction]
	// Get new head and poll tail
	head := this.body.Last()
	i, j := head.i+dir[0], head.j+dir[1]
	index := i*this.width + j
	tail := this.body.PollFirst()
	tailIndex := tail.i*this.width + tail.j
	delete(this.bodyMap, tailIndex)
	// Valid
	if i < 0 || i >= this.height || j < 0 || j >= this.width || this.bodyMap[index] {
		return -1
	}
	// If valid, put new head to body
	this.body.OfferLast(Pos{i, j})
	this.bodyMap[index] = true
	// Eat. If food, we add tail back to body
	if this.foodPtr < len(this.food) && i == this.food[this.foodPtr][0] && j == this.food[this.foodPtr][1] {
		this.foodPtr++
		this.body.OfferFirst(tail)
		this.bodyMap[tailIndex] = true
	}
	return this.body.Len() - 1
}

type Pos struct {
	i int
	j int
}

type deque []Pos

func (q *deque) Len() int {
	return len(*q)
}
func (q *deque) OfferLast(p Pos) {
	*q = append(*q, p)
}
func (q *deque) PollFirst() Pos {
	p := (*q)[0]
	*q = (*q)[1:]
	return p
}
func (q *deque) OfferFirst(p Pos) {
	*q = append([]Pos{p}, (*q)...)
}
func (q *deque) Last() Pos {
	return (*q)[len(*q)-1]
}

/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */
