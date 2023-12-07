package psychube

type Psychube struct {
	Atk              float64
	AtkPercent       float64
	DmgBonus         float64
	IncantationMight float64
	UltimateMight    float64
	CritRate         float64
	CritDmg          float64
	AdditionalEffect Stat
}

type Stat struct {
	Atk              float64
	AtkPercent       float64
	DmgBonus         float64
	IncantationMight float64
	UltimateMight    float64
	CritRate         float64
	CritDmg          float64
}

var ThunderousApplause = Psychube{
	Atk:      330,
	CritRate: 0.16,
	CritDmg:  0.16,
}

/*
For each enemy target defeated by the carrier, Ultimate Might +4% for the carrier. Stacks up to 4 times.
*/
var Hopscotch = Psychube{
	Atk:              370,
	IncantationMight: 0.18,
}

/*
After the carrier casts an Ultimate, Incantation Might of the next incantation +20%.
*/
var BraveNewWorld = Psychube{
	Atk:           370,
	UltimateMight: 0.18,
	AdditionalEffect: Stat{
		IncantationMight: 0.2,
	},
}

var HisBoundenDuty = Psychube{
	Atk:      410,
	DmgBonus: 0.12,
}
