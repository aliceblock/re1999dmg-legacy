package damage_calculator

import (
	"fmt"

	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/character"
	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg-legacy/damage_calculator/psychube"
)

func CharlieDmgCalculate(calParams CalParams) []DamageResponse {
	actionsCount := 19
	damageResponse := []DamageResponse{}

	enemyCritDef := 0.1

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
		{
			Ideas: []resonance.IdeaAmount{
				{Idea: resonance.CharlieBaseIdea, Amount: 1},
				{Idea: resonance.C4IOTIdea, Amount: 2},
				{Idea: resonance.C4LIdea, Amount: 2},
				{Idea: resonance.C4SIdea, Amount: 3},
				{Idea: resonance.C4JIdea, Amount: 1},
				{Idea: resonance.C3Idea, Amount: 3},
				{Idea: resonance.C2Idea, Amount: 1},
				{Idea: resonance.C1Idea, Amount: 1},
			},
		},
	}

	character.Charlie.SetInsightLevel(character.Insight3L60)

	calculatorForBoundenDuty := DamageCalculator{
		Character:         character.Charlie,
		Psychube:          &psychube.HisBoundenDuty,
		Resonance:         &resonances[calParams.ResonanceIndex],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.EnemyHit)
	skill2Damages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2ExtraDamages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)
	ultimateDamages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage := basicCalculateExpectTotalDmg(skill1ExtraDamages, skill2ExtraDamages, ultimateDamages) - skill2ExtraDamages[character.Star2] + skill2Damages[character.Star2]
	if calParams.ShowDamagePerAction {
		expectTotalDamage = expectTotalDamage / float64(actionsCount)
	}

	fmt.Printf("---------\nCharlie His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName := calculatorForBoundenDuty.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForHopscotch := DamageCalculator{
		Character:         character.Charlie,
		Psychube:          &psychube.Hopscotch,
		Resonance:         &resonances[calParams.ResonanceIndex],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2ExtraDamages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff1Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.PsychubeAmp].UltimateMight() * 1}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff2Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.PsychubeAmp].UltimateMight() * 2}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff3Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.PsychubeAmp].UltimateMight() * 3}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff4Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.PsychubeAmp].UltimateMight() * 4}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1ExtraDamages, skill2ExtraDamages, ultimateDamages) - skill2ExtraDamages[character.Star2] + skill2Damages[character.Star2]
	if calParams.ShowDamagePerAction {
		expectTotalDamage = expectTotalDamage / float64(actionsCount)
	}
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
	expectTotalBuffDamage := skill1ExtraDamages[character.Star1]*2 + skill1ExtraDamages[character.Star3]*2 + skill2Damages[character.Star2]*1 + skill2ExtraDamages[character.Star2]*2 + ultimateDamages[character.Star1]*1 + ultimateBuff1Damages[character.Star1]*1 + ultimateBuff2Damages[character.Star1]*1
	if calParams.ShowDamagePerAction {
		expectTotalBuffDamage = expectTotalBuffDamage / float64(actionsCount)
	}

	fmt.Printf("---------\nCharlie Hopscotch Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2])
	fmt.Printf("\nUltimate: %.2f with Hopscotch buff (%.2f, %.2f, %.2f, %.2f)", ultimateDamages[0], ultimateBuff1Damages[0], ultimateBuff2Damages[0], ultimateBuff3Damages[0], ultimateBuff4Damages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)
	fmt.Printf("\nExpect with buff total damage: %.2f", expectTotalBuffDamage)

	psychubeName = calculatorForHopscotch.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	psychubeName = calculatorForHopscotch.Psychube.Name() + " Buff"
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalBuffDamage, 2),
	})

	fmt.Println()

	calculatorForThunder := DamageCalculator{
		Character:         character.Charlie,
		Psychube:          &psychube.ThunderousApplause,
		Resonance:         &resonances[calParams.ResonanceIndex],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill2, calParams.EnemyHit)
	skill2ExtraDamages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1ExtraDamages, skill2ExtraDamages, ultimateDamages) - skill2ExtraDamages[character.Star2] + skill2Damages[character.Star2]
	if calParams.ShowDamagePerAction {
		expectTotalDamage = expectTotalDamage / float64(actionsCount)
	}

	fmt.Printf("\nCrit Rate: %.2f\n", calculatorForThunder.GetTotalCritRate()*100)
	fmt.Printf("---------\nCharlie Thunderous Applause Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
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

	calculatorForBraveNewWorld := DamageCalculator{
		Character:         character.Charlie,
		Psychube:          &psychube.BraveNewWorld,
		Resonance:         &resonances[calParams.ResonanceIndex],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.EnemyHit)
	skill1ExtraBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill2, calParams.EnemyHit)
	skill2ExtraDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)
	skill2ExtraBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = skill1ExtraDamages[character.Star1]*2 + skill2Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill2ExtraBuffDamages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill1ExtraBuffDamages[character.Star3]*1 + ultimateDamages[character.Star1]*1 + skill1ExtraBuffDamages[character.Star3]*1 + skill2ExtraDamages[character.Star2]*1
	if calParams.ShowDamagePerAction {
		expectTotalDamage = expectTotalDamage / float64(actionsCount)
	}

	fmt.Printf("---------\nCharlie Brave New World Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 1 with BNW: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2], skill1ExtraBuffDamages[0], skill1ExtraBuffDamages[1], skill1ExtraBuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2])
	fmt.Printf("\nSkill 2 with BNW: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2], skill2ExtraBuffDamages[0], skill2ExtraBuffDamages[1], skill2ExtraBuffDamages[2])
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

	calculatorForYearningDesire := DamageCalculator{
		Character:         character.Charlie,
		Psychube:          &psychube.YearningDesire,
		Resonance:         &resonances[calParams.ResonanceIndex],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1BuffDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.EnemyHit)
	skill1ExtraBuffDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2BuffDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill2, calParams.EnemyHit)
	skill2ExtraDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)
	skill2ExtraBuffDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	ultimateBuffDamages := calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1ExtraBuffDamages, skill2ExtraBuffDamages, ultimateBuffDamages) - skill2ExtraBuffDamages[character.Star2] + skill2BuffDamages[character.Star2]
	if calParams.ShowDamagePerAction {
		expectTotalDamage = expectTotalDamage / float64(actionsCount)
	}

	fmt.Printf("---------\nCharlie Yearning Desire Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 1 with buff: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2], skill1ExtraBuffDamages[0], skill1ExtraBuffDamages[1], skill1ExtraBuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill2Damages[0], skill2Damages[1], skill2Damages[2], skill2ExtraDamages[0], skill2ExtraDamages[1], skill2ExtraDamages[2])
	fmt.Printf("\nSkill 2 with buff: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2], skill2ExtraBuffDamages[0], skill2ExtraBuffDamages[1], skill2ExtraBuffDamages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nUltimate with Buff: %.2f", ultimateBuffDamages[0])
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
