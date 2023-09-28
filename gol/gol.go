package main

import "fmt"

func neighbour(p golParams, x, y int, world [][]byte) int {
	count := 0
	for i := -1; i < 1; i++ {
		for j := -1; j < 1; j++ {
			if i != 0 && j != 0 && x+i > -1 && x+i < p.imageWidth && y+j > -1 && y+j < p.imageHeight {
				if world[x+i][y+j] != 0 {
					fmt.Println(world[x+i][y+j])
				}
				if world[x+i][y+j] == 255 {
					count++
				}
			}
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
				fmt.Println(i, j)
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
