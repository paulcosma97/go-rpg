package char

import "game/src/types"

type Character struct {
	Id              string           `json:"id"`
	Name            string           `json:"name"`
	MaxHealth       uint16           `json:"maxHealth"`
	Health          uint16           `json:"health"`
	MinDamage       uint16           `json:"minDamage"`
	MaxDamage       uint16           `json:"maxDamage"`
	DamageType      types.DamageType `json:"damageType"`
	Spells          []*types.Spell   `json:"spells"`
	Crit            float32          `json:"crit"`
	CritMultiplier  float32          `json:"critMultiplier"`
	HitChance       float32          `json:"hitChance"`
	Multistrike     float32          `json:"multistrike"`
	Block           float32          `json:"block"`
	Leech           float32          `json:"leech"`
	MagicalDefense  float32          `json:"magicalDefense"`
	PhysicalDefense float32          `json:"physicalDefense"`
}

func (c *Character) Id() string {
	return c.Id
}

func (c *Character) Name() string {
	return c.Name
}

func (c *Character) MinDamage() uint16 {
	return c.MinDamage
}

func (c *Character) SetMinDamage(v uint16) {
	c.MinDamage = v
}

func (c *Character) MaxDamage() uint16 {
	return c.MaxDamage
}

func (c *Character) SetMaxDamage(v uint16) {
	c.MaxDamage = v
}

func (c *Character) DamageType() types.DamageType {
	return c.DamageType
}

func (c *Character) SetDamaype(v types.DamageType) {
	c.DamageType = v
}

func (c *Character) MaxHealth() uint16 {
	return c.MaxHealth
}

func (c *Character) Health() uint16 {
	return c.Health
}

func (c *Character) SetName(n string) {
	c.Name = n
}

func (c *Character) SetMaxHealth(h uint16) {
	c.MaxHealth = h
}

func (c *Character) SetHealth(h uint16) {
	c.Health = h
}

func (c *Character) Spells() []*types.Spell {
	return c.Spells
}

func (c *Character) SetSpells(v []*types.Spell) {
	c.Spells = v
}

func (c *Character) Crit() float32 {
	return c.Crit
}

func (c *Character) SetCrit(v float32) {
	c.Crit = v
}

func (c *Character) CritMultiplier() float32 {
	return c.CritMultiplier
}

func (c *Character) SetCritMultiplier(v float32) {
	c.CritMultiplier = v
}

func (c *Character) HitChance() float32 {
	return c.HitChance
}

func (c *Character) SetHitChance(v float32) {
	c.HitChance = v
}

func (c *Character) Block() float32 {
	return c.Block
}

func (c *Character) SetBlock(v float32) {
	c.Block = v
}

func (c *Character) Leech() float32 {
	return c.Leech
}

func (c *Character) SetLeech(v float32) {
	c.Leech = v
}

func (c *Character) Multistrike() float32 {
	return c.Multistrike
}

func (c *Character) SetMultistrike(v float32) {
	c.Multistrike = v
}

func (c *Character) MagicalDefense() float32 {
	return c.MagicalDefense
}

func (c *Character) SetMagicalDefense(v float32) {
	c.MagicalDefense = v
}

func (c *Character) PhysicalDefense() float32 {
	return c.PhysicalDefense
}

func (c *Character) SetPhysicalDefense(v float32) {
	c.PhysicalDefense = v
}
