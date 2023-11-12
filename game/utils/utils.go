package utils

import (
	"math"
	"slices"
	"vexaworld/game/consts"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Clamp returns f clamped to [low, high]
func Clamp(f, low, high float64) float64 {
	if f < low {
		return low
	}
	if f > high {
		return high
	}
	return f
}

func SnapToGrid(number, step int) int {
	return int(math.Floor(float64(number+(step/2)) / float64(step)))
}

func Lerp(a, b float64, t float64) float64 {
	return a + t*(b-a)
}

func EaseLerp(a, b float64, t float64) float64 {
	t = t*2 - 1
	if t < 0 {
		return a + (b-a)*(t*t+1)/2
	} else {
		t = t * 2
		return a + (b-a)*(1-t*t)/2
	}
}

func Positive(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func BlockPosToChunkPos(x, y int) (int, int) {
	chunkX := x / consts.CHUNK_SIZE
	chunkY := y / consts.CHUNK_SIZE

	// Handle negative coordinates by rounding toward negative infinity
	if x < 0 && x%consts.CHUNK_SIZE != 0 {
		chunkX--
	}
	if y < 0 && y%consts.CHUNK_SIZE != 0 {
		chunkY--
	}

	return chunkX, chunkY
}

func BlockPosToLocalPos(x, y int) (int, int) {
	blockX := Positive(x % consts.CHUNK_SIZE)
	blockY := Positive(y % consts.CHUNK_SIZE)
	if x < 0 && blockX != 0 {
		blockX = consts.CHUNK_SIZE - blockX
	}
	if y < 0 && blockY != 0 {
		blockY = consts.CHUNK_SIZE - blockY
	}
	return blockX, blockY
}

func GetNeighbors(radius int) [][2]int {
	var neighbors [][2]int
	for i := -radius; i <= radius; i++ {
		for j := -radius; j <= radius; j++ {
			// Skip the center point
			if i == 0 && j == 0 {
				continue
			}
			neighbors = append(neighbors, [2]int{i, j})
		}
	}
	return neighbors
}

// thx: https://stackoverflow.com/a/72484396
func SliceDelete[T comparable](collection []T, el T) []T {
	idx := SliceFind(collection, el)
	if idx > -1 {
		return slices.Delete(collection, idx, idx+1)
	}
	return collection
}

func SliceFind[T comparable](collection []T, el T) int {
	for i := range collection {
		if collection[i] == el {
			return i
		}
	}
	return -1
}

func WaitForClock(clockVar *float64, ticks int) bool {
	*clockVar += 1
	if *clockVar < float64(ticks) {
		return false
	}
	*clockVar = 0
	return true
}
