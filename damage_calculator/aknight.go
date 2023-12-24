package damage_calculator

import (
	"fmt"

	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/character"
	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/psychube"
)

func AKnightDmgCalculate(calParams CalParams) []DamageResponse {
	damageResponse := []DamageResponse{}

	enemyCritDef := 0.1

	resonance := resonance.Resonance{
		Ideas: []resonance.IdeaAmount{
			{Idea: resonance.AKnightBaseIdea, Amount: 1},
			{Idea: resonance.C4LIdea, Amount: 3},
			{Idea: resonance.C4IOTIdea, Amount: 6},
			{Idea: resonance.C4SIdea, Amount: 2},
		},
	}

	character.AKnight.SetInsightLevel(character.Insight3L60)

	balancePleaseDmgBonus := 0.0
	if !calParams.AfflatusAdvantage {
		balancePleaseDmgBonus += psychube.BalancePlease.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()
	}
	calculatorForBalancePlease := DamageCalculator{
		Character:         character.AKnight,
		Psychube:          &psychube.BalancePlease,
		Resonance:         &resonance,
		BuffDmgBonus:      balancePleaseDmgBonus,
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages := calculatorForBalancePlease.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill2Damages := calculatorForBalancePlease.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	ultimateDamages := calculatorForBalancePlease.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage := aKnightCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nA Knight Balance, Please Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName := calculatorForBalancePlease.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForBraveNewWorld := DamageCalculator{
		Character:         character.AKnight,
		Psychube:          &psychube.BraveNewWorld,
		Resonance:         &resonance,
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.BalancePlease.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.BalancePlease.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = skill1Damages[character.Star1] + skill1Damages[character.Star2] + skill2Damages[character.Star2]*3 + ultimateDamages[character.Star1]*3 + skill1BuffDamages[character.Star2] + skill1BuffDamages[character.Star3]

	fmt.Printf("---------\nA Knight Brave New World Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName = calculatorForBraveNewWorld.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForBoundenDuty := DamageCalculator{
		Character:         character.AKnight,
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
	expectTotalDamage = aKnightCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nA Knight His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
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

	calculatorForHopscotch := DamageCalculator{
		Character:         character.AKnight,
		Psychube:          &psychube.Hopscotch,
		Resonance:         &resonance,
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff1Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{UltimateMight: 0.04 * 1}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff2Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{UltimateMight: 0.04 * 2}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff3Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{UltimateMight: 0.04 * 3}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff4Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{UltimateMight: 0.04 * 4}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = aKnightCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nA Knight Hopsotch Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f with Hopscotch buff (%.2f, %.2f, %.2f, %.2f)", ultimateDamages[0], ultimateBuff1Damages[0], ultimateBuff2Damages[0], ultimateBuff3Damages[0], ultimateBuff4Damages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName = calculatorForHopscotch.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForYearningDesire := DamageCalculator{
		Character:         character.AKnight,
		Psychube:          &psychube.YearningDesire,
		Resonance:         &resonance,
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1BuffDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2BuffDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	ultimateBuffDamages := calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = aKnightCalculateExpectTotalDmg(skill1BuffDamages, skill2BuffDamages, ultimateBuffDamages)

	fmt.Printf("---------\nA Knight Yearning Desire Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Buff %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with Buff %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2])
	fmt.Printf("\nUltimate: %.2f (with Buff %.2f)", ultimateDamages[0], ultimateBuffDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName = calculatorForYearningDesire.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	return damageResponse
}
