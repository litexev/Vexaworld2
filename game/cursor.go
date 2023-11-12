package game

import (
	"vexaworld/game/consts"
	"vexaworld/game/context"

	"github.com/hajimehoshi/ebiten/v2"
)

var imgCursorMain *ebiten.Image
var imgCursorPlace *ebiten.Image
var imgCursorDelete *ebiten.Image
var cursorImg *ebiten.Image

func LoadCursorImages(ctx *context.Context) {
	imgCursorMain = ctx.Cache.GetImage("assets/cursor/main.png")
	imgCursorPlace = ctx.Cache.GetImage("assets/cursor/place.png")
	imgCursorDelete = ctx.Cache.GetImage("assets/cursor/delete.png")
}
func DrawCustomCursor(ctx *context.Context, screen *ebiten.Image) {
	switch ctx.CursorState {
	case consts.CURSOR_MAIN:
		cursorImg = imgCursorMain
	case consts.CURSOR_PLACE:
		cursorImg = imgCursorPlace
	case consts.CURSOR_DELETE:
		cursorImg = imgCursorDelete
	}
	op := &ebiten.DrawImageOptions{
		Filter: ebiten.FilterNearest,
	}
	cursorX, cursorY := ebiten.CursorPosition()
	op.GeoM.Translate(float64(cursorX), float64(cursorY))
	screen.DrawImage(cursorImg, op)
}
