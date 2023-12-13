package main

import (
	"fmt"

	DmgCal "github.com/aliceblock/re1999dmg/damage_calculator"
	"github.com/aliceblock/re1999dmg/damage_calculator/character"
	"github.com/aliceblock/re1999dmg/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg/damage_calculator/psychube"
)

func main() {
	calculatorFunc := calculator[Charlie]
	calculatorFunc(CalParams{
		enemyHit:       1,
		psychubeAmp:    psychube.Amp1,
		resonanceIndex: 1,
		enemyDef:       600.0,
	})
}

type CharacterIndex int16

const (
	Regulus CharacterIndex = 0
	AKnight CharacterIndex = 1
	Lilya   CharacterIndex = 2
	Eagle   CharacterIndex = 3
	Jessica CharacterIndex = 4
	Charlie CharacterIndex = 5
)

var calculator = map[CharacterIndex]func(enemyHit CalParams){
	Regulus: regulusDmgCalculate,
	AKnight: aKnightDmgCalculate,
	Lilya:   lilyaDmgCalculate,
	Eagle:   eagleDmgCalculate,
	Jessica: jessicaDmgCalculate,
	Charlie: charlieDmgCalculate,
}

type CalParams struct {
	enemyHit            int16
	psychubeAmp         psychube.Amplification
	resonanceIndex      int16
	afflatusAdvantage   bool
	enemyDef            float64
	applyAnAnLeeBuff    bool
	applyBkbBuff        bool
	applyConfusion      bool
	applyToothFairyBuff bool
	applySenseWeakness  bool
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
func regulusDmgCalculate(calParams CalParams) {
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
	if calParams.applyAnAnLeeBuff {
		anAnleeDmgBonus = 0.15
	}
	dmgBonus += anAnleeDmgBonus

	bkbDmgTakenPlus := 0.0
	if calParams.applyBkbBuff {
		bkbDmgTakenPlus = -0.15
	}
	enemyDamageTakenReduction += bkbDmgTakenPlus

	confusionCritResistRateDown := 0.0
	if calParams.applyConfusion {
		confusionCritResistRateDown = 0.25
	}
	critRate += confusionCritResistRateDown

	tfCritResistRateDown := 0.0
	tfCritDefDown := 0.0
	if calParams.applyToothFairyBuff {
		tfCritResistRateDown = 0.15
		tfCritDefDown = 0.15
	}
	critRate += tfCritResistRateDown
	enemyCritDef -= tfCritDefDown

	resonance := resonance.Resonance{
		Ideas: []resonance.IdeaAmount{
			{Idea: resonance.RegulusBaseIdea, Amount: 1},
			{Idea: resonance.C4LIdea, Amount: 2},
			{Idea: resonance.C4IOTIdea, Amount: 2},
			{Idea: resonance.C4SIdea, Amount: 3},
			{Idea: resonance.C4JIdea, Amount: 1},
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
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill1, calParams.enemyHit)
	skill2Damages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	skill2BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill2, calParams.enemyHit)
	ultimateDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, calParams.enemyHit)

	fmt.Printf("---------\nRegulus Brave New World Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])

	skill1RestlessDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill1, calParams.enemyHit)
	skill1RestlessBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight(), CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill1, calParams.enemyHit)
	skill2RestlessDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2, calParams.enemyHit)
	skill2RestlessBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight(), CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2, calParams.enemyHit)
	ultimateRestlessDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Ultimate, calParams.enemyHit)

	fmt.Printf("\nSkill 1 with Restless Heart: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill1RestlessDamages[0], skill1RestlessDamages[1], skill1RestlessDamages[2], skill1RestlessBuffDamages[0], skill1RestlessBuffDamages[1], skill1RestlessBuffDamages[2])
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill2RestlessDamages[0], skill2RestlessDamages[1], skill2RestlessDamages[2], skill2RestlessBuffDamages[0], skill2RestlessBuffDamages[1], skill2RestlessBuffDamages[2])
	fmt.Printf("\nUltimate with Restless Heart: %.2f", ultimateRestlessDamages[0])

	expectTotalDamage := skill1RestlessDamages[character.Star2]*1 + skill2Damages[character.Star1]*1 + ultimateRestlessDamages[character.Star1]*3 + skill1RestlessBuffDamages[character.Star2]*2

	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForBoundenDuty := DmgCal.DamageCalculator{
		Character:                 character.Regulus,
		Psychube:                  &psychube.HisBoundenDuty,
		Resonance:                 &resonance,
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill2Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, calParams.enemyHit)

	fmt.Printf("---------\nRegulus His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])

	skill1RestlessDamages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill1, calParams.enemyHit)
	skill2RestlessDamages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2, calParams.enemyHit)
	ultimateRestlessDamages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + regulusCritRateBonus)}, character.Ultimate, calParams.enemyHit)

	fmt.Printf("\nSkill 1 with Restless Heart: %.2f, %.2f, %.2f", skill1RestlessDamages[0], skill1RestlessDamages[1], skill1RestlessDamages[2])
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f, %.2f, %.2f", skill2RestlessDamages[0], skill2RestlessDamages[1], skill2RestlessDamages[2])
	fmt.Printf("\nUltimate with Restless Heart: %.2f", ultimateRestlessDamages[0])

	expectTotalDamage = skill1RestlessDamages[character.Star2]*3 + skill2Damages[character.Star1]*1 + ultimateRestlessDamages[character.Star1]*3

	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForThunder := DmgCal.DamageCalculator{
		Character:                 character.Regulus,
		Psychube:                  &psychube.ThunderousApplause,
		Resonance:                 &resonance,
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.Skill1, calParams.enemyHit)
	skill2Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, calParams.enemyHit)

	fmt.Printf("---------\nRegulus Thunderous Applause Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])

	skill1RestlessDamages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg() + DmgCal.ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate()+regulusCritRateBonus)}, character.Skill1, calParams.enemyHit)
	skill2RestlessDamages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2, calParams.enemyHit)
	ultimateRestlessDamages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate() + regulusCritRateBonus)}, character.Ultimate, calParams.enemyHit)

	fmt.Printf("\nSkill 1 with Restless Heart: %.2f, %.2f, %.2f", skill1RestlessDamages[0], skill1RestlessDamages[1], skill1RestlessDamages[2])
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f, %.2f, %.2f", skill2RestlessDamages[0], skill2RestlessDamages[1], skill2RestlessDamages[2])
	fmt.Printf("\nUltimate with Restless Heart: %.2f", ultimateRestlessDamages[0])

	expectTotalDamage = skill1RestlessDamages[character.Star2]*3 + skill2Damages[character.Star1]*1 + ultimateRestlessDamages[character.Star1]*3

	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)
}

func aKnightDmgCalculate(calParams CalParams) {
	critRate := 0.0

	enemyDef := 600.0
	enemyCritDef := 0.1

	// Buff/Debuff
	dmgBonus := 0.0
	enemyDefReduction := 0.0
	enemyDamageTakenReduction := 0.0

	anAnleeDmgBonus := 0.0
	if calParams.applyAnAnLeeBuff {
		anAnleeDmgBonus = 0.15
	}
	dmgBonus += anAnleeDmgBonus

	bkbDefReduction := 0.0
	bkbDmgTakenPlus := 0.0
	if calParams.applyBkbBuff {
		bkbDefReduction = 0.15
		bkbDmgTakenPlus = -0.15
	}
	enemyDefReduction += bkbDefReduction
	enemyDamageTakenReduction += bkbDmgTakenPlus

	confusionCritResistRateDown := 0.0
	if calParams.applyConfusion {
		confusionCritResistRateDown = 0.25
	}
	critRate += confusionCritResistRateDown

	tfCritResistRateDown := 0.0
	tfCritDefDown := 0.0
	if calParams.applyToothFairyBuff {
		tfCritResistRateDown = 0.15
		tfCritDefDown = 0.15
	}
	critRate += tfCritResistRateDown
	enemyCritDef -= tfCritDefDown

	resonance := resonance.Resonance{
		Ideas: []resonance.IdeaAmount{
			{Idea: resonance.AKnightBaseIdea, Amount: 1},
			{Idea: resonance.C4LIdea, Amount: 3},
			{Idea: resonance.C4IOTIdea, Amount: 6},
			{Idea: resonance.C4SIdea, Amount: 2},
		},
	}

	character.AKnight.SetInsightLevel(character.Insight3L60)

	balancePleaseDmgBonus := dmgBonus
	if !calParams.afflatusAdvantage {
		balancePleaseDmgBonus += psychube.BalancePlease.AdditionalEffect()[calParams.psychubeAmp].DmgBonus()
	}
	calculatorForBalancePlease := DmgCal.DamageCalculator{
		Character:                 character.AKnight,
		Psychube:                  &psychube.BalancePlease,
		Resonance:                 &resonance,
		BuffDmgBonus:              balancePleaseDmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages := calculatorForBalancePlease.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill2Damages := calculatorForBalancePlease.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	ultimateDamages := calculatorForBalancePlease.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, calParams.enemyHit)
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
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BalancePlease.AdditionalEffect()[calParams.psychubeAmp].DmgBonus()}, character.Skill1, calParams.enemyHit)
	skill2Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	skill2BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BalancePlease.AdditionalEffect()[calParams.psychubeAmp].DmgBonus()}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, calParams.enemyHit)
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
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill2Damages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, calParams.enemyHit)
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
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill2Damages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, calParams.enemyHit)
	ultimateBuff1Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{UltimateMight: 0.04 * 1}, character.Ultimate, calParams.enemyHit)
	ultimateBuff2Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{UltimateMight: 0.04 * 2}, character.Ultimate, calParams.enemyHit)
	ultimateBuff3Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{UltimateMight: 0.04 * 3}, character.Ultimate, calParams.enemyHit)
	ultimateBuff4Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{UltimateMight: 0.04 * 4}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage = aKnightCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nA Knight Hopsotch Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f with Hopscotch buff (%.2f, %.2f, %.2f, %.2f)", ultimateDamages[0], ultimateBuff1Damages[0], ultimateBuff2Damages[0], ultimateBuff3Damages[0], ultimateBuff4Damages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)
}

func lilyaDmgCalculate(calParams CalParams) {
	critRate := 0.1

	enemyDef := 600.0
	enemyCritDef := 0.1

	// Buff/Debuff
	dmgBonus := 0.0
	enemyDefReduction := 0.0
	enemyDamageTakenReduction := 0.0

	anAnleeDmgBonus := 0.0
	if calParams.applyAnAnLeeBuff {
		anAnleeDmgBonus = 0.15
	}
	dmgBonus += anAnleeDmgBonus

	bkbDefReduction := 0.0
	bkbDmgTakenPlus := 0.0
	if calParams.applyBkbBuff {
		bkbDefReduction = 0.15
		bkbDmgTakenPlus = -0.15
	}
	enemyDefReduction += bkbDefReduction
	enemyDamageTakenReduction += bkbDmgTakenPlus

	confusionCritResistRateDown := 0.0
	if calParams.applyConfusion {
		confusionCritResistRateDown = 0.25
	}
	critRate += confusionCritResistRateDown

	tfCritResistRateDown := 0.0
	tfCritDefDown := 0.0
	if calParams.applyToothFairyBuff {
		tfCritResistRateDown = 0.15
		tfCritDefDown = 0.15
	}
	critRate += tfCritResistRateDown
	enemyCritDef -= tfCritDefDown

	resonances := []resonance.Resonance{
		{
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.LilyaBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 2},
				{Idea: resonance.C4LIdea, Amount: 2},
				{Idea: resonance.C4SIdea, Amount: 1},
				{Idea: resonance.C4JIdea, Amount: 3},
				{Idea: resonance.C3Idea, Amount: 3},
				{Idea: resonance.C2Idea, Amount: 1},
				{Idea: resonance.C1Idea, Amount: 1},
			},
		}, {
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.LilyaBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 5},
				{Idea: resonance.C4LIdea, Amount: 2},
				{Idea: resonance.C4SIdea, Amount: 1},
				{Idea: resonance.C4JIdea, Amount: 3},
			},
		}, {
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.LilyaBaseIdea, Amount: 1},
				{Idea: resonance.C4LIdea, Amount: 3},
				{Idea: resonance.C4IOTIdea, Amount: 6},
				{Idea: resonance.C4SIdea, Amount: 2},
			},
		},
	}

	character.Lilya.SetInsightLevel(character.Insight3L60)

	calculatorForReso1 := DmgCal.DamageCalculator{
		Character:                 character.Lilya,
		Psychube:                  &psychube.ThunderousApplause,
		Resonance:                 &resonances[0],
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate + psychube.ThunderousApplause.CritRate(), // Insight 3
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages := calculatorForReso1.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages := calculatorForReso1.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.ExtraAction1, calParams.enemyHit)
	totalCritRate := calculatorForReso1.GetTotalCritRate()
	fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
	for index, damage := range skill1ExtraDamages {
		if totalCritRate >= 1.0 {
			skill1Damages[index] += damage
		} else {
			skill1Damages[index] += damage * totalCritRate
		}
	}
	skill2Damages := calculatorForReso1.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	ultimateDamages := calculatorForReso1.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage := basicCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nLilya Thunderous Applause resonance 1 Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForReso2 := DmgCal.DamageCalculator{
		Character:                 character.Lilya,
		Psychube:                  &psychube.ThunderousApplause,
		Resonance:                 &resonances[1],
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate + psychube.ThunderousApplause.CritRate(), // Insight 3
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForReso2.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages = calculatorForReso2.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.ExtraAction1, calParams.enemyHit)
	totalCritRate = calculatorForReso2.GetTotalCritRate()
	fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
	for index, damage := range skill1ExtraDamages {
		if totalCritRate >= 1.0 {
			skill1Damages[index] += damage
		} else {
			skill1Damages[index] += damage * totalCritRate
		}
	}
	skill2Damages = calculatorForReso2.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForReso2.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nLilya Thunderous Applause resonance 2 Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForReso3 := DmgCal.DamageCalculator{
		Character:                 character.Lilya,
		Psychube:                  &psychube.ThunderousApplause,
		Resonance:                 &resonances[2],
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate + psychube.ThunderousApplause.CritRate(), // Insight 3
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForReso3.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages = calculatorForReso3.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.ExtraAction1, calParams.enemyHit)
	totalCritRate = calculatorForReso3.GetTotalCritRate()
	fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
	for index, damage := range skill1ExtraDamages {
		if totalCritRate >= 1.0 {
			skill1Damages[index] += damage
		} else {
			skill1Damages[index] += damage * totalCritRate
		}
	}
	skill2Damages = calculatorForReso3.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForReso3.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nLilya Thunderous Applause resonance 3 Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForReso1HBD := DmgCal.DamageCalculator{
		Character:                 character.Lilya,
		Psychube:                  &psychube.HisBoundenDuty,
		Resonance:                 &resonances[0],
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForReso1HBD.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages = calculatorForReso1HBD.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	totalCritRate = calculatorForReso1HBD.GetTotalCritRate()
	fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
	for index, damage := range skill1ExtraDamages {
		if totalCritRate >= 1.0 {
			skill1Damages[index] += damage
		} else {
			skill1Damages[index] += damage * totalCritRate
		}
	}
	skill2Damages = calculatorForReso1HBD.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForReso1HBD.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nLilya His Bounden Duty resonance 1 Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForReso2HBD := DmgCal.DamageCalculator{
		Character:                 character.Lilya,
		Psychube:                  &psychube.HisBoundenDuty,
		Resonance:                 &resonances[1],
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForReso2HBD.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages = calculatorForReso2HBD.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	totalCritRate = calculatorForReso2HBD.GetTotalCritRate()
	fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
	for index, damage := range skill1ExtraDamages {
		if totalCritRate >= 1.0 {
			skill1Damages[index] += damage
		} else {
			skill1Damages[index] += damage * totalCritRate
		}
	}
	skill2Damages = calculatorForReso2HBD.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForReso2HBD.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nLilya His Bounden Duty resonance 2 Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForReso1Lux := DmgCal.DamageCalculator{
		Character:                 character.Lilya,
		Psychube:                  &psychube.LuxuriousLeisure,
		Resonance:                 &resonances[0],
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1Lux1Damages := calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Skill1, calParams.enemyHit)
	skill1Lux2Damages := calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Skill1, calParams.enemyHit)
	skill1Lux3Damages := calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages = calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	skill1ExtraLux1Damages := calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	skill1ExtraLux2Damages := calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	skill1ExtraLux3Damages := calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	totalCritRate = calculatorForReso1Lux.GetTotalCritRate()
	fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
	for index, damage := range skill1ExtraDamages {
		if totalCritRate >= 1.0 {
			skill1Damages[index] += damage
		} else {
			skill1Damages[index] += damage * totalCritRate
		}
	}
	for index, damage := range skill1ExtraLux1Damages {
		if totalCritRate >= 1.0 {
			skill1Lux1Damages[index] += damage
		} else {
			skill1Lux1Damages[index] += damage * totalCritRate
		}
	}
	for index, damage := range skill1ExtraLux2Damages {
		if totalCritRate >= 1.0 {
			skill1Lux2Damages[index] += damage
		} else {
			skill1Lux2Damages[index] += damage * totalCritRate
		}
	}
	for index, damage := range skill1ExtraLux3Damages {
		if totalCritRate >= 1.0 {
			skill1Lux3Damages[index] += damage
		} else {
			skill1Lux3Damages[index] += damage * totalCritRate
		}
	}
	skill2Damages = calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	skill2Lux1Damages := calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Skill2, calParams.enemyHit)
	skill2Lux2Damages := calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Skill2, calParams.enemyHit)
	skill2Lux3Damages := calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2}, character.Ultimate, calParams.enemyHit)
	ultimateLux1Damages := calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Ultimate, calParams.enemyHit)
	ultimateLux2Damages := calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Ultimate, calParams.enemyHit)
	ultimateLux3Damages := calculatorForReso1Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage = skill1Damages[character.Star1]*2 + skill2Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + ultimateLux1Damages[character.Star1]*1 + skill2Lux2Damages[character.Star3]*1 + skill1Lux2Damages[character.Star3] + ultimateLux2Damages[character.Star1]*1
	/*
		Skill1(1) x2
		Skill2(2) x1
		Ultimate x1
		Skill2(2) x1
		Ultimate x1 + Skill1(3) x1
		Ultimate x1
		Skill1(3) x1
		Skill2(2) x1
	*/
	expectTotalDamage = skill1Damages[character.Star1]*2 + skill2Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill2Lux1Damages[character.Star2]*1 + ultimateLux1Damages[character.Star1]*1 + skill1Lux2Damages[character.Star3]*1 + ultimateLux2Damages[character.Star1]*1 + skill1Lux3Damages[character.Star3]*1 + skill2Lux2Damages[character.Star2]*1

	fmt.Printf("---------\nLilya Luxurious Leisure resonance 1 Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 1 Buff 1 stack: %.2f, %.2f, %.2f", skill1Lux1Damages[0], skill1Lux1Damages[1], skill1Lux1Damages[2])
	fmt.Printf("\nSkill 1 Buff 2 stack: %.2f, %.2f, %.2f", skill1Lux2Damages[0], skill1Lux2Damages[1], skill1Lux2Damages[2])
	fmt.Printf("\nSkill 1 Buff 3 stack: %.2f, %.2f, %.2f", skill1Lux3Damages[0], skill1Lux3Damages[1], skill1Lux3Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nSkill 2 Buff 1 stack: %.2f, %.2f, %.2f", skill2Lux1Damages[0], skill2Lux1Damages[1], skill2Lux1Damages[2])
	fmt.Printf("\nSkill 2 Buff 2 stack: %.2f, %.2f, %.2f", skill2Lux2Damages[0], skill2Lux2Damages[1], skill2Lux2Damages[2])
	fmt.Printf("\nSkill 2 Buff 3 stack: %.2f, %.2f, %.2f", skill2Lux3Damages[0], skill2Lux3Damages[1], skill2Lux3Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nUltimate Buff 1 stack: %.2f", ultimateLux1Damages[0])
	fmt.Printf("\nUltimate Buff 2 stack: %.2f", ultimateLux2Damages[0])
	fmt.Printf("\nUltimate Buff 3 stack: %.2f", ultimateLux3Damages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForReso2Lux := DmgCal.DamageCalculator{
		Character:                 character.Lilya,
		Psychube:                  &psychube.LuxuriousLeisure,
		Resonance:                 &resonances[1],
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1Lux1Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Skill1, calParams.enemyHit)
	skill1Lux2Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Skill1, calParams.enemyHit)
	skill1Lux3Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	skill1ExtraLux1Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	skill1ExtraLux2Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	skill1ExtraLux3Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	totalCritRate = calculatorForReso2Lux.GetTotalCritRate()
	fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
	for index, damage := range skill1ExtraDamages {
		if totalCritRate >= 1.0 {
			skill1Damages[index] += damage
		} else {
			skill1Damages[index] += damage * totalCritRate
		}
	}
	for index, damage := range skill1ExtraLux1Damages {
		if totalCritRate >= 1.0 {
			skill1Lux1Damages[index] += damage
		} else {
			skill1Lux1Damages[index] += damage * totalCritRate
		}
	}
	for index, damage := range skill1ExtraLux2Damages {
		if totalCritRate >= 1.0 {
			skill1Lux2Damages[index] += damage
		} else {
			skill1Lux2Damages[index] += damage * totalCritRate
		}
	}
	for index, damage := range skill1ExtraLux3Damages {
		if totalCritRate >= 1.0 {
			skill1Lux3Damages[index] += damage
		} else {
			skill1Lux3Damages[index] += damage * totalCritRate
		}
	}
	skill2Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	skill2Lux1Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Skill2, calParams.enemyHit)
	skill2Lux2Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Skill2, calParams.enemyHit)
	skill2Lux3Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2}, character.Ultimate, calParams.enemyHit)
	ultimateLux1Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Ultimate, calParams.enemyHit)
	ultimateLux2Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Ultimate, calParams.enemyHit)
	ultimateLux3Damages = calculatorForReso2Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage = skill1Damages[character.Star1]*2 + skill2Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + ultimateLux1Damages[character.Star1]*1 + skill2Lux2Damages[character.Star3]*1 + skill1Lux2Damages[character.Star3] + ultimateLux2Damages[character.Star1]*1
	/*
		Skill1(1) x2
		Skill2(2) x1
		Ultimate x1
		Skill2(2) x1
		Ultimate x1 + Skill1(3) x1
		Ultimate x1
		Skill1(3) x1
		Skill2(2) x1
	*/
	expectTotalDamage = skill1Damages[character.Star1]*2 + skill2Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill2Lux1Damages[character.Star2]*1 + ultimateLux1Damages[character.Star1]*1 + skill1Lux2Damages[character.Star3]*1 + ultimateLux2Damages[character.Star1]*1 + skill1Lux3Damages[character.Star3]*1 + skill2Lux2Damages[character.Star2]*1

	fmt.Printf("---------\nLilya Luxurious Leisure resonance 2 Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 1 Buff 1 stack: %.2f, %.2f, %.2f", skill1Lux1Damages[0], skill1Lux1Damages[1], skill1Lux1Damages[2])
	fmt.Printf("\nSkill 1 Buff 2 stack: %.2f, %.2f, %.2f", skill1Lux2Damages[0], skill1Lux2Damages[1], skill1Lux2Damages[2])
	fmt.Printf("\nSkill 1 Buff 3 stack: %.2f, %.2f, %.2f", skill1Lux3Damages[0], skill1Lux3Damages[1], skill1Lux3Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nSkill 2 Buff 1 stack: %.2f, %.2f, %.2f", skill2Lux1Damages[0], skill2Lux1Damages[1], skill2Lux1Damages[2])
	fmt.Printf("\nSkill 2 Buff 2 stack: %.2f, %.2f, %.2f", skill2Lux2Damages[0], skill2Lux2Damages[1], skill2Lux2Damages[2])
	fmt.Printf("\nSkill 2 Buff 3 stack: %.2f, %.2f, %.2f", skill2Lux3Damages[0], skill2Lux3Damages[1], skill2Lux3Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nUltimate Buff 1 stack: %.2f", ultimateLux1Damages[0])
	fmt.Printf("\nUltimate Buff 2 stack: %.2f", ultimateLux2Damages[0])
	fmt.Printf("\nUltimate Buff 3 stack: %.2f", ultimateLux3Damages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForReso3Lux := DmgCal.DamageCalculator{
		Character:                 character.Lilya,
		Psychube:                  &psychube.LuxuriousLeisure,
		Resonance:                 &resonances[2],
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1Lux1Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Skill1, calParams.enemyHit)
	skill1Lux2Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Skill1, calParams.enemyHit)
	skill1Lux3Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	skill1ExtraLux1Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	skill1ExtraLux2Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	skill1ExtraLux3Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.ExtraAction1, calParams.enemyHit)
	totalCritRate = calculatorForReso3Lux.GetTotalCritRate()
	fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
	for index, damage := range skill1ExtraDamages {
		if totalCritRate >= 1.0 {
			skill1Damages[index] += damage
		} else {
			skill1Damages[index] += damage * totalCritRate
		}
	}
	for index, damage := range skill1ExtraLux1Damages {
		if totalCritRate >= 1.0 {
			skill1Lux1Damages[index] += damage
		} else {
			skill1Lux1Damages[index] += damage * totalCritRate
		}
	}
	for index, damage := range skill1ExtraLux2Damages {
		if totalCritRate >= 1.0 {
			skill1Lux2Damages[index] += damage
		} else {
			skill1Lux2Damages[index] += damage * totalCritRate
		}
	}
	for index, damage := range skill1ExtraLux3Damages {
		if totalCritRate >= 1.0 {
			skill1Lux3Damages[index] += damage
		} else {
			skill1Lux3Damages[index] += damage * totalCritRate
		}
	}
	skill2Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	skill2Lux1Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Skill2, calParams.enemyHit)
	skill2Lux2Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Skill2, calParams.enemyHit)
	skill2Lux3Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2}, character.Ultimate, calParams.enemyHit)
	ultimateLux1Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Ultimate, calParams.enemyHit)
	ultimateLux2Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Ultimate, calParams.enemyHit)
	ultimateLux3Damages = calculatorForReso3Lux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage = skill1Damages[character.Star1]*2 + skill2Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + ultimateLux1Damages[character.Star1]*1 + skill2Lux2Damages[character.Star3]*1 + skill1Lux2Damages[character.Star3] + ultimateLux2Damages[character.Star1]*1
	/*
		Skill1(1) x2
		Skill2(2) x1
		Ultimate x1
		Skill2(2) x1
		Ultimate x1 + Skill1(3) x1
		Ultimate x1
		Skill1(3) x1
		Skill2(2) x1
	*/
	expectTotalDamage = skill1Damages[character.Star1]*2 + skill2Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill2Lux1Damages[character.Star2]*1 + ultimateLux1Damages[character.Star1]*1 + skill1Lux2Damages[character.Star3]*1 + ultimateLux2Damages[character.Star1]*1 + skill1Lux3Damages[character.Star3]*1 + skill2Lux2Damages[character.Star2]*1

	fmt.Printf("---------\nLilya Luxurious Leisure resonance 2 Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 1 Buff 1 stack: %.2f, %.2f, %.2f", skill1Lux1Damages[0], skill1Lux1Damages[1], skill1Lux1Damages[2])
	fmt.Printf("\nSkill 1 Buff 2 stack: %.2f, %.2f, %.2f", skill1Lux2Damages[0], skill1Lux2Damages[1], skill1Lux2Damages[2])
	fmt.Printf("\nSkill 1 Buff 3 stack: %.2f, %.2f, %.2f", skill1Lux3Damages[0], skill1Lux3Damages[1], skill1Lux3Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nSkill 2 Buff 1 stack: %.2f, %.2f, %.2f", skill2Lux1Damages[0], skill2Lux1Damages[1], skill2Lux1Damages[2])
	fmt.Printf("\nSkill 2 Buff 2 stack: %.2f, %.2f, %.2f", skill2Lux2Damages[0], skill2Lux2Damages[1], skill2Lux2Damages[2])
	fmt.Printf("\nSkill 2 Buff 3 stack: %.2f, %.2f, %.2f", skill2Lux3Damages[0], skill2Lux3Damages[1], skill2Lux3Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nUltimate Buff 1 stack: %.2f", ultimateLux1Damages[0])
	fmt.Printf("\nUltimate Buff 2 stack: %.2f", ultimateLux2Damages[0])
	fmt.Printf("\nUltimate Buff 3 stack: %.2f", ultimateLux3Damages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)
}

func eagleDmgCalculate(calParams CalParams) {
	critRate := 0.0

	enemyCritDef := 0.1

	// Buff/Debuff
	dmgBonus := 0.0
	enemyDefReduction := 0.0
	enemyDamageTakenReduction := 0.0

	anAnleeDmgBonus := 0.0
	if calParams.applyAnAnLeeBuff {
		anAnleeDmgBonus = 0.15
	}
	dmgBonus += anAnleeDmgBonus

	bkbDefReduction := 0.0
	bkbDmgTakenPlus := 0.0
	if calParams.applyBkbBuff {
		bkbDefReduction = 0.15
		bkbDmgTakenPlus = -0.15
	}
	enemyDefReduction += bkbDefReduction
	enemyDamageTakenReduction += bkbDmgTakenPlus

	confusionCritResistRateDown := 0.0
	if calParams.applyConfusion {
		confusionCritResistRateDown = 0.25
	}
	critRate += confusionCritResistRateDown

	tfCritResistRateDown := 0.0
	tfCritDefDown := 0.0
	if calParams.applyToothFairyBuff {
		tfCritResistRateDown = 0.15
		tfCritDefDown = 0.15
	}
	critRate += tfCritResistRateDown
	enemyCritDef -= tfCritDefDown

	senseDefReduction := 0.0
	senseCritDefDown := 0.0
	if calParams.applySenseWeakness {
		senseDefReduction = 0.2
		senseCritDefDown = 0.2
	}
	enemyDefReduction += senseDefReduction
	enemyCritDef -= senseCritDefDown

	resonance := resonance.Resonance{
		Ideas: []resonance.IdeaAmount{
			{Idea: resonance.EagleBaseIdea, Amount: 1},
			{Idea: resonance.C4IOTIdea, Amount: 2},
			{Idea: resonance.C4LIdea, Amount: 2},
			{Idea: resonance.C4SIdea, Amount: 1},
			{Idea: resonance.C4JIdea, Amount: 3},
			{Idea: resonance.C3Idea, Amount: 3},
			{Idea: resonance.C2Idea, Amount: 1},
			{Idea: resonance.C1Idea, Amount: 1},
		},
	}

	character.Eagle.SetInsightLevel(character.Insight2L50)

	calculatorForBoundenDuty := DmgCal.DamageCalculator{
		Character:                 character.Eagle,
		Psychube:                  &psychube.HisBoundenDuty,
		Resonance:                 &resonance,
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  calParams.enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.enemyHit)
	skill2Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.enemyHit)
	ultimateDamages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0)}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage := basicCalculateExpectTotalDmg(skill1ExtraDamages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nEagle His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForHopscotch := DmgCal.DamageCalculator{
		Character:                 character.Eagle,
		Psychube:                  &psychube.Hopscotch,
		Resonance:                 &resonance,
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  calParams.enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.enemyHit)
	skill2Damages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0)}, character.Ultimate, calParams.enemyHit)
	ultimateBuff1Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.psychubeAmp].UltimateMight() * 1}, character.Ultimate, calParams.enemyHit)
	ultimateBuff2Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.psychubeAmp].UltimateMight() * 2}, character.Ultimate, calParams.enemyHit)
	ultimateBuff3Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.psychubeAmp].UltimateMight() * 3}, character.Ultimate, calParams.enemyHit)
	ultimateBuff4Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.psychubeAmp].UltimateMight() * 4}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1ExtraDamages, skill2Damages, ultimateDamages)
	/*
		Skill1(1) x2
		Skill2(2) x1
		Ultimate x1
		Skill2(2) x1
		Ultimate x1 + Skill1(3) x1
		Ultimate x1
		Skill1(3) x1
		Skill2(2) x1
	*/
	expectTotalBuffDamage := skill1ExtraDamages[character.Star1]*2 + skill1ExtraDamages[character.Star3]*2 + skill2Damages[character.Star2]*3 + ultimateDamages[character.Star1]*1 + ultimateBuff1Damages[character.Star1]*1 + ultimateBuff2Damages[character.Star1]*1

	fmt.Printf("---------\nEagle Hopscotch Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f with Hopscotch buff (%.2f, %.2f, %.2f, %.2f)", ultimateDamages[0], ultimateBuff1Damages[0], ultimateBuff2Damages[0], ultimateBuff3Damages[0], ultimateBuff4Damages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)
	fmt.Printf("\nExpect with buff total damage: %.2f", expectTotalBuffDamage)

	fmt.Println()

	calculatorForThunder := DmgCal.DamageCalculator{
		Character:                 character.Eagle,
		Psychube:                  &psychube.ThunderousApplause,
		Resonance:                 &resonance,
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  calParams.enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.enemyHit)
	skill2Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate()+1.0) + psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1ExtraDamages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nEagle Thunderous Applause Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForBraveNewWorld := DmgCal.DamageCalculator{
		Character:                 character.Eagle,
		Psychube:                  &psychube.BraveNewWorld,
		Resonance:                 &resonance,
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  calParams.enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.enemyHit)
	skill1ExtraBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill1, calParams.enemyHit)
	skill2Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.enemyHit)
	skill2BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0)}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage = skill1ExtraDamages[character.Star1]*2 + skill2Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill2BuffDamages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill1ExtraBuffDamages[character.Star3]*1 + ultimateDamages[character.Star1]*1 + skill1ExtraBuffDamages[character.Star3]*1 + skill2Damages[character.Star2]*1

	fmt.Printf("---------\nEagle Brave New World Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 1 with BNW: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2], skill1ExtraBuffDamages[0], skill1ExtraBuffDamages[1], skill1ExtraBuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nSkill 2 with BNW: %.2f, %.2f, %.2f", skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForLux := DmgCal.DamageCalculator{
		Character:                 character.Eagle,
		Psychube:                  &psychube.LuxuriousLeisure,
		Resonance:                 &resonance,
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  calParams.enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1Lux1Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Skill1, calParams.enemyHit)
	skill1Lux2Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Skill1, calParams.enemyHit)
	skill1Lux3Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages = calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.enemyHit)
	skill1ExtraLux1Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Skill1, calParams.enemyHit)
	skill1ExtraLux2Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Skill1, calParams.enemyHit)
	skill1ExtraLux3Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Skill1, calParams.enemyHit)
	skill2Damages = calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.enemyHit)
	skill2Lux1Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Skill2, calParams.enemyHit)
	skill2Lux2Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Skill2, calParams.enemyHit)
	skill2Lux3Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0)}, character.Ultimate, calParams.enemyHit)
	ultimateLux1Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Ultimate, calParams.enemyHit)
	ultimateLux2Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Ultimate, calParams.enemyHit)
	ultimateLux3Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Ultimate, calParams.enemyHit)
	/*
		Skill1(1) x2
		Skill2(2) x1
		Ultimate x1
		Skill2(2) x1
		Ultimate x1 + Skill1(3) x1
		Ultimate x1
		Skill1(3) x1
		Skill2(2) x1
	*/
	expectTotalDamage = skill1ExtraDamages[character.Star1]*2 + skill2Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill2Lux1Damages[character.Star2]*1 + ultimateLux1Damages[character.Star1]*1 + skill1ExtraLux2Damages[character.Star3]*1 + ultimateLux2Damages[character.Star1]*1 + skill1ExtraLux3Damages[character.Star3]*1 + skill2Lux2Damages[character.Star2]*1

	fmt.Printf("---------\nEagle Luxurious Leisure Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 1 Lux1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Lux1Damages[0], skill1Lux1Damages[1], skill1Lux1Damages[2], skill1ExtraLux1Damages[0], skill1ExtraLux1Damages[1], skill1ExtraLux1Damages[2])
	fmt.Printf("\nSkill 1 Lux2: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Lux2Damages[0], skill1Lux2Damages[1], skill1Lux2Damages[2], skill1ExtraLux2Damages[0], skill1ExtraLux2Damages[1], skill1ExtraLux2Damages[2])
	fmt.Printf("\nSkill 1 Lux3: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Lux3Damages[0], skill1Lux3Damages[1], skill1Lux3Damages[2], skill1ExtraLux3Damages[0], skill1ExtraLux3Damages[1], skill1ExtraLux3Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nSkill 2 Lux1: %.2f, %.2f, %.2f", skill2Lux1Damages[0], skill2Lux1Damages[1], skill2Lux1Damages[2])
	fmt.Printf("\nSkill 2 Lux2: %.2f, %.2f, %.2f", skill2Lux2Damages[0], skill2Lux2Damages[1], skill2Lux2Damages[2])
	fmt.Printf("\nSkill 2 Lux3: %.2f, %.2f, %.2f", skill2Lux3Damages[0], skill2Lux3Damages[1], skill2Lux3Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nUltimate Lux1: %.2f", ultimateLux1Damages[0])
	fmt.Printf("\nUltimate Lux2: %.2f", ultimateLux2Damages[0])
	fmt.Printf("\nUltimate Lux3: %.2f", ultimateLux3Damages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()
}

func jessicaDmgCalculate(calParams CalParams) {
	critRate := 0.0

	enemyCritDef := 0.1

	// Buff/Debuff
	dmgBonus := 0.2 // Insight 1
	enemyDefReduction := 0.0
	enemyDamageTakenReduction := 0.0

	anAnleeDmgBonus := 0.0
	if calParams.applyAnAnLeeBuff {
		anAnleeDmgBonus = 0.15
	}
	dmgBonus += anAnleeDmgBonus

	bkbDefReduction := 0.0
	bkbDmgTakenPlus := 0.0
	if calParams.applyBkbBuff {
		bkbDefReduction = 0.15
		bkbDmgTakenPlus = -0.15
	}
	enemyDefReduction += bkbDefReduction
	enemyDamageTakenReduction += bkbDmgTakenPlus

	confusionCritResistRateDown := 0.0
	if calParams.applyConfusion {
		confusionCritResistRateDown = 0.25
	}
	critRate += confusionCritResistRateDown

	tfCritResistRateDown := 0.0
	tfCritDefDown := 0.0
	if calParams.applyToothFairyBuff {
		tfCritResistRateDown = 0.15
		tfCritDefDown = 0.15
	}
	critRate += tfCritResistRateDown
	enemyCritDef -= tfCritDefDown

	senseDefReduction := 0.0
	senseCritDefDown := 0.0
	if calParams.applySenseWeakness {
		senseDefReduction = 0.2
		senseCritDefDown = 0.2
	}
	enemyDefReduction += senseDefReduction
	enemyCritDef -= senseCritDefDown

	resonances := []resonance.Resonance{
		{
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.JessicaBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 6},
				{Idea: resonance.C4SIdea, Amount: 2},
				{Idea: resonance.C3Idea, Amount: 3},
				{Idea: resonance.C1Idea, Amount: 2},
			},
		},
		{
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.JessicaBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 2},
				{Idea: resonance.C4LIdea, Amount: 2},
				{Idea: resonance.C4SIdea, Amount: 3},
				{Idea: resonance.C4JIdea, Amount: 1},
				{Idea: resonance.C3Idea, Amount: 3},
				{Idea: resonance.C2Idea, Amount: 1},
				{Idea: resonance.C1Idea, Amount: 1},
			},
		},
		{
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.JessicaBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 6},
				{Idea: resonance.C4LIdea, Amount: 2},
				{Idea: resonance.C4SIdea, Amount: 3},
			},
		},
	}

	character.Jessica.SetInsightLevel(character.Insight3L60)

	calculatorForBoundenDuty := DmgCal.DamageCalculator{
		Character:                 character.Jessica,
		Psychube:                  &psychube.HisBoundenDuty,
		Resonance:                 &resonances[calParams.resonanceIndex],
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  calParams.enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1Extra1Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.enemyHit)
	skill1Extra2Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.enemyHit)
	skill1Extra3Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.enemyHit)
	skill2Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	skill2ExtraDamages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill2, calParams.enemyHit)
	ultimateDamages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, calParams.enemyHit)
	poisonDamage := calculatorForBoundenDuty.CalculateGenesisDamage(DmgCal.DamageCalculatorInfo{}, 0.3)
	expectTotalDamage := basicCalculateExpectTotalDmg(skill1Extra2Damages, skill2ExtraDamages, ultimateDamages) + poisonDamage*(6+6+3)

	fmt.Printf("---------\nJessica His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 1 Extra 1: %.2f, %.2f, %.2f", skill1Extra1Damages[0], skill1Extra1Damages[1], skill1Extra1Damages[2])
	fmt.Printf("\nSkill 1 Extra 2: %.2f, %.2f, %.2f", skill1Extra2Damages[0], skill1Extra2Damages[1], skill1Extra2Damages[2])
	fmt.Printf("\nSkill 1 Extra 3: %.2f, %.2f, %.2f", skill1Extra3Damages[0], skill1Extra3Damages[1], skill1Extra3Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (Extra: %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nPoison stack/round: %.2f", poisonDamage)
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForHopscotch := DmgCal.DamageCalculator{
		Character:                 character.Jessica,
		Psychube:                  &psychube.Hopscotch,
		Resonance:                 &resonances[calParams.resonanceIndex],
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  calParams.enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1Extra1Damages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.enemyHit)
	skill1Extra2Damages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.enemyHit)
	skill1Extra3Damages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.enemyHit)
	skill2Damages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	skill2ExtraDamages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, calParams.enemyHit)
	ultimateBuff1Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.psychubeAmp].UltimateMight() * 1}, character.Ultimate, calParams.enemyHit)
	ultimateBuff2Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.psychubeAmp].UltimateMight() * 2}, character.Ultimate, calParams.enemyHit)
	ultimateBuff3Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.psychubeAmp].UltimateMight() * 3}, character.Ultimate, calParams.enemyHit)
	ultimateBuff4Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.psychubeAmp].UltimateMight() * 4}, character.Ultimate, calParams.enemyHit)
	poisonDamage = calculatorForHopscotch.CalculateGenesisDamage(DmgCal.DamageCalculatorInfo{}, 0.3)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1Extra2Damages, skill2ExtraDamages, ultimateDamages) + poisonDamage*(6+6+3)
	/*
		Skill1(1) x2
		Skill2(2) x1
		Ultimate x1
		Skill2(2) x1
		Ultimate x1 + Skill1(3) x1
		Ultimate x1
		Skill1(3) x1
		Skill2(2) x1
	*/
	expectTotalBuffDamage := skill1Extra2Damages[character.Star1]*2 + skill1Extra2Damages[character.Star3]*2 + skill2ExtraDamages[character.Star2]*3 + ultimateDamages[character.Star1]*1 + ultimateBuff1Damages[character.Star1]*1 + ultimateBuff2Damages[character.Star1]*1 + poisonDamage*(6+6+3)

	fmt.Printf("---------\nJessica Hopscotch Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 1 Extra 1: %.2f, %.2f, %.2f", skill1Extra1Damages[0], skill1Extra1Damages[1], skill1Extra1Damages[2])
	fmt.Printf("\nSkill 1 Extra 2: %.2f, %.2f, %.2f", skill1Extra2Damages[0], skill1Extra2Damages[1], skill1Extra2Damages[2])
	fmt.Printf("\nSkill 1 Extra 3: %.2f, %.2f, %.2f", skill1Extra3Damages[0], skill1Extra3Damages[1], skill1Extra3Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (Extra: %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2])
	fmt.Printf("\nUltimate: %.2f with Hopscotch buff (%.2f, %.2f, %.2f, %.2f)", ultimateDamages[0], ultimateBuff1Damages[0], ultimateBuff2Damages[0], ultimateBuff3Damages[0], ultimateBuff4Damages[0])
	fmt.Printf("\nPoison stack/round: %.2f", poisonDamage)
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)
	fmt.Printf("\nExpect with buff total damage: %.2f", expectTotalBuffDamage)

	fmt.Println()

	calculatorForBraveNewWorld := DmgCal.DamageCalculator{
		Character:                 character.Jessica,
		Psychube:                  &psychube.BraveNewWorld,
		Resonance:                 &resonances[calParams.resonanceIndex],
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  calParams.enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill1, calParams.enemyHit)
	skill1Extra1Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.enemyHit)
	skill1Extra1BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 0, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill1, calParams.enemyHit)
	skill1Extra2Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.enemyHit)
	skill1Extra2BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 1, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill1, calParams.enemyHit)
	skill1Extra3Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.enemyHit)
	skill1Extra3BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 2, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill1, calParams.enemyHit)
	skill2Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	skill2BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill2, calParams.enemyHit)
	skill2ExtraDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill2, calParams.enemyHit)
	skill2ExtraBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, calParams.enemyHit)
	poisonDamage = calculatorForBraveNewWorld.CalculateGenesisDamage(DmgCal.DamageCalculatorInfo{}, 0.3)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1Extra2Damages, skill2ExtraDamages, ultimateDamages) + poisonDamage*(6+6+3)
	/*
		Skill1(1) x2
		Skill2(2) x1
		Ultimate x1
		Skill2(2) x1
		Ultimate x1 + Skill1(3) x1
		Ultimate x1
		Skill1(3) x1
		Skill2(2) x1
	*/
	expectTotalDamage = skill1Extra2Damages[character.Star1]*2 + skill2ExtraDamages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill2ExtraBuffDamages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill1Extra2BuffDamages[character.Star3]*1 + ultimateDamages[character.Star1]*1 + skill1Extra2BuffDamages[character.Star3]*1 + skill2ExtraDamages[character.Star2]*1 + poisonDamage*(6+6+3)

	fmt.Printf("---------\nJessica Brave New World Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with BNW %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2])
	fmt.Printf("\nSkill 1 Extra 1: %.2f, %.2f, %.2f (with BNW %.2f, %.2f, %.2f)", skill1Extra1Damages[0], skill1Extra1Damages[1], skill1Extra1Damages[2], skill1Extra1BuffDamages[0], skill1Extra1BuffDamages[1], skill1Extra1BuffDamages[2])
	fmt.Printf("\nSkill 1 Extra 2: %.2f, %.2f, %.2f (with BNW %.2f, %.2f, %.2f)", skill1Extra2Damages[0], skill1Extra2Damages[1], skill1Extra2Damages[2], skill1Extra2BuffDamages[0], skill1Extra2BuffDamages[1], skill1Extra2BuffDamages[2])
	fmt.Printf("\nSkill 1 Extra 3: %.2f, %.2f, %.2f (with BNW %.2f, %.2f, %.2f)", skill1Extra3Damages[0], skill1Extra3Damages[1], skill1Extra3Damages[2], skill1Extra3BuffDamages[0], skill1Extra3BuffDamages[1], skill1Extra3BuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with BNW: %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2])
	fmt.Printf("\nSkill 2 Extra: %.2f, %.2f, %.2f (with BNW: %.2f, %.2f, %.2f)", skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2], skill2ExtraBuffDamages[0], skill2ExtraBuffDamages[1], skill2ExtraBuffDamages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nPoison stack/round: %.2f", poisonDamage)
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	calculatorForBlasphemer := DmgCal.DamageCalculator{
		Character:                 character.Jessica,
		Psychube:                  &psychube.BlasphemerOfNight,
		Resonance:                 &resonances[calParams.resonanceIndex],
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  calParams.enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages = calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1BuffDamages = calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.psychubeAmp].DmgBonus()}, character.Skill1, calParams.enemyHit)
	skill1Extra1Damages = calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.enemyHit)
	skill1Extra1BuffDamages = calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 0, BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.psychubeAmp].DmgBonus()}, character.Skill1, calParams.enemyHit)
	skill1Extra2Damages = calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.enemyHit)
	skill1Extra2BuffDamages = calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 1, BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.psychubeAmp].DmgBonus()}, character.Skill1, calParams.enemyHit)
	skill1Extra3Damages = calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.enemyHit)
	skill1Extra3BuffDamages = calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 2, BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.psychubeAmp].DmgBonus()}, character.Skill1, calParams.enemyHit)
	skill2Damages = calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	skill2BuffDamages = calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.psychubeAmp].DmgBonus()}, character.Skill2, calParams.enemyHit)
	skill2ExtraDamages = calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill2, calParams.enemyHit)
	skill2ExtraBuffDamages = calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.psychubeAmp].DmgBonus()}, character.Skill2, calParams.enemyHit)
	ultimateDamages = calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Ultimate, calParams.enemyHit)
	ultimateBuffDamages := calculatorForBlasphemer.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.psychubeAmp].DmgBonus()}, character.Ultimate, calParams.enemyHit)
	poisonDamage = calculatorForBlasphemer.CalculateGenesisDamage(DmgCal.DamageCalculatorInfo{}, 0.3)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1Extra2BuffDamages, skill2ExtraBuffDamages, ultimateBuffDamages) + poisonDamage*(6+6+3)

	fmt.Printf("---------\nJessica Blasphemer Of Night Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Buff %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2])
	fmt.Printf("\nSkill 1 Extra 1: %.2f, %.2f, %.2f (with Buff %.2f, %.2f, %.2f)", skill1Extra1Damages[0], skill1Extra1Damages[1], skill1Extra1Damages[2], skill1Extra1BuffDamages[0], skill1Extra1BuffDamages[1], skill1Extra1BuffDamages[2])
	fmt.Printf("\nSkill 1 Extra 2: %.2f, %.2f, %.2f (with Buff %.2f, %.2f, %.2f)", skill1Extra2Damages[0], skill1Extra2Damages[1], skill1Extra2Damages[2], skill1Extra2BuffDamages[0], skill1Extra2BuffDamages[1], skill1Extra2BuffDamages[2])
	fmt.Printf("\nSkill 1 Extra 3: %.2f, %.2f, %.2f (with Buff %.2f, %.2f, %.2f)", skill1Extra3Damages[0], skill1Extra3Damages[1], skill1Extra3Damages[2], skill1Extra3BuffDamages[0], skill1Extra3BuffDamages[1], skill1Extra3BuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with Buff: %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2])
	fmt.Printf("\nSkill 2 Extra: %.2f, %.2f, %.2f (with Buff: %.2f, %.2f, %.2f)", skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2], skill2ExtraBuffDamages[0], skill2ExtraBuffDamages[1], skill2ExtraBuffDamages[2])
	fmt.Printf("\nUltimate: %.2f (with Buff: %.2f)", ultimateDamages[0], ultimateBuffDamages[0])
	fmt.Printf("\nPoison stack/round: %.2f", poisonDamage)
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()
}

func charlieDmgCalculate(calParams CalParams) {
	critRate := 0.0

	enemyCritDef := 0.1

	// Buff/Debuff
	dmgBonus := 0.2 // Insight 3
	enemyDefReduction := 0.0
	enemyDamageTakenReduction := 0.0

	anAnleeDmgBonus := 0.0
	if calParams.applyAnAnLeeBuff {
		anAnleeDmgBonus = 0.15
	}
	dmgBonus += anAnleeDmgBonus

	bkbDefReduction := 0.0
	bkbDmgTakenPlus := 0.0
	if calParams.applyBkbBuff {
		bkbDmgTakenPlus = -0.15
	}
	enemyDefReduction += bkbDefReduction
	enemyDamageTakenReduction += bkbDmgTakenPlus

	confusionCritResistRateDown := 0.0
	if calParams.applyConfusion {
		confusionCritResistRateDown = 0.25
	}
	critRate += confusionCritResistRateDown

	tfCritResistRateDown := 0.0
	tfCritDefDown := 0.0
	if calParams.applyToothFairyBuff {
		tfCritResistRateDown = 0.15
		tfCritDefDown = 0.15
	}
	critRate += tfCritResistRateDown
	enemyCritDef -= tfCritDefDown

	resonances := []resonance.Resonance{
		{
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.CharlieBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 6},
				{Idea: resonance.C4SIdea, Amount: 1},
				{Idea: resonance.C4JIdea, Amount: 1},
				{Idea: resonance.C3Idea, Amount: 3},
				{Idea: resonance.C2Idea, Amount: 1},
				{Idea: resonance.C1Idea, Amount: 1},
			},
		},
		{
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.CharlieBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 6},
				{Idea: resonance.C4LIdea, Amount: 2},
				{Idea: resonance.C4SIdea, Amount: 3},
			},
		},
	}

	character.Charlie.SetInsightLevel(character.Insight3L60)

	calculatorForBoundenDuty := DmgCal.DamageCalculator{
		Character:                 character.Charlie,
		Psychube:                  &psychube.HisBoundenDuty,
		Resonance:                 &resonances[calParams.resonanceIndex],
		BuffDmgBonus:              dmgBonus,
		EnemyDef:                  calParams.enemyDef,
		EnemyDefReduction:         enemyDefReduction,
		EnemyDamageTakenReduction: enemyDamageTakenReduction,
		CritRate:                  critRate,
		EnemyCritDef:              enemyCritDef,
		AfflatusAdvantage:         calParams.afflatusAdvantage,
	}

	skill1Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	skill1ExtraDamages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.enemyHit)
	skill2Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill2, calParams.enemyHit)
	skill2ExtraDamages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill2, calParams.enemyHit)
	ultimateDamages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0)}, character.Ultimate, calParams.enemyHit)
	expectTotalDamage := basicCalculateExpectTotalDmg(skill1ExtraDamages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nEagle His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	fmt.Println()

	// calculatorForHopscotch := DmgCal.DamageCalculator{
	// 	Character:                 character.Eagle,
	// 	Psychube:                  &psychube.Hopscotch,
	// 	Resonance:                 &resonance,
	// 	BuffDmgBonus:              dmgBonus,
	// 	EnemyDef:                  calParams.enemyDef,
	// 	EnemyDefReduction:         enemyDefReduction,
	// 	EnemyDamageTakenReduction: enemyDamageTakenReduction,
	// 	CritRate:                  critRate,
	// 	EnemyCritDef:              enemyCritDef,
	// 	AfflatusAdvantage:         calParams.afflatusAdvantage,
	// }

	// skill1Damages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	// skill1ExtraDamages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.enemyHit)
	// skill2Damages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.enemyHit)
	// ultimateDamages = calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0)}, character.Ultimate, calParams.enemyHit)
	// ultimateBuff1Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.psychubeAmp].UltimateMight() * 1}, character.Ultimate, calParams.enemyHit)
	// ultimateBuff2Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.psychubeAmp].UltimateMight() * 2}, character.Ultimate, calParams.enemyHit)
	// ultimateBuff3Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.psychubeAmp].UltimateMight() * 3}, character.Ultimate, calParams.enemyHit)
	// ultimateBuff4Damages := calculatorForHopscotch.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.psychubeAmp].UltimateMight() * 4}, character.Ultimate, calParams.enemyHit)
	// expectTotalDamage = basicCalculateExpectTotalDmg(skill1ExtraDamages, skill2Damages, ultimateDamages)
	// /*
	// 	Skill1(1) x2
	// 	Skill2(2) x1
	// 	Ultimate x1
	// 	Skill2(2) x1
	// 	Ultimate x1 + Skill1(3) x1
	// 	Ultimate x1
	// 	Skill1(3) x1
	// 	Skill2(2) x1
	// */
	// expectTotalBuffDamage := skill1ExtraDamages[character.Star1]*2 + skill1ExtraDamages[character.Star3]*2 + skill2Damages[character.Star2]*3 + ultimateDamages[character.Star1]*1 + ultimateBuff1Damages[character.Star1]*1 + ultimateBuff2Damages[character.Star1]*1

	// fmt.Printf("---------\nEagle Hopscotch Final Damage:")
	// fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	// fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	// fmt.Printf("\nUltimate: %.2f with Hopscotch buff (%.2f, %.2f, %.2f, %.2f)", ultimateDamages[0], ultimateBuff1Damages[0], ultimateBuff2Damages[0], ultimateBuff3Damages[0], ultimateBuff4Damages[0])
	// fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)
	// fmt.Printf("\nExpect with buff total damage: %.2f", expectTotalBuffDamage)

	// fmt.Println()

	// calculatorForThunder := DmgCal.DamageCalculator{
	// 	Character:                 character.Eagle,
	// 	Psychube:                  &psychube.ThunderousApplause,
	// 	Resonance:                 &resonance,
	// 	BuffDmgBonus:              dmgBonus,
	// 	EnemyDef:                  calParams.enemyDef,
	// 	EnemyDefReduction:         enemyDefReduction,
	// 	EnemyDamageTakenReduction: enemyDamageTakenReduction,
	// 	CritRate:                  critRate,
	// 	EnemyCritDef:              enemyCritDef,
	// 	AfflatusAdvantage:         calParams.afflatusAdvantage,
	// }

	// skill1Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.Skill1, calParams.enemyHit)
	// skill1ExtraDamages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.enemyHit)
	// skill2Damages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.enemyHit)
	// ultimateDamages = calculatorForThunder.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate()+1.0) + psychube.ThunderousApplause.AdditionalEffect()[calParams.psychubeAmp].CritDmg()}, character.Ultimate, calParams.enemyHit)
	// expectTotalDamage = basicCalculateExpectTotalDmg(skill1ExtraDamages, skill2Damages, ultimateDamages)

	// fmt.Printf("---------\nEagle Thunderous Applause Final Damage:")
	// fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	// fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	// fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	// fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	// fmt.Println()

	// calculatorForBraveNewWorld := DmgCal.DamageCalculator{
	// 	Character:                 character.Eagle,
	// 	Psychube:                  &psychube.BraveNewWorld,
	// 	Resonance:                 &resonance,
	// 	BuffDmgBonus:              dmgBonus,
	// 	EnemyDef:                  calParams.enemyDef,
	// 	EnemyDefReduction:         enemyDefReduction,
	// 	EnemyDamageTakenReduction: enemyDamageTakenReduction,
	// 	CritRate:                  critRate,
	// 	EnemyCritDef:              enemyCritDef,
	// 	AfflatusAdvantage:         calParams.afflatusAdvantage,
	// }

	// skill1Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	// skill1BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill1, calParams.enemyHit)
	// skill1ExtraDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.enemyHit)
	// skill1ExtraBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill1, calParams.enemyHit)
	// skill2Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.enemyHit)
	// skill2BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.psychubeAmp].IncantationMight()}, character.Skill2, calParams.enemyHit)
	// ultimateDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0)}, character.Ultimate, calParams.enemyHit)
	// expectTotalDamage = skill1ExtraDamages[character.Star1]*2 + skill2Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill2BuffDamages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill1ExtraBuffDamages[character.Star3]*1 + ultimateDamages[character.Star1]*1 + skill1ExtraBuffDamages[character.Star3]*1 + skill2Damages[character.Star2]*1

	// fmt.Printf("---------\nEagle Brave New World Final Damage:")
	// fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	// fmt.Printf("\nSkill 1 with BNW: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2], skill1ExtraBuffDamages[0], skill1ExtraBuffDamages[1], skill1ExtraBuffDamages[2])
	// fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	// fmt.Printf("\nSkill 2 with BNW: %.2f, %.2f, %.2f", skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2])
	// fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	// fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	// fmt.Println()

	// calculatorForLux := DmgCal.DamageCalculator{
	// 	Character:                 character.Eagle,
	// 	Psychube:                  &psychube.LuxuriousLeisure,
	// 	Resonance:                 &resonance,
	// 	BuffDmgBonus:              dmgBonus,
	// 	EnemyDef:                  calParams.enemyDef,
	// 	EnemyDefReduction:         enemyDefReduction,
	// 	EnemyDamageTakenReduction: enemyDamageTakenReduction,
	// 	CritRate:                  critRate,
	// 	EnemyCritDef:              enemyCritDef,
	// 	AfflatusAdvantage:         calParams.afflatusAdvantage,
	// }

	// skill1Damages = calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{}, character.Skill1, calParams.enemyHit)
	// skill1Lux1Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Skill1, calParams.enemyHit)
	// skill1Lux2Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Skill1, calParams.enemyHit)
	// skill1Lux3Damages := calculatorForBoundenDuty.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Skill1, calParams.enemyHit)
	// skill1ExtraDamages = calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.enemyHit)
	// skill1ExtraLux1Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Skill1, calParams.enemyHit)
	// skill1ExtraLux2Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Skill1, calParams.enemyHit)
	// skill1ExtraLux3Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Skill1, calParams.enemyHit)
	// skill2Damages = calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.enemyHit)
	// skill2Lux1Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Skill2, calParams.enemyHit)
	// skill2Lux2Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Skill2, calParams.enemyHit)
	// skill2Lux3Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{PenetrationRate: 0.4, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Skill2, calParams.enemyHit)
	// ultimateDamages = calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0)}, character.Ultimate, calParams.enemyHit)
	// ultimateLux1Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 1}, character.Ultimate, calParams.enemyHit)
	// ultimateLux2Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 2}, character.Ultimate, calParams.enemyHit)
	// ultimateLux3Damages := calculatorForLux.CalculateFinalDamage(DmgCal.DamageCalculatorInfo{CritRate: 1.0, CritDmg: DmgCal.ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.psychubeAmp].DmgBonus() * 3}, character.Ultimate, calParams.enemyHit)
	// /*
	// 	Skill1(1) x2
	// 	Skill2(2) x1
	// 	Ultimate x1
	// 	Skill2(2) x1
	// 	Ultimate x1 + Skill1(3) x1
	// 	Ultimate x1
	// 	Skill1(3) x1
	// 	Skill2(2) x1
	// */
	// expectTotalDamage = skill1ExtraDamages[character.Star1]*2 + skill2Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill2Lux1Damages[character.Star2]*1 + ultimateLux1Damages[character.Star1]*1 + skill1ExtraLux2Damages[character.Star3]*1 + ultimateLux2Damages[character.Star1]*1 + skill1ExtraLux3Damages[character.Star3]*1 + skill2Lux2Damages[character.Star2]*1

	// fmt.Printf("---------\nEagle Luxurious Leisure Final Damage:")
	// fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	// fmt.Printf("\nSkill 1 Lux1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Lux1Damages[0], skill1Lux1Damages[1], skill1Lux1Damages[2], skill1ExtraLux1Damages[0], skill1ExtraLux1Damages[1], skill1ExtraLux1Damages[2])
	// fmt.Printf("\nSkill 1 Lux2: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Lux2Damages[0], skill1Lux2Damages[1], skill1Lux2Damages[2], skill1ExtraLux2Damages[0], skill1ExtraLux2Damages[1], skill1ExtraLux2Damages[2])
	// fmt.Printf("\nSkill 1 Lux3: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Lux3Damages[0], skill1Lux3Damages[1], skill1Lux3Damages[2], skill1ExtraLux3Damages[0], skill1ExtraLux3Damages[1], skill1ExtraLux3Damages[2])
	// fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	// fmt.Printf("\nSkill 2 Lux1: %.2f, %.2f, %.2f", skill2Lux1Damages[0], skill2Lux1Damages[1], skill2Lux1Damages[2])
	// fmt.Printf("\nSkill 2 Lux2: %.2f, %.2f, %.2f", skill2Lux2Damages[0], skill2Lux2Damages[1], skill2Lux2Damages[2])
	// fmt.Printf("\nSkill 2 Lux3: %.2f, %.2f, %.2f", skill2Lux3Damages[0], skill2Lux3Damages[1], skill2Lux3Damages[2])
	// fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	// fmt.Printf("\nUltimate Lux1: %.2f", ultimateLux1Damages[0])
	// fmt.Printf("\nUltimate Lux2: %.2f", ultimateLux2Damages[0])
	// fmt.Printf("\nUltimate Lux3: %.2f", ultimateLux3Damages[0])
	// fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	// fmt.Println()
}

/*
Skill1(1) x2
Skill1(3) x2
Skill2(2) x3
Ultimate x3
*/
func basicCalculateExpectTotalDmg(skill1Damages []float64, skill2Damages []float64, ultimateDamages []float64) float64 {
	return skill1Damages[character.Star1]*2 + skill1Damages[character.Star3]*2 + skill2Damages[character.Star2]*3 + ultimateDamages[character.Star1]*3
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
