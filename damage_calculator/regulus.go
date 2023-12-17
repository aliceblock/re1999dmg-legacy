package damage_calculator

import (
	"fmt"

	"github.com/aliceblock/re1999dmg/damage_calculator/character"
	"github.com/aliceblock/re1999dmg/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg/damage_calculator/psychube"
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
func RegulusDmgCalculate(calParams CalParams) []DamageResponse {
	damageResponse := make([]DamageResponse, 0)

	enemyCritDef := 0.1

	// Additional Bonus
	regulusCritRateBonus := 0.5

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

	calculatorForBraveNewWorld := DamageCalculator{
		Character:         character.Regulus,
		Psychube:          &psychube.BraveNewWorld,
		Resonance:         &resonance,
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
	skill2Damages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill2, calParams.EnemyHit)
	ultimateDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)

	fmt.Printf("---------\nRegulus Brave New World Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])

	skill1RestlessDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill1, calParams.EnemyHit)
	skill1RestlessBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight(), CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill1, calParams.EnemyHit)
	skill2RestlessDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2, calParams.EnemyHit)
	skill2RestlessBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight(), CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2, calParams.EnemyHit)
	ultimateRestlessDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForBraveNewWorld.GetTotalCritRate() + regulusCritRateBonus)}, character.Ultimate, calParams.EnemyHit)

	fmt.Printf("\nSkill 1 with Restless Heart: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill1RestlessDamages[0], skill1RestlessDamages[1], skill1RestlessDamages[2], skill1RestlessBuffDamages[0], skill1RestlessBuffDamages[1], skill1RestlessBuffDamages[2])
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill2RestlessDamages[0], skill2RestlessDamages[1], skill2RestlessDamages[2], skill2RestlessBuffDamages[0], skill2RestlessBuffDamages[1], skill2RestlessBuffDamages[2])
	fmt.Printf("\nUltimate with Restless Heart: %.2f", ultimateRestlessDamages[0])

	expectTotalDamage := skill1RestlessDamages[character.Star2]*1 + skill2Damages[character.Star1]*1 + ultimateRestlessDamages[character.Star1]*3 + skill1RestlessBuffDamages[character.Star2]*2

	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName := calculatorForBraveNewWorld.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForBoundenDuty := DamageCalculator{
		Character:         character.Regulus,
		Psychube:          &psychube.HisBoundenDuty,
		Resonance:         &resonance,
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)

	fmt.Printf("---------\nRegulus His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])

	skill1RestlessDamages = calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill1, calParams.EnemyHit)
	skill2RestlessDamages = calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2, calParams.EnemyHit)
	ultimateRestlessDamages = calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + regulusCritRateBonus)}, character.Ultimate, calParams.EnemyHit)

	fmt.Printf("\nSkill 1 with Restless Heart: %.2f, %.2f, %.2f", skill1RestlessDamages[0], skill1RestlessDamages[1], skill1RestlessDamages[2])
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f, %.2f, %.2f", skill2RestlessDamages[0], skill2RestlessDamages[1], skill2RestlessDamages[2])
	fmt.Printf("\nUltimate with Restless Heart: %.2f", ultimateRestlessDamages[0])

	expectTotalDamage = skill1RestlessDamages[character.Star2]*3 + skill2Damages[character.Star1]*1 + ultimateRestlessDamages[character.Star1]*3

	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName = calculatorForBoundenDuty.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForThunder := DamageCalculator{
		Character:         character.Regulus,
		Psychube:          &psychube.ThunderousApplause,
		Resonance:         &resonance,
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)

	fmt.Printf("---------\nRegulus Thunderous Applause Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])

	skill1RestlessDamages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg() + ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate()+regulusCritRateBonus)}, character.Skill1, calParams.EnemyHit)
	skill2RestlessDamages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2, calParams.EnemyHit)
	ultimateRestlessDamages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate() + regulusCritRateBonus)}, character.Ultimate, calParams.EnemyHit)

	fmt.Printf("\nSkill 1 with Restless Heart: %.2f, %.2f, %.2f", skill1RestlessDamages[0], skill1RestlessDamages[1], skill1RestlessDamages[2])
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f, %.2f, %.2f", skill2RestlessDamages[0], skill2RestlessDamages[1], skill2RestlessDamages[2])
	fmt.Printf("\nUltimate with Restless Heart: %.2f", ultimateRestlessDamages[0])

	expectTotalDamage = skill1RestlessDamages[character.Star2]*3 + skill2Damages[character.Star1]*1 + ultimateRestlessDamages[character.Star1]*3

	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName = calculatorForThunder.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForYearning := DamageCalculator{
		Character:         character.Regulus,
		Psychube:          &psychube.YearningDesire,
		Resonance:         &resonance,
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1BuffDamages = calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2BuffDamages = calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	ultimateBuffDamages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Ultimate, calParams.EnemyHit)

	fmt.Printf("---------\nRegulus Yearning Desire Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Buff %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with Buff %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2])
	fmt.Printf("\nUltimate: %.2f (with Buff %.2f)", ultimateDamages[0], ultimateBuffDamages[0])

	skill1RestlessDamages = calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForYearning.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill1, calParams.EnemyHit)
	skill1RestlessBuffDamages = calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForYearning.GetTotalCritRate() + regulusCritRateBonus), BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill1, calParams.EnemyHit)
	skill2RestlessDamages = calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate() + regulusCritRateBonus)}, character.Skill2, calParams.EnemyHit)
	skill2RestlessBuffDamages = calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate() + regulusCritRateBonus), BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill2, calParams.EnemyHit)
	ultimateRestlessDamages = calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate() + regulusCritRateBonus)}, character.Ultimate, calParams.EnemyHit)
	ultimateRestlessBuffDamages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{CritRate: regulusCritRateBonus, CritDmg: ExcessCritDmgBonus(calculatorForThunder.GetTotalCritRate() + regulusCritRateBonus), BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Ultimate, calParams.EnemyHit)

	fmt.Printf("\nSkill 1 with Restless Heart: %.2f, %.2f, %.2f (with Buff %.2f, %.2f, %.2f)", skill1RestlessDamages[0], skill1RestlessDamages[1], skill1RestlessDamages[2], skill1RestlessBuffDamages[0], skill1RestlessBuffDamages[1], skill1RestlessBuffDamages[2])
	fmt.Printf("\nSkill 2 with Restless Heart: %.2f, %.2f, %.2f (with Buff %.2f, %.2f, %.2f)", skill2RestlessDamages[0], skill2RestlessDamages[1], skill2RestlessDamages[2], skill2RestlessBuffDamages[0], skill2RestlessBuffDamages[1], skill2RestlessBuffDamages[2])
	fmt.Printf("\nUltimate with Restless Heart: %.2f (with Buff %.2f)", ultimateRestlessDamages[0], ultimateRestlessBuffDamages[0])

	expectTotalDamage = skill1RestlessBuffDamages[character.Star2]*3 + skill2BuffDamages[character.Star1]*1 + ultimateRestlessBuffDamages[character.Star1]*3

	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName = calculatorForYearning.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	return damageResponse
}
