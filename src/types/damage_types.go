package types

type DamageType = string

type _DamageTypes struct {
	Magical      DamageType
	Physical     DamageType
	PurePhysical DamageType
	PureMagical  DamageType
}

var DamageTypes _DamageTypes = _DamageTypes{
	Magical:      "Magical",
	Physical:     "Physical",
	PurePhysical: "PurePhysical",
	PureMagical:  "PureMagical",
}
