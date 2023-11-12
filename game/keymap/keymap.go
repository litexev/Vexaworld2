/* Actions and default key bindings */
package keymap

import (
	input "github.com/quasilyte/ebitengine-input"
)

const (
	ActionMoveLeft input.Action = iota
	ActionMoveRight
	ActionJump
	ActionToggleFPS
	ActionPlace
	ActionModifierDestroy
	ActionCopyBlock
)

var Keymap = input.Keymap{
	ActionMoveLeft:        {input.KeyGamepadLeft, input.KeyLeft, input.KeyA},
	ActionMoveRight:       {input.KeyGamepadRight, input.KeyRight, input.KeyD},
	ActionJump:            {input.KeyGamepadUp, input.KeyUp, input.KeySpace, input.KeyW},
	ActionToggleFPS:       {input.KeyF},
	ActionPlace:           {input.KeyMouseLeft},
	ActionModifierDestroy: {input.KeyControlLeft, input.KeyControlRight},
	ActionCopyBlock:       {input.KeyMouseMiddle},
}
