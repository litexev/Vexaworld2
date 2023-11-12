/* Main game loop */
package game

import (
	"embed"
	"fmt"
	"log"
	"math"
	"strconv"

	"vexaworld/game/context"
	"vexaworld/game/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
)

var FileSystem embed.FS
var InputHandler *input.Handler
var CanvasWidth float64
var CanvasHeight float64
var WindowScale = 2.0

type Game struct {
	World       *world.World
	Context     *context.Context
	InputSystem input.System
	FileSystem  *embed.FS
}

func (g *Game) Update() error {
	g.InputSystem.Update()
	g.World.Update(g.Context)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.World.Draw(g.Context, screen)
	fps := int(math.Round(ebiten.ActualFPS()))
	ebitenutil.DebugPrint(screen, "V7 PROTO "+strconv.Itoa(fps)+"\n"+strconv.Itoa(len(g.World.Chunks))+" chunks; ")
	DrawCustomCursor(g.Context, screen)
}

func (g *Game) Layout(w, h int) (int, int) { panic("not implemented") }

func (g *Game) LayoutF(logicWinWidth, logicWinHeight float64) (float64, float64) {
	deviceScale := ebiten.DeviceScaleFactor()
	scaleFactor := math.Floor(ebiten.DeviceScaleFactor() * WindowScale)
	canvasWidth := math.Ceil(logicWinWidth * deviceScale / float64(scaleFactor))
	canvasHeight := math.Ceil(logicWinHeight * deviceScale / float64(scaleFactor))
	g.Context.Width, g.Context.Height = int(math.Round(canvasWidth)), int(math.Round(canvasHeight))
	return canvasWidth, canvasHeight
}

func StartGame(fs *embed.FS) {
	game := Game{}
	game.FileSystem = fs

	fmt.Println("init inputsystem")
	game.InputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	game.Context = context.NewContext(fs, &game.InputSystem)

	LoadCursorImages(game.Context)

	// load test world
	game.World = world.CreateTestWorld(game.Context)

	// debug unlimit framerate
	ebiten.SetVsyncEnabled(true)
	// ebiten.SetTPS(-1)

	// run game
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
