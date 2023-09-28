package main

func neighbour(p golParams, world [][]byte, x, y int) int {
	count := 0
	if x != 0 {
		if world[x-1][y] == 255 {
			count++
		}
		if y != 0 {
			if world[x-1][y-1] == 255 {
				count++
			}
		}
		if y != p.imageHeight-1 {
			if world[x-1][y+1] == 255 {
				count++
			}
		}
	}
	if x != p.imageWidth-1 {
		if world[x+1][y] == 255 {
			count++
		}
		if y != 0 {
			if world[x+1][y-1] == 255 {
				count++
			}
		}
		if y != p.imageHeight-1 {
			if world[x+1][y+1] == 255 {
				count++
			}
		}
	}
	if y != 0 {
		if world[x][y-1] == 255 {
			count++
		}
	}
	if y != p.imageHeight-1 {
		if world[x][y+1] == 255 {
			count++
		}
	}
	return count

}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	staticWorld := world
	for i := 0; i < p.imageWidth-1; i++ {
		for j := 0; j < p.imageHeight; j++ {
			lives := neighbour(p, staticWorld, i, j)
			if staticWorld[i][j] == 255 {
				if lives == 2 || lives == 3 {
					world[i][j] = 255
				}
				if lives < 2 || lives > 3 {
					world[i][j] = 0
				}
			}
			if staticWorld[i][j] == 0 {
				if lives == 3 {
					world[i][j] = 255
				}
			}
		}
	}
	return world

}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var alives []cell
	for i := 0; i < p.imageWidth; i++ {
		for j := 0; j < p.imageHeight; j++ {
			if world[i][j] == 255 {
				alives = append(alives, cell{j, i})
			}
		}
	}

	return alives
}
