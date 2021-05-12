package spell

import (
	"game/src/types"
)

type Spell struct {
	Name        string
	Effects     []*types.Effect
	MinDamage   uint16
	MaxDamage   uint16
	Cooldown    uint8
	MaxCooldown uint8
	DamageType  *types.DamageType
	OnCast      func(caster *types.Character, target *types.Character)
}

func (x *Spell) Name() string {
	return x.Name
}

func (x *Spell) SetName(v string) {
	x.Name = v
}

func (x *Spell) Effects() []*types.Effect {
	return x.Effects
}

func (x *Spell) SetEffects(v []*types.Effect) {
	x.Effects = v
}

func (x *Spell) MinDamage() uint16 {
	return x.MinDamage
}

func (x *Spell) SetMinDamage(v uint16) {
	x.MinDamage = v
}

func (x *Spell) MaxDamage() uint16 {
	return x.MaxDamage
}

func (x *Spell) SetMaxDamage(v uint16) {
	x.MaxDamage = v
}

func (x *Spell) DamageType() *types.DamageType {
	return x.DamageType
}

func (x *Spell) SetDamaype(v *types.DamageType) {
	x.DamageType = v
}

func (x *Spell) Cooldown() uint8 {
	return x.Cooldown
}

func (x *Spell) SetCooldown(v uint8) {
	x.Cooldown = v
}

func (x *Spell) MaxCooldown() uint8 {
	return x.MaxCooldown
}

func (x *Spell) SetMaxCooldown(v uint8) {
	x.MaxCooldown = v
}
