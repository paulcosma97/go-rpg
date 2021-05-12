package effect

import (
	"game/src/game/types"
)

type Effect struct {
	Type     types.EffectType
	Duration uint8

	OnApply  types.EffectEventHandler
	OnExpire types.EffectEventHandler
}

func (e *Effect) Type() types.EffectType {
	return e.Type
}

func (e *Effect) SetType(v types.EffectType) {
	e.Type = v
}

func (e *Effect) Duration() uint8 {
	return e.Duration
}

func (e *Effect) SetDuration(v uint8) {
	e.Duration = v
}

func (e *Effect) EventHandler_OnApply() types.EffectEventHandler {
	return e.OnApply
}

func (e *Effect) SetEventHandler_OnApply(v types.EffectEventHandler) {
	e.OnApply = v
}

func (e *Effect) EventHandler_OnExpire() types.EffectEventHandler {
	return e.OnExpire
}

func (e *Effect) SetEventHandler_OnExpire(v types.EffectEventHandler) {
	e.OnExpire = v
}
