package psychube

type Psychube struct {
	// Fields are private (start with a lowercase letter)
	atk              float64
	atkPercent       float64
	dmgBonus         float64
	incantationMight float64
	ultimateMight    float64
	critRate         float64
	critDmg          float64
	additionalEffect *Stat
}

type Stat struct {
	atk              float64
	atkPercent       float64
	dmgBonus         float64
	incantationMight float64
	ultimateMight    float64
	critRate         float64
	critDmg          float64
}

// Getter methods for Psychube fields

func (p *Psychube) Atk() float64 {
	return p.atk
}

func (p *Psychube) AtkPercent() float64 {
	return p.atkPercent
}

func (p *Psychube) DmgBonus() float64 {
	return p.dmgBonus
}

func (p *Psychube) IncantationMight() float64 {
	return p.incantationMight
}

func (p *Psychube) UltimateMight() float64 {
	return p.ultimateMight
}

func (p *Psychube) CritRate() float64 {
	return p.critRate
}

func (p *Psychube) CritDmg() float64 {
	return p.critDmg
}

func (p *Psychube) AdditionalEffect() *Stat {
	return p.additionalEffect
}

// Getter methods for Stat fields

func (s *Stat) Atk() float64 {
	return s.atk
}

func (s *Stat) AtkPercent() float64 {
	return s.atkPercent
}

func (s *Stat) DmgBonus() float64 {
	return s.dmgBonus
}

func (s *Stat) IncantationMight() float64 {
	return s.incantationMight
}

func (s *Stat) UltimateMight() float64 {
	return s.ultimateMight
}

func (s *Stat) CritRate() float64 {
	return s.critRate
}

func (s *Stat) CritDmg() float64 {
	return s.critDmg
}

var ThunderousApplause = Psychube{
	atk:      330,
	critRate: 0.16,
	additionalEffect: &Stat{
		critDmg: 0.16,
	},
}

/*
For each enemy target defeated by the carrier, Ultimate Might +4% for the carrier. Stacks up to 4 times.
*/
var Hopscotch = Psychube{
	atk:              370,
	incantationMight: 0.18,
	additionalEffect: &Stat{
		ultimateMight: 0.04,
	},
}

/*
After the carrier casts an Ultimate, Incantation Might of the next incantation +20%.
*/
var BraveNewWorld = Psychube{
	atk:           370,
	ultimateMight: 0.18,
	additionalEffect: &Stat{
		incantationMight: 0.2,
	},
}

var HisBoundenDuty = Psychube{
	atk:      410,
	dmgBonus: 0.12,
}

/*
When the wearer attacks, if they do not have Afflatus advantage/disadvantage, the damage is increased by 6%. If the wearer's Afflatus is Spirit or Intellect, this effect is increased to 12%.
*/
var BalancePlease = Psychube{
	atk:              380,
	incantationMight: 0.18,
	additionalEffect: &Stat{
		dmgBonus: 0.12,
	},
}
