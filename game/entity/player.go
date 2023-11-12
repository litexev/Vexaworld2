/* Vexaworld main player with controls */
package entity

import (
	"strconv"
	"vexaworld/game/consts"
	"vexaworld/game/context"
	"vexaworld/game/keymap"
	"vexaworld/game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	PhysEntity
	AvatarId  int
	imageIdle *ebiten.Image
	imageJump *ebiten.Image
	imageRun  *ebiten.Image
	state     int
	runClock  float64
	walkCycle bool
}

func CreatePlayer(world IWorld, x float64, y float64, avatarId int) *Player {
	return &Player{
		PhysEntity: *CreatePhysEntity(world, x, y, -1),
		AvatarId:   avatarId,
	}
}

const (
	PlayerStateIdle = iota
	PlayerStateJumping
	PlayerStateRunning
)

func (p *Player) Init(ctx *context.Context) {
	p.Entity.Init(ctx)
	p.SetAvatar(ctx, p.AvatarId)
	p.SetState(PlayerStateIdle)
}

func (p *Player) SetAvatar(ctx *context.Context, avatarId int) {
	p.AvatarId = avatarId
	p.imageIdle = ctx.Cache.GetImage("assets/avatar/" + strconv.Itoa(p.AvatarId) + "/main.png")
	p.imageJump = ctx.Cache.GetImage("assets/avatar/" + strconv.Itoa(p.AvatarId) + "/jump.png")
	p.imageRun = ctx.Cache.GetImage("assets/avatar/" + strconv.Itoa(p.AvatarId) + "/walk.png")
}
func (p *Player) Update(ctx *context.Context) {
	p.PhysEntity.Update(ctx)

	if p.state == PlayerStateJumping && p.Grounded {
		p.SetState(PlayerStateIdle)
	}

	// Walk cycle
	if p.state == PlayerStateRunning {
		if p.VelX != 0 {
			p.WalkCycle()
		} else {
			p.SetImage(p.imageIdle)
			p.walkCycle = false
		}
	}

	RunKeyPressed := false
	if ctx.Input.ActionIsPressed(keymap.ActionMoveLeft) {
		RunKeyPressed = true
		p.Flipped = true
		p.VelX = -consts.SIDE_VEL
		if p.Grounded {
			p.SetState(PlayerStateRunning)
		}
	}
	if ctx.Input.ActionIsPressed(keymap.ActionMoveRight) {
		RunKeyPressed = true
		p.Flipped = false
		p.VelX = consts.SIDE_VEL
		if p.Grounded {
			p.SetState(PlayerStateRunning)
		}
	}

	if ctx.Input.ActionIsPressed(keymap.ActionJump) {
		if p.Grounded {
			p.VelY = -consts.JUMP_VEL
			p.Grounded = false
			p.SetState(PlayerStateJumping)
		}
	}

	if p.state == PlayerStateRunning && !RunKeyPressed {
		p.SetState(PlayerStateIdle)
	}
}

func (p *Player) SetState(state int) {
	p.state = state
	switch state {
	case PlayerStateIdle:
		p.SetImage(p.imageIdle)
	case PlayerStateJumping:
		p.SetImage(p.imageJump)
	case PlayerStateRunning:
		if p.walkCycle {
			p.SetImage(p.imageRun)
		}
	}
}

func (p *Player) WalkCycle() {
	if utils.WaitForClock(&p.runClock, 6) {
		p.walkCycle = !p.walkCycle
		if p.walkCycle {
			p.SetImage(p.imageRun)
		} else {
			p.SetImage(p.imageIdle)
		}
	}
}
