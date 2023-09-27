package main

func neighbour(p golParams, i, j int, world [][]byte) int {
	count := 0

	if i == 0 && j == 0 {
		return 0
	}

	if i != 0 {
		if world[i-1][j] == 255 {
			count++
		}
	}
	if i != p.imageWidth {
		if world[i+1][j] == 255 {
			count++
		}
	}
	if j != p.imageHeight {
		if world[i][j+1] == 255 {
			count++
		}
	}
	if j != 0 {
		if world[i][j-1] == 255 {
			count++
		}
	}

	return count
}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	var staticWorld = world
	for i := 0; i < p.imageWidth; i++ {
		for j := 0; j < p.imageHeight; j++ {
			liveNeighbour := neighbour(p, i, j, staticWorld)
			if staticWorld[i][j] == 0 {
				if liveNeighbour == 3 {
					world[i][j] = 255
				}
			}
			if staticWorld[i][j] == 255 {
				if liveNeighbour == 2 || liveNeighbour == 3 {
					world[i][j] = 255
				}
				if liveNeighbour < 2 || liveNeighbour > 3 {
					world[i][j] = 0
				}
			}

		}

	}
	return world

}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	return []cell{}
}
