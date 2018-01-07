package main

type direction int

const (
	left direction = iota
	up
	right
	down
)

func buildGrid3a(goal int) (int, int) {
	var x, y, minX, minY, maxX, maxY int

	dir := left
	for i := 1; i < goal; i++ {
		switch dir {
		case left:
			x++
			if x > maxX {
				dir = up
				maxX = x
			}
			break
		case up:
			y++
			if y > maxY {
				dir = right
				maxY = y
			}
			break
		case right:
			x--
			if x < minX {
				dir = down
				minX = x
			}
			break
		case down:
			y--
			if y < minY {
				dir = left
				minY = y
			}
			break
		}
	}
	return x, y
}

func buildGrid3b(goal, max int) int {
	x, y, minX, minY, maxX, maxY := 300, 300, 300, 300, 300, 300

	grid := [600][600]int{}

	grid[x][y] = 1

	dir := left
	for i := 1; i < goal; i++ {
		switch dir {
		case left:
			x++
			if x > maxX {
				dir = up
				maxX = x
			}
			break
		case up:
			y++
			if y > maxY {
				dir = right
				maxY = y
			}
			break
		case right:
			x--
			if x < minX {
				dir = down
				minX = x
			}
			break
		case down:
			y--
			if y < minY {
				dir = left
				minY = y
			}
			break
		}

		v := getValue2DArray(grid, x, y)
		grid[x][y] = v
		if v > max {
			return v
		}
	}
	return grid[x][y]
}

func getValue2DArray(grid [600][600]int, x, y int) int {
	result := 0

	result += grid[x+1][y]
	result += grid[x+1][y+1]
	result += grid[x][y+1]
	result += grid[x-1][y+1]
	result += grid[x-1][y]
	result += grid[x-1][y-1]
	result += grid[x][y-1]
	result += grid[x+1][y-1]

	return result
}
