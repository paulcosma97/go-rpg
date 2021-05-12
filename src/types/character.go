package types

type Character interface {
	Id() string

	Name() string
	SetName(string)

	MaxHealth() uint16
	SetMaxHealth(uint16)

	Health() uint16
	SetHealth(uint16)

	Spells() []*Spell
	SetSpells([]*Spell)

	MinDamage() uint16
	SetMinDamage(uint16)

	MaxDamage() uint16
	SetMaxDamage(uint16)

	DamageType() DamageType
	SetDamaype(DamageType)

	Crit() float32
	SetCrit(float32)

	CritMultiplier() float32
	SetCritMultiplier(float32)

	Multistrike() float32
	SetMultistrike(float32)

	Block() float32
	SetBlock(float32)

	HitChance() float32
	SetHitChance(float32)

	Leech() float32
	SetLeech(float32)

	MagicalDefense() float32
	SetMagicalDefense(float32)

	PhysicalDefense() float32
	SetPhysicalDefense(float32)
}
