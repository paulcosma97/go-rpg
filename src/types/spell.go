package types

type Spell interface {
	Name() string
	SetName(string)

	Effects() []*Effect
	SetEffects([]*Effect)

	MinDamage() uint16
	SetMinDamage(uint16)

	MaxDamage() uint16
	SetMaxDamage(uint16)

	DamageType() *DamageType
	SetDamaype(*DamageType)

	Cooldown() uint8
	SetCooldown(uint8)

	MaxCooldown() uint8
	SetMaxCooldown(uint8)
}
