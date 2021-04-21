package spell

import (
	"game/src/game/spell/effect"
	tgame "game/src/game/types"
)

type Spell struct {
	Name           string
	Mana           string
	Effects        []effect.Effect
	MinDamage      uint16
	MaxDamage      uint16
	CritRate       float32
	CritMultiplier float32
	HitChance      float32
	OnCast         func(caster *tgame.Character, e *effect.Effect, target *tgame.Character)
}

func (x *Spell) GetName() string {
	return x.Name
}

func (x *Spell) SetName(v string) {
	x.Name = v
}

func (x *Spell) GetMana() string {
	return x.Mana
}

func (x *Spell) SetMana(v string) {
	x.Mana = v
}

func (x *Spell) GetEffects() []effect.Effect {
	return x.Effects
}

func (x *Spell) SetEffects(v []effect.Effect) {
	x.Effects = v
}

func (x *Spell) GetMinDamage() uint16 {
	return x.MinDamage
}

func (x *Spell) SetMinDamage(v uint16) {
	x.MinDamage = v
}

func (x *Spell) GetMaxDamage() uint16 {
	return x.MaxDamage
}

func (x *Spell) SetMaxDamage(v uint16) {
	x.MaxDamage = v
}

func (x *Spell) GetCritRate() float32 {
	return x.CritRate
}

func (x *Spell) SetCritRate(v float32) {
	x.CritRate = v
}

func (x *Spell) GetCritMultiplier() float32 {
	return x.CritMultiplier
}

func (x *Spell) SetCritMultiplier(v float32) {
	x.CritMultiplier = v
}

func (x *Spell) GetHitChance() float32 {
	return x.HitChance
}

func (x *Spell) SetHitChance(v float32) {
	x.HitChance = v
}
