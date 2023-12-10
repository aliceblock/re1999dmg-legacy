package damage_calculator

import (
	"math"

	"github.com/aliceblock/re1999dmg/damage_calculator/character"
	"github.com/aliceblock/re1999dmg/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg/damage_calculator/psychube"
)

// DamageCalculator struct represents the parameters for damage calculation.
type DamageCalculator struct {
	Character                 *character.Character
	Psychube                  *psychube.Psychube
	Resonance                 *resonance.Resonance
	EnemyDef                  float64
	EnemyDefBonus             float64
	EnemyDamageTakenReduction float64
	BuffDmgBonus              float64
	EnemyDefReduction         float64
	PenetrationRate           float64
	DmgMight                  float64
	CritRate                  float64
	CritDmg                   float64
	EnemyCritDef              float64
	AfflatusAdvantage         bool
}

// CalculateFinalDamage calculates the final damage using the defined formula.
func (d *DamageCalculator) CalculateFinalDamage(additionalInfo DamageCalculatorInfo, skill character.SkillIndex) []float64 {
	var finalDamages []float64

	resonanceStats := d.Resonance.GetResonanceStats()

	// Calculate Effective Attack Value
	effectiveAttackValue := d.Character.Atk()*(1+resonanceStats.AtkPercent()+d.Psychube.AtkPercent()) + resonanceStats.Atk() + d.Psychube.Atk()

	// Calculate Attack and Defense Factor
	attackDefenseFactor := effectiveAttackValue*(1+d.Character.Insight().AtkPercent) - d.EnemyDef*(1+d.EnemyDefBonus-d.EnemyDefReduction-additionalInfo.EnemyDefReduction)*(1-d.PenetrationRate-additionalInfo.PenetrationRate)

	// Check if the result is less than the specified threshold
	if attackDefenseFactor < effectiveAttackValue*(1+d.Character.Insight().AtkPercent)*0.1 {
		attackDefenseFactor = effectiveAttackValue * (1 + d.Character.Insight().AtkPercent) * 0.1
	}

	// Calculate DMG Bonus
	posDmgBonus := d.Character.Insight().DmgBonus + resonanceStats.DmgBonus() + d.Psychube.DmgBonus() + d.BuffDmgBonus + additionalInfo.BuffDmgBonus
	negDmgBonus := d.EnemyDamageTakenReduction + additionalInfo.EnemyDamageTakenReduction
	dmgBonus := math.Max(1+posDmgBonus-negDmgBonus, 0.3)

	// Calculate Incantation/Ultimate/Ritual Might
	dmgMight := d.DmgMight
	if skill == character.Skill1 || skill == character.Skill2 {
		dmgMight += d.Psychube.IncantationMight() + additionalInfo.IncantationMight
	}
	if skill == character.Ultimate {
		dmgMight += d.Psychube.UltimateMight() + additionalInfo.UltimateMight
	}
	incantationUltimateRitualMight := math.Max(1+dmgMight, 0)

	// Calculate Critical Bonus
	criticalBonus := math.Max(1+d.Character.CritDmg()+resonanceStats.CritDmg()+d.Psychube.CritDmg()+d.CritDmg+additionalInfo.CritDmg-d.EnemyCritDef-additionalInfo.EnemyCritDef, 1.1)

	// Calculate Afflatus Bonus
	afflatusBonus := 1.0
	if d.AfflatusAdvantage {
		afflatusBonus = 1.3
	}

	// Calculate Final Damage
	for _, skillInfo := range d.Character.Skills(skill) {
		var skillMultiplier float64 = skillInfo.Multiplier
		var finalDamage float64
		critRate := d.Character.CritRate() + resonanceStats.CritRate() + d.Psychube.CritRate() + d.CritRate + additionalInfo.CritRate
		if critRate >= 1 {
			finalDamage = attackDefenseFactor * dmgBonus * incantationUltimateRitualMight * criticalBonus * afflatusBonus * skillMultiplier
		} else {
			finalDamage = (attackDefenseFactor*dmgBonus*incantationUltimateRitualMight*critRate*criticalBonus*afflatusBonus + attackDefenseFactor*dmgBonus*incantationUltimateRitualMight*(1-critRate)*afflatusBonus) * skillMultiplier
		}
		finalDamages = append(finalDamages, finalDamage)
	}

	return finalDamages
}

func (d *DamageCalculator) GetTotalCritRate() float64 {
	resonanceStats := d.Resonance.GetResonanceStats()
	return d.Character.CritRate() + resonanceStats.CritRate() + d.Psychube.CritRate() + d.CritRate
}

func ExcessCritDmgBonus(critRate float64) float64 {
	if critRate > 1.0 {
		return critRate - 1.0
	}
	return 0.0
}

type DamageCalculatorInfo struct {
	BuffDmgBonus              float64
	EnemyDefReduction         float64
	EnemyDamageTakenReduction float64
	PenetrationRate           float64
	IncantationMight          float64
	UltimateMight             float64
	CritRate                  float64
	CritDmg                   float64
	EnemyCritDef              float64
}
