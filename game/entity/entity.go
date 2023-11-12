/* Base entity (drawing code is in render.go) */
package entity

import (
	"vexaworld/game/consts"
	"vexaworld/game/context"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	World        IWorld
	X            float64
	Y            float64
	Visible      bool
	shouldRender bool
	ImagePath    string
	Image        *ebiten.Image
	Opacity      float32
	imageW       int
	imageH       int
	Flipped      bool
}
type IWorld interface {
	GetBlock(x, y int) int
	GetGravity() float64
	SetBlock(x, y, block int)
}

func CreateEntity(world IWorld, x float64, y float64, opacity float32, visible bool) *Entity {
	return &Entity{
		World:        world,
		X:            x,
		Y:            y,
		Visible:      visible,
		shouldRender: false,
		Opacity:      opacity,
	}
}

func (e *Entity) Init(ctx *context.Context) {
	if e.Visible && e.Image == nil && e.ImagePath != "" {
		image := ctx.Cache.GetImage(e.ImagePath)
		e.SetImage(image)
	}
}

func (e *Entity) Update(ctx *context.Context) {
	// stub
}

func (e *Entity) SetImage(image *ebiten.Image) {
	e.Image = image
	e.imageW = e.Image.Bounds().Max.X
	e.imageH = e.Image.Bounds().Max.Y
	if !e.shouldRender {
		e.shouldRender = true
	}
}

func (e *Entity) Draw(ctx *context.Context, screen *ebiten.Image) {
	if !e.Visible || !e.shouldRender {
		return
	}
	op := &ebiten.DrawImageOptions{
		Filter: ebiten.FilterNearest,
	}
	drawX := int(e.X*consts.BLOCK_SIZE) + int(consts.BLOCK_SIZE-e.imageW) + ctx.ViewOffsetX
	drawY := int(e.Y*consts.BLOCK_SIZE) + int(consts.BLOCK_SIZE-e.imageH) + ctx.ViewOffsetY
	if e.Flipped {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(consts.BLOCK_SIZE, 0)
	}
	op.GeoM.Translate(float64(drawX), float64(drawY))
	if e.Opacity < 1 {
		op.ColorScale.ScaleAlpha(e.Opacity)
	}
	screen.DrawImage(e.Image, op)
}
