package data

import (
	char "game/src/game/character"
	"game/src/game/spell"
	"game/src/types"
)

const (
	Character_SonicSamurai = "character_sonic_samurai"
	Spell_WarpStop         = "spell_warp_stop"
)

func NewCharacter_SonicSamurai(p *types.Player) *types.Character {
	sonicSamurai := &char.Character{
		Id:         Character_SonicSamurai,
		Name:       "Sonic Samurai",
		MaxHealth:  13000,
		Health:     13000,
		MinDamage:  550,
		MaxDamage:  750,
		DamageType: types.DamageTypes.Physical,
		Spells:     make([]*types.Spell, 5),
	}

	spellWarpStop := &spell.Spell{
		Name:        "Warp Stop",
		MaxCooldown: 5,
		Cooldown:    5,
		MinDamage:   700,
		MaxDamage:   900,
		DamageType:  &types.DamageTypes.Physical,
		Effects:     nil,
		OnCast: func(caster *Character, target *Character) {

		},
	}

	sonicSamurai.Spells[0] = spellWarpStop

	return out
}
