package effect

import (
	tgame "game/src/game/character"
)

type EffectType = bool

var EffectHarm EffectType = false
var EffectHeal EffectType = true

type Effect struct {
	Type     EffectType
	Duration uint8

	OnApply  func(caster *tgame.Character, e *Effect, target *tgame.Character)
	OnExpire func(caster *tgame.Character, e *Effect, target *tgame.Character)
}
