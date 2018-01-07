package main

import (
	"fmt"
)

func day14a(input string) int {
	count := 0
	for i := 0; i < 128; i++ {
		cur := fmt.Sprintf("%s-%d", input, i)
		hash := knotHash(cur)
		count += countSetBits(hash)
	}
	return count
}

func day14b(input string) int {
	d := disk{}

	for i := 0; i < 128; i++ {
		cur := fmt.Sprintf("%s-%d", input, i)
		d.addRow(i, knotHash(cur))
	}

	count := d.countGroups()

	return count
}

func countSetBits(bytes []byte) int {
	count := 0

	for _, b := range bytes {
		for b != 0 {
			count += int(b & 1)
			b >>= 1
		}
	}
	return count
}

type disk struct {
	grid  [128][128]bool
	group []diskLocation
}

type diskLocation struct {
	x int
	y int
}

func (d *disk) addRow(row int, bytes []byte) {
	for i := 0; i < 16; i++ {
		offset := i * 8
		d.grid[row][7+offset] = (bytes[i] & 1) == 1
		d.grid[row][6+offset] = (bytes[i] & 2) == 2
		d.grid[row][5+offset] = (bytes[i] & 4) == 4
		d.grid[row][4+offset] = (bytes[i] & 8) == 8
		d.grid[row][3+offset] = (bytes[i] & 16) == 16
		d.grid[row][2+offset] = (bytes[i] & 32) == 32
		d.grid[row][1+offset] = (bytes[i] & 64) == 64
		d.grid[row][0+offset] = (bytes[i] & 128) == 128
	}
}

func (d disk) print() {
	fmt.Println()
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if d.grid[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (d *disk) countGroups() int {
	count := 0
	for x := 0; x < 128; x++ {
		for y := 0; y < 128; y++ {
			if d.grid[x][y] {
				d.clearGroup(x, y)
				count++
			}
		}
	}
	return count
}

func (d *disk) clearGroup(x, y int) {
	d.group = []diskLocation{}

	d.findGroup(x, y)

	for _, l := range d.group {
		d.grid[l.x][l.y] = false
	}
}

func (d *disk) findGroup(x, y int) {
	if x < 0 || x > 127 || y < 0 || y > 127 {
		return
	}

	if !d.grid[x][y] {
		return
	}

	loc := diskLocation{x, y}

	for _, l := range d.group {
		if loc.x == l.x && loc.y == l.y {
			return
		}
	}

	d.group = append(d.group, loc)

	d.findGroup(x-1, y)
	d.findGroup(x+1, y)
	d.findGroup(x, y-1)
	d.findGroup(x, y+1)
}
