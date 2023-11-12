/* The image cache stores loaded images */
package cache

import (
	"io/fs"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ImageCache struct {
	fs     fs.FS
	images map[string]*ebiten.Image
	files  map[string]fs.File
	mu     sync.RWMutex
}

var Cache ImageCache

func CreateCache(fileSystem fs.FS) *ImageCache {
	return &ImageCache{
		fs:     fileSystem,
		images: make(map[string]*ebiten.Image),
		mu:     sync.RWMutex{},
	}
}

func (ic *ImageCache) GetImage(filePath string) *ebiten.Image {
	ic.mu.RLock()
	cachedImage, exists := ic.images[filePath]
	ic.mu.RUnlock()

	if exists {
		return cachedImage
	}

	image, _, err := ebitenutil.NewImageFromFileSystem(ic.fs, filePath)
	if err != nil {
		return ic.GetImage("assets/error.png")
	}

	ic.mu.Lock()
	ic.images[filePath] = image
	ic.mu.Unlock()

	return image
}

func (ic *ImageCache) GetFile(filePath string) fs.File {
	ic.mu.RLock()
	cachedFile, exists := ic.files[filePath]
	ic.mu.RUnlock()

	if exists {
		return cachedFile
	}

	file, err := ic.fs.Open(filePath)
	if err != nil {
		return nil
	}

	ic.mu.Lock()
	ic.files[filePath] = file
	ic.mu.Unlock()

	return file
}
