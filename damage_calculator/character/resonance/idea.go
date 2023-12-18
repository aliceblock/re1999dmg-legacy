package resonance

type Idea struct {
	Atk        float64
	AtkPercent float64
	DmgBonus   float64
	CritRate   float64
	CritDmg    float64
}

var RegulusBaseIdea = Idea{
	Atk:      142,
	DmgBonus: 0.06,
}

var AKnightBaseIdea = Idea{
	Atk:      141,
	DmgBonus: 0.02,
	CritDmg:  0.06,
}

var LilyaBaseIdea = Idea{
	Atk:      133,
	CritRate: 0.06,
	CritDmg:  0.08,
}

var EagleBaseIdea = Idea{
	Atk:      110,
	CritRate: 0.05,
	CritDmg:  0.08,
}

var JessicaBaseIdea = Idea{
	Atk:      137,
	DmgBonus: 0.06,
}

var CharlieBaseIdea = Idea{
	Atk:      141,
	DmgBonus: 0.06,
}

var BkornblumeBaseIdea = Idea{
	Atk:      136,
	DmgBonus: 0.02,
}

var CenturionBaseIdea = Idea{
	Atk:      138,
	CritRate: 0.06,
	CritDmg:  0.08,
}

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

var C4IOTIdea = Idea{
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
