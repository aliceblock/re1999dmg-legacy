package damage_calculator

import (
	"fmt"

	"github.com/aliceblock/re1999dmg/damage_calculator/character"
	"github.com/aliceblock/re1999dmg/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg/damage_calculator/psychube"
)

func EagleDmgCalculate(calParams CalParams) []DamageResponse {
	damageResponse := []DamageResponse{}

	enemyCritDef := 0.1

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

	calculatorForBoundenDuty := DamageCalculator{
		Character:         character.Eagle,
		Psychube:          &psychube.HisBoundenDuty,
		Resonance:         &resonance,
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.EnemyHit)
	skill2Damages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.EnemyHit)
	ultimateDamages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0)}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage := basicCalculateExpectTotalDmg(skill1ExtraDamages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nEagle His Bounden Duty Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
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
		Character:         character.Eagle,
		Psychube:          &psychube.Hopscotch,
		Resonance:         &resonance,
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0)}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff1Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.PsychubeAmp].UltimateMight() * 1}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff2Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.PsychubeAmp].UltimateMight() * 2}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff3Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.PsychubeAmp].UltimateMight() * 3}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff4Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.PsychubeAmp].UltimateMight() * 4}, character.Ultimate, calParams.EnemyHit)
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
		Character:         character.Eagle,
		Psychube:          &psychube.ThunderousApplause,
		Resonance:         &resonance,
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForThunder.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate()+1.0) + psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1ExtraDamages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nEagle Thunderous Applause Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
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
		Character:         character.Eagle,
		Psychube:          &psychube.BraveNewWorld,
		Resonance:         &resonance,
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
	skill2Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.EnemyHit)
	skill2BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{PenetrationRate: 0.4, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0)}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = skill1ExtraDamages[character.Star1]*2 + skill2Damages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill2BuffDamages[character.Star2]*1 + ultimateDamages[character.Star1]*1 + skill1ExtraBuffDamages[character.Star3]*1 + ultimateDamages[character.Star1]*1 + skill1ExtraBuffDamages[character.Star3]*1 + skill2Damages[character.Star2]*1

	fmt.Printf("---------\nEagle Brave New World Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 1 with BNW: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2], skill1ExtraBuffDamages[0], skill1ExtraBuffDamages[1], skill1ExtraBuffDamages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nSkill 2 with BNW: %.2f, %.2f, %.2f", skill2BuffDamages[0], skill2BuffDamages[1], skill2BuffDamages[2])
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

	calculatorForLux := DamageCalculator{
		Character:         character.Eagle,
		Psychube:          &psychube.LuxuriousLeisure,
		Resonance:         &resonance,
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1Lux1Damages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Skill1, calParams.EnemyHit)
	skill1Lux2Damages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Skill1, calParams.EnemyHit)
	skill1Lux3Damages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.EnemyHit)
	skill1ExtraLux1Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Skill1, calParams.EnemyHit)
	skill1ExtraLux2Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Skill1, calParams.EnemyHit)
	skill1ExtraLux3Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.EnemyHit)
	skill2Lux1Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{PenetrationRate: 0.4, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Skill2, calParams.EnemyHit)
	skill2Lux2Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{PenetrationRate: 0.4, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Skill2, calParams.EnemyHit)
	skill2Lux3Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{PenetrationRate: 0.4, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0)}, character.Ultimate, calParams.EnemyHit)
	ultimateLux1Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Ultimate, calParams.EnemyHit)
	ultimateLux2Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Ultimate, calParams.EnemyHit)
	ultimateLux3Damages := calculatorForLux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Ultimate, calParams.EnemyHit)
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

	psychubeName = calculatorForLux.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForYearningDesire := DamageCalculator{
		Character:         character.Eagle,
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

	skill1ExtraDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill1, calParams.EnemyHit)
	skill1ExtraBuffDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{PenetrationRate: 0.4}, character.Skill2, calParams.EnemyHit)
	skill2BuffDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{PenetrationRate: 0.4, BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0)}, character.Ultimate, calParams.EnemyHit)
	ultimateBuffDamages := calculatorForYearningDesire.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 1.0, CritDmg: ExcessCritDmgBonus(calculatorForBoundenDuty.GetTotalCritRate() + 1.0), BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1ExtraBuffDamages, skill2BuffDamages, ultimateBuffDamages)

	fmt.Printf("---------\nEagle Yearning Desire Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1Damages[0], skill1Damages[1], skill1Damages[2], skill1ExtraDamages[0], skill1ExtraDamages[1], skill1ExtraDamages[2])
	fmt.Printf("\nSkill 1 with Buff: %.2f, %.2f, %.2f (with Extra %.2f, %.2f, %.2f)", skill1BuffDamages[0], skill1BuffDamages[1], skill1BuffDamages[2], skill1ExtraBuffDamages[0], skill1ExtraBuffDamages[1], skill1ExtraBuffDamages[2])
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
