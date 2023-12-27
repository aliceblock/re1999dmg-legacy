package damage_calculator

import (
	"fmt"

	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/character"
	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/psychube"
)

func EternityDmgCalculate(calParams CalParams) []DamageResponse {
	actionsCount := 12
	damageResponse := []DamageResponse{}

	insight3DmgBonus := 0.05
	enemyCritDef := 0.1

	resonances := []resonance.Resonance{
		{
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.EternityBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 6},
				{Idea: resonance.C4LIdea, Amount: 3},
				{Idea: resonance.C4SIdea, Amount: 2},
			},
		},
		{
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.EternityBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 1},
				{Idea: resonance.C4LIdea, Amount: 3},
				{Idea: resonance.C4JIdea, Amount: 2},
				{Idea: resonance.C4SIdea, Amount: 2},
				{Idea: resonance.C3Idea, Amount: 3},
				{Idea: resonance.C2Idea, Amount: 1},
				{Idea: resonance.C1Idea, Amount: 1},
			},
		},
	}

	character.Eternity.SetInsightLevel(character.Insight3L60)

	for i := 0; i < len(resonances); i++ {
		calculator := DamageCalculator{
			Character:         character.Eternity,
			Psychube:          &psychube.HisBoundenDuty,
			Resonance:         &resonances[i],
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Turn1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1}, character.Skill1, calParams.EnemyHit)
		skill1Turn2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2}, character.Skill1, calParams.EnemyHit)
		skill1Turn3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3}, character.Skill1, calParams.EnemyHit)
		skill1Turn4Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4}, character.Skill1, calParams.EnemyHit)
		skill1Turn5Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5}, character.Skill1, calParams.EnemyHit)

		skill2Turn1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1}, character.Skill2, calParams.EnemyHit)
		skill2Turn2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2}, character.Skill2, calParams.EnemyHit)
		skill2Turn3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3}, character.Skill2, calParams.EnemyHit)
		skill2Turn4Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4}, character.Skill2, calParams.EnemyHit)
		skill2Turn5Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5}, character.Skill2, calParams.EnemyHit)

		ultimateTurn1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn4Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn5Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5}, character.Ultimate, calParams.EnemyHit)

		expectTotalDamage := skill2Turn2Damages[character.Star2] + skill1Turn2Damages[character.Star2] + ultimateTurn4Damages[character.Star1] + skill1Turn5Damages[character.Star2] + skill2Turn5Damages[character.Star2] + skill1Turn5Damages[character.Star2] + ultimateTurn5Damages[character.Star1] + skill2Turn5Damages[character.Star2]
		if calParams.ShowDamagePerAction {
			expectTotalDamage = expectTotalDamage / float64(actionsCount)
		}

		fmt.Printf("---------\nEternity His Bounden Duty Final Damage:")
		fmt.Printf("\nSkill 1 Turn 1: %.2f, %.2f, %.2f", skill1Turn1Damages[0], skill1Turn1Damages[1], skill1Turn1Damages[2])
		fmt.Printf("\nSkill 1 Turn 2: %.2f, %.2f, %.2f", skill1Turn2Damages[0], skill1Turn2Damages[1], skill1Turn2Damages[2])
		fmt.Printf("\nSkill 1 Turn 3: %.2f, %.2f, %.2f", skill1Turn3Damages[0], skill1Turn3Damages[1], skill1Turn3Damages[2])
		fmt.Printf("\nSkill 1 Turn 4: %.2f, %.2f, %.2f", skill1Turn4Damages[0], skill1Turn4Damages[1], skill1Turn4Damages[2])
		fmt.Printf("\nSkill 1 Turn 5: %.2f, %.2f, %.2f", skill1Turn5Damages[0], skill1Turn5Damages[1], skill1Turn5Damages[2])
		fmt.Printf("\nSkill 2 Turn 1: %.2f, %.2f, %.2f", skill2Turn1Damages[0], skill2Turn1Damages[1], skill2Turn1Damages[2])
		fmt.Printf("\nSkill 2 Turn 2: %.2f, %.2f, %.2f", skill2Turn2Damages[0], skill2Turn2Damages[1], skill2Turn2Damages[2])
		fmt.Printf("\nSkill 2 Turn 3: %.2f, %.2f, %.2f", skill2Turn3Damages[0], skill2Turn3Damages[1], skill2Turn3Damages[2])
		fmt.Printf("\nSkill 2 Turn 4: %.2f, %.2f, %.2f", skill2Turn4Damages[0], skill2Turn4Damages[1], skill2Turn4Damages[2])
		fmt.Printf("\nSkill 2 Turn 5: %.2f, %.2f, %.2f", skill2Turn5Damages[0], skill2Turn5Damages[1], skill2Turn5Damages[2])
		fmt.Printf("\nUltimate Turn 1: %.2f", ultimateTurn1Damages[0])
		fmt.Printf("\nUltimate Turn 2: %.2f", ultimateTurn2Damages[0])
		fmt.Printf("\nUltimate Turn 3: %.2f", ultimateTurn3Damages[0])
		fmt.Printf("\nUltimate Turn 4: %.2f", ultimateTurn4Damages[0])
		fmt.Printf("\nUltimate Turn 5: %.2f", ultimateTurn5Damages[0])
		fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

		psychubeName := calculator.Psychube.Name()
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
		calculator := DamageCalculator{
			Character:         character.Eternity,
			Psychube:          &psychube.YearningDesire,
			Resonance:         &resonances[i],
			BuffDmgBonus:      psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus(),
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Turn1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1}, character.Skill1, calParams.EnemyHit)
		skill1Turn2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2}, character.Skill1, calParams.EnemyHit)
		skill1Turn3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3}, character.Skill1, calParams.EnemyHit)
		skill1Turn4Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4}, character.Skill1, calParams.EnemyHit)
		skill1Turn5Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5}, character.Skill1, calParams.EnemyHit)

		skill2Turn1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1}, character.Skill2, calParams.EnemyHit)
		skill2Turn2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2}, character.Skill2, calParams.EnemyHit)
		skill2Turn3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3}, character.Skill2, calParams.EnemyHit)
		skill2Turn4Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4}, character.Skill2, calParams.EnemyHit)
		skill2Turn5Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5}, character.Skill2, calParams.EnemyHit)

		ultimateTurn1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn4Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn5Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5}, character.Ultimate, calParams.EnemyHit)

		expectTotalDamage := skill2Turn2Damages[character.Star2] + skill1Turn2Damages[character.Star2] + ultimateTurn4Damages[character.Star1] + skill1Turn5Damages[character.Star2] + skill2Turn5Damages[character.Star2] + skill1Turn5Damages[character.Star2] + ultimateTurn5Damages[character.Star1] + skill2Turn5Damages[character.Star2]
		if calParams.ShowDamagePerAction {
			expectTotalDamage = expectTotalDamage / float64(actionsCount)
		}

		fmt.Printf("---------\nEternity Yearning Desire Final Damage:")
		fmt.Printf("\nSkill 1 Turn 1: %.2f, %.2f, %.2f", skill1Turn1Damages[0], skill1Turn1Damages[1], skill1Turn1Damages[2])
		fmt.Printf("\nSkill 1 Turn 2: %.2f, %.2f, %.2f", skill1Turn2Damages[0], skill1Turn2Damages[1], skill1Turn2Damages[2])
		fmt.Printf("\nSkill 1 Turn 3: %.2f, %.2f, %.2f", skill1Turn3Damages[0], skill1Turn3Damages[1], skill1Turn3Damages[2])
		fmt.Printf("\nSkill 1 Turn 4: %.2f, %.2f, %.2f", skill1Turn4Damages[0], skill1Turn4Damages[1], skill1Turn4Damages[2])
		fmt.Printf("\nSkill 1 Turn 5: %.2f, %.2f, %.2f", skill1Turn5Damages[0], skill1Turn5Damages[1], skill1Turn5Damages[2])
		fmt.Printf("\nSkill 2 Turn 1: %.2f, %.2f, %.2f", skill2Turn1Damages[0], skill2Turn1Damages[1], skill2Turn1Damages[2])
		fmt.Printf("\nSkill 2 Turn 2: %.2f, %.2f, %.2f", skill2Turn2Damages[0], skill2Turn2Damages[1], skill2Turn2Damages[2])
		fmt.Printf("\nSkill 2 Turn 3: %.2f, %.2f, %.2f", skill2Turn3Damages[0], skill2Turn3Damages[1], skill2Turn3Damages[2])
		fmt.Printf("\nSkill 2 Turn 4: %.2f, %.2f, %.2f", skill2Turn4Damages[0], skill2Turn4Damages[1], skill2Turn4Damages[2])
		fmt.Printf("\nSkill 2 Turn 5: %.2f, %.2f, %.2f", skill2Turn5Damages[0], skill2Turn5Damages[1], skill2Turn5Damages[2])
		fmt.Printf("\nUltimate Turn 1: %.2f", ultimateTurn1Damages[0])
		fmt.Printf("\nUltimate Turn 2: %.2f", ultimateTurn2Damages[0])
		fmt.Printf("\nUltimate Turn 3: %.2f", ultimateTurn3Damages[0])
		fmt.Printf("\nUltimate Turn 4: %.2f", ultimateTurn4Damages[0])
		fmt.Printf("\nUltimate Turn 5: %.2f", ultimateTurn5Damages[0])
		fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

		psychubeName := calculator.Psychube.Name()
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
		calculator := DamageCalculator{
			Character:         character.Eternity,
			Psychube:          &psychube.Hopscotch,
			Resonance:         &resonances[i],
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Turn1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1}, character.Skill1, calParams.EnemyHit)
		skill1Turn2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2}, character.Skill1, calParams.EnemyHit)
		skill1Turn3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3}, character.Skill1, calParams.EnemyHit)
		skill1Turn4Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4}, character.Skill1, calParams.EnemyHit)
		skill1Turn5Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5}, character.Skill1, calParams.EnemyHit)

		skill2Turn1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1}, character.Skill2, calParams.EnemyHit)
		skill2Turn2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2}, character.Skill2, calParams.EnemyHit)
		skill2Turn3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3}, character.Skill2, calParams.EnemyHit)
		skill2Turn4Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4}, character.Skill2, calParams.EnemyHit)
		skill2Turn5Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5}, character.Skill2, calParams.EnemyHit)

		ultimateTurn1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn4Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn5Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5}, character.Ultimate, calParams.EnemyHit)

		expectTotalDamage := skill2Turn2Damages[character.Star2] + skill1Turn2Damages[character.Star2] + ultimateTurn4Damages[character.Star1] + skill1Turn5Damages[character.Star2] + skill2Turn5Damages[character.Star2] + skill1Turn5Damages[character.Star2] + ultimateTurn5Damages[character.Star1] + skill2Turn5Damages[character.Star2]
		if calParams.ShowDamagePerAction {
			expectTotalDamage = expectTotalDamage / float64(actionsCount)
		}

		fmt.Printf("---------\nEternity Hopscotch Final Damage:")
		fmt.Printf("\nSkill 1 Turn 1: %.2f, %.2f, %.2f", skill1Turn1Damages[0], skill1Turn1Damages[1], skill1Turn1Damages[2])
		fmt.Printf("\nSkill 1 Turn 2: %.2f, %.2f, %.2f", skill1Turn2Damages[0], skill1Turn2Damages[1], skill1Turn2Damages[2])
		fmt.Printf("\nSkill 1 Turn 3: %.2f, %.2f, %.2f", skill1Turn3Damages[0], skill1Turn3Damages[1], skill1Turn3Damages[2])
		fmt.Printf("\nSkill 1 Turn 4: %.2f, %.2f, %.2f", skill1Turn4Damages[0], skill1Turn4Damages[1], skill1Turn4Damages[2])
		fmt.Printf("\nSkill 1 Turn 5: %.2f, %.2f, %.2f", skill1Turn5Damages[0], skill1Turn5Damages[1], skill1Turn5Damages[2])
		fmt.Printf("\nSkill 2 Turn 1: %.2f, %.2f, %.2f", skill2Turn1Damages[0], skill2Turn1Damages[1], skill2Turn1Damages[2])
		fmt.Printf("\nSkill 2 Turn 2: %.2f, %.2f, %.2f", skill2Turn2Damages[0], skill2Turn2Damages[1], skill2Turn2Damages[2])
		fmt.Printf("\nSkill 2 Turn 3: %.2f, %.2f, %.2f", skill2Turn3Damages[0], skill2Turn3Damages[1], skill2Turn3Damages[2])
		fmt.Printf("\nSkill 2 Turn 4: %.2f, %.2f, %.2f", skill2Turn4Damages[0], skill2Turn4Damages[1], skill2Turn4Damages[2])
		fmt.Printf("\nSkill 2 Turn 5: %.2f, %.2f, %.2f", skill2Turn5Damages[0], skill2Turn5Damages[1], skill2Turn5Damages[2])
		fmt.Printf("\nUltimate Turn 1: %.2f", ultimateTurn1Damages[0])
		fmt.Printf("\nUltimate Turn 2: %.2f", ultimateTurn2Damages[0])
		fmt.Printf("\nUltimate Turn 3: %.2f", ultimateTurn3Damages[0])
		fmt.Printf("\nUltimate Turn 4: %.2f", ultimateTurn4Damages[0])
		fmt.Printf("\nUltimate Turn 5: %.2f", ultimateTurn5Damages[0])
		fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

		psychubeName := calculator.Psychube.Name()
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
		calculator := DamageCalculator{
			Character:         character.Eternity,
			Psychube:          &psychube.BraveNewWorld,
			Resonance:         &resonances[i],
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Turn1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1}, character.Skill1, calParams.EnemyHit)
		skill1Turn2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2}, character.Skill1, calParams.EnemyHit)
		skill1Turn3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3}, character.Skill1, calParams.EnemyHit)
		skill1Turn4Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4}, character.Skill1, calParams.EnemyHit)
		skill1Turn5Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5}, character.Skill1, calParams.EnemyHit)
		skill1Turn1BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
		skill1Turn2BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
		skill1Turn3BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
		skill1Turn4BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
		skill1Turn5BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)

		skill2Turn1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1}, character.Skill2, calParams.EnemyHit)
		skill2Turn2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2}, character.Skill2, calParams.EnemyHit)
		skill2Turn3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3}, character.Skill2, calParams.EnemyHit)
		skill2Turn4Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4}, character.Skill2, calParams.EnemyHit)
		skill2Turn5Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5}, character.Skill2, calParams.EnemyHit)
		skill2Turn1BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill2, calParams.EnemyHit)
		skill2Turn2BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill2, calParams.EnemyHit)
		skill2Turn3BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill2, calParams.EnemyHit)
		skill2Turn4BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill2, calParams.EnemyHit)
		skill2Turn5BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill2, calParams.EnemyHit)

		ultimateTurn1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 1}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 2}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 3}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn4Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 4}, character.Ultimate, calParams.EnemyHit)
		ultimateTurn5Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight3DmgBonus * 5}, character.Ultimate, calParams.EnemyHit)

		expectTotalDamage := skill2Turn2Damages[character.Star2] + skill1Turn2Damages[character.Star2] + ultimateTurn4Damages[character.Star1] + skill1Turn5BuffDamages[character.Star2] + skill2Turn5Damages[character.Star2] + skill1Turn5Damages[character.Star2] + ultimateTurn5Damages[character.Star1] + skill2Turn5BuffDamages[character.Star2]
		if calParams.ShowDamagePerAction {
			expectTotalDamage = expectTotalDamage / float64(actionsCount)
		}

		fmt.Printf("---------\nEternity Brave New World Final Damage:")
		fmt.Printf("\nSkill 1 Turn 1: %.2f, %.2f, %.2f (with Buff: %.2f, %.2f, %.2f)", skill1Turn1Damages[0], skill1Turn1Damages[1], skill1Turn1Damages[2], skill1Turn1BuffDamages[0], skill1Turn1BuffDamages[1], skill1Turn1BuffDamages[2])
		fmt.Printf("\nSkill 1 Turn 2: %.2f, %.2f, %.2f (with Buff: %.2f, %.2f, %.2f)", skill1Turn2Damages[0], skill1Turn2Damages[1], skill1Turn2Damages[2], skill1Turn2BuffDamages[0], skill1Turn2BuffDamages[1], skill1Turn2BuffDamages[2])
		fmt.Printf("\nSkill 1 Turn 3: %.2f, %.2f, %.2f (with Buff: %.2f, %.2f, %.2f)", skill1Turn3Damages[0], skill1Turn3Damages[1], skill1Turn3Damages[2], skill1Turn3BuffDamages[0], skill1Turn3BuffDamages[1], skill1Turn3BuffDamages[2])
		fmt.Printf("\nSkill 1 Turn 4: %.2f, %.2f, %.2f (with Buff: %.2f, %.2f, %.2f)", skill1Turn4Damages[0], skill1Turn4Damages[1], skill1Turn4Damages[2], skill1Turn4BuffDamages[0], skill1Turn4BuffDamages[1], skill1Turn4BuffDamages[2])
		fmt.Printf("\nSkill 1 Turn 5: %.2f, %.2f, %.2f (with Buff: %.2f, %.2f, %.2f)", skill1Turn5Damages[0], skill1Turn5Damages[1], skill1Turn5Damages[2], skill1Turn5BuffDamages[0], skill1Turn5BuffDamages[1], skill1Turn5BuffDamages[2])
		fmt.Printf("\nSkill 2 Turn 1: %.2f, %.2f, %.2f (with Buff: %.2f, %.2f, %.2f)", skill2Turn1Damages[0], skill2Turn1Damages[1], skill2Turn1Damages[2], skill2Turn1BuffDamages[0], skill2Turn1BuffDamages[1], skill2Turn1BuffDamages[2])
		fmt.Printf("\nSkill 2 Turn 2: %.2f, %.2f, %.2f (with Buff: %.2f, %.2f, %.2f)", skill2Turn2Damages[0], skill2Turn2Damages[1], skill2Turn2Damages[2], skill2Turn2BuffDamages[0], skill2Turn2BuffDamages[1], skill2Turn2BuffDamages[2])
		fmt.Printf("\nSkill 2 Turn 3: %.2f, %.2f, %.2f (with Buff: %.2f, %.2f, %.2f)", skill2Turn3Damages[0], skill2Turn3Damages[1], skill2Turn3Damages[2], skill2Turn3BuffDamages[0], skill2Turn3BuffDamages[1], skill2Turn3BuffDamages[2])
		fmt.Printf("\nSkill 2 Turn 4: %.2f, %.2f, %.2f (with Buff: %.2f, %.2f, %.2f)", skill2Turn4Damages[0], skill2Turn4Damages[1], skill2Turn4Damages[2], skill2Turn4BuffDamages[0], skill2Turn4BuffDamages[1], skill2Turn4BuffDamages[2])
		fmt.Printf("\nSkill 2 Turn 5: %.2f, %.2f, %.2f (with Buff: %.2f, %.2f, %.2f)", skill2Turn5Damages[0], skill2Turn5Damages[1], skill2Turn5Damages[2], skill2Turn5BuffDamages[0], skill2Turn5BuffDamages[1], skill2Turn5BuffDamages[2])
		fmt.Printf("\nUltimate Turn 1: %.2f", ultimateTurn1Damages[0])
		fmt.Printf("\nUltimate Turn 2: %.2f", ultimateTurn2Damages[0])
		fmt.Printf("\nUltimate Turn 3: %.2f", ultimateTurn3Damages[0])
		fmt.Printf("\nUltimate Turn 4: %.2f", ultimateTurn4Damages[0])
		fmt.Printf("\nUltimate Turn 5: %.2f", ultimateTurn5Damages[0])
		fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

		psychubeName := calculator.Psychube.Name()
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
