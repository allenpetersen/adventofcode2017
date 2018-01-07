package main

func day17a(skip int) int {
	list := make([]int, 2019, 2019)
	end := 0
	c := 0

	for i := 1; i <= 2017; i++ {
		end++
		c = (c + skip) % end
		if end-c > 0 {
			move(list, c, end-c)
		}
		c++
		list[c] = i
	}
	return list[c+1]
}

func move(list []int, pos, size int) {
	for i := pos + size + 1; i > pos; i-- {
		list[i] = list[i-1]
	}
}
