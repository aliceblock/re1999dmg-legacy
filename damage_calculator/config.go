package damage_calculator

type CharacterIndex string

const (
	Regulus      CharacterIndex = "regulus"
	AKnight      CharacterIndex = "aknight"
	Lilya        CharacterIndex = "lilya"
	Eagle        CharacterIndex = "eagle"
	Jessica      CharacterIndex = "jessica"
	Charlie      CharacterIndex = "charlie"
	Bkornblume   CharacterIndex = "bkornblume"
	BkornblumeP2 CharacterIndex = "bkornblume_p2"
	BkornblumeP5 CharacterIndex = "bkornblume_p5"
)

var Calculator = map[CharacterIndex]func(CalParams) []DamageResponse{
	Regulus:      RegulusDmgCalculate,
	AKnight:      AKnightDmgCalculate,
	Lilya:        LilyaDmgCalculate,
	Eagle:        EagleDmgCalculate,
	Jessica:      JessicaDmgCalculate,
	Charlie:      CharlieDmgCalculate,
	Bkornblume:   BkornblumeDmgCalculate,
	BkornblumeP2: BkornblumeP2DmgCalculate,
	BkornblumeP5: BkornblumeP5DmgCalculate,
}
