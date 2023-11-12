/* The world contains all chunks and entities */
package world

import (
	"image/color"
	"math"
	"vexaworld/game/chunk"
	"vexaworld/game/consts"
	"vexaworld/game/context"
	"vexaworld/game/entity"
	"vexaworld/game/keymap"
	"vexaworld/game/utils"
	"vexaworld/game/vis"

	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	Chunks      map[[2]int]*chunk.Chunk
	Entities    []IEntity
	BgColor     color.RGBA
	Gravity     float64
	Player      entity.Player
	BlockPlacer entity.BlockPlacer
	viewOffsetX float64
	viewOffsetY float64
}

type IEntity interface {
	Init()
	Update(ctx *context.Context)
	SetImageId(id int)
	Draw(screen *ebiten.Image)
}

var targetX float64 = 0
var targetY float64 = 0
var LimitFPS = true

func (w *World) Update(ctx *context.Context) {
	for _, entity := range w.Entities {
		entity.Update(ctx)
	}
	w.Player.Update(ctx)
	w.BlockPlacer.Update(ctx)
	w.UpdateViewOffset(ctx)
	w.LoadChunksAroundPlayer(2)

	// fps toggle
	if ctx.Input.ActionIsJustPressed(keymap.ActionToggleFPS) {
		LimitFPS = !LimitFPS
		if LimitFPS {
			ebiten.SetVsyncEnabled(false)
			ebiten.SetTPS(60)
		} else {
			ebiten.SetVsyncEnabled(false)
			ebiten.SetTPS(-1)
		}
	}
}

func (w *World) Draw(ctx *context.Context, screen *ebiten.Image) {
	screen.Fill(w.BgColor)
	viewVis := w.CalculateViewVis(ctx)
	for _, chunk := range w.Chunks {
		if !viewVis.Overlaps(chunk.GetVis()) {
			continue
		}
		chunk.Draw(ctx, screen, &viewVis)
	}
	for _, entity := range w.Entities {
		entity.Draw(screen)
	}
	w.Player.Draw(ctx, screen)
	w.BlockPlacer.Draw(ctx, screen)
}

func (w *World) GetBlock(x, y int) int {
	// fmt.Println("world GetBlock", x, y)
	chunk := w.GetChunkContainingPos(x, y)
	if chunk == nil {
		return 0
	}
	blockX, blockY := utils.BlockPosToLocalPos(x, y)
	return chunk.GetBlock(blockX, blockY)
}

func (w *World) SetBlock(x, y int, block int) {
	chunkX, chunkY := utils.BlockPosToChunkPos(x, y)
	chunk := w.GetChunk(chunkX, chunkY)
	if chunk == nil {
		return
	}
	blockX, blockY := utils.BlockPosToLocalPos(x, y)
	// fmt.Println("setblock", chunkX, chunkY, blockX, blockY)
	chunk.SetBlock(blockX, blockY, block)
}

func (w *World) GetChunkContainingPos(x, y int) *chunk.Chunk {
	chunkX, chunkY := utils.BlockPosToChunkPos(x, y)
	return w.GetChunk(chunkX, chunkY)
}

func (w *World) GetChunk(x, y int) *chunk.Chunk {
	chunk, exists := w.Chunks[[2]int{x, y}]
	if !exists {
		return nil
	}
	// re-enable this eventually
	// newChunk := w.GenerateChunk(x, y)
	// return newChunk
	return chunk
}

func (w *World) GenerateChunk(x, y int) *chunk.Chunk {
	if w.GetChunk(x, y) != nil {
		panic("trying to generate chunk that exists")
	}
	w.AddChunk(GenerateChunk(x, y))
	newChunk := w.GetChunk(x, y)
	// fmt.Println("generated chunk", newChunk.X, newChunk.Y)
	return newChunk
}
func (w *World) LoadChunksAroundPlayer(radius int) {
	curChunkX, curChunkY := utils.BlockPosToChunkPos(int(w.Player.X), int(w.Player.Y))
	neighbors := utils.GetNeighbors(radius)

	chunksToKeep := make(map[[2]int]struct{})
	chunksToKeep[[2]int{curChunkX, curChunkY}] = struct{}{}
	for _, neighbor := range neighbors {
		newX := curChunkX + neighbor[0]
		newY := curChunkY + neighbor[1]

		// Add neighboring chunk to the chunksToKeep map
		chunksToKeep[[2]int{newX, newY}] = struct{}{}

		chunk := w.GetChunk(newX, newY)
		if chunk == nil {
			w.GenerateChunk(newX, newY)
		}
	}

	// Unload chunks that are not in chunksToKeep
	for _, chunk := range w.Chunks {
		if _, keep := chunksToKeep[[2]int{chunk.X, chunk.Y}]; !keep {
			w.UnloadChunk(chunk.X, chunk.Y)
		}
	}
}

func (w *World) UnloadChunk(x, y int) {
	// delete(w.Chunks, [2]int{x, y})
}
func (w *World) AddChunk(chunk *chunk.Chunk) {
	w.Chunks[[2]int{chunk.X, chunk.Y}] = chunk
}
func (w *World) GetGravity() float64 {
	return w.Gravity
}

func (w *World) UpdateViewOffset(ctx *context.Context) {
	targetX = -w.Player.X*consts.BLOCK_SIZE + float64(ctx.Width)/4
	targetY = -w.Player.Y*consts.BLOCK_SIZE + float64(ctx.Height)/1.6
	w.viewOffsetX += (targetX - w.viewOffsetX) * 0.1
	w.viewOffsetY += (targetY - w.viewOffsetY) * 0.02
	ctx.ViewOffsetX = int(w.viewOffsetX)
	ctx.ViewOffsetY = int(w.viewOffsetY)
}

func (w *World) CalculateViewVis(ctx *context.Context) vis.BlockRange {
	startX := float64(-ctx.ViewOffsetX) / float64(consts.BLOCK_SIZE)
	startY := float64(-ctx.ViewOffsetY) / float64(consts.BLOCK_SIZE)
	endX := float64(ctx.Width-ctx.ViewOffsetX) / float64(consts.BLOCK_SIZE)
	endY := float64(ctx.Height-ctx.ViewOffsetY) / float64(consts.BLOCK_SIZE)
	//return int(math.Floor(startX)), int(math.Floor(startY)), int(math.Ceil(endX)), int(math.Ceil(endY))
	return vis.BlockRange{
		TopX:    int(math.Floor(startX)),
		TopY:    int(math.Floor(startY)),
		BottomX: int(math.Ceil(endX)),
		BottomY: int(math.Ceil(endY)),
	}
}

func CreateTestWorld(ctx *context.Context) *World {
	world := World{
		Chunks:   make(map[[2]int]*chunk.Chunk),
		Entities: []IEntity{},
		BgColor:  color.RGBA{27, 27, 33, 255},
		Gravity:  consts.DEFAULT_GRAVITY,
	}
	world.GenerateChunk(0, 0)

	player := entity.CreatePlayer(&world, 8, 3, 1)
	world.Player = *player
	world.Player.Init(ctx)

	BlockPlacer := entity.CreateBlockPlacer(&world)
	world.BlockPlacer = *BlockPlacer
	world.BlockPlacer.Init(ctx)

	return &world
}
