/* Procedural world generation */
package world

import (
	"vexaworld/game/chunk"
	"vexaworld/game/consts"
)

func WorldGen(x int, y int) int {
	// grass
	if y == 7 {
		return 1
	}
	if y > 7 {
		return 2
	}
	// blue blocks
	if x == 2 && y == 6 || x == 2 && y == 5 || x == 5 && y == 6 {
		return 4
	}
	return 0
}

func GenerateChunk(x int, y int) *chunk.Chunk {
	startX := x * consts.CHUNK_SIZE
	startY := y * consts.CHUNK_SIZE
	grid := [consts.CHUNK_SIZE][consts.CHUNK_SIZE]int{}
	for dx := 0; dx < consts.CHUNK_SIZE; dx++ {
		for dy := 0; dy < consts.CHUNK_SIZE; dy++ {
			grid[dx][dy] = WorldGen(startX+dx, startY+dy)
		}
	}
	return &chunk.Chunk{
		Blocks: grid,
		X:      x,
		Y:      y,
	}
}
