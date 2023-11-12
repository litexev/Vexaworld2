package entity

import (
	"strconv"
	"vexaworld/game/consts"
	"vexaworld/game/context"
	"vexaworld/game/keymap"
	"vexaworld/game/utils"

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
	b.SetBlock(ctx, 3)
	b.SetImage(b.imagePicker)
	b.shouldRender = true
}

func (b *BlockPlacer) Update(ctx *context.Context) {
	b.Entity.Update(ctx)
	mouseX, mouseY := ebiten.CursorPosition()
	placeKey := ctx.Input.ActionIsPressed(keymap.ActionPlace)
	placeKeyFirstFrame := ctx.Input.ActionIsJustPressed(keymap.ActionPlace)
	destroyKey := ctx.Input.ActionIsPressed(keymap.ActionModifierDestroy)
	copyKey := ctx.Input.ActionIsJustPressed(keymap.ActionCopyBlock)
	if copyKey {
		b.changeNextFrame = true
	}
	if destroyKey {
		ctx.CursorState = consts.CURSOR_DELETE
	} else if placeKey {
		ctx.CursorState = consts.CURSOR_PLACE
	} else {
		ctx.CursorState = consts.CURSOR_MAIN
	}

	x := utils.SnapToGrid(mouseX-ctx.ViewOffsetX, consts.BLOCK_SIZE)
	y := utils.SnapToGrid(mouseY-ctx.ViewOffsetY, consts.BLOCK_SIZE)
	// middle click to copy block
	if x == b.lastX && y == b.lastY && !placeKeyFirstFrame && !b.changeNextFrame {
		return
	}
	b.lastX, b.lastY = x, y
	b.changeNextFrame = false

	block := b.World.GetBlock(x, y)
	placeMode := block == 0
	if placeMode && !destroyKey {
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

	if copyKey && !placeMode {
		b.SetBlock(ctx, block)
	}

	b.X, b.Y = float64(x), float64(y)
}

func (b *BlockPlacer) SetBlock(ctx *context.Context, block int) {
	b.BlockToPlace = block
	b.imageBlock = ctx.Cache.GetImage("assets/blocks/" + strconv.Itoa(block) + ".png")
}
