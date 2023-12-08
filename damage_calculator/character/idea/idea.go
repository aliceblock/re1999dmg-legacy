package idea

type Idea struct {
	Atk        float64 `json:"atk"`
	AtkPercent float64 `json:"atkPercent"`
	DmgBonus   float64 `json:"dmgBonus"`
	CritRate   float64 `json:"critRate"`
	CritDmg    float64 `json:"critDmg"`
}

var BaseZIdea = Idea{
	Atk:      141,
	CritDmg:  0.06,
	DmgBonus: 0.02,
}

// var BaseUIdea = Idea{
// 	Atk:      142,
// 	DmgBonus: 0.06,
// }

var RegulusBaseIdea = Idea{
	Atk:      142,
	DmgBonus: 0.06,
}

// var BasePlusIdea = Idea{
// 	Atk:      142,
// 	DmgBonus: 0.06,
// }

var C1Idea = Idea{
	DmgBonus: 0.005,
}

var C2Idea = Idea{
	CritDmg: 0.03,
}

var C3Idea = Idea{
	AtkPercent: 0.015,
	CritRate:   0.025,
}

var C4IIdea = Idea{
	AtkPercent: 0.03,
}

var C4TIdea = Idea{
	AtkPercent: 0.03,
}

var C4OIdea = Idea{
	AtkPercent: 0.03,
}

var C4SIdea = Idea{
	CritRate: 0.025,
	DmgBonus: 0.03,
}

var C4LIdea = Idea{
	CritRate: 0.035,
	DmgBonus: 0.025,
}

var C4JIdea = Idea{
	CritRate: 0.06,
	CritDmg:  0.035,
}
