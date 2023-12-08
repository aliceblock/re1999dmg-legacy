package character

import (
	"math"

	"github.com/aliceblock/re1999dmg/damage_calculator/character/idea"
)

type Character struct {
	Atk        float64                `json:"atk"`
	AtkPercent float64                `json:"atkPercent"`
	DmgBonus   float64                `json:"dmgBonus"`
	CritRate   float64                `json:"critRate"`
	CritDmg    float64                `json:"critDmg"`
	Resonance  Resonance              `json:"resonance"`
	Skill      map[SkillIndex][]Skill `json:"skill"`
}

func (c *Character) GetResonanceStats() ResonanceStats {
	stats := ResonanceStats{}
	for _, ideaAmount := range c.Resonance.Ideas {
		if ideaAmount.Amount <= 0 {
			continue
		}
		stats.Atk += ideaAmount.Idea.Atk * float64(ideaAmount.Amount)
		stats.AtkPercent += ideaAmount.Idea.AtkPercent * float64(ideaAmount.Amount)
		stats.CritRate += ideaAmount.Idea.CritRate * float64(ideaAmount.Amount)
		stats.CritDmg += ideaAmount.Idea.CritDmg * float64(ideaAmount.Amount)
		stats.DmgBonus += ideaAmount.Idea.DmgBonus * float64(ideaAmount.Amount)
	}
	stats.Atk = toFixed(stats.Atk, 3)
	stats.AtkPercent = toFixed(stats.AtkPercent, 3)
	stats.DmgBonus = toFixed(stats.DmgBonus, 3)
	stats.CritRate = toFixed(stats.CritRate, 3)
	stats.CritDmg = toFixed(stats.CritDmg, 3)
	return stats
}

type Resonance struct {
	Ideas []IdeaAmount `json:"ideas"`
}

type IdeaAmount struct {
	Idea   idea.Idea `json:"idea"`
	Amount int16     `json:"amount"`
}

type ResonanceStats struct {
	Atk        float64 `json:"atk"`
	AtkPercent float64 `json:"atkPercent"`
	DmgBonus   float64 `json:"dmgBonus"`
	CritRate   float64 `json:"critRate"`
	CritDmg    float64 `json:"critDmg"`
}

type Skill struct {
	Multiplier float64 `json:"multiplier"`
}

type SkillIndex int16

const (
	Skill1   SkillIndex = 0
	Skill2              = 1
	Ultimate            = 2
)

var RegulusStat = Character{
	Atk:      1186.0,
	CritRate: 0.172,
	CritDmg:  0.558 + 0.15, // Base + Insight 2
	Resonance: Resonance{
		Ideas: []IdeaAmount{
			{Idea: idea.BaseTIdea, Amount: 1},
			{Idea: idea.C4LIdea, Amount: 2},
			{Idea: idea.C4IIdea, Amount: 1},
			{Idea: idea.C4SIdea, Amount: 3},
			{Idea: idea.C4JIdea, Amount: 1},
			{Idea: idea.C4TIdea, Amount: 1},
			{Idea: idea.C3Idea, Amount: 3},
			{Idea: idea.C2Idea, Amount: 1},
			{Idea: idea.C1Idea, Amount: 1},
		},
	},
	Skill: map[SkillIndex][]Skill{
		Skill1: {
			{Multiplier: 2.0},
			{Multiplier: 3.0},
			{Multiplier: 5.0},
		},
		Skill2: {
			{Multiplier: 1.5},
			{Multiplier: 1.75},
			{Multiplier: 2.75},
		},
		Ultimate: {
			{Multiplier: 3.0},
		},
	},
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
