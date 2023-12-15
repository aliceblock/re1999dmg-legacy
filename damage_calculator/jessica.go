package damage_calculator

import (
	"fmt"

	"github.com/aliceblock/re1999dmg/damage_calculator/character"
	"github.com/aliceblock/re1999dmg/damage_calculator/character/resonance"
	"github.com/aliceblock/re1999dmg/damage_calculator/psychube"
)

func JessicaDmgCalculate(calParams CalParams) []DamageResponse {
	damageResponse := []DamageResponse{}

	enemyCritDef := 0.1

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

	calculatorForBoundenDuty := DamageCalculator{
		Character:         character.Jessica,
		Psychube:          &psychube.HisBoundenDuty,
		Resonance:         &resonances[calParams.ResonanceIndex],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1Extra1Damages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.EnemyHit)
	skill1Extra2Damages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.EnemyHit)
	skill1Extra3Damages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.EnemyHit)
	skill2Damages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2ExtraDamages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)
	ultimateDamages := calculatorForBoundenDuty.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	poisonDamage := calculatorForBoundenDuty.CalculateGenesisDamage(DamageCalculatorInfo{}, 0.3)
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
		Character:         character.Jessica,
		Psychube:          &psychube.Hopscotch,
		Resonance:         &resonances[calParams.ResonanceIndex],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1Extra1Damages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.EnemyHit)
	skill1Extra2Damages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.EnemyHit)
	skill1Extra3Damages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2ExtraDamages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff1Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.PsychubeAmp].UltimateMight() * 1}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff2Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.PsychubeAmp].UltimateMight() * 2}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff3Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.PsychubeAmp].UltimateMight() * 3}, character.Ultimate, calParams.EnemyHit)
	ultimateBuff4Damages := calculatorForHopscotch.CalculateFinalDamage(DamageCalculatorInfo{UltimateMight: psychube.Hopscotch.AdditionalEffect()[calParams.PsychubeAmp].UltimateMight() * 4}, character.Ultimate, calParams.EnemyHit)
	poisonDamage = calculatorForHopscotch.CalculateGenesisDamage(DamageCalculatorInfo{}, 0.3)
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

	psychubeName = calculatorForHopscotch.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForBraveNewWorld := DamageCalculator{
		Character:         character.Jessica,
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
	skill1Extra1Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.EnemyHit)
	skill1Extra1BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 0, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
	skill1Extra2Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.EnemyHit)
	skill1Extra2BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 1, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
	skill1Extra3Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.EnemyHit)
	skill1Extra3BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 2, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2BuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill2, calParams.EnemyHit)
	skill2ExtraDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)
	skill2ExtraBuffDamages := calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, IncantationMight: psychube.BraveNewWorld.AdditionalEffect()[calParams.PsychubeAmp].IncantationMight()}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForBraveNewWorld.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	poisonDamage = calculatorForBraveNewWorld.CalculateGenesisDamage(DamageCalculatorInfo{}, 0.3)
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

	psychubeName = calculatorForBraveNewWorld.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	fmt.Println()

	calculatorForBlasphemer := DamageCalculator{
		Character:         character.Jessica,
		Psychube:          &psychube.BlasphemerOfNight,
		Resonance:         &resonances[calParams.ResonanceIndex],
		Buff:              &calParams.Buff,
		Debuff:            &calParams.Debuff,
		EnemyDef:          calParams.EnemyDef,
		EnemyCritDef:      enemyCritDef,
		AfflatusAdvantage: calParams.AfflatusAdvantage,
	}

	skill1Damages = calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill1, calParams.EnemyHit)
	skill1BuffDamages = calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill1, calParams.EnemyHit)
	skill1Extra1Damages = calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 0}, character.Skill1, calParams.EnemyHit)
	skill1Extra1BuffDamages = calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 0, BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill1, calParams.EnemyHit)
	skill1Extra2Damages = calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 1}, character.Skill1, calParams.EnemyHit)
	skill1Extra2BuffDamages = calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 1, BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill1, calParams.EnemyHit)
	skill1Extra3Damages = calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 2}, character.Skill1, calParams.EnemyHit)
	skill1Extra3BuffDamages = calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, ExtraDamageStack: 2, BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill1, calParams.EnemyHit)
	skill2Damages = calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{}, character.Skill2, calParams.EnemyHit)
	skill2BuffDamages = calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill2, calParams.EnemyHit)
	skill2ExtraDamages = calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true}, character.Skill2, calParams.EnemyHit)
	skill2ExtraBuffDamages = calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{HasExtraDamage: true, BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Skill2, calParams.EnemyHit)
	ultimateDamages = calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{}, character.Ultimate, calParams.EnemyHit)
	ultimateBuffDamages := calculatorForBlasphemer.CalculateFinalDamage(DamageCalculatorInfo{BuffDmgBonus: psychube.BlasphemerOfNight.AdditionalEffect()[calParams.PsychubeAmp].DmgBonus()}, character.Ultimate, calParams.EnemyHit)
	poisonDamage = calculatorForBlasphemer.CalculateGenesisDamage(DamageCalculatorInfo{}, 0.3)
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

	psychubeName = calculatorForBlasphemer.Psychube.Name()
	if calParams.PsychubeAmp > 0 {
		psychubeName += fmt.Sprintf(" (A%d)", calParams.PsychubeAmp+1)
	}
	damageResponse = append(damageResponse, DamageResponse{
		Name:   psychubeName,
		Damage: toFixed(expectTotalDamage, 2),
	})

	return damageResponse
}
