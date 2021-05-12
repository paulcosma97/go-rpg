package types

type EffectType = bool

var EffectHarm EffectType = false
var EffectHeal EffectType = true

type EffectEventHandler = func(caster *Character, e *Effect, target *Character)

type Effect interface {
	SetType(EffectType)
	Type() EffectType
	
	SetDuration(uint8)
	Duration() uint8

	SetEventHandler_OnApply(EffectEventHandler)
	EventHandler_OnApply() EffectEventHandler

	SetEventHandler_OnExpire(EffectEventHandler)
	EventHandler_OnExpire() EffectEventHandler
}