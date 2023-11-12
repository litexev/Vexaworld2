package render

import (
	"strconv"
	"vexaworld/game/consts"
	"vexaworld/game/context"

	"github.com/hajimehoshi/ebiten/v2"
)

func RenderBlock(ctx *context.Context, screen *ebiten.Image, x int, y int, block int) {
	imagePath := "assets/blocks/" + strconv.Itoa(block) + ".png"
	image := ctx.Cache.GetImage(imagePath)
	op := &ebiten.DrawImageOptions{
		Filter: ebiten.FilterNearest,
	}
	drawX := (x * consts.BLOCK_SIZE) + ctx.ViewOffsetX
	drawY := (y * consts.BLOCK_SIZE) + ctx.ViewOffsetY
	op.GeoM.Translate(float64(drawX), float64(drawY))
	screen.DrawImage(image, op)
}
