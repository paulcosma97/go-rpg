package tgame

type Character interface {
	GetName() string
	GetMaxHealth() uint16
	GetHealth() uint16
	GetMaxMana() uint16
	GetMana() uint16

	SetName(n string)
	SetMaxHealth(h uint16)
	SetMaxMana(m uint16)
	SetHealth(h uint16)
	SetMana(m uint16)
}
