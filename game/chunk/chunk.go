/* Chunks divide up the world into tiles for efficiency */
package chunk

import (
	"vexaworld/game/consts"
	"vexaworld/game/context"
	"vexaworld/game/render"
	"vexaworld/game/vis"

	"github.com/hajimehoshi/ebiten/v2"
)

type Chunk struct {
	Blocks [consts.CHUNK_SIZE][consts.CHUNK_SIZE]int
	X      int
	Y      int
	New    bool
}

func (c *Chunk) Draw(ctx *context.Context, screen *ebiten.Image, viewVis *vis.BlockRange) {
	// draw chunk borders (debug)
	// vector.StrokeRect(screen, float32(c.X*consts.CHUNK_SIZE*consts.BLOCK_SIZE)+float32(ctx.ViewOffsetX), float32(c.Y*consts.CHUNK_SIZE*consts.BLOCK_SIZE)+float32(ctx.ViewOffsetY), consts.CHUNK_SIZE*consts.BLOCK_SIZE, consts.CHUNK_SIZE*consts.BLOCK_SIZE, 1, color.White, false)
	chunkStartX := c.X * consts.CHUNK_SIZE
	chunkStartY := c.Y * consts.CHUNK_SIZE
	for dx := 0; dx < consts.CHUNK_SIZE; dx++ {
		globalX := chunkStartX + dx
		if !viewVis.ContainsX(globalX) {
			continue
		}
		for dy := 0; dy < consts.CHUNK_SIZE; dy++ {
			block := c.Blocks[dx][dy]
			if block == 0 {
				continue
			}
			globalY := chunkStartY + dy
			if !viewVis.ContainsY(globalY) {
				continue
			}

			// fmt.Println("chunk.go: calling renderblock")
			render.RenderBlock(ctx, screen, globalX, globalY, block)
		}
	}
}

func (c *Chunk) GetBlock(x int, y int) int {
	// fmt.Println("chunk GetBlock", x, y)
	return c.Blocks[x][y]
}

func (c *Chunk) GetVis() vis.BlockRange {
	return vis.BlockRange{
		TopX:    c.X * consts.CHUNK_SIZE,
		TopY:    c.Y * consts.CHUNK_SIZE,
		BottomX: c.X*consts.CHUNK_SIZE + consts.CHUNK_SIZE,
		BottomY: c.Y*consts.CHUNK_SIZE + consts.CHUNK_SIZE,
	}
}

func (c *Chunk) SetBlock(x int, y int, block int) {
	c.Blocks[x][y] = block
}
