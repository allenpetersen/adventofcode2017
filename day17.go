package main

func day17a(skip int) int {
	//list := make([]int, 50000002, 50000002)
	end := 0
	c := 0

	last := 0
	for i := 1; i <= 50000000; i++ {
		end++
		c = (c + skip) % end
		// fmt.Printf("i:%d c:%d\n", i, c)
		// if end-c > 0 {
		// 	move(list, c, end-c)
		// }
		// fmt.Println(list)
		if c == 0 {
			last = i
		}
		c++
		//list[c] = i
		//fmt.Printf("\ni:%d c:%d - %v\n", i, c, list)
	}
	//fmt.Println(list)
	//return list[1]
	return last
}

func move(list []int, pos, size int) {
	for i := pos + size + 1; i > pos; i-- {
		list[i] = list[i-1]
	}
}
