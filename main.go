package main

import (
	"fmt"

	DmgCal "github.com/aliceblock/re1999dmg/damage_calculator"
	"github.com/aliceblock/re1999dmg/damage_calculator/character"
	"github.com/aliceblock/re1999dmg/damage_calculator/psychube"
)

func main() {
	regulusDmgCalculate(false, false, false, false, false)
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
func regulusDmgCalculate(afflatusAdvantage bool, applyAnAnLeeBuff bool, applyBkbBuff bool, applyConfusion bool, applyToothFairyBuff bool) {
	critRate := 0.0

	enemyDef := 600.0
	enemyCritDef := 0.1

	// Additional Bonus
	regulusCritRateBonus := 0.5

	// Buff/Debuff
	dmgBonus := 0.0
	enemyDamageTakenReduction := 0.0

	anAnleeDmgBonus := 0.0
	if applyAnAnLeeBuff {
		anAnleeDmgBonus = 0.15
	}
	dmgBonus += anAnleeDmgBonus

	bkbDmgTakenPlus := 0.0
	if applyBkbBuff {
		bkbDmgTakenPlus = -0.15
	}
	enemyDamageTakenReduction += bkbDmgTakenPlus

	confusionCritResistRateDown := 0.0
	if applyConfusion {
		confusionCritResistRateDown = 0.25
	}
	critRate += confusionCritResistRateDown

	tfCritResistRateDown := 0.0
	tfCritDefDown := 0.0
	if applyToothFairyBuff {
		tfCritResistRateDown = 0.15
		tfCritDefDown = -0.15
	}
	critRate += tfCritResistRateDown
	enemyCritDef += tfCritDefDown

	// calculatorForBraveNewWorld := DmgCal.DamageCalculator{
	// 	DamageCalculatorInfo: DmgCal.DamageCalculatorInfo{
	// 		Character:                         character.RegulusStat,
	// 		BaseAttackStats:                   baseAtk,
	// 		ResonanceAttackPercentage:         resoAtkPercent,
	// 		PsychubeAttackPercentage:          0,
	// 		ResonanceFixedAttackStats:         resoAtkFixed,
	// 		PsychubeFixedAttackStats:          psychube.BraveNewWorld.Atk,
	// 		DamageBonus:                       anAnleeDmgBonus,
	// 		EnemyDefenseValue:                 enemyDef,
	// 		DefenseBonus:                      0,
	// 		DefenseReduction:                  bkbDefDown,
	// 		PenetrationRate:                   0,
	// 		CasterDamageIncrease:              resonanceStats.DmgBonus,
	// 		TargetDamageTakenReduction:        bkbDmgTakenPlus,
	// 		IncantationUltimateRitualBonusDmg: 0,
	// 		CasterCriticalRate:                baseCrit,
	// 		CasterCriticalDamageBonus:         char.CritDmg + resonanceStats.CritDmg + regulusExcessCritDmgBonus(baseCrit),
	// 		TargetCriticalDefense:             enemyCritDef + tfCritDefDown,
	// 		AfflatusAdvantage:                 afflatusAdvantage,
	// 	},
	// }

	// calculatorForBraveNewWorld.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	// finalDamageBraveNewWorldSkill1 := calculatorForBraveNewWorld.CalculateFinalDamage()
	// calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = psychube.BraveNewWorld.AdditionalEffect.IncantationMight
	// finalDamageBraveNewWorldBuffSkill1 := calculatorForBraveNewWorld.CalculateFinalDamage()

	// calculatorForBraveNewWorld.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	// calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = 0
	// finalDamageBraveNewWorldSkill2 := calculatorForBraveNewWorld.CalculateFinalDamage()
	// calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = psychube.BraveNewWorld.AdditionalEffect.IncantationMight
	// finalDamageBraveNewWorldBuffSkill2 := calculatorForBraveNewWorld.CalculateFinalDamage()

	// calculatorForBraveNewWorld.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	// calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = psychube.BraveNewWorld.UltimateMight
	// finalDamageBraveNewWorldUlt := calculatorForBraveNewWorld.CalculateFinalDamage()

	// // Reset
	// calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = 0
	// calculatorForBraveNewWorld.CasterCriticalRate = baseCrit + regulusCritRateBonus
	// calculatorForBraveNewWorld.CasterCriticalDamageBonus = char.CritDmg + resonanceStats.CritDmg + regulusExcessCritDmgBonus(baseCrit+regulusCritRateBonus)

	// calculatorForBraveNewWorld.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	// finalDamageBraveNewWorldWithRestlessHeartSkill1 := calculatorForBraveNewWorld.CalculateFinalDamage()
	// calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = psychube.BraveNewWorld.AdditionalEffect.IncantationMight
	// finalDamageBraveNewWorldWithRestlessHeartBuffSkill1 := calculatorForBraveNewWorld.CalculateFinalDamage()

	// calculatorForBraveNewWorld.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	// calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = 0
	// finalDamageBraveNewWorldWithRestlessHeartSkill2 := calculatorForBraveNewWorld.CalculateFinalDamage()
	// calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = psychube.BraveNewWorld.AdditionalEffect.IncantationMight
	// finalDamageBraveNewWorldWithRestlessHeartBuffSkill2 := calculatorForBraveNewWorld.CalculateFinalDamage()

	// calculatorForBraveNewWorld.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	// calculatorForBraveNewWorld.IncantationUltimateRitualBonusDmg = psychube.BraveNewWorld.UltimateMight
	// finalDamageBraveNewWorldWithRestlessHeartUlt := calculatorForBraveNewWorld.CalculateFinalDamage()

	// fmt.Printf("\n---------\nRegulus Brave New World Final Damage:")
	// fmt.Printf("\nSkill 1: %.2f (with BNW Buff %.2f)", finalDamageBraveNewWorldSkill1, finalDamageBraveNewWorldBuffSkill1)
	// fmt.Printf("\nSkill 2: %.2f (with BNW Buff %.2f)", finalDamageBraveNewWorldSkill2, finalDamageBraveNewWorldBuffSkill2)
	// fmt.Printf("\nUltimate: %.2f", finalDamageBraveNewWorldUlt)
	// fmt.Printf("\nSkill 1 with Restless Heart: %.2f (with BNW Buff %.2f)", finalDamageBraveNewWorldWithRestlessHeartSkill1, finalDamageBraveNewWorldWithRestlessHeartBuffSkill1)
	// fmt.Printf("\nSkill 2 with Restless Heart: %.2f (with BNW Buff %.2f)", finalDamageBraveNewWorldWithRestlessHeartSkill2, finalDamageBraveNewWorldWithRestlessHeartBuffSkill2)
	// fmt.Printf("\nUltimate with Restless Heart: %.2f", finalDamageBraveNewWorldWithRestlessHeartUlt)

	// fmt.Println()

	calculatorForBoundenDuty := DmgCal.DamageCalculator{
		Character:                 character.Regulus,
		Psychube:                  psychube.HisBoundenDuty,
		EnemyDef:                  enemyDef,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         afflatusAdvantage,
	}

	skill1Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1)
	skill2Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2)
	ultimateDamages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate)

	fmt.Printf("---------\nRegulus His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])

	skill1Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: regulusExcessCritDmgBonus(calculatorForBoundenDuty.CritRate + calculatorForBoundenDuty.Character.CritRate + calculatorForBoundenDuty.Character.Insight.CritRate + calculatorForBoundenDuty.Psychube.CritRate() + regulusCritRateBonus)}, character.Skill1)
	skill2Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: regulusExcessCritDmgBonus(calculatorForBoundenDuty.CritRate + calculatorForBoundenDuty.Character.CritRate + calculatorForBoundenDuty.Character.Insight.CritRate + calculatorForBoundenDuty.Psychube.CritRate() + regulusCritRateBonus)}, character.Skill2)
	ultimateDamages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: regulusExcessCritDmgBonus(calculatorForBoundenDuty.CritRate + calculatorForBoundenDuty.Character.CritRate + calculatorForBoundenDuty.Character.Insight.CritRate + calculatorForBoundenDuty.Psychube.CritRate() + regulusCritRateBonus)}, character.Ultimate)

	fmt.Printf("\nSkill 1 with Restless Heart: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate with Restless Heart: %.2f", ultimateDamages[0])

	// fmt.Println()

	// calculatorForThunder := DmgCal.DamageCalculator{
	// 	DamageCalculatorInfo: DmgCal.DamageCalculatorInfo{
	// 		BaseAttackStats:                   baseAtk,
	// 		ResonanceAttackPercentage:         resoAtkPercent,
	// 		PsychubeAttackPercentage:          0,
	// 		ResonanceFixedAttackStats:         resoAtkFixed,
	// 		PsychubeFixedAttackStats:          psychube.ThunderousApplause.Atk,
	// 		DamageBonus:                       anAnleeDmgBonus,
	// 		EnemyDefenseValue:                 enemyDef,
	// 		DefenseBonus:                      0,
	// 		DefenseReduction:                  bkbDefDown,
	// 		PenetrationRate:                   0,
	// 		CasterDamageIncrease:              resonanceStats.DmgBonus,
	// 		TargetDamageTakenReduction:        bkbDmgTakenPlus,
	// 		IncantationUltimateRitualBonusDmg: 0,
	// 		CasterCriticalRate:                baseCrit + psychube.ThunderousApplause.CritRate,
	// 		CasterCriticalDamageBonus:         char.CritDmg + resonanceStats.CritDmg + psychube.ThunderousApplause.CritDmg + regulusExcessCritDmgBonus(baseCrit+psychube.ThunderousApplause.CritRate),
	// 		TargetCriticalDefense:             0.1,
	// 		AfflatusAdvantage:                 afflatusAdvantage,
	// 	},
	// }

	// calculatorForThunder.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	// finalDamageThunderSkill1 := calculatorForThunder.CalculateFinalDamage()
	// calculatorForThunder.CasterCriticalDamageBonus = char.CritDmg + resonanceStats.CritDmg + regulusExcessCritDmgBonus(baseCrit+psychube.ThunderousApplause.CritRate+regulusCritRateBonus)
	// calculatorForThunder.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	// finalDamageThunderSkill2 := calculatorForThunder.CalculateFinalDamage()
	// calculatorForThunder.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	// finalDamageThunderUlt := calculatorForThunder.CalculateFinalDamage()

	// calculatorForThunder.CasterCriticalRate = baseCrit + psychube.ThunderousApplause.CritRate + regulusCritRateBonus
	// calculatorForThunder.CasterCriticalDamageBonus = char.CritDmg + resonanceStats.CritDmg + psychube.ThunderousApplause.CritDmg + regulusExcessCritDmgBonus(baseCrit+psychube.ThunderousApplause.CritRate+regulusCritRateBonus)

	// calculatorForThunder.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	// finalDamageThunderWithRestlessHeartSkill1 := calculatorForThunder.CalculateFinalDamage()
	// calculatorForThunder.CasterCriticalDamageBonus = char.CritDmg + resonanceStats.CritDmg + regulusExcessCritDmgBonus(baseCrit+psychube.ThunderousApplause.CritRate+regulusCritRateBonus)
	// calculatorForThunder.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	// finalDamageThunderWithRestlessHeartSkill2 := calculatorForThunder.CalculateFinalDamage()
	// calculatorForThunder.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	// finalDamageThunderWithRestlessHeartUlt := calculatorForThunder.CalculateFinalDamage()

	// fmt.Printf("---------\nRegulus Thunder Final Damage:")
	// fmt.Printf("\nSkill 1: %.2f", finalDamageThunderSkill1)
	// fmt.Printf("\nSkill 2: %.2f", finalDamageThunderSkill2)
	// fmt.Printf("\nUltimate: %.2f", finalDamageThunderUlt)
	// fmt.Printf("\nSkill 1 with Restless Heart: %.2f", finalDamageThunderWithRestlessHeartSkill1)
	// fmt.Printf("\nSkill 2 with Restless Heart: %.2f", finalDamageThunderWithRestlessHeartSkill2)
	// fmt.Printf("\nUltimate with Restless Heart: %.2f", finalDamageThunderWithRestlessHeartUlt)

	// fmt.Println()

	// calculatorForHop := DmgCal.DamageCalculator{
	// 	DamageCalculatorInfo: DmgCal.DamageCalculatorInfo{
	// 		BaseAttackStats:                   baseAtk,
	// 		ResonanceAttackPercentage:         resoAtkPercent,
	// 		PsychubeAttackPercentage:          0,
	// 		ResonanceFixedAttackStats:         resoAtkFixed,
	// 		PsychubeFixedAttackStats:          psychube.Hopscotch.Atk,
	// 		DamageBonus:                       anAnleeDmgBonus,
	// 		EnemyDefenseValue:                 enemyDef,
	// 		DefenseBonus:                      0,
	// 		DefenseReduction:                  bkbDefDown,
	// 		PenetrationRate:                   0,
	// 		CasterDamageIncrease:              resonanceStats.DmgBonus,
	// 		TargetDamageTakenReduction:        bkbDmgTakenPlus,
	// 		IncantationUltimateRitualBonusDmg: 0,
	// 		CasterCriticalRate:                baseCrit,
	// 		CasterCriticalDamageBonus:         char.CritDmg + resonanceStats.CritDmg + regulusExcessCritDmgBonus(baseCrit),
	// 		TargetCriticalDefense:             0.1,
	// 		AfflatusAdvantage:                 afflatusAdvantage,
	// 	},
	// }

	// calculatorForHop.IncantationUltimateRitualBonusDmg = psychube.Hopscotch.IncantationMight
	// calculatorForHop.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	// finalDamageHopSkill1 := calculatorForHop.CalculateFinalDamage()
	// calculatorForHop.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	// finalDamageHopSkill2 := calculatorForHop.CalculateFinalDamage()
	// calculatorForHop.IncantationUltimateRitualBonusDmg = 0
	// calculatorForHop.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	// finalDamageHopUlt := calculatorForHop.CalculateFinalDamage()

	// calculatorForHop.CasterCriticalRate = baseCrit + regulusCritRateBonus
	// calculatorForHop.CasterCriticalDamageBonus = char.CritDmg + resonanceStats.CritDmg + regulusExcessCritDmgBonus(baseCrit+regulusCritRateBonus)

	// calculatorForHop.IncantationUltimateRitualBonusDmg = psychube.Hopscotch.IncantationMight
	// calculatorForHop.SkillMultiplier = char.Skill[character.Skill1][1].Multiplier
	// finalDamageHopWithRestlessHeartSkill1 := calculatorForHop.CalculateFinalDamage()
	// calculatorForHop.SkillMultiplier = char.Skill[character.Skill2][1].Multiplier
	// finalDamageHopWithRestlessHeartSkill2 := calculatorForHop.CalculateFinalDamage()
	// calculatorForHop.IncantationUltimateRitualBonusDmg = 0
	// calculatorForHop.SkillMultiplier = char.Skill[character.Ultimate][0].Multiplier
	// finalDamageHopWithRestlessHeartUlt := calculatorForHop.CalculateFinalDamage()

	// fmt.Printf("---------\nRegulus Hopscotch Final Damage:")
	// fmt.Printf("\nSkill 1: %.2f", finalDamageHopSkill1)
	// fmt.Printf("\nSkill 2: %.2f", finalDamageHopSkill2)
	// fmt.Printf("\nUltimate: %.2f", finalDamageHopUlt)
	// fmt.Printf("\nSkill 1 with Restless Heart: %.2f", finalDamageHopWithRestlessHeartSkill1)
	// fmt.Printf("\nSkill 2 with Restless Heart: %.2f", finalDamageHopWithRestlessHeartSkill2)
	// fmt.Printf("\nUltimate with Restless Heart: %.2f", finalDamageHopWithRestlessHeartUlt)
}

func regulusExcessCritDmgBonus(critRate float64) float64 {
	if critRate > 1.0 {
		return critRate - 1.0
	}
	return 0.0
}
