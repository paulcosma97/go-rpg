type Spell interface {
	GetName() string
	SetName(v string)

	GetMana() string
	SetMana(v string)

	GetEffects() []effect.Effect
	SetEffects(v []effect.Effect)

	GetMinDamage() uint16
	SetMinDamage(v uint16)

	GetMaxDamage() uint16
	SetMaxDamage(v uint16)

	GetCritRate() float32
	SetCritRate(v float32)

	GetCritMultiplier() float32
	SetCritMultiplier(v float32)

	GetHitChance() float32
	SetHitChance(v float32)
}