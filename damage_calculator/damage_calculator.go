package damage_calculator

import (
	"math"

	"github.com/aliceblock/re1999dmg/damage_calculator/character"
)

// DamageCalculator struct represents the parameters for damage calculation.
type DamageCalculator struct {
	DamageCalculatorInfo
}

// CalculateFinalDamage calculates the final damage using the defined formula.
func (d *DamageCalculator) CalculateFinalDamage(damageCalculatorInfo DamageCalculatorInfo, skill character.SkillIndex) []float64 {
	finalDamages := make([]float64, 3)

	// Calculate Effective Attack Value
	effectiveAttackValue := d.BaseAttackStats*(1+d.ResonanceAttackPercentage+d.PsychubeAttackPercentage) +
		d.ResonanceFixedAttackStats + d.PsychubeFixedAttackStats

	// Calculate Attack and Defense Factor
	attackDefenseFactor := effectiveAttackValue*(1+d.DamageBonus) -
		d.EnemyDefenseValue*(1+d.DefenseBonus-d.DefenseReduction)*(1-d.PenetrationRate)

	// Check if the result is less than the specified threshold
	if attackDefenseFactor < effectiveAttackValue*(1+d.ResonanceAttackPercentage)*0.1 {
		attackDefenseFactor = effectiveAttackValue * (1 + d.ResonanceAttackPercentage) * 0.1
	}

	// Calculate DMG Bonus
	dmgBonus := math.Max(1+d.CasterDamageIncrease-d.TargetDamageTakenReduction, 0.3)

	// Calculate Incantation/Ultimate/Ritual Might
	incantationUltimateRitualMight := math.Max(1+d.IncantationUltimateRitualBonusDmg, 0)

	// Calculate Critical Bonus
	criticalBonus := math.Max(1+d.CasterCriticalDamageBonus-d.TargetCriticalDefense, 1.1)

	// Calculate Afflatus Bonus
	afflatusBonus := 1.0
	if d.AfflatusAdvantage {
		afflatusBonus = 1.3
	}

	// Calculate Final Damage
	for _, skillInfo := range d.Character.Skill[skill] {
		var skillMultiplier float64 = skillInfo.Multiplier
		var finalDamage float64
		if d.CasterCriticalRate >= 1 {
			finalDamage = attackDefenseFactor * dmgBonus * incantationUltimateRitualMight * criticalBonus * afflatusBonus * skillMultiplier
		} else {
			finalDamage = (attackDefenseFactor*dmgBonus*incantationUltimateRitualMight*d.CasterCriticalRate*criticalBonus*afflatusBonus + attackDefenseFactor*dmgBonus*incantationUltimateRitualMight*(1-d.CasterCriticalRate)*afflatusBonus) * skillMultiplier
		}
		finalDamages = append(finalDamages, finalDamage)
	}

	return finalDamages
}

type DamageCalculatorInfo struct {
	Character                         character.Character
	BaseAttackStats                   float64
	ResonanceAttackPercentage         float64
	PsychubeAttackPercentage          float64
	ResonanceFixedAttackStats         float64
	PsychubeFixedAttackStats          float64
	DamageBonus                       float64
	EnemyDefenseValue                 float64
	DefenseBonus                      float64
	DefenseReduction                  float64
	PenetrationRate                   float64
	CasterDamageIncrease              float64
	TargetDamageTakenReduction        float64
	IncantationUltimateRitualBonusDmg float64
	CasterCriticalRate                float64
	CasterCriticalDamageBonus         float64
	TargetCriticalDefense             float64
	AfflatusAdvantage                 bool
}
