package main

import (
	"fmt"

	DmgCal "github.com/aliceblock/re1999dmg/damage_calculator"
	"github.com/aliceblock/re1999dmg/damage_calculator/character"
	"github.com/aliceblock/re1999dmg/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg/damage_calculator/psychube"
)

var calculator = map[CharacterIndex]func(int16, bool, bool, bool, bool, bool){
	Regulus: regulusDmgCalculate,
	AKnight: aKnightDmgCalculate,
}

func main() {
	calculatorFunc := calculator[AKnight]
	calculatorFunc(1, false, false, false, false, false)
}

type CharacterIndex int16

const (
	Regulus CharacterIndex = 0
	AKnight                = 1
)

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
func regulusDmgCalculate(enemyHit int16, afflatusAdvantage bool, applyAnAnLeeBuff bool, applyBkbBuff bool, applyConfusion bool, applyToothFairyBuff bool) {
	critRate := 0.0

	enemyDef := 600.0
	enemyCritDef := 0.1

	// Additional Bonus
	regulusCritRateBonus := 0.5

	// Buff/Debuff
	dmgBonus := 0.0
	enemyDefReduction := 0.0
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
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         afflatusAdvantage,
	}

	skill1Damages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, enemyHit)
	skill1WithBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect().IncantationMight()}, character.Skill1, enemyHit)
	skill2Damages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, enemyHit)
	skill2WithBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect().IncantationMight()}, character.Skill2, enemyHit)
	ultimateDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, enemyHit)

	fmt.Printf("---------\nRegulus Brave New World Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1WithBuffDamages[0], skill1WithBuffDamages[1], skill1WithBuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2WithBuffDamages[0], skill2WithBuffDamages[1], skill2WithBuffDamages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])

	skill1Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill1, enemyHit)
	skill1WithBuffDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect().IncantationMight(), CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill1, enemyHit)
	skill2Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2, enemyHit)
	skill2WithBuffDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect().IncantationMight(), CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2, enemyHit)
	ultimateDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Ultimate, enemyHit)

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

	skill1Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, enemyHit)
	skill2Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, enemyHit)
	ultimateDamages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, enemyHit)

	fmt.Printf("---------\nRegulus His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])

	skill1Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill1, enemyHit)
	skill2Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2, enemyHit)
	ultimateDamages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + regulusCritRateBonus)}, character.Ultimate, enemyHit)

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

	skill1Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect().CritDmg()}, character.Skill1, enemyHit)
	skill2Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, enemyHit)
	ultimateDamages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, enemyHit)

	fmt.Printf("---------\nRegulus Thunderous Applause Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])

	skill1Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: psychube.ThunderousApplause.AdditionalEffect().CritDmg() + DmgCal.ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate()+regulusCritRateBonus)}, character.Skill1, enemyHit)
	skill2Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2, enemyHit)
	ultimateDamages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate() + regulusCritRateBonus)}, character.Ultimate, enemyHit)

	fmt.Printf("\nSkill 1 with Restless Heart: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate with Restless Heart: %.2f", ultimateDamages[0])
}

func aKnightDmgCalculate(enemyHit int16, afflatusAdvantage bool, applyAnAnLeeBuff bool, applyBkbBuff bool, applyConfusion bool, applyToothFairyBuff bool) {
	critRate := 0.0

	enemyDef := 600.0
	enemyCritDef := 0.1

	// Buff/Debuff
	dmgBonus := 0.0
	enemyDefReduction := 0.0
	enemyDamageTakenReduction := 0.0

	anAnleeDmgBonus := 0.0
	if applyAnAnLeeBuff {
		anAnleeDmgBonus = 0.15
	}
	dmgBonus += anAnleeDmgBonus

	bkbDefReduction := 0.0
	bkbDmgTakenPlus := 0.0
	if applyBkbBuff {
		bkbDefReduction = 0.15
		bkbDmgTakenPlus = -0.15
	}
	enemyDefReduction += bkbDefReduction
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
			{Idea: resonance.AKnightBaseIdea, Amount: 1},
			{Idea: resonance.C4LIdea, Amount: 3},
			{Idea: resonance.C4IIdea, Amount: 2},
			{Idea: resonance.C4SIdea, Amount: 2},
			{Idea: resonance.C4TIdea, Amount: 2},
			{Idea: resonance.C4OIdea, Amount: 2},
		},
	}

	character.AKnight.SetInsightLevel(character.Insight3L60)

	balancePleaseDmgBonus := dmgBonus
	if !afflatusAdvantage {
		balancePleaseDmgBonus += psychube.BalancePlease.AdditionalEffect().DmgBonus()
	}
	calculatorForBalancePlease := DmgCal.DamageCalculator{
		Character:                 character.AKnight,
		Psychube:                  &psychube.BalancePlease,
		Resonance:                 &resonance,
		BuffDmgBonus:              balancePleaseDmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         afflatusAdvantage,
	}

	skill1Damages := calculatorForBalancePlease.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, enemyHit)
	skill2Damages := calculatorForBalancePlease.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, enemyHit)
	ultimateDamages := calculatorForBalancePlease.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, enemyHit)
	expectTotalDamage := aKnightCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nA Knight Balance, Please Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForBraveNewWorld := DmgCal.DamageCalculator{
		Character:                 character.AKnight,
		Psychube:                  &psychube.BraveNewWorld,
		Resonance:                 &resonance,
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         afflatusAdvantage,
	}

	skill1Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, enemyHit)
	skill1BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect().IncantationMight()}, character.Skill1, enemyHit)
	skill2Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, enemyHit)
	skill2BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect().IncantationMight()}, character.Skill2, enemyHit)
	ultimateDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, enemyHit)
	expectTotalDamage = skill1Damages[character.Star1] + skill1Damages[character.Star2] + skill2Damages[character.Star2]*3 + ultimateDamages[character.Star1]*3 + skill1BuffDamages[character.Star2] + skill1BuffDamages[character.Star3]

	fmt.Printf("---------\nA Knight Brave New World Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForBoundenDuty := DmgCal.DamageCalculator{
		Character:                 character.AKnight,
		Psychube:                  &psychube.HisBoundenDuty,
		Resonance:                 &resonance,
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         afflatusAdvantage,
	}

	skill1Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, enemyHit)
	skill2Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, enemyHit)
	ultimateDamages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, enemyHit)
	expectTotalDamage = aKnightCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nA Knight His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForHopscotch := DmgCal.DamageCalculator{
		Character:                 character.AKnight,
		Psychube:                  &psychube.Hopscotch,
		Resonance:                 &resonance,
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         afflatusAdvantage,
	}

	skill1Damages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, enemyHit)
	skill2Damages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, enemyHit)
	ultimateDamages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, enemyHit)
	ultimateBuff1Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{UltimateMight: 0.04 * 1}, character.Ultimate, enemyHit)
	ultimateBuff2Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{UltimateMight: 0.04 * 2}, character.Ultimate, enemyHit)
	ultimateBuff3Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{UltimateMight: 0.04 * 3}, character.Ultimate, enemyHit)
	ultimateBuff4Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{UltimateMight: 0.04 * 4}, character.Ultimate, enemyHit)
	expectTotalDamage = aKnightCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nA Knight Hopsotch Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f with Hopscotch buff (%.2f, %.2f, %.2f, %.2f)", ultimateDamages[0], ultimateBuff1Damages[0], ultimateBuff2Damages[0], ultimateBuff3Damages[0], ultimateBuff4Damages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)
}

/*
Skill1(1) x1
Skill1(2) x2
Skill1(3) x1
Skill2(2) x3
Ultimate x3
*/
func aKnightCalculateExpectTotalDmg(skill1Damages []float64, skill2Damages []float64, ultimateDamages []float64) float64 {
	return skill1Damages[character.Star1] + skill1Damages[character.Star2]*2 + skill1Damages[character.Star3] + skill2Damages[character.Star2]*3 + ultimateDamages[character.Star1]*3
}
