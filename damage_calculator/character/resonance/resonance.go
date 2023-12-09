package resonance

import "math"

type Resonance struct {
	Ideas []IdeaAmount `json:"ideas"`
}

func (r *Resonance) GetResonanceStats() *ResonanceStats {
	stats := &ResonanceStats{}
	for _, ideaAmount := range r.Ideas {
		if ideaAmount.Amount <= 0 {
			continue
		}
		stats.atk += ideaAmount.Idea.Atk * float64(ideaAmount.Amount)
		stats.atkPercent += ideaAmount.Idea.AtkPercent * float64(ideaAmount.Amount)
		stats.critRate += ideaAmount.Idea.CritRate * float64(ideaAmount.Amount)
		stats.critDmg += ideaAmount.Idea.CritDmg * float64(ideaAmount.Amount)
		stats.dmgBonus += ideaAmount.Idea.DmgBonus * float64(ideaAmount.Amount)
	}
	stats.atk = toFixed(stats.atk, 3)
	stats.atkPercent = toFixed(stats.atkPercent, 3)
	stats.dmgBonus = toFixed(stats.dmgBonus, 3)
	stats.critRate = toFixed(stats.critRate, 3)
	stats.critDmg = toFixed(stats.critDmg, 3)
	return stats
}

type IdeaAmount struct {
	Idea   Idea
	Amount int16
}

type ResonanceStats struct {
	atk        float64
	atkPercent float64
	dmgBonus   float64
	critRate   float64
	critDmg    float64
}

func (r *ResonanceStats) Atk() float64 {
	return r.atk
}

func (r *ResonanceStats) AtkPercent() float64 {
	return r.atkPercent
}

func (r *ResonanceStats) DmgBonus() float64 {
	return r.dmgBonus
}

func (r *ResonanceStats) CritRate() float64 {
	return r.critRate
}

func (r *ResonanceStats) CritDmg() float64 {
	return r.critDmg
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
