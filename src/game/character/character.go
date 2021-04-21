package char

type Character struct {
	Name      string `json:"name"`
	MaxHealth uint16 `json:"maxHealth"`
	Health    uint16 `json:"health"`
	MaxMana   uint16 `json:"maxMana"`
	Mana      uint16 `json:"mana"`
}

func (c *Character) GetName() string {
	return c.Name
}

func (c *Character) GetMaxHealth() uint16 {
	return c.MaxHealth
}

func (c *Character) GetHealth() uint16 {
	return c.Health
}

func (c *Character) GetMaxMana() uint16 {
	return c.MaxMana
}

func (c *Character) GetMana() uint16 {
	return c.Mana
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

func (c *Character) SetMaxMana(m uint16) {
	c.MaxMana = m
}

func (c *Character) SetMana(m uint16) {
	c.Mana = m
}
