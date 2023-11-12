/* Vexaphys Physics Engine and Entity */
package entity

import (
	"math"
	"vexaworld/game/context"
	"vexaworld/game/utils"
)

type PhysEntity struct {
	Entity
	VelX     float64
	VelY     float64
	Grounded bool
	Gravity  float64
}

type Blocker struct {
	X int
	Y int
}

func CreatePhysEntity(world IWorld, x float64, y float64, imageId int) *PhysEntity {
	return &PhysEntity{
		Entity:   *CreateEntity(world, x, y, 1, true),
		VelX:     0,
		VelY:     0,
		Gravity:  world.GetGravity(),
		Grounded: false,
	}
}

func (e *PhysEntity) Update(ctx *context.Context) {
	newX := e.X
	newY := e.Y
	canMoveX := true
	canMoveY := true
	xBlocker := Blocker{}
	yBlocker := Blocker{}
	neighbors := [8][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	//fmt.Println("velocity:", e.VelX, e.VelY)

	// Slow existing velocity over time
	e.VelX *= 1 - (0.2)
	e.VelY = utils.Clamp(e.VelY, -1, 1)

	// Apply gravity
	//if !e.Grounded {
	e.VelY += e.Gravity
	//}

	// Calculate new position
	newX += e.VelX
	newY += e.VelY

	// Calculate grid position
	gridX := int(math.Round(e.X))
	gridY := int(math.Round(e.Y))

	// Collision check
	for _, neighbor := range neighbors {
		neighborX := gridX + neighbor[0]
		neighborY := gridY + neighbor[1]
		block := e.World.GetBlock(neighborX, neighborY)
		if block != 0 {
			blockedX := AABB(newX, e.Y, float64(neighborX), float64(neighborY))
			blockedY := AABB(e.X, newY, float64(neighborX), float64(neighborY))
			if blockedX {
				canMoveX = false
				xBlocker = Blocker{
					X: neighborX,
					Y: neighborY,
				}
			}
			if blockedY {
				canMoveY = false
				yBlocker = Blocker{
					X: neighborX,
					Y: neighborY,
				}
			}
		}
	}

	//fmt.Println("pos:", e.X, e.Y, "vel:", e.VelX, e.VelY, "canMove:", canMoveX, canMoveY)

	// Apply new position or snap to the colliding block
	if canMoveX {
		e.X = newX
	} else {
		e.VelX = 0
		// Snap to left / right
		if e.X < float64(xBlocker.X) {
			e.X = float64(xBlocker.X) - 1 // left
		} else {
			e.X = float64(xBlocker.X) + 1 // right
		}
	}
	if canMoveY {
		e.Y = newY
		e.Grounded = false
	} else {
		e.VelY = 0
		// Snap to top / bottom
		if e.Y < float64(yBlocker.Y) {
			e.Y = float64(yBlocker.Y) - 1 // top
		} else {
			e.Y = float64(yBlocker.Y) + 1 // bottom
		}
		// Grounded check
		if float64(yBlocker.Y) > e.Y {
			e.Grounded = true
		} else {
			e.Grounded = false
		}
	}
}

func AABB(blockX, blockY, newX, newY float64) bool {
	if blockX < newX+1 && blockX+1 > newX && blockY < newY+1 && blockY+1 > newY {
		return true
	}
	return false
}
