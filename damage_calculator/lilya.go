package damage_calculator

import (
	"fmt"

	"github.com/aliceblock/re1999dmg/damage_calculator/character"
	"github.com/aliceblock/re1999dmg/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg/damage_calculator/psychube"
)

func LilyaDmgCalculate(calParams CalParams) []DamageResponse {
	damageResponse := []DamageResponse{}

	enemyCritDef := 0.1

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

	calculatorForReso1 := DamageCalculator{
		Character:         character.Lilya,
		Psychube:          &psychube.ThunderousApplause,
		Resonance:         &resonances[0],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		CritRate:          psychube.ThunderousApplause.CritRate(), // Insight 3
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages := calculatorForReso1.CalculateFinalDamage(DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages := calculatorForReso1.CalculateFinalDamage(DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.ExtraAction1, calParams.EnemyHit)
	totalCritRate := calculatorForReso1.GetTotalCritRate()
	fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
	for index, damage := range skill1ExtraDamages {
		if totalCritRate >= 1.0 {
			skill1Damages[index] += damage
		} else {
			skill1Damages[index] += damage * totalCritRate
		}
	}
	skill2Damages := calculatorForReso1.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	ultimateDamages := calculatorForReso1.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage := basicCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nLilya Thunderous Applause resonance 1 Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName := calculatorForReso1.Psychube.Name() + " Reso1"
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForReso2 := DamageCalculator{
		Character:         character.Lilya,
		Psychube:          &psychube.ThunderousApplause,
		Resonance:         &resonances[1],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		CritRate:          psychube.ThunderousApplause.CritRate(), // Insight 3
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForReso2.CalculateFinalDamage(DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForReso2.CalculateFinalDamage(DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.ExtraAction1, calParams.EnemyHit)
	totalCritRate = calculatorForReso2.GetTotalCritRate()
	fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
	for index, damage := range skill1ExtraDamages {
		if totalCritRate >= 1.0 {
			skill1Damages[index] += damage
		} else {
			skill1Damages[index] += damage * totalCritRate
		}
	}
	skill2Damages = calculatorForReso2.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForReso2.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nLilya Thunderous Applause resonance 2 Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName = calculatorForReso2.Psychube.Name() + " Reso2"
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForReso3 := DamageCalculator{
		Character:         character.Lilya,
		Psychube:          &psychube.ThunderousApplause,
		Resonance:         &resonances[2],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		CritRate:          psychube.ThunderousApplause.CritRate(), // Insight 3
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForReso3.CalculateFinalDamage(DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForReso3.CalculateFinalDamage(DamageCalculatorInfo{CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.ExtraAction1, calParams.EnemyHit)
	totalCritRate = calculatorForReso3.GetTotalCritRate()
	fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
	for index, damage := range skill1ExtraDamages {
		if totalCritRate >= 1.0 {
			skill1Damages[index] += damage
		} else {
			skill1Damages[index] += damage * totalCritRate
		}
	}
	skill2Damages = calculatorForReso3.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForReso3.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2, CritDmg: psychube.ThunderousApplause.AdditionalEffect()[calParams.PsychubeAmp].CritDmg()}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nLilya Thunderous Applause resonance 3 Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName = calculatorForReso3.Psychube.Name() + " Reso3"
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForReso1HBD := DamageCalculator{
		Character:         character.Lilya,
		Psychube:          &psychube.HisBoundenDuty,
		Resonance:         &resonances[0],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForReso1HBD.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForReso1HBD.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
	totalCritRate = calculatorForReso1HBD.GetTotalCritRate()
	fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
	for index, damage := range skill1ExtraDamages {
		if totalCritRate >= 1.0 {
			skill1Damages[index] += damage
		} else {
			skill1Damages[index] += damage * totalCritRate
		}
	}
	skill2Damages = calculatorForReso1HBD.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForReso1HBD.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nLilya His Bounden Duty resonance 1 Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName = calculatorForReso1HBD.Psychube.Name() + " Reso1"
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForReso2HBD := DamageCalculator{
		Character:         character.Lilya,
		Psychube:          &psychube.HisBoundenDuty,
		Resonance:         &resonances[1],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForReso2HBD.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForReso2HBD.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
	totalCritRate = calculatorForReso2HBD.GetTotalCritRate()
	fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
	for index, damage := range skill1ExtraDamages {
		if totalCritRate >= 1.0 {
			skill1Damages[index] += damage
		} else {
			skill1Damages[index] += damage * totalCritRate
		}
	}
	skill2Damages = calculatorForReso2HBD.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForReso2HBD.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2}, character.Ultimate, calParams.EnemyHit)
	expectTotalDamage = basicCalculateExpectTotalDmg(skill1Damages, skill2Damages, ultimateDamages)

	fmt.Printf("---------\nLilya His Bounden Duty resonance 2 Final Damage:")
	fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
	fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
	fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
	fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

	psychubeName = calculatorForReso2HBD.Psychube.Name() + " Reso2"
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForReso1Lux := DamageCalculator{
		Character:         character.Lilya,
		Psychube:          &psychube.LuxuriousLeisure,
		Resonance:         &resonances[0],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1Lux1Damages := calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Skill1, calParams.EnemyHit)
	skill1Lux2Damages := calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Skill1, calParams.EnemyHit)
	skill1Lux3Damages := calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
	skill1ExtraLux1Damages := calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
	skill1ExtraLux2Damages := calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
	skill1ExtraLux3Damages := calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
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
	skill2Damages = calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2Lux1Damages := calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Skill2, calParams.EnemyHit)
	skill2Lux2Damages := calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Skill2, calParams.EnemyHit)
	skill2Lux3Damages := calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2}, character.Ultimate, calParams.EnemyHit)
	ultimateLux1Damages := calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Ultimate, calParams.EnemyHit)
	ultimateLux2Damages := calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Ultimate, calParams.EnemyHit)
	ultimateLux3Damages := calculatorForReso1Lux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Ultimate, calParams.EnemyHit)
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

	psychubeName = calculatorForReso1Lux.Psychube.Name() + " Reso1"
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForReso2Lux := DamageCalculator{
		Character:         character.Lilya,
		Psychube:          &psychube.LuxuriousLeisure,
		Resonance:         &resonances[1],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1Lux1Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Skill1, calParams.EnemyHit)
	skill1Lux2Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Skill1, calParams.EnemyHit)
	skill1Lux3Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
	skill1ExtraLux1Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
	skill1ExtraLux2Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
	skill1ExtraLux3Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
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
	skill2Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2Lux1Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Skill2, calParams.EnemyHit)
	skill2Lux2Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Skill2, calParams.EnemyHit)
	skill2Lux3Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2}, character.Ultimate, calParams.EnemyHit)
	ultimateLux1Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Ultimate, calParams.EnemyHit)
	ultimateLux2Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Ultimate, calParams.EnemyHit)
	ultimateLux3Damages = calculatorForReso2Lux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Ultimate, calParams.EnemyHit)
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

	psychubeName = calculatorForReso2Lux.Psychube.Name() + " Reso2"
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForReso3Lux := DamageCalculator{
		Character:         character.Lilya,
		Psychube:          &psychube.LuxuriousLeisure,
		Resonance:         &resonances[2],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1Lux1Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Skill1, calParams.EnemyHit)
	skill1Lux2Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Skill1, calParams.EnemyHit)
	skill1Lux3Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Skill1, calParams.EnemyHit)
	skill1ExtraDamages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
	skill1ExtraLux1Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
	skill1ExtraLux2Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
	skill1ExtraLux3Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
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
	skill2Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2Lux1Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Skill2, calParams.EnemyHit)
	skill2Lux2Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Skill2, calParams.EnemyHit)
	skill2Lux3Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2}, character.Ultimate, calParams.EnemyHit)
	ultimateLux1Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 1}, character.Ultimate, calParams.EnemyHit)
	ultimateLux2Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 2}, character.Ultimate, calParams.EnemyHit)
	ultimateLux3Damages = calculatorForReso3Lux.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.LuxuriousLeisure.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus() * 3}, character.Ultimate, calParams.EnemyHit)
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

	psychubeName = calculatorForReso3Lux.Psychube.Name() + " Reso3"
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	for i := 0; i < 3; i++ {
		calculatorForYearning := DamageCalculator{
			Character:         character.Lilya,
			Psychube:          &psychube.YearningDesire,
			Resonance:         &resonances[i],
			Buff:              &calParams.Buff,
			Debuff:            &calParams.Debuff,
			EnemyDef:          calParams.EnemyDef,
			EnemyCritDef:      enemyCritDef,
			AfflatusAdvantage: calParams.AfflatusAdvantage,
		}

		skill1Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
		skill1BuffDamages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill1, calParams.EnemyHit)
		skill1ExtraDamages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{}, character.ExtraAction1, calParams.EnemyHit)
		skill1ExtraBuffDamages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.ExtraAction1, calParams.EnemyHit)
		totalCritRate := calculatorForYearning.GetTotalCritRate()
		fmt.Printf("\nCrit rate: %d\n", int16(totalCritRate*100))
		for index, damage := range skill1ExtraDamages {
			if totalCritRate >= 1.0 {
				skill1Damages[index] += damage
			} else {
				skill1Damages[index] += damage * totalCritRate
			}
		}
		for index, damage := range skill1ExtraBuffDamages {
			if totalCritRate >= 1.0 {
				skill1BuffDamages[index] += damage
			} else {
				skill1BuffDamages[index] += damage * totalCritRate
			}
		}
		skill2Damages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
		skill2BuffDamages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill2, calParams.EnemyHit)
		ultimateDamages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2}, character.Ultimate, calParams.EnemyHit)
		ultimateBuffDamages := calculatorForYearning.CalculateFinalDamage(DamageCalculatorInfo{CritRate: 0.2, BuffDmgBonus: psychube.YearningDesire.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Ultimate, calParams.EnemyHit)
		expectTotalDamage := basicCalculateExpectTotalDmg(skill1BuffDamages, skill2BuffDamages, ultimateBuffDamages)

		fmt.Printf("---------\nLilya Yearning Desire resonance %d Final Damage:", i+1)
		fmt.Printf("\nSkill 1: %.2f, %.2f, %.2f", skill1Damages[0], skill1Damages[1], skill1Damages[2])
		fmt.Printf("\nSkill 2: %.2f, %.2f, %.2f", skill2Damages[0], skill2Damages[1], skill2Damages[2])
		fmt.Printf("\nUltimate: %.2f", ultimateDamages[0])
		fmt.Printf("\nExpect total damage: %.2f", expectTotalDamage)

		psychubeName = calculatorForYearning.Psychube.Name() + fmt.Sprintf(" Reso%d", i+1)
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
