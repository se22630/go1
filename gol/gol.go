package main

func neighbour(p golParams, alives []cell, y, x int) int {
	count := 0
	z := cell{x, y}
	var left, right, up, down int = 0, 0, 0, 0

	if x == 0 {
		left = p.imageWidth - 1
	} else {
		left = x - 1
	}
	if x == p.imageWidth-1 {
		right = 0
	} else {
		right = x + 1
	}
	if y == 0 {
		up = p.imageHeight - 1
	} else {
		up = y - 1
	}
	if y == p.imageHeight-1 {
		down = 0
	} else {
		down = y + 1
	}

	if cell.in(z, alives) {
	}
	if cell.in(cell{right, y}, alives) {
		count++
	}
	if cell.in(cell{right, down}, alives) {
		count++
	}
	if cell.in(cell{right, up}, alives) {
		count++
	}
	if cell.in(cell{left, y}, alives) {
		count++
	}
	if cell.in(cell{left, down}, alives) {
		count++
	}
	if cell.in(cell{left, up}, alives) {
		count++
	}
	if cell.in(cell{x, down}, alives) {
		count++
	}
	if cell.in(cell{x, up}, alives) {
		count++
	}
	return count
}

//{4 5} {5 6} {3 7} {4 7} {5 7} {15 15}
func calculateNextState(p golParams, world [][]byte) [][]byte {
	alives := calculateAliveCells(p, world)
	for i := 0; i < p.imageWidth; i++ {
		for j := 0; j < p.imageHeight; j++ {
			count := neighbour(p, alives, i, j)
			if world[i][j] == 255 {
				if count != 2 && count != 3 {
					world[i][j] = 0
				}
			} else {
				if count == 3 {
					world[i][j] = 255
				}
			}

		}
	}
	return world
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var alives []cell
	//world[6][4] = 255
	for i := 0; i < p.imageWidth; i++ {
		for j := 0; j < p.imageHeight; j++ {
			if world[i][j] == 255 {
				alives = append(alives, cell{j, i})
			}
		}
	}
	//alives = append(alives, cell{0, 15})
	//fmt.Println(alives)

	return alives
}
