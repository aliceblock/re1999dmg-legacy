package main

import (
	"fmt"
	"math"
)

// DamageCalculator struct represents the parameters for damage calculation.
type DamageCalculator struct {
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
	SkillMultiplier                   float64
}

// CalculateFinalDamage calculates the final damage using the defined formula.
func (d *DamageCalculator) CalculateFinalDamage() float64 {
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
	var finalDamage float64
	if d.CasterCriticalRate >= 1 {
		finalDamage = attackDefenseFactor * dmgBonus * incantationUltimateRitualMight * criticalBonus * afflatusBonus * d.SkillMultiplier
	} else {
		finalDamage = (attackDefenseFactor*dmgBonus*incantationUltimateRitualMight*d.CasterCriticalRate*criticalBonus*afflatusBonus + attackDefenseFactor*dmgBonus*incantationUltimateRitualMight*(1-d.CasterCriticalRate)*afflatusBonus) * d.SkillMultiplier

	}

	// Echo every calculation step
	// fmt.Printf("Effective Attack Value: %.2f\n", effectiveAttackValue)
	// fmt.Printf("Attack and Defense Factor: %.2f\n", attackDefenseFactor)
	// fmt.Printf("DMG Bonus: %.2f\n", dmgBonus)
	// fmt.Printf("Incantation/Ultimate/Ritual Might: %.2f\n", incantationUltimateRitualMight)
	// fmt.Printf("Critical Bonus: %.2f\n", criticalBonus)
	// fmt.Printf("Afflatus Bonus: %.2f\n", afflatusBonus)
	// fmt.Printf("Skill Multiplier: %.2f\n", d.SkillMultiplier)

	return finalDamage
}

func main() {
	aKnightDmgCalculate()
	regulusDmgCalculate()
}

func aKnightDmgCalculate() {
	baseAtk := 1176.0
	resoAtkPercent := 0.18
	resoAtkFixed := 141.0
	anAnleeDmgBonus := 0.15
	anAnleeDmgBonus = 0.0
	enemyDef := 600.0
	bkbDefDown := 0.15
	bkbDefDown = 0.0
	bkbDmgTakenPlus := -0.15
	bkbDmgTakenPlus = 0.0
	afflatusAdvantage := true
	baseCrit := 0.27

	calculatorForHop := DamageCalculator{
		BaseAttackStats:                   baseAtk,
		ResonanceAttackPercentage:         resoAtkPercent,
		PsychubeAttackPercentage:          0.0,
		ResonanceFixedAttackStats:         resoAtkFixed,
		PsychubeFixedAttackStats:          370,
		DamageBonus:                       anAnleeDmgBonus,
		EnemyDefenseValue:                 enemyDef,
		DefenseBonus:                      0,
		DefenseReduction:                  bkbDefDown,
		PenetrationRate:                   0,
		CasterDamageIncrease:              0.155,
		TargetDamageTakenReduction:        bkbDmgTakenPlus,
		IncantationUltimateRitualBonusDmg: 0.18,
		CasterCriticalRate:                baseCrit,
		CasterCriticalDamageBonus:         0.533,
		TargetCriticalDefense:             0.1,
		AfflatusAdvantage:                 afflatusAdvantage,
		SkillMultiplier:                   2.75,
	}

	finalDamageHopSkill1 := calculatorForHop.CalculateFinalDamage()
	calculatorForHop.SkillMultiplier = 1.5
	finalDamageHopSkill2 := calculatorForHop.CalculateFinalDamage()
	calculatorForHop.SkillMultiplier = 5.5
	calculatorForHop.IncantationUltimateRitualBonusDmg = 0
	finalDamageHopUlt := calculatorForHop.CalculateFinalDamage()
	calculatorForHop.IncantationUltimateRitualBonusDmg = 0.04
	finalDamageHopUlt1 := calculatorForHop.CalculateFinalDamage()
	calculatorForHop.IncantationUltimateRitualBonusDmg = 0.08
	finalDamageHopUlt2 := calculatorForHop.CalculateFinalDamage()
	calculatorForHop.IncantationUltimateRitualBonusDmg = 0.12
	finalDamageHopUlt3 := calculatorForHop.CalculateFinalDamage()
	calculatorForHop.IncantationUltimateRitualBonusDmg = 0.16
	finalDamageHopUlt4 := calculatorForHop.CalculateFinalDamage()

	fmt.Printf("\n---------\nA Knight Hopscotch:")
	fmt.Printf("\nSkill 1: %.2f", finalDamageHopSkill1)
	fmt.Printf("\nSkill 2: %.2f", finalDamageHopSkill2)
	fmt.Printf("\nUltimate: %.2f", finalDamageHopUlt)
	fmt.Printf("\nUltimate+1: %.2f", finalDamageHopUlt1)
	fmt.Printf("\nUltimate+2: %.2f", finalDamageHopUlt2)
	fmt.Printf("\nUltimate+3: %.2f", finalDamageHopUlt3)
	fmt.Printf("\nUltimate+4: %.2f", finalDamageHopUlt4)

	fmt.Println()

	calculatorForBoundenDuty := DamageCalculator{
		BaseAttackStats:                   baseAtk,
		ResonanceAttackPercentage:         resoAtkPercent,
		PsychubeAttackPercentage:          0.0,
		ResonanceFixedAttackStats:         resoAtkFixed,
		PsychubeFixedAttackStats:          410,
		DamageBonus:                       anAnleeDmgBonus,
		EnemyDefenseValue:                 enemyDef,
		DefenseBonus:                      0,
		DefenseReduction:                  bkbDefDown,
		PenetrationRate:                   0,
		CasterDamageIncrease:              0.275,
		TargetDamageTakenReduction:        bkbDmgTakenPlus,
		IncantationUltimateRitualBonusDmg: 0,
		CasterCriticalRate:                baseCrit,
		CasterCriticalDamageBonus:         0.533,
		TargetCriticalDefense:             0.1,
		AfflatusAdvantage:                 afflatusAdvantage,
		SkillMultiplier:                   2.75,
	}

	finalDamageHisBoundenDutySkill1 := calculatorForBoundenDuty.CalculateFinalDamage()
	calculatorForBoundenDuty.SkillMultiplier = 1.5
	finalDamageHisBoundenDutySkill2 := calculatorForBoundenDuty.CalculateFinalDamage()
	calculatorForBoundenDuty.SkillMultiplier = 5.5
	finalDamageHisBoundenDutyUlt := calculatorForBoundenDuty.CalculateFinalDamage()

	fmt.Printf("---------\nA Knight His Bounden Duty:")
	fmt.Printf("\nSkill 1: %.2f", finalDamageHisBoundenDutySkill1)
	fmt.Printf("\nSkill 2: %.2f", finalDamageHisBoundenDutySkill2)
	fmt.Printf("\nUltimate: %.2f", finalDamageHisBoundenDutyUlt)
}

func regulusDmgCalculate() {
	baseAtk := 1186.0
	resoAtkPercent := 0.105
	resoAtkFixed := 142.0
	anAnleeDmgBonus := 0.15
	anAnleeDmgBonus = 0.0
	enemyDef := 600.0
	bkbDefDown := 0.15
	bkbDefDown = 0.0
	bkbDmgTakenPlus := -0.15
	bkbDmgTakenPlus = 0.0
	afflatusAdvantage := false
	baseCrit := 0.452
	regulusCritRateBonus := 0.5
	psychubeCritRate := 0.16
	resoDmgBonus := 0.205
	hisBoundenDutyDmgBonus := 0.12
	braveNewWorldUltBonus := 0.18

	calculatorForThunder := DamageCalculator{
		BaseAttackStats:                   baseAtk,
		ResonanceAttackPercentage:         resoAtkPercent,
		PsychubeAttackPercentage:          0.0,
		ResonanceFixedAttackStats:         resoAtkFixed,
		PsychubeFixedAttackStats:          330,
		DamageBonus:                       anAnleeDmgBonus,
		EnemyDefenseValue:                 enemyDef,
		DefenseBonus:                      0,
		DefenseReduction:                  bkbDefDown,
		PenetrationRate:                   0,
		CasterDamageIncrease:              resoDmgBonus,
		TargetDamageTakenReduction:        bkbDmgTakenPlus,
		IncantationUltimateRitualBonusDmg: 0,
		CasterCriticalRate:                baseCrit + psychubeCritRate,
		CasterCriticalDamageBonus:         0.558 + 0.065 + 0.15 + 0.16 + regulusExcessCritDmgBonus(baseCrit+psychubeCritRate),
		TargetCriticalDefense:             0.1,
		AfflatusAdvantage:                 afflatusAdvantage,
		SkillMultiplier:                   3.0,
	}

	finalDamageThunderSkill1 := calculatorForThunder.CalculateFinalDamage()
	calculatorForThunder.SkillMultiplier = 1.75
	finalDamageThunderSkill2 := calculatorForThunder.CalculateFinalDamage()
	calculatorForThunder.SkillMultiplier = 3.0
	finalDamageThunderUlt := calculatorForThunder.CalculateFinalDamage()

	calculatorForThunder.CasterCriticalRate = baseCrit + psychubeCritRate + regulusCritRateBonus
	calculatorForThunder.CasterCriticalDamageBonus = 0.558 + 0.065 + 0.15 + 0.16 + regulusExcessCritDmgBonus(baseCrit+psychubeCritRate+regulusCritRateBonus)
	calculatorForThunder.SkillMultiplier = 3.0
	finalDamageThunderWithRestlessHeartSkill1 := calculatorForThunder.CalculateFinalDamage()
	calculatorForThunder.SkillMultiplier = 1.75
	finalDamageThunderWithRestlessHeartSkill2 := calculatorForThunder.CalculateFinalDamage()
	calculatorForThunder.SkillMultiplier = 3.0
	finalDamageThunderWithRestlessHeartUlt := calculatorForThunder.CalculateFinalDamage()

	fmt.Printf("\n---------\nRegulus Thunder Final Damage:")
	fmt.Printf("\nSkill 1: %.2f", finalDamageThunderSkill1)
	fmt.Printf("\nSkill 2: %.2f", finalDamageThunderSkill2)
	fmt.Printf("\nUltimate: %.2f", finalDamageThunderUlt)
	fmt.Printf("\nSkill 1 with Restless Heart: %.2f", finalDamageThunderWithRestlessHeartSkill1)
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f", finalDamageThunderWithRestlessHeartSkill2)
	fmt.Printf("\nUltimate with Restless Heart: %.2f", finalDamageThunderWithRestlessHeartUlt)

	fmt.Println()

	calculatorForBraveNewWorld := DamageCalculator{
		BaseAttackStats:                   baseAtk,
		ResonanceAttackPercentage:         resoAtkPercent,
		PsychubeAttackPercentage:          0.0,
		ResonanceFixedAttackStats:         resoAtkFixed,
		PsychubeFixedAttackStats:          370,
		DamageBonus:                       anAnleeDmgBonus,
		EnemyDefenseValue:                 enemyDef,
		DefenseBonus:                      0,
		DefenseReduction:                  bkbDefDown,
		PenetrationRate:                   0,
		CasterDamageIncrease:              resoDmgBonus,
		TargetDamageTakenReduction:        bkbDmgTakenPlus,
		IncantationUltimateRitualBonusDmg: 0,
		CasterCriticalRate:                baseCrit,
		CasterCriticalDamageBonus:         0.558 + 0.065 + 0.15 + regulusExcessCritDmgBonus(baseCrit),
		TargetCriticalDefense:             0.1,
		AfflatusAdvantage:                 afflatusAdvantage,
		SkillMultiplier:                   3.0,
	}

	finalDamageBraveNewWorldSkill1 := calculatorForBraveNewWorld.CalculateFinalDamage()
	calculatorForBraveNewWorld.SkillMultiplier = 1.75
	finalDamageBraveNewWorldSkill2 := calculatorForBraveNewWorld.CalculateFinalDamage()
	calculatorForBraveNewWorld.SkillMultiplier = 3.0
	calculatorForBraveNewWorld.CasterDamageIncrease = resoDmgBonus + braveNewWorldUltBonus
	finalDamageBraveNewWorldUlt := calculatorForBraveNewWorld.CalculateFinalDamage()

	calculatorForBraveNewWorld.CasterDamageIncrease = resoDmgBonus
	calculatorForBraveNewWorld.CasterCriticalRate = baseCrit + regulusCritRateBonus
	calculatorForBraveNewWorld.CasterCriticalDamageBonus = 0.558 + 0.065 + 0.15 + 0.16 + regulusExcessCritDmgBonus(baseCrit+regulusCritRateBonus)
	calculatorForBraveNewWorld.SkillMultiplier = 3.0
	finalDamageBraveNewWorldWithRestlessHeartSkill1 := calculatorForBraveNewWorld.CalculateFinalDamage()
	calculatorForBraveNewWorld.SkillMultiplier = 1.75
	finalDamageBraveNewWorldWithRestlessHeartSkill2 := calculatorForBraveNewWorld.CalculateFinalDamage()
	calculatorForBraveNewWorld.SkillMultiplier = 3.0
	calculatorForBraveNewWorld.CasterDamageIncrease = resoDmgBonus + braveNewWorldUltBonus
	finalDamageBraveNewWorldWithRestlessHeartUlt := calculatorForBraveNewWorld.CalculateFinalDamage()

	// TODO: Add Brave New World Buff Triggered

	fmt.Printf("---------\nRegulus Brave New World Final Damage:")
	fmt.Printf("\nSkill 1: %.2f", finalDamageBraveNewWorldSkill1)
	fmt.Printf("\nSkill 2: %.2f", finalDamageBraveNewWorldSkill2)
	fmt.Printf("\nUltimate: %.2f", finalDamageBraveNewWorldUlt)
	fmt.Printf("\nSkill 1 with Restless Heart: %.2f", finalDamageBraveNewWorldWithRestlessHeartSkill1)
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f", finalDamageBraveNewWorldWithRestlessHeartSkill2)
	fmt.Printf("\nUltimate with Restless Heart: %.2f", finalDamageBraveNewWorldWithRestlessHeartUlt)

	fmt.Println()

	calculatorForBoundenDuty := DamageCalculator{
		BaseAttackStats:                   baseAtk,
		ResonanceAttackPercentage:         resoAtkPercent,
		PsychubeAttackPercentage:          0.0,
		ResonanceFixedAttackStats:         resoAtkFixed,
		PsychubeFixedAttackStats:          410,
		DamageBonus:                       anAnleeDmgBonus,
		EnemyDefenseValue:                 enemyDef,
		DefenseBonus:                      0,
		DefenseReduction:                  bkbDefDown,
		PenetrationRate:                   0,
		CasterDamageIncrease:              resoDmgBonus + hisBoundenDutyDmgBonus,
		TargetDamageTakenReduction:        bkbDmgTakenPlus,
		IncantationUltimateRitualBonusDmg: 0,
		CasterCriticalRate:                baseCrit,
		CasterCriticalDamageBonus:         0.558 + 0.065 + 0.15 + regulusExcessCritDmgBonus(baseCrit),
		TargetCriticalDefense:             0.1,
		AfflatusAdvantage:                 afflatusAdvantage,
		SkillMultiplier:                   3.0,
	}

	finalDamageBoundenDutySkill1 := calculatorForBoundenDuty.CalculateFinalDamage()
	calculatorForBoundenDuty.SkillMultiplier = 1.75
	finalDamageBoundenDutySkill2 := calculatorForBoundenDuty.CalculateFinalDamage()
	calculatorForBoundenDuty.SkillMultiplier = 3.0
	finalDamageBoundenDutyUlt := calculatorForBoundenDuty.CalculateFinalDamage()

	calculatorForBoundenDuty.CasterCriticalRate = baseCrit + regulusCritRateBonus
	calculatorForBoundenDuty.CasterCriticalDamageBonus = 0.558 + 0.065 + 0.15 + 0.16 + regulusExcessCritDmgBonus(baseCrit+regulusCritRateBonus)
	calculatorForBoundenDuty.SkillMultiplier = 3.0
	finalDamageBoundenDutyWithRestlessHeartSkill1 := calculatorForBoundenDuty.CalculateFinalDamage()
	calculatorForBoundenDuty.SkillMultiplier = 1.75
	finalDamageBoundenDutyWithRestlessHeartSkill2 := calculatorForBoundenDuty.CalculateFinalDamage()
	calculatorForBoundenDuty.SkillMultiplier = 3.0
	finalDamageBoundenDutyWithRestlessHeartUlt := calculatorForBoundenDuty.CalculateFinalDamage()

	fmt.Printf("---------\nRegulus His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f", finalDamageBoundenDutySkill1)
	fmt.Printf("\nSkill 2: %.2f", finalDamageBoundenDutySkill2)
	fmt.Printf("\nUltimate: %.2f", finalDamageBoundenDutyUlt)
	fmt.Printf("\nSkill 1 with Restless Heart: %.2f", finalDamageBoundenDutyWithRestlessHeartSkill1)
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f", finalDamageBoundenDutyWithRestlessHeartSkill2)
	fmt.Printf("\nUltimate with Restless Heart: %.2f", finalDamageBoundenDutyWithRestlessHeartUlt)
}

func regulusExcessCritDmgBonus(critRate float64) float64 {
	if critRate > 1.0 {
		return critRate - 1.0
	}
	return 0.0
}
