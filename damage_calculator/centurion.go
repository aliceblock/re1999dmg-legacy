package damage_calculator

import (
	"fmt"

	"github.com/aliceblock/re1999dmg/damage_calculator/character"
	"github.com/aliceblock/re1999dmg/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg/damage_calculator/psychube"
)

func CenturionDmgCalculate(calParams CalParams) []DamageResponse {
	damageResponse := []DamageResponse{}

	insight1DmgBonus := 0.06
	enemyCritDef := 0.1

	resonances := []resonance.Resonance{
		{
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.CenturionBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 6},
				{Idea: resonance.C4LIdea, Amount: 2},
				{Idea: resonance.C4JIdea, Amount: 2},
				{Idea: resonance.C4SIdea, Amount: 1},
			},
		},
		{
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.CenturionBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 2},
				{Idea: resonance.C4LIdea, Amount: 2},
				{Idea: resonance.C4JIdea, Amount: 3},
				{Idea: resonance.C4SIdea, Amount: 1},
				{Idea: resonance.C3Idea, Amount: 3},
				{Idea: resonance.C2Idea, Amount: 1},
				{Idea: resonance.C1Idea, Amount: 1},
			},
		},
	}

	character.Centurion.SetInsightLevel(character.Insight3L60)

	for i := 0; i < len(resonances); i++ {
		calculatorForHisBounden := DamageCalculator{
			Character:         character.Centurion,
			Psychube:          &psychube.HisBoundenDuty,
			Resonance:         &resonances[i],
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Damages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
		skill1Moxie1Damages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 1, HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.EnemyHit)
		skill1Moxie2Damages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 2, HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.EnemyHit)
		skill1Moxie3Damages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 3, HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.EnemyHit)
		skill1Moxie4Damages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 4, HasExtraDamage: true, ExtraDamageStack: 3}, character.Skill1, calParams.EnemyHit)
		skill1Moxie5Damages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 5, HasExtraDamage: true, ExtraDamageStack: 4}, character.Skill1, calParams.EnemyHit)

		skill2Damages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
		skill2Moxie1Damages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 1}, character.Skill2, calParams.EnemyHit)
		skill2Moxie2Damages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 2}, character.Skill2, calParams.EnemyHit)
		skill2Moxie3Damages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 3}, character.Skill2, calParams.EnemyHit)
		skill2Moxie4Damages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 4}, character.Skill2, calParams.EnemyHit)
		skill2Moxie5Damages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 5}, character.Skill2, calParams.EnemyHit)

		ultimateDamages := calculatorForHisBounden.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 5}, character.Ultimate, calParams.EnemyHit)

		expectTotalDamage := skill2Moxie2Damages[character.Star2] + skill1Moxie4Damages[character.Star1] + skill2Moxie5Damages[character.Star1] + skill1Moxie5Damages[character.Star3] + skill2Moxie5Damages[character.Star3] + skill1Moxie5Damages[character.Star3] + skill2Moxie5Damages[character.Star1] + skill1Moxie5Damages[character.Star2] + skill1Moxie5Damages[character.Star2] + ultimateDamages[character.Star1]

		fmt.Printf("---------\nCenturion His Bounden Duty Final Damage:")
		fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
		fmt.Printf("\nSkill 1 Moxie 1: %.2f, %.2f, %.2f", skill1Moxie1Damages[0], skill1Moxie1Damages[1], skill1Moxie1Damages[2])
		fmt.Printf("\nSkill 1 Moxie 2: %.2f, %.2f, %.2f", skill1Moxie2Damages[0], skill1Moxie2Damages[1], skill1Moxie2Damages[2])
		fmt.Printf("\nSkill 1 Moxie 3: %.2f, %.2f, %.2f", skill1Moxie3Damages[0], skill1Moxie3Damages[1], skill1Moxie3Damages[2])
		fmt.Printf("\nSkill 1 Moxie 4: %.2f, %.2f, %.2f", skill1Moxie4Damages[0], skill1Moxie4Damages[1], skill1Moxie4Damages[2])
		fmt.Printf("\nSkill 1 Moxie 5: %.2f, %.2f, %.2f", skill1Moxie5Damages[0], skill1Moxie5Damages[1], skill1Moxie5Damages[2])
		fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
		fmt.Printf("\nSkill 2 Moxie 1: %.2f, %.2f, %.2f", skill2Moxie1Damages[0], skill2Moxie1Damages[1], skill2Moxie1Damages[2])
		fmt.Printf("\nSkill 2 Moxie 2: %.2f, %.2f, %.2f", skill2Moxie2Damages[0], skill2Moxie2Damages[1], skill2Moxie2Damages[2])
		fmt.Printf("\nSkill 2 Moxie 3: %.2f, %.2f, %.2f", skill2Moxie3Damages[0], skill2Moxie3Damages[1], skill2Moxie3Damages[2])
		fmt.Printf("\nSkill 2 Moxie 4: %.2f, %.2f, %.2f", skill2Moxie4Damages[0], skill2Moxie4Damages[1], skill2Moxie4Damages[2])
		fmt.Printf("\nSkill 2 Moxie 5: %.2f, %.2f, %.2f", skill2Moxie5Damages[0], skill2Moxie5Damages[1], skill2Moxie5Damages[2])
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
		calculatorForThunder := DamageCalculator{
			Character:         character.Centurion,
			Psychube:          &psychube.ThunderousApplause,
			Resonance:         &resonances[i],
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Damages := calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill1, calParams.EnemyHit)
		skill1Moxie1Damages := calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 1, HasExtraDamage: true, ExtraDamageStack: 0, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill1, calParams.EnemyHit)
		skill1Moxie2Damages := calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 2, HasExtraDamage: true, ExtraDamageStack: 1, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill1, calParams.EnemyHit)
		skill1Moxie3Damages := calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 3, HasExtraDamage: true, ExtraDamageStack: 2, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill1, calParams.EnemyHit)
		skill1Moxie4Damages := calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 4, HasExtraDamage: true, ExtraDamageStack: 3, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill1, calParams.EnemyHit)
		skill1Moxie5Damages := calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 5, HasExtraDamage: true, ExtraDamageStack: 4, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill1, calParams.EnemyHit)

		skill2Damages := calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
		skill2Moxie1Damages := calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 1}, character.Skill2, calParams.EnemyHit)
		skill2Moxie2Damages := calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 2}, character.Skill2, calParams.EnemyHit)
		skill2Moxie3Damages := calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 3}, character.Skill2, calParams.EnemyHit)
		skill2Moxie4Damages := calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 4}, character.Skill2, calParams.EnemyHit)
		skill2Moxie5Damages := calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 5}, character.Skill2, calParams.EnemyHit)

		ultimateDamages := calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 5}, character.Ultimate, calParams.EnemyHit)

		expectTotalDamage := skill2Moxie2Damages[character.Star2] + skill1Moxie4Damages[character.Star1] + skill2Moxie5Damages[character.Star1] + skill1Moxie5Damages[character.Star3] + skill2Moxie5Damages[character.Star3] + skill1Moxie5Damages[character.Star3] + skill2Moxie5Damages[character.Star1] + skill1Moxie5Damages[character.Star2] + skill1Moxie5Damages[character.Star2] + ultimateDamages[character.Star1]

		fmt.Printf("---------\nCenturion Thunderous Applause Final Damage:")
		fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
		fmt.Printf("\nSkill 1 Moxie 1: %.2f, %.2f, %.2f", skill1Moxie1Damages[0], skill1Moxie1Damages[1], skill1Moxie1Damages[2])
		fmt.Printf("\nSkill 1 Moxie 2: %.2f, %.2f, %.2f", skill1Moxie2Damages[0], skill1Moxie2Damages[1], skill1Moxie2Damages[2])
		fmt.Printf("\nSkill 1 Moxie 3: %.2f, %.2f, %.2f", skill1Moxie3Damages[0], skill1Moxie3Damages[1], skill1Moxie3Damages[2])
		fmt.Printf("\nSkill 1 Moxie 4: %.2f, %.2f, %.2f", skill1Moxie4Damages[0], skill1Moxie4Damages[1], skill1Moxie4Damages[2])
		fmt.Printf("\nSkill 1 Moxie 5: %.2f, %.2f, %.2f", skill1Moxie5Damages[0], skill1Moxie5Damages[1], skill1Moxie5Damages[2])
		fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
		fmt.Printf("\nSkill 2 Moxie 1: %.2f, %.2f, %.2f", skill2Moxie1Damages[0], skill2Moxie1Damages[1], skill2Moxie1Damages[2])
		fmt.Printf("\nSkill 2 Moxie 2: %.2f, %.2f, %.2f", skill2Moxie2Damages[0], skill2Moxie2Damages[1], skill2Moxie2Damages[2])
		fmt.Printf("\nSkill 2 Moxie 3: %.2f, %.2f, %.2f", skill2Moxie3Damages[0], skill2Moxie3Damages[1], skill2Moxie3Damages[2])
		fmt.Printf("\nSkill 2 Moxie 4: %.2f, %.2f, %.2f", skill2Moxie4Damages[0], skill2Moxie4Damages[1], skill2Moxie4Damages[2])
		fmt.Printf("\nSkill 2 Moxie 5: %.2f, %.2f, %.2f", skill2Moxie5Damages[0], skill2Moxie5Damages[1], skill2Moxie5Damages[2])
		fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
		fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

		psychubeName := calculatorForThunder.Psychube.Name()
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
			Character:         character.Centurion,
			Psychube:          &psychube.YearningDesire,
			Resonance:         &resonances[i],
			BuffDmgBonus:      psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus(),
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
		skill1Moxie1Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 1, HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.EnemyHit)
		skill1Moxie2Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 2, HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.EnemyHit)
		skill1Moxie3Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 3, HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.EnemyHit)
		skill1Moxie4Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 4, HasExtraDamage: true, ExtraDamageStack: 3}, character.Skill1, calParams.EnemyHit)
		skill1Moxie5Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 5, HasExtraDamage: true, ExtraDamageStack: 4}, character.Skill1, calParams.EnemyHit)

		skill2Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
		skill2Moxie1Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 1}, character.Skill2, calParams.EnemyHit)
		skill2Moxie2Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 2}, character.Skill2, calParams.EnemyHit)
		skill2Moxie3Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 3}, character.Skill2, calParams.EnemyHit)
		skill2Moxie4Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 4}, character.Skill2, calParams.EnemyHit)
		skill2Moxie5Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 5}, character.Skill2, calParams.EnemyHit)

		ultimateDamages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus * 5}, character.Ultimate, calParams.EnemyHit)

		expectTotalDamage := skill2Moxie2Damages[character.Star2] + skill1Moxie4Damages[character.Star1] + skill2Moxie5Damages[character.Star1] + skill1Moxie5Damages[character.Star3] + skill2Moxie5Damages[character.Star3] + skill1Moxie5Damages[character.Star3] + skill2Moxie5Damages[character.Star1] + skill1Moxie5Damages[character.Star2] + skill1Moxie5Damages[character.Star2] + ultimateDamages[character.Star1]

		fmt.Printf("---------\nCenturion Yearning Desire Final Damage:")
		fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
		fmt.Printf("\nSkill 1 Moxie 1: %.2f, %.2f, %.2f", skill1Moxie1Damages[0], skill1Moxie1Damages[1], skill1Moxie1Damages[2])
		fmt.Printf("\nSkill 1 Moxie 2: %.2f, %.2f, %.2f", skill1Moxie2Damages[0], skill1Moxie2Damages[1], skill1Moxie2Damages[2])
		fmt.Printf("\nSkill 1 Moxie 3: %.2f, %.2f, %.2f", skill1Moxie3Damages[0], skill1Moxie3Damages[1], skill1Moxie3Damages[2])
		fmt.Printf("\nSkill 1 Moxie 4: %.2f, %.2f, %.2f", skill1Moxie4Damages[0], skill1Moxie4Damages[1], skill1Moxie4Damages[2])
		fmt.Printf("\nSkill 1 Moxie 5: %.2f, %.2f, %.2f", skill1Moxie5Damages[0], skill1Moxie5Damages[1], skill1Moxie5Damages[2])
		fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
		fmt.Printf("\nSkill 2 Moxie 1: %.2f, %.2f, %.2f", skill2Moxie1Damages[0], skill2Moxie1Damages[1], skill2Moxie1Damages[2])
		fmt.Printf("\nSkill 2 Moxie 2: %.2f, %.2f, %.2f", skill2Moxie2Damages[0], skill2Moxie2Damages[1], skill2Moxie2Damages[2])
		fmt.Printf("\nSkill 2 Moxie 3: %.2f, %.2f, %.2f", skill2Moxie3Damages[0], skill2Moxie3Damages[1], skill2Moxie3Damages[2])
		fmt.Printf("\nSkill 2 Moxie 4: %.2f, %.2f, %.2f", skill2Moxie4Damages[0], skill2Moxie4Damages[1], skill2Moxie4Damages[2])
		fmt.Printf("\nSkill 2 Moxie 5: %.2f, %.2f, %.2f", skill2Moxie5Damages[0], skill2Moxie5Damages[1], skill2Moxie5Damages[2])
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
