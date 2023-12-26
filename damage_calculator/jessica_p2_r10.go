package damage_calculator

import (
	"fmt"

	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/character"
	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/psychube"
)

func JessicaP2R10DmgCalculate(calParams CalParams) []DamageResponse {
	damageResponse := []DamageResponse{}

	insight1DmgBonus := 0.2 // Damage to poisoned enemy
	enemyCritDef := 0.1

	resonances := []resonance.Resonance{
		// {
		// 	Ideas: []resonance.IdeaAmount{
		// 		{Idea: resonance.JessicaBaseIdea, Amount: 1},
		// 		{Idea: resonance.C4IOTIdea, Amount: 6},
		// 		{Idea: resonance.C4SIdea, Amount: 2},
		// 		{Idea: resonance.C3Idea, Amount: 3},
		// 		{Idea: resonance.C1Idea, Amount: 2},
		// 	},
		// },
		// {
		// 	Ideas: []resonance.IdeaAmount{
		// 		{Idea: resonance.JessicaBaseIdea, Amount: 1},
		// 		{Idea: resonance.C4IOTIdea, Amount: 2},
		// 		{Idea: resonance.C4LIdea, Amount: 2},
		// 		{Idea: resonance.C4SIdea, Amount: 3},
		// 		{Idea: resonance.C4JIdea, Amount: 1},
		// 		{Idea: resonance.C3Idea, Amount: 3},
		// 		{Idea: resonance.C2Idea, Amount: 1},
		// 		{Idea: resonance.C1Idea, Amount: 1},
		// 	},
		// },
		{
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.JessicaBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 6},
				{Idea: resonance.C4LIdea, Amount: 2},
				{Idea: resonance.C4SIdea, Amount: 3},
			},
		},
	}
	//        Turn: 1 2  3   4   5  6 7  8    9  10 11 12
	// P0-1 Poison: 2 1  2   1   0  2 1  0    2   1 0  2
	// P2   Poison: 3 2 1|3  2   1  3 2  1    3   2 1  3
	// Ulti Poison: 0 0  0   2   1  0 0  2    1   0 0  2
	//     Total 1: 2 1  2  1|2  1  2 1 2|2 1|1|2 1 0 2|2 = 14
	//     Total 2: 3 2 1|3 2|2 1|1 3 2 1|2  1|3  2 1 2|3 = 18
	//     Skill 1: 0 0  2   0   0  2 0  0    2   0 0  2
	//     Skill 2: 2 0  0   2   0  0 2  0    0   2 0  0
	character.JessicaP2.SetInsightLevel(character.Insight3L60)

	for i := 0; i < len(resonances); i++ {
		calculator := DamageCalculator{
			Character:         character.JessicaP2,
			Psychube:          &psychube.HisBoundenDuty,
			Resonance:         &resonances[i],
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
		skill1Extra1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.EnemyHit)
		skill1Extra2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.EnemyHit)
		skill1Extra3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.EnemyHit)

		skill2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
		skill2ExtraDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)

		ultimateDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
		ultimateI1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus}, character.Ultimate, calParams.EnemyHit)

		poisonDamage := calculator.CalculateGenesisDamage(DamageCalculatorInfo{}, 0.3) * 15
		expectTotalDamage := skill2ExtraDamages[character.Star2] + skill1Extra2Damages[character.Star2] + ultimateI1Damages[character.Star1] + skill2ExtraDamages[character.Star2] + skill1Extra1Damages[character.Star2] + skill2ExtraDamages[character.Star2] + ultimateI1Damages[character.Star1] + skill1Extra2Damages[character.Star2] + skill2ExtraDamages[character.Star2] + poisonDamage

		fmt.Printf("---------\nJessica His Bounden Duty Final Damage:")
		fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
		fmt.Printf("\nSkill 1 Extra 1: %.2f, %.2f, %.2f", skill1Extra1Damages[0], skill1Extra1Damages[1], skill1Extra1Damages[2])
		fmt.Printf("\nSkill 1 Extra 2: %.2f, %.2f, %.2f", skill1Extra2Damages[0], skill1Extra2Damages[1], skill1Extra2Damages[2])
		fmt.Printf("\nSkill 1 Extra 3: %.2f, %.2f, %.2f", skill1Extra3Damages[0], skill1Extra3Damages[1], skill1Extra3Damages[2])
		fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
		fmt.Printf("\nSkill 2 Extra: %.2f, %.2f, %.2f", skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2])
		fmt.Printf("\nUltimate: %.2f (I1 Buff: %.2f)", ultimateDamages[0], ultimateI1Damages[0])
		fmt.Printf("\nPoison stack/round: %.2f", poisonDamage/12.0)
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
			Character:         character.JessicaP2,
			Psychube:          &psychube.Hopscotch,
			Resonance:         &resonances[i],
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
		skill1Extra1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.EnemyHit)
		skill1Extra2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.EnemyHit)
		skill1Extra3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.EnemyHit)

		skill2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
		skill2ExtraDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)

		ultimateDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
		ultimateI1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus}, character.Ultimate, calParams.EnemyHit)

		poisonDamage := calculator.CalculateGenesisDamage(DamageCalculatorInfo{}, 0.3) * 15
		expectTotalDamage := skill2ExtraDamages[character.Star2] + skill1Extra2Damages[character.Star2] + ultimateI1Damages[character.Star1] + skill2ExtraDamages[character.Star2] + skill1Extra1Damages[character.Star2] + skill2ExtraDamages[character.Star2] + ultimateI1Damages[character.Star1] + skill1Extra2Damages[character.Star2] + skill2ExtraDamages[character.Star2] + poisonDamage

		fmt.Printf("---------\nJessica Hopscotch Final Damage:")
		fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
		fmt.Printf("\nSkill 1 Extra 1: %.2f, %.2f, %.2f", skill1Extra1Damages[0], skill1Extra1Damages[1], skill1Extra1Damages[2])
		fmt.Printf("\nSkill 1 Extra 2: %.2f, %.2f, %.2f", skill1Extra2Damages[0], skill1Extra2Damages[1], skill1Extra2Damages[2])
		fmt.Printf("\nSkill 1 Extra 3: %.2f, %.2f, %.2f", skill1Extra3Damages[0], skill1Extra3Damages[1], skill1Extra3Damages[2])
		fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
		fmt.Printf("\nSkill 2 Extra: %.2f, %.2f, %.2f", skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2])
		fmt.Printf("\nUltimate: %.2f (I1 Buff: %.2f)", ultimateDamages[0], ultimateI1Damages[0])
		fmt.Printf("\nPoison stack/round: %.2f", poisonDamage/12.0)
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
			Character:         character.JessicaP2,
			Psychube:          &psychube.YearningDesire,
			Resonance:         &resonances[i],
			BuffDmgBonus:      psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus(),
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
		skill1Extra1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.EnemyHit)
		skill1Extra2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.EnemyHit)
		skill1Extra3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.EnemyHit)

		skill2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
		skill2ExtraDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)

		ultimateDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
		ultimateI1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus}, character.Ultimate, calParams.EnemyHit)

		poisonDamage := calculator.CalculateGenesisDamage(DamageCalculatorInfo{}, 0.3) * 15
		expectTotalDamage := skill2ExtraDamages[character.Star2] + skill1Extra2Damages[character.Star2] + ultimateI1Damages[character.Star1] + skill2ExtraDamages[character.Star2] + skill1Extra1Damages[character.Star2] + skill2ExtraDamages[character.Star2] + ultimateI1Damages[character.Star1] + skill1Extra2Damages[character.Star2] + skill2ExtraDamages[character.Star2] + poisonDamage

		fmt.Printf("---------\nJessica Yearning Desire Final Damage:")
		fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
		fmt.Printf("\nSkill 1 Extra 1: %.2f, %.2f, %.2f", skill1Extra1Damages[0], skill1Extra1Damages[1], skill1Extra1Damages[2])
		fmt.Printf("\nSkill 1 Extra 2: %.2f, %.2f, %.2f", skill1Extra2Damages[0], skill1Extra2Damages[1], skill1Extra2Damages[2])
		fmt.Printf("\nSkill 1 Extra 3: %.2f, %.2f, %.2f", skill1Extra3Damages[0], skill1Extra3Damages[1], skill1Extra3Damages[2])
		fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
		fmt.Printf("\nSkill 2 Extra: %.2f, %.2f, %.2f", skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2])
		fmt.Printf("\nUltimate: %.2f (I1 Buff: %.2f)", ultimateDamages[0], ultimateI1Damages[0])
		fmt.Printf("\nPoison stack/round: %.2f", poisonDamage/12.0)
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
			Character:         character.JessicaP2,
			Psychube:          &psychube.BraveNewWorld,
			Resonance:         &resonances[i],
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
		skill1Extra1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.EnemyHit)
		skill1Extra2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.EnemyHit)
		skill1Extra3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.EnemyHit)
		skill1BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
		skill1Extra1BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 0, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
		skill1Extra2BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 1, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
		skill1Extra3BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 2, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)

		skill2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
		skill2ExtraDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)
		skill2BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill2, calParams.EnemyHit)
		skill2ExtraBuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill2, calParams.EnemyHit)

		ultimateDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
		ultimateI1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus}, character.Ultimate, calParams.EnemyHit)

		poisonDamage := calculator.CalculateGenesisDamage(DamageCalculatorInfo{}, 0.3) * 15
		expectTotalDamage := skill2ExtraDamages[character.Star2] + skill1Extra2Damages[character.Star2] + ultimateI1Damages[character.Star1] + skill2ExtraBuffDamages[character.Star2] + skill1Extra1Damages[character.Star2] + skill2ExtraDamages[character.Star2] + ultimateI1Damages[character.Star1] + skill1Extra2BuffDamages[character.Star2] + skill2ExtraDamages[character.Star2] + poisonDamage

		fmt.Printf("---------\nJessica Brave New World Final Damage:")
		fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (With BNW Buff: %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2])
		fmt.Printf("\nSkill 1 Extra 1: %.2f, %.2f, %.2f (With BNW Buff: %.2f, %.2f, %.2f)", skill1Extra1Damages[0], skill1Extra1Damages[1], skill1Extra1Damages[2], skill1Extra1BuffDamages[0], skill1Extra1BuffDamages[1], skill1Extra1BuffDamages[2])
		fmt.Printf("\nSkill 1 Extra 2: %.2f, %.2f, %.2f (With BNW Buff: %.2f, %.2f, %.2f)", skill1Extra2Damages[0], skill1Extra2Damages[1], skill1Extra2Damages[2], skill1Extra2BuffDamages[0], skill1Extra2BuffDamages[1], skill1Extra2BuffDamages[2])
		fmt.Printf("\nSkill 1 Extra 3: %.2f, %.2f, %.2f (With BNW Buff: %.2f, %.2f, %.2f)", skill1Extra3Damages[0], skill1Extra3Damages[1], skill1Extra3Damages[2], skill1Extra3BuffDamages[0], skill1Extra3BuffDamages[1], skill1Extra3BuffDamages[2])
		fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (With BNW Buff: %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2])
		fmt.Printf("\nSkill 2 Extra: %.2f, %.2f, %.2f (With BNW Buff: %.2f, %.2f, %.2f)", skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2], skill2ExtraBuffDamages[0], skill2ExtraBuffDamages[1], skill2ExtraBuffDamages[2])
		fmt.Printf("\nUltimate: %.2f (I1 Buff: %.2f)", ultimateDamages[0], ultimateI1Damages[0])
		fmt.Printf("\nPoison stack/round: %.2f", poisonDamage/12.0)
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
			Character:         character.JessicaP2,
			Psychube:          &psychube.BlasphemerOfNight,
			Resonance:         &resonances[i],
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
		skill1Extra1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.EnemyHit)
		skill1Extra2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.EnemyHit)
		skill1Extra3Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.EnemyHit)
		skill1Extra1BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus + psychube.BlasphemerOfNight.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus(), HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.EnemyHit)
		skill1Extra2BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus + psychube.BlasphemerOfNight.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus(), HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.EnemyHit)
		skill1Extra3BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus + psychube.BlasphemerOfNight.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus(), HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.EnemyHit)

		skill2Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
		skill2ExtraDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus, HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)
		skill2ExtraBuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus + psychube.BlasphemerOfNight.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus(), HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)

		ultimateDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
		ultimateI1Damages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus}, character.Ultimate, calParams.EnemyHit)
		ultimateI1BuffDamages := calculator.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: insight1DmgBonus + psychube.BlasphemerOfNight.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Ultimate, calParams.EnemyHit)

		poisonDamage := calculator.CalculateGenesisDamage(DamageCalculatorInfo{}, 0.3) * 15
		expectTotalDamage := skill2ExtraDamages[character.Star2] + skill1Extra2Damages[character.Star2] + ultimateI1Damages[character.Star1] + skill2ExtraBuffDamages[character.Star2] + skill1Extra1Damages[character.Star2] + skill2ExtraDamages[character.Star2] + ultimateI1Damages[character.Star1] + skill1Extra2BuffDamages[character.Star2] + skill2ExtraDamages[character.Star2] + poisonDamage

		fmt.Printf("---------\nJessica Blasphemer of Night Final Damage:")
		fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
		fmt.Printf("\nSkill 1 Extra 1: %.2f, %.2f, %.2f (With Buff: %.2f, %.2f, %.2f)", skill1Extra1Damages[0], skill1Extra1Damages[1], skill1Extra1Damages[2], skill1Extra1BuffDamages[0], skill1Extra1BuffDamages[1], skill1Extra1BuffDamages[2])
		fmt.Printf("\nSkill 1 Extra 2: %.2f, %.2f, %.2f (With Buff: %.2f, %.2f, %.2f)", skill1Extra2Damages[0], skill1Extra2Damages[1], skill1Extra2Damages[2], skill1Extra2BuffDamages[0], skill1Extra2BuffDamages[1], skill1Extra2BuffDamages[2])
		fmt.Printf("\nSkill 1 Extra 3: %.2f, %.2f, %.2f (With Buff: %.2f, %.2f, %.2f)", skill1Extra3Damages[0], skill1Extra3Damages[1], skill1Extra3Damages[2], skill1Extra3BuffDamages[0], skill1Extra3BuffDamages[1], skill1Extra3BuffDamages[2])
		fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
		fmt.Printf("\nSkill 2 Extra: %.2f, %.2f, %.2f (With Buff: %.2f, %.2f, %.2f)", skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2], skill2ExtraBuffDamages[0], skill2ExtraBuffDamages[1], skill2ExtraBuffDamages[2])
		fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
		fmt.Printf("\nUltimate I1: %.2f (With Buff: %.2f)", ultimateI1Damages[0], ultimateI1BuffDamages[0])
		fmt.Printf("\nPoison stack/round: %.2f", poisonDamage/12.0)
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
