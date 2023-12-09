package main

import (
	"fmt"

	DmgCal "github.com/aliceblock/re1999dmg/damage_calculator"
	"github.com/aliceblock/re1999dmg/damage_calculator/character"
	"github.com/aliceblock/re1999dmg/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg/damage_calculator/psychube"
)

func main() {
	regulusDmgCalculate(false, false, false, false, false)
}

/*
1 :2 Move Merge
2 :5 Use 2 Incantation Skill1(2) + Skill2(1)
3 :5 Wait
4 :1 Ultimate
5 :3 Move Merge
6 :5 Use 1 Incantation Skill1(2)
7 :5 Wait
8 :1 Ultimate
9 :3 Move Merge
10:5 Use 1 Incantation Skill2(2)
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

	resonance := resonance.Resonance{
		Ideas: []resonance.IdeaAmount{
			{Idea: resonance.RegulusBaseIdea, Amount: 1},
			{Idea: resonance.C4LIdea, Amount: 2},
			{Idea: resonance.C4IIdea, Amount: 1},
			{Idea: resonance.C4SIdea, Amount: 3},
			{Idea: resonance.C4JIdea, Amount: 1},
			{Idea: resonance.C4TIdea, Amount: 1},
			{Idea: resonance.C3Idea, Amount: 3},
			{Idea: resonance.C2Idea, Amount: 1},
			{Idea: resonance.C1Idea, Amount: 1},
		},
	}

	character.Regulus.SetInsightLevel(character.Insight3L60)

	calculatorForBraveNewWorld := DmgCal.DamageCalculator{
		Character:                 character.Regulus,
		Psychube:                  &psychube.BraveNewWorld,
		Resonance:                 &resonance,
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         afflatusAdvantage,
	}

	skill1Damages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1)
	skill1WithBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect().IncantationMight()}, character.Skill1)
	skill2Damages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2)
	skill2WithBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect().IncantationMight()}, character.Skill2)
	ultimateDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate)

	fmt.Printf("---------\nRegulus Brave New World Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1WithBuffDamages[0], skill1WithBuffDamages[1], skill1WithBuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2WithBuffDamages[0], skill2WithBuffDamages[1], skill2WithBuffDamages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])

	skill1Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill1)
	skill1WithBuffDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect().IncantationMight(), CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill1)
	skill2Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2)
	skill2WithBuffDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect().IncantationMight(), CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2)
	ultimateDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Ultimate)

	fmt.Printf("\nSkill 1 with Restless Heart: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1WithBuffDamages[0], skill1WithBuffDamages[1], skill1WithBuffDamages[2])
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2WithBuffDamages[0], skill2WithBuffDamages[1], skill2WithBuffDamages[2])
	fmt.Printf("\nUltimate with Restless Heart: %.2f", ultimateDamages[0])

	fmt.Println()

	calculatorForBoundenDuty := DmgCal.DamageCalculator{
		Character:                 character.Regulus,
		Psychube:                  &psychube.HisBoundenDuty,
		Resonance:                 &resonance,
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         afflatusAdvantage,
	}

	skill1Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1)
	skill2Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2)
	ultimateDamages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate)

	fmt.Printf("---------\nRegulus His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])

	skill1Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill1)
	skill2Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2)
	ultimateDamages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + regulusCritRateBonus)}, character.Ultimate)

	fmt.Printf("\nSkill 1 with Restless Heart: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate with Restless Heart: %.2f", ultimateDamages[0])

	fmt.Println()

	calculatorForThunder := DmgCal.DamageCalculator{
		Character:                 character.Regulus,
		Psychube:                  &psychube.ThunderousApplause,
		Resonance:                 &resonance,
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         afflatusAdvantage,
	}

	skill1Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect().CritDmg()}, character.Skill1)
	skill2Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2)
	ultimateDamages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate)

	fmt.Printf("---------\nRegulus Thunderous Applause Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])

	skill1Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: psychube.ThunderousApplause.AdditionalEffect().CritDmg() + DmgCal.ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate()+regulusCritRateBonus)}, character.Skill1)
	skill2Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2)
	ultimateDamages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate() + regulusCritRateBonus)}, character.Ultimate)

	fmt.Printf("\nSkill 1 with Restless Heart: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate with Restless Heart: %.2f", ultimateDamages[0])
}
