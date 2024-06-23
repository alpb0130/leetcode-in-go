package dfs

import "fmt"

// This is the robot's control interface.
// You should not implement it, or speculate about its implementation
type Robot struct {
}

// Returns true if the cell in front is open and robot moves into the cell.
// Returns false if the cell in front is blocked and robot stays in the current cell.
func (robot *Robot) Move() bool { return false }

// Robot will stay in the same cell after calling TurnLeft/TurnRight.
// Each turn will be 90 degrees.
func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}

// Clean the current cell.
func (robot *Robot) Clean() {}

// Start with the original position and do dfs search. We need a map to record whether a
// position has been visited or not. The position is relative position.
// Time: O(m*n)
// Space: O(m*n)
func cleanRoom(robot *Robot) {
	isVisited := map[string]bool{}
	dfs(robot, 0, 0, -1, 0, isVisited)
}

func dfs(robot *Robot, i, j int, iDelta, jDelta int, isVisited map[string]bool) {
	isVisited[getHash(i, j)] = true
	robot.Clean()
	for k := 0; k < 4; k++ {
		newI, newJ := i+iDelta, j+jDelta
		if !isVisited[getHash(newI, newJ)] && robot.Move() {
			dfs(robot, newI, newJ, iDelta, jDelta, isVisited)
			// Go back to the original position and direction
			robot.TurnLeft()
			robot.TurnLeft()
			robot.Move()
			robot.TurnLeft()
			robot.TurnLeft()
		}
		robot.TurnLeft()
		iDelta, jDelta = -jDelta, iDelta
	}
}

func getHash(i, j int) string {
	return fmt.Sprintf("%d#%d", i, j)
}
