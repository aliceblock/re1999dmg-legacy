package damage_calculator

import (
	"fmt"

	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/character"
	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/psychube"
)

func BkornblumeP2DmgCalculate(calParams CalParams) []DamageResponse {
	damageResponse := []DamageResponse{}

	insight1DmgBonus := 0.2
	enemyCritDef := 0.1

	resonances := []resonance.Resonance{
		// {
		// 	Ideas: []resonance.IdeaAmount{
		// 		{Idea: resonance.BkornblumeBaseIdea, Amount: 1},
		// 		{Idea: resonance.C4IOTIdea, Amount: 6},
		// 		{Idea: resonance.C4LIdea, Amount: 3},
		// 		{Idea: resonance.C4SIdea, Amount: 2},
		// 	},
		// },
		{
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.BkornblumeBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 1},
				{Idea: resonance.C4LIdea, Amount: 3},
				{Idea: resonance.C4SIdea, Amount: 2},
				{Idea: resonance.C4JIdea, Amount: 2},
				{Idea: resonance.C3Idea, Amount: 3},
				{Idea: resonance.C2Idea, Amount: 1},
				{Idea: resonance.C1Idea, Amount: 1},
			},
		},
	}

	character.BkornblumeP2.SetInsightLevel(character.Insight3L60)

	for i := 0; i < len(resonances); i++ {
		calculatorForBraveNewWorld := DamageCalculator{
			Character:         character.BkornblumeP2,
			Psychube:          &psychube.BraveNewWorld,
			Resonance:         &resonances[i],
			BuffDmgBonus:      insight1DmgBonus,
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Damages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.EnemyHit)
		skill1BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
		ultimateDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
		expectTotalDamage := skill1Damages[character.Star1]*1 + ultimateDamages[character.Star1]*1 + skill1BuffDamages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill1BuffDamages[character.Star3]*1 + ultimateDamages[character.Star1]*1 + skill1BuffDamages[character.Star3]*1

		fmt.Printf("---------\nBkornblume Brave New World Final Damage:")
		fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with BNW Buff %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2])
		fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
		fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

		psychubeName := calculatorForBraveNewWorld.Psychube.Name()
		if len(resonances) > 1 {
			psychubeName += fmt.Sprintf(" Reso %d", i+1)
		}
		if calParams.PsychubeAmp > 0 {
			psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
		}
		damageResponse = append(damageResponse, DamageResponse{
			Name:   psychubeName,
			Damage: toFixed(expectTotalDamage, 2),
		})

		fmt.Println()
	}

	for i := 0; i < len(resonances); i++ {
		calculatorForHisBounden := DamageCalculator{
			Character:         character.BkornblumeP2,
			Psychube:          &psychube.HisBoundenDuty,
			Resonance:         &resonances[i],
			BuffDmgBonus:      insight1DmgBonus,
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Damages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.EnemyHit)
		ultimateDamages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
		expectTotalDamage := skill1Damages[character.Star1]*1 + ultimateDamages[character.Star1]*1 + skill1Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill1Damages[character.Star3]*1 + ultimateDamages[character.Star1]*1 + skill1Damages[character.Star3]*1

		fmt.Printf("---------\nBkornblume His Bounden Duty Final Damage:")
		fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
		fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
		fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

		psychubeName := calculatorForHisBounden.Psychube.Name()
		if len(resonances) > 1 {
			psychubeName += fmt.Sprintf(" Reso %d", i+1)
		}
		if calParams.PsychubeAmp > 0 {
			psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
		}
		damageResponse = append(damageResponse, DamageResponse{
			Name:   psychubeName,
			Damage: toFixed(expectTotalDamage, 2),
		})

		fmt.Println()
	}

	for i := 0; i < len(resonances); i++ {
		calculatorForLux := DamageCalculator{
			Character:         character.BkornblumeP2,
			Psychube:          &psychube.LuxuriousLeisure,
			Resonance:         &resonances[i],
			BuffDmgBonus:      insight1DmgBonus,
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.EnemyHit)
		skill1Buff1Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Skill1, calParams.EnemyHit)
		skill1Buff2Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Skill1, calParams.EnemyHit)
		skill1Buff3Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Skill1, calParams.EnemyHit)
		skill1Buff4Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 4}, character.Skill1, calParams.EnemyHit)
		ultimateDamages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
		ultimateBuff1Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Ultimate, calParams.EnemyHit)
		ultimateBuff2Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Ultimate, calParams.EnemyHit)
		ultimateBuff3Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Ultimate, calParams.EnemyHit)
		ultimateBuff4Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 4}, character.Ultimate, calParams.EnemyHit)
		expectTotalDamage := skill1Damages[character.Star1]*1 + ultimateDamages[character.Star1]*1 + skill1Buff1Damages[character.Star2]*1 + ultimateBuff1Damages[character.Star1]*1 + skill1Buff2Damages[character.Star3]*1 + ultimateBuff2Damages[character.Star1]*1 + skill1Buff3Damages[character.Star3]*1

		fmt.Printf("---------\nBkornblume Luxurious Leisure Final Damage:")
		fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
		fmt.Printf("\nSkill 1 Buff 1: %.2f, %.2f, %.2f", skill1Buff1Damages[0], skill1Buff1Damages[1], skill1Buff1Damages[2])
		fmt.Printf("\nSkill 1 Buff 2: %.2f, %.2f, %.2f", skill1Buff2Damages[0], skill1Buff2Damages[1], skill1Buff2Damages[2])
		fmt.Printf("\nSkill 1 Buff 3: %.2f, %.2f, %.2f", skill1Buff3Damages[0], skill1Buff3Damages[1], skill1Buff3Damages[2])
		fmt.Printf("\nSkill 1 Buff 4: %.2f, %.2f, %.2f", skill1Buff4Damages[0], skill1Buff4Damages[1], skill1Buff4Damages[2])
		fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
		fmt.Printf("\nUltimate Buff 1: %.2f", ultimateBuff1Damages[0])
		fmt.Printf("\nUltimate Buff 2: %.2f", ultimateBuff2Damages[0])
		fmt.Printf("\nUltimate Buff 3: %.2f", ultimateBuff3Damages[0])
		fmt.Printf("\nUltimate Buff 4: %.2f", ultimateBuff4Damages[0])
		fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

		psychubeName := calculatorForLux.Psychube.Name()
		if len(resonances) > 1 {
			psychubeName += fmt.Sprintf(" Reso %d", i+1)
		}
		if calParams.PsychubeAmp > 0 {
			psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
		}
		damageResponse = append(damageResponse, DamageResponse{
			Name:   psychubeName,
			Damage: toFixed(expectTotalDamage, 2),
		})

		fmt.Println()
	}

	for i := 0; i < len(resonances); i++ {
		calculatorForYearning := DamageCalculator{
			Character:         character.BkornblumeP2,
			Psychube:          &psychube.YearningDesire,
			Resonance:         &resonances[i],
			BuffDmgBonus:      insight1DmgBonus,
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill1, calParams.EnemyHit)
		ultimateDamages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Ultimate, calParams.EnemyHit)
		expectTotalDamage := skill1Damages[character.Star1]*1 + ultimateDamages[character.Star1]*1 + skill1Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill1Damages[character.Star3]*1 + ultimateDamages[character.Star1]*1 + skill1Damages[character.Star3]*1

		fmt.Printf("---------\nBkornblume Yearning Desire Final Damage:")
		fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
		fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
		fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

		psychubeName := calculatorForYearning.Psychube.Name()
		if len(resonances) > 1 {
			psychubeName += fmt.Sprintf(" Reso %d", i+1)
		}
		if calParams.PsychubeAmp > 0 {
			psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
		}
		damageResponse = append(damageResponse, DamageResponse{
			Name:   psychubeName,
			Damage: toFixed(expectTotalDamage, 2),
		})

		fmt.Println()
	}

	return damageResponse
}
