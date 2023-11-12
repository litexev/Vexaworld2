package entity

import (
	"math"
	"vexaworld/game/consts"
	"vexaworld/game/context"
	"vexaworld/game/keymap"

	"github.com/hajimehoshi/ebiten/v2"
)

type BlockPlacer struct {
	Entity
	imageBlock      *ebiten.Image
	imagePicker     *ebiten.Image
	BlockToPlace    int
	DestroyMode     bool
	lastX           int
	lastY           int
	changeNextFrame bool
}

func CreateBlockPlacer(world IWorld) *BlockPlacer {
	return &BlockPlacer{
		Entity:       *CreateEntity(world, 0, 0, 0.5, true),
		BlockToPlace: 3,
	}
}

func (b *BlockPlacer) Init(ctx *context.Context) {
	b.Entity.Init(ctx)
	b.imagePicker = ctx.Cache.GetImage("assets/picker.png")
	b.imageBlock = ctx.Cache.GetImage("assets/blocks/3.png")
	b.SetImage(b.imagePicker)
	b.shouldRender = true
}

func snapToGrid(num, size float64) int {
	return int(math.Floor(num / size))
}

func (b *BlockPlacer) Update(ctx *context.Context) {
	b.Entity.Update(ctx)
	mouseX, mouseY := ebiten.CursorPosition()
	placeKey := ctx.Input.ActionIsPressed(keymap.ActionPlace)
	destroyKey := ctx.Input.ActionIsPressed(keymap.ActionModifierDestroy)
	if destroyKey {
		ctx.CursorState = consts.CURSOR_DELETE
	} else if placeKey {
		ctx.CursorState = consts.CURSOR_PLACE
	} else {
		ctx.CursorState = consts.CURSOR_MAIN
	}
	x := snapToGrid(float64(mouseX-int(ctx.ViewOffsetX)), float64(consts.BLOCK_SIZE))
	y := snapToGrid(float64(mouseY-int(ctx.ViewOffsetY)), float64(consts.BLOCK_SIZE))
	if x == b.lastX && y == b.lastY && !ctx.Input.ActionIsJustPressed(keymap.ActionPlace) && !b.changeNextFrame {
		return
	}
	b.lastX, b.lastY = x, y
	b.changeNextFrame = false

	placeMode := b.World.GetBlock(x, y) == 0
	if placeMode {
		b.SetImage(b.imageBlock)
	} else {
		b.SetImage(b.imagePicker)
	}
	if placeKey && placeMode && !destroyKey {
		b.World.SetBlock(x, y, b.BlockToPlace)
		b.changeNextFrame = true
	}

	if placeKey && destroyKey {
		b.World.SetBlock(x, y, 0)
		b.changeNextFrame = true
	}

	b.X, b.Y = float64(x), float64(y)
}
