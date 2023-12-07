package main

import (
	"fmt"

	DmgCal "github.com/aliceblock/re1999dmg/damage_calculator"
	"github.com/aliceblock/re1999dmg/damage_calculator/character"
	"github.com/aliceblock/re1999dmg/damage_calculator/psychube"
)

func main() {
	// aKnightDmgCalculate(false, false, false)
	regulusDmgCalculate(character.RegulusStat, false, false, false)
}

func aKnightDmgCalculate(afflatusAdvantage bool, applyAnAnLeeBuff bool, applyBkbBuff bool) {
	baseAtk := 1176.0
	resoAtkPercent := 0.18
	resoAtkFixed := 141.0
	anAnleeDmgBonus := 0.0
	if applyAnAnLeeBuff {
		anAnleeDmgBonus = 0.15
	}
	enemyDef := 600.0
	bkbDefDown := 0.0
	bkbDmgTakenPlus := 0.0
	if applyBkbBuff {
		bkbDmgTakenPlus = -0.15
	}
	baseCrit := 0.27

	calculatorForHop := DmgCal.DamageCalculator{
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

	calculatorForBoundenDuty := DmgCal.DamageCalculator{
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

/*
1 :2 Move Merge
2 :5 Use 2 Incantation
3 :5 Wait
4 :1 Ultimate
5 :3 Move Merge
6 :5 Use 1 Incantation
7 :5 Wait
8 :1 Ultimate
9 :3 Move Merge
10:5 Use 1 Incantation
11:5 Wait
12:1 Ultimate
*/
func regulusDmgCalculate(char character.Character, afflatusAdvantage bool, applyAnAnLeeBuff bool, applyBkbBuff bool) {
	resonanceStats := char.GetResonanceStats()

	baseAtk := char.Atk
	resoAtkPercent := resonanceStats.AtkPercent
	resoAtkFixed := resonanceStats.Atk
	baseCrit := char.CritRate + resonanceStats.CritRate

	enemyDef := 600.0

	// Additional Bonus
	regulusCritRateBonus := 0.5

	anAnleeDmgBonus := 0.0
	if applyAnAnLeeBuff {
		anAnleeDmgBonus = 0.15
	}
	bkbDefDown := 0.0
	bkbDmgTakenPlus := 0.0
	if applyBkbBuff {
		bkbDefDown = 0.15
		bkbDmgTakenPlus = -0.15
	}

	calculatorForBraveNewWorld := DmgCal.DamageCalculator{
		BaseAttackStats:                   baseAtk,
		ResonanceAttackPercentage:         resoAtkPercent,
		PsychubeAttackPercentage:          0,
		ResonanceFixedAttackStats:         resoAtkFixed,
		PsychubeFixedAttackStats:          psychube.BraveNewWorld.Atk,
		DamageBonus:                       anAnleeDmgBonus,
		EnemyDefenseValue:                 enemyDef,
		DefenseBonus:                      0,
		DefenseReduction:                  bkbDefDown,
		PenetrationRate:                   0,
		CasterDamageIncrease:              resonanceStats.DmgBonus,
		TargetDamageTakenReduction:        bkbDmgTakenPlus,
		IncantationUltimateRitualBonusDmg: 0,
		CasterCriticalRate:                baseCrit,
		CasterCriticalDamageBonus:         char.CritDmg + resonanceStats.CritDmg + regulusExcessCritDmgBonus(baseCrit),
		TargetCriticalDefense:             0.1,
		AfflatusAdvantage:                 afflatusAdvantage,
	}

	calculatorForBraveNewWorld.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	finalDamageBraveNewWorldSkill1 := calculatorForBraveNewWorld.CalculateFinalDamage()
	calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = psychube.BraveNewWorld.AdditionalEffect.IncantationMight
	finalDamageBraveNewWorldBuffSkill1 := calculatorForBraveNewWorld.CalculateFinalDamage()

	calculatorForBraveNewWorld.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = 0
	finalDamageBraveNewWorldSkill2 := calculatorForBraveNewWorld.CalculateFinalDamage()
	calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = psychube.BraveNewWorld.AdditionalEffect.IncantationMight
	finalDamageBraveNewWorldBuffSkill2 := calculatorForBraveNewWorld.CalculateFinalDamage()

	calculatorForBraveNewWorld.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = psychube.BraveNewWorld.UltimateMight
	finalDamageBraveNewWorldUlt := calculatorForBraveNewWorld.CalculateFinalDamage()

	// Reset
	calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = 0
	calculatorForBraveNewWorld.CasterCriticalRate = baseCrit + regulusCritRateBonus
	calculatorForBraveNewWorld.CasterCriticalDamageBonus = char.CritDmg + resonanceStats.CritDmg + regulusExcessCritDmgBonus(baseCrit+regulusCritRateBonus)

	calculatorForBraveNewWorld.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	finalDamageBraveNewWorldWithRestlessHeartSkill1 := calculatorForBraveNewWorld.CalculateFinalDamage()
	calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = psychube.BraveNewWorld.AdditionalEffect.IncantationMight
	finalDamageBraveNewWorldWithRestlessHeartBuffSkill1 := calculatorForBraveNewWorld.CalculateFinalDamage()

	calculatorForBraveNewWorld.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = 0
	finalDamageBraveNewWorldWithRestlessHeartSkill2 := calculatorForBraveNewWorld.CalculateFinalDamage()
	calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = psychube.BraveNewWorld.AdditionalEffect.IncantationMight
	finalDamageBraveNewWorldWithRestlessHeartBuffSkill2 := calculatorForBraveNewWorld.CalculateFinalDamage()

	calculatorForBraveNewWorld.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = psychube.BraveNewWorld.UltimateMight
	finalDamageBraveNewWorldWithRestlessHeartUlt := calculatorForBraveNewWorld.CalculateFinalDamage()

	fmt.Printf("\n---------\nRegulus Brave New World Final Damage:")
	fmt.Printf("\nSkill 1: %.2f (with BNW Buff %.2f)", finalDamageBraveNewWorldSkill1, finalDamageBraveNewWorldBuffSkill1)
	fmt.Printf("\nSkill 2: %.2f (with BNW Buff %.2f)", finalDamageBraveNewWorldSkill2, finalDamageBraveNewWorldBuffSkill2)
	fmt.Printf("\nUltimate: %.2f", finalDamageBraveNewWorldUlt)
	fmt.Printf("\nSkill 1 with Restless Heart: %.2f (with BNW Buff %.2f)", finalDamageBraveNewWorldWithRestlessHeartSkill1, finalDamageBraveNewWorldWithRestlessHeartBuffSkill1)
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f (with BNW Buff %.2f)", finalDamageBraveNewWorldWithRestlessHeartSkill2, finalDamageBraveNewWorldWithRestlessHeartBuffSkill2)
	fmt.Printf("\nUltimate with Restless Heart: %.2f", finalDamageBraveNewWorldWithRestlessHeartUlt)

	fmt.Println()

	calculatorForThunder := DmgCal.DamageCalculator{
		BaseAttackStats:                   baseAtk,
		ResonanceAttackPercentage:         resoAtkPercent,
		PsychubeAttackPercentage:          0,
		ResonanceFixedAttackStats:         resoAtkFixed,
		PsychubeFixedAttackStats:          psychube.ThunderousApplause.Atk,
		DamageBonus:                       anAnleeDmgBonus,
		EnemyDefenseValue:                 enemyDef,
		DefenseBonus:                      0,
		DefenseReduction:                  bkbDefDown,
		PenetrationRate:                   0,
		CasterDamageIncrease:              resonanceStats.DmgBonus,
		TargetDamageTakenReduction:        bkbDmgTakenPlus,
		IncantationUltimateRitualBonusDmg: 0,
		CasterCriticalRate:                baseCrit + psychube.ThunderousApplause.CritRate,
		CasterCriticalDamageBonus:         char.CritDmg + resonanceStats.CritDmg + psychube.ThunderousApplause.CritDmg + regulusExcessCritDmgBonus(baseCrit+psychube.ThunderousApplause.CritRate),
		TargetCriticalDefense:             0.1,
		AfflatusAdvantage:                 afflatusAdvantage,
	}

	calculatorForThunder.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	finalDamageThunderSkill1 := calculatorForThunder.CalculateFinalDamage()
	calculatorForThunder.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	finalDamageThunderSkill2 := calculatorForThunder.CalculateFinalDamage()
	calculatorForThunder.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	finalDamageThunderUlt := calculatorForThunder.CalculateFinalDamage()

	calculatorForThunder.CasterCriticalRate = baseCrit + psychube.ThunderousApplause.CritRate + regulusCritRateBonus
	calculatorForThunder.CasterCriticalDamageBonus = char.CritDmg + resonanceStats.CritDmg + psychube.ThunderousApplause.CritDmg + regulusExcessCritDmgBonus(baseCrit+psychube.ThunderousApplause.CritRate+regulusCritRateBonus)

	calculatorForThunder.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	finalDamageThunderWithRestlessHeartSkill1 := calculatorForThunder.CalculateFinalDamage()
	calculatorForThunder.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	finalDamageThunderWithRestlessHeartSkill2 := calculatorForThunder.CalculateFinalDamage()
	calculatorForThunder.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	finalDamageThunderWithRestlessHeartUlt := calculatorForThunder.CalculateFinalDamage()

	fmt.Printf("---------\nRegulus Thunder Final Damage:")
	fmt.Printf("\nSkill 1: %.2f", finalDamageThunderSkill1)
	fmt.Printf("\nSkill 2: %.2f", finalDamageThunderSkill2)
	fmt.Printf("\nUltimate: %.2f", finalDamageThunderUlt)
	fmt.Printf("\nSkill 1 with Restless Heart: %.2f", finalDamageThunderWithRestlessHeartSkill1)
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f", finalDamageThunderWithRestlessHeartSkill2)
	fmt.Printf("\nUltimate with Restless Heart: %.2f", finalDamageThunderWithRestlessHeartUlt)

	fmt.Println()

	calculatorForBoundenDuty := DmgCal.DamageCalculator{
		BaseAttackStats:                   baseAtk,
		ResonanceAttackPercentage:         resoAtkPercent,
		PsychubeAttackPercentage:          0,
		ResonanceFixedAttackStats:         resoAtkFixed,
		PsychubeFixedAttackStats:          psychube.HisBoundenDuty.Atk,
		DamageBonus:                       anAnleeDmgBonus,
		EnemyDefenseValue:                 enemyDef,
		DefenseBonus:                      0,
		DefenseReduction:                  bkbDefDown,
		PenetrationRate:                   0,
		CasterDamageIncrease:              resonanceStats.DmgBonus + psychube.HisBoundenDuty.DmgBonus,
		TargetDamageTakenReduction:        bkbDmgTakenPlus,
		IncantationUltimateRitualBonusDmg: 0,
		CasterCriticalRate:                baseCrit,
		CasterCriticalDamageBonus:         char.CritDmg + resonanceStats.CritDmg + regulusExcessCritDmgBonus(baseCrit),
		TargetCriticalDefense:             0.1,
		AfflatusAdvantage:                 afflatusAdvantage,
	}

	calculatorForBoundenDuty.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	finalDamageBoundenDutySkill1 := calculatorForBoundenDuty.CalculateFinalDamage()
	calculatorForBoundenDuty.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	finalDamageBoundenDutySkill2 := calculatorForBoundenDuty.CalculateFinalDamage()
	calculatorForBoundenDuty.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	finalDamageBoundenDutyUlt := calculatorForBoundenDuty.CalculateFinalDamage()

	calculatorForBoundenDuty.CasterCriticalRate = baseCrit + regulusCritRateBonus
	calculatorForBoundenDuty.CasterCriticalDamageBonus = char.CritDmg + resonanceStats.CritDmg + regulusExcessCritDmgBonus(baseCrit+regulusCritRateBonus)
	calculatorForBoundenDuty.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	finalDamageBoundenDutyWithRestlessHeartSkill1 := calculatorForBoundenDuty.CalculateFinalDamage()
	calculatorForBoundenDuty.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	finalDamageBoundenDutyWithRestlessHeartSkill2 := calculatorForBoundenDuty.CalculateFinalDamage()
	calculatorForBoundenDuty.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	finalDamageBoundenDutyWithRestlessHeartUlt := calculatorForBoundenDuty.CalculateFinalDamage()

	fmt.Printf("---------\nRegulus His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f", finalDamageBoundenDutySkill1)
	fmt.Printf("\nSkill 2: %.2f", finalDamageBoundenDutySkill2)
	fmt.Printf("\nUltimate: %.2f", finalDamageBoundenDutyUlt)
	fmt.Printf("\nSkill 1 with Restless Heart: %.2f", finalDamageBoundenDutyWithRestlessHeartSkill1)
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f", finalDamageBoundenDutyWithRestlessHeartSkill2)
	fmt.Printf("\nUltimate with Restless Heart: %.2f", finalDamageBoundenDutyWithRestlessHeartUlt)

	fmt.Println()

	calculatorForHop := DmgCal.DamageCalculator{
		BaseAttackStats:                   baseAtk,
		ResonanceAttackPercentage:         resoAtkPercent,
		PsychubeAttackPercentage:          0,
		ResonanceFixedAttackStats:         resoAtkFixed,
		PsychubeFixedAttackStats:          psychube.Hopscotch.Atk,
		DamageBonus:                       anAnleeDmgBonus,
		EnemyDefenseValue:                 enemyDef,
		DefenseBonus:                      0,
		DefenseReduction:                  bkbDefDown,
		PenetrationRate:                   0,
		CasterDamageIncrease:              resonanceStats.DmgBonus,
		TargetDamageTakenReduction:        bkbDmgTakenPlus,
		IncantationUltimateRitualBonusDmg: 0,
		CasterCriticalRate:                baseCrit,
		CasterCriticalDamageBonus:         char.CritDmg + resonanceStats.CritDmg + regulusExcessCritDmgBonus(baseCrit),
		TargetCriticalDefense:             0.1,
		AfflatusAdvantage:                 afflatusAdvantage,
	}

	calculatorForHop.IncantationUltimateRitualBonusDmg = psychube.Hopscotch.IncantationMight
	calculatorForHop.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	finalDamageHopSkill1 := calculatorForHop.CalculateFinalDamage()
	calculatorForHop.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	finalDamageHopSkill2 := calculatorForHop.CalculateFinalDamage()
	calculatorForHop.IncantationUltimateRitualBonusDmg = 0
	calculatorForHop.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	finalDamageHopUlt := calculatorForHop.CalculateFinalDamage()

	calculatorForHop.CasterCriticalRate = baseCrit + regulusCritRateBonus
	calculatorForHop.CasterCriticalDamageBonus = char.CritDmg + resonanceStats.CritDmg + regulusExcessCritDmgBonus(baseCrit+regulusCritRateBonus)

	calculatorForHop.IncantationUltimateRitualBonusDmg = psychube.Hopscotch.IncantationMight
	calculatorForHop.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	finalDamageHopWithRestlessHeartSkill1 := calculatorForHop.CalculateFinalDamage()
	calculatorForHop.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	finalDamageHopWithRestlessHeartSkill2 := calculatorForHop.CalculateFinalDamage()
	calculatorForHop.IncantationUltimateRitualBonusDmg = 0
	calculatorForHop.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	finalDamageHopWithRestlessHeartUlt := calculatorForHop.CalculateFinalDamage()

	fmt.Printf("---------\nRegulus Hopscotch Final Damage:")
	fmt.Printf("\nSkill 1: %.2f", finalDamageHopSkill1)
	fmt.Printf("\nSkill 2: %.2f", finalDamageHopSkill2)
	fmt.Printf("\nUltimate: %.2f", finalDamageHopUlt)
	fmt.Printf("\nSkill 1 with Restless Heart: %.2f", finalDamageHopWithRestlessHeartSkill1)
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f", finalDamageHopWithRestlessHeartSkill2)
	fmt.Printf("\nUltimate with Restless Heart: %.2f", finalDamageHopWithRestlessHeartUlt)
}

func regulusExcessCritDmgBonus(critRate float64) float64 {
	if critRate > 1.0 {
		return critRate - 1.0
	}
	return 0.0
}
