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
	Buff                      *Buff
	Debuff                    *Debuff
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
func (d *DamageCalculator) CalculateFinalDamage(additionalInfo DamageCalculatorInfo, skill character.SkillIndex, enemyHit int16) []float64 {
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
		enemyHit := enemyHit
		if enemyHit > skillInfo.EnemyHit {
			enemyHit = skillInfo.EnemyHit
		}
		var skillMultiplier float64 = skillInfo.Multiplier
		var finalDamage float64

		if additionalInfo.HasExtraDamage {
			skillMultiplier += skillInfo.ExtraMultiplier[additionalInfo.ExtraDamageStack]
		}

		critRate := d.Character.CritRate() + resonanceStats.CritRate() + d.Psychube.CritRate() + d.CritRate + additionalInfo.CritRate
		if critRate >= 1 {
			finalDamage = attackDefenseFactor * dmgBonus * incantationUltimateRitualMight * criticalBonus * afflatusBonus * skillMultiplier * float64(enemyHit)
		} else {
			finalDamage = (attackDefenseFactor*dmgBonus*incantationUltimateRitualMight*critRate*criticalBonus*afflatusBonus + attackDefenseFactor*dmgBonus*incantationUltimateRitualMight*(1-critRate)*afflatusBonus) * skillMultiplier * float64(enemyHit)
		}
		finalDamages = append(finalDamages, finalDamage)
	}

	return finalDamages
}

func (d *DamageCalculator) CalculateGenesisDamage(additionalInfo DamageCalculatorInfo, skillMultiplier float64) float64 {
	resonanceStats := d.Resonance.GetResonanceStats()
	genesisDamage := (d.Character.Atk()*(1+resonanceStats.AtkPercent()+d.Psychube.AtkPercent()) + resonanceStats.Atk() + d.Psychube.Atk()) * (1 + d.Character.Insight().AtkPercent) * (1 + additionalInfo.BuffDmgBonus) * skillMultiplier
	return genesisDamage
}

func (d *DamageCalculator) GetTotalCritRate() float64 {
	resonanceStats := d.Resonance.GetResonanceStats()
	return d.Character.CritRate() + resonanceStats.CritRate() + d.Psychube.CritRate() + d.CritRate
}

func (d *DamageCalculator) GetBuffDebuffValue() BuffDebuffResult {
	dmgBonus := 0.0
	dmgTaken := 0.0
	defReduction := 0.0
	critResistDown := 0.0
	critDefDown := 0.0

	switch d.Buff.Sonetto {
	case 1:
		dmgBonus = 0.15
	case 2:
		dmgBonus = 0.2
	case 3:
		dmgBonus = 0.25
	}
	switch d.Buff.AnAnLee {
	case 1:
		if dmgBonus < 0.15 {
			dmgBonus = 0.15
		}
	case 2:
		if dmgBonus < 0.2 {
			dmgBonus = 0.2
		}
	case 3:
		if dmgBonus < 0.2 {
			dmgBonus = 0.3
		}
	}
	if d.Buff.Necrologist {
		if dmgBonus < 0.3 {
			dmgBonus = 0.3
		}
	}
	switch d.Debuff.Bkornblume {
	case 1:
		dmgTaken = 0.15
	case 2:
		dmgTaken = 0.2
	case 3:
		dmgTaken = 0.25
	}
	switch d.Debuff.BabyBlueSkill2 {
	case 1:
		if dmgTaken < 0.15 {
			dmgTaken = 0.15
		}
	case 2:
		if dmgTaken < 0.2 {
			dmgTaken = 0.2
		}
	case 3:
		if dmgTaken < 0.25 {
			dmgTaken = 0.25
		}
	}
	if d.Debuff.Confusion > 0 {
		critResistDown += float64(d.Debuff.Confusion) * 0.25
	}
	if d.Debuff.ToothFairy {
		critResistDown += 0.15
		critDefDown += 0.15
	}
	if d.Character.DamageType() == character.RealityDamage {
		switch d.Debuff.Bkornblume {
		case 1:
			defReduction = 0.15
		case 2:
			defReduction = 0.2
		case 3:
			defReduction = 0.25
		}
		if d.Debuff.SenseWeakness {
			defReduction += 0.2
			critDefDown += 0.2
		}
	}
	if d.Character.DamageType() == character.MentalDamage {
		switch d.Debuff.BabyBlueSKill1 {
		case 2:
			defReduction = 0.25
		case 3:
			defReduction = 0.35
		}
	}

	return BuffDebuffResult{
		DmgBonus:       dmgBonus,
		DmgTaken:       dmgTaken,
		DefReduction:   defReduction,
		CritResistDown: critResistDown,
		CritDefDown:    critDefDown,
	}
}

func ExcessCritDmgBonus(critRate float64) float64 {
	if critRate > 1.0 {
		return critRate - 1.0
	}
	return 0.0
}

type Buff struct {
	AnAnLee     int16
	Sonetto     int16
	Necrologist bool
}

type Debuff struct {
	Bkornblume     int16
	BabyBlueSKill1 int16
	BabyBlueSkill2 int16
	Confusion      int16
	ToothFairy     bool
	SenseWeakness  bool
}

type BuffDebuffResult struct {
	DmgBonus       float64
	DmgTaken       float64
	DefReduction   float64
	CritResistDown float64
	CritDefDown    float64
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
	HasExtraDamage            bool
	ExtraDamageStack          int16
}
