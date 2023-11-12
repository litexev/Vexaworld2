/* Vexaworld Launcher */
package main

import (
	"embed"
	"image"
	"vexaworld/game"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assetFS embed.FS

func main() {
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Vexaworld 2")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	icon24, _ := assetFS.Open("assets/icon24.png")
	icon36, _ := assetFS.Open("assets/icon36.png")
	img24, _, _ := image.Decode(icon24)
	img36, _, _ := image.Decode(icon36)
	ebiten.SetWindowIcon([]image.Image{img24, img36})

	game.StartGame(&assetFS)
}
