package psychube

type Amplification int16

const (
	Amp1 Amplification = 0
	Amp2 Amplification = 1
	Amp3 Amplification = 2
	Amp4 Amplification = 3
	Amp5 Amplification = 4
)

type Psychube struct {
	// Fields are private (start with a lowercase letter)
	atk              float64
	atkPercent       float64
	dmgBonus         float64
	incantationMight float64
	ultimateMight    float64
	critRate         float64
	critDmg          float64
	additionalEffect map[Amplification]*Stat
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

func (p *Psychube) AdditionalEffect() map[Amplification]*Stat {
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
	additionalEffect: map[Amplification]*Stat{
		Amp1: {critDmg: 0.16},
		Amp2: {critDmg: 0.2},
		Amp3: {critDmg: 0.24},
		Amp4: {critDmg: 0.28},
		Amp5: {critDmg: 0.32},
	},
}

/*
For each enemy target defeated by the carrier, Ultimate Might +4% for the carrier. Stacks up to 4 times.
*/
var Hopscotch = Psychube{
	atk:              370,
	incantationMight: 0.18,
	additionalEffect: map[Amplification]*Stat{
		Amp1: {ultimateMight: 0.04},
		Amp2: {ultimateMight: 0.05},
		Amp3: {ultimateMight: 0.06},
		Amp4: {ultimateMight: 0.07},
		Amp5: {ultimateMight: 0.08},
	},
}

/*
After the carrier casts an Ultimate, Incantation Might of the next incantation +20%.
*/
var BraveNewWorld = Psychube{
	atk:           370,
	ultimateMight: 0.18,
	additionalEffect: map[Amplification]*Stat{
		Amp1: {incantationMight: 0.2},
		Amp2: {incantationMight: 0.25},
		Amp3: {incantationMight: 0.30},
		Amp4: {incantationMight: 0.35},
		Amp5: {incantationMight: 0.4},
	},
}

var HisBoundenDuty = Psychube{
	atk:      410,
	dmgBonus: 0.12,
}

var BlasphemerOfNight = Psychube{
	atk:              360,
	incantationMight: 0.18,
	additionalEffect: map[Amplification]*Stat{
		Amp1: {dmgBonus: 0.12},
		Amp2: {dmgBonus: 0.15},
		Amp3: {dmgBonus: 0.18},
		Amp4: {dmgBonus: 0.21},
		Amp5: {dmgBonus: 0.24},
	},
}

var LuxuriousLeisure = Psychube{
	atk:           380,
	ultimateMight: 0.18,
	additionalEffect: map[Amplification]*Stat{
		Amp1: {dmgBonus: 0.05},
		Amp2: {dmgBonus: 0.06},
		Amp3: {dmgBonus: 0.07},
		Amp4: {dmgBonus: 0.08},
		Amp5: {dmgBonus: 0.09},
	},
}

/*
When the wearer attacks, if they do not have Afflatus advantage/disadvantage, the damage is increased by 6%. If the wearer's Afflatus is Spirit or Intellect, this effect is increased to 12%.
*/
var BalancePlease = Psychube{
	atk:              380,
	incantationMight: 0.18,
	additionalEffect: map[Amplification]*Stat{
		Amp1: {dmgBonus: 0.12},
		Amp2: {dmgBonus: 0.15},
		Amp3: {dmgBonus: 0.18},
		Amp4: {dmgBonus: 0.21},
		Amp5: {dmgBonus: 0.24},
	},
}
