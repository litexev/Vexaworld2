package context

import (
	"io/fs"
	"vexaworld/game/cache"
	"vexaworld/game/keymap"

	input "github.com/quasilyte/ebitengine-input"
)

type Context struct {
	Input       *input.Handler
	Cache       *cache.ImageCache
	ViewOffsetX int
	ViewOffsetY int
	Width       int
	Height      int
}

func NewContext(fileSystem fs.FS, inputSystem *input.System) *Context {
	ctx := &Context{
		Cache: cache.CreateCache(fileSystem),
	}
	ctx.Input = inputSystem.NewHandler(0, keymap.Keymap)
	return ctx
}

func (c *Context) Update() {
	// stub
}
