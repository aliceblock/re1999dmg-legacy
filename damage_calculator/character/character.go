package character

type SkillIndex int16

const (
	Skill1       SkillIndex = 0
	Skill2       SkillIndex = 1
	Ultimate     SkillIndex = 2
	ExtraAction1 SkillIndex = 3
	ExtraAction2 SkillIndex = 4
	ExtraAction3 SkillIndex = 5
)

type SKillStarIndex int16

const (
	Star1 SKillStarIndex = 0
	Star2 SKillStarIndex = 1
	Star3 SKillStarIndex = 2
)

type CharacterInsightLevel int16

const (
	Insight2L50 CharacterInsightLevel = 0
	Insight3L1  CharacterInsightLevel = 1
	Insight3L60 CharacterInsightLevel = 2
)

type DamageType string

const (
	RealityDamage DamageType = "reality"
	MentalDamage  DamageType = "mental"
)

type Character struct {
	InsightLevel *CharacterInsightLevel
	damageType   DamageType
	insight      Insight
	stat         map[CharacterInsightLevel]Stat
	skill        map[SkillIndex][]Skill
}

func (c *Character) DamageType() DamageType {
	return c.damageType
}

func (c *Character) Atk() float64 {
	return c.stat[*c.InsightLevel].Atk
}

func (c *Character) CritRate() float64 {
	return c.stat[*c.InsightLevel].CritRate + c.Insight().CritRate
}

func (c *Character) CritDmg() float64 {
	return c.stat[*c.InsightLevel].CritDmg + c.Insight().CritDmg
}

func (c *Character) Insight() Insight {
	return c.insight
}

func (c *Character) Skills(skillIndex SkillIndex) []Skill {
	return c.skill[skillIndex]
}

func (c *Character) Skill(skillIndex SkillIndex, star SKillStarIndex) Skill {
	if skillIndex == Ultimate {
		return c.skill[skillIndex][Star1]
	}
	return c.skill[skillIndex][star]
}

func (c *Character) SetInsightLevel(insightLevel CharacterInsightLevel) {
	c.InsightLevel = &insightLevel
}

func MakeCharacter(insightLevel *CharacterInsightLevel, damageType DamageType, stat map[CharacterInsightLevel]Stat, insight Insight, skill map[SkillIndex][]Skill) *Character {
	char := new(Character)
	if insightLevel == nil {
		char.InsightLevel = new(CharacterInsightLevel)
		*char.InsightLevel = Insight3L60
	} else {
		char.InsightLevel = insightLevel
	}
	char.damageType = damageType
	char.stat = stat
	char.insight = insight
	char.skill = skill
	return char
}

type Insight struct {
	AtkPercent float64
	DmgBonus   float64
	CritRate   float64
	CritDmg    float64
}

type Stat struct {
	Atk      float64
	CritRate float64
	CritDmg  float64
}

type Skill struct {
	Multiplier      float64
	EnemyHit        int16
	ExtraAction     SkillIndex
	ExtraMultiplier []float64
}

var Regulus = MakeCharacter(
	nil,
	MentalDamage,
	map[CharacterInsightLevel]Stat{
		Insight2L50: {
			Atk:      1009.0,
			CritRate: 0.1556,
			CritDmg:  0.533,
		},
		Insight3L1: {
			Atk:      1046.0,
			CritRate: 0.172,
			CritDmg:  0.5575,
		},
		Insight3L60: {
			Atk:      1186.0,
			CritRate: 0.172,
			CritDmg:  0.5575,
		},
	},
	Insight{
		CritDmg: 0.15,
	},
	map[SkillIndex][]Skill{
		Skill1: {
			{Multiplier: 2.0, EnemyHit: 1},
			{Multiplier: 3.0, EnemyHit: 1},
			{Multiplier: 5.0, EnemyHit: 1},
		},
		Skill2: {
			{Multiplier: 1.5, EnemyHit: 2},
			{Multiplier: 1.75, EnemyHit: 2},
			{Multiplier: 2.75, EnemyHit: 2},
		},
		Ultimate: {
			{Multiplier: 3.0, EnemyHit: 4},
		},
	},
)

var AKnight = MakeCharacter(
	nil,
	RealityDamage,
	map[CharacterInsightLevel]Stat{
		Insight2L50: {
			Atk:      1000.0,
			CritRate: 0.1153,
			CritDmg:  0.4725,
		},
		Insight3L1: {
			Atk:      1037.0,
			CritRate: 0.1273,
			CritDmg:  0.4905,
		},
		Insight3L60: {
			Atk:      1176.0,
			CritRate: 0.1273,
			CritDmg:  0.4905,
		},
	},
	Insight{},
	map[SkillIndex][]Skill{
		Skill1: {
			{Multiplier: 1.8, EnemyHit: 1},
			{Multiplier: 2.5, EnemyHit: 1},
			{Multiplier: 4.5, EnemyHit: 1},
		},
		Skill2: {
			{Multiplier: 1.5, EnemyHit: 2},
			{Multiplier: 1.5, EnemyHit: 2},
			{Multiplier: 2.25, EnemyHit: 2},
		},
		Ultimate: {
			{Multiplier: 4.0, EnemyHit: 4},
		},
	},
)

var Lilya = MakeCharacter(
	nil,
	RealityDamage,
	map[CharacterInsightLevel]Stat{
		Insight2L50: {
			Atk:      944.0,
			CritRate: 0.1556,
			CritDmg:  0.533,
		},
		Insight3L1: {
			Atk:      979.0,
			CritRate: 0.172,
			CritDmg:  0.5575,
		},
		Insight3L60: {
			Atk:      1110.0,
			CritRate: 0.172,
			CritDmg:  0.5575,
		},
	},
	Insight{
		CritRate: 0.1,
	},
	map[SkillIndex][]Skill{
		Skill1: {
			{Multiplier: 1.6, EnemyHit: 1, ExtraAction: ExtraAction1},
			{Multiplier: 2.4, EnemyHit: 1, ExtraAction: ExtraAction1},
			{Multiplier: 4.0, EnemyHit: 1, ExtraAction: ExtraAction1},
		},
		Skill2: {
			{Multiplier: 1.5, EnemyHit: 2},
			{Multiplier: 2.25, EnemyHit: 2},
			{Multiplier: 3.75, EnemyHit: 2},
		},
		Ultimate: {
			{Multiplier: 7.0, EnemyHit: 1},
		},
		ExtraAction1: {
			{Multiplier: 0.8, EnemyHit: 1},
			{Multiplier: 1.2, EnemyHit: 1},
			{Multiplier: 2.0, EnemyHit: 1},
		},
	},
)

var Eagle = MakeCharacter(
	nil,
	RealityDamage,
	map[CharacterInsightLevel]Stat{
		Insight2L50: {
			Atk:      919.0,
			CritRate: 0.1486,
			CritDmg:  0.5225,
		},
		Insight3L1: {
			Atk:      919.0,
			CritRate: 0.1486,
			CritDmg:  0.5225,
		},
		Insight3L60: {
			Atk:      919.0,
			CritRate: 0.1486,
			CritDmg:  0.5225,
		},
	},
	Insight{
		CritDmg: 0.15,
	},
	map[SkillIndex][]Skill{
		Skill1: {
			{Multiplier: 1.8, EnemyHit: 1, ExtraMultiplier: []float64{0.4}},
			{Multiplier: 2.5, EnemyHit: 1, ExtraMultiplier: []float64{0.6}},
			{Multiplier: 4.5, EnemyHit: 1, ExtraMultiplier: []float64{1.0}},
		},
		Skill2: {
			{Multiplier: 1.2, EnemyHit: 2},
			{Multiplier: 1.8, EnemyHit: 2},
			{Multiplier: 3.0, EnemyHit: 2},
		},
		Ultimate: {
			{Multiplier: 4.0, EnemyHit: 4},
		},
	},
)

var Jessica = MakeCharacter(
	nil,
	RealityDamage,
	map[CharacterInsightLevel]Stat{
		Insight2L50: {
			Atk:      1143.0,
			CritRate: 0.127,
			CritDmg:  0.491,
		},
		Insight3L1: {
			Atk:      1143.0,
			CritRate: 0.127,
			CritDmg:  0.491,
		},
		Insight3L60: {
			Atk:      1143.0,
			CritRate: 0.127,
			CritDmg:  0.491,
		},
	},
	Insight{
		DmgBonus: 0.08,
	},
	map[SkillIndex][]Skill{
		Skill1: {
			{Multiplier: 1.8, EnemyHit: 1, ExtraMultiplier: []float64{0.4, 0.6, 0.8}},
			{Multiplier: 2.7, EnemyHit: 1, ExtraMultiplier: []float64{0.6, 0.9, 1.2}},
			{Multiplier: 4.5, EnemyHit: 1, ExtraMultiplier: []float64{1.0, 1.5, 2.0}},
		},
		Skill2: {
			{Multiplier: 1.35, EnemyHit: 2, ExtraMultiplier: []float64{0.3}},
			{Multiplier: 2.0, EnemyHit: 2, ExtraMultiplier: []float64{0.45}},
			{Multiplier: 3.35, EnemyHit: 2, ExtraMultiplier: []float64{0.75}},
		},
		Ultimate: {
			{Multiplier: 3.5, EnemyHit: 1},
		},
	},
)
var Changeling = Jessica
