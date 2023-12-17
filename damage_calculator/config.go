package damage_calculator

type CharacterIndex string

const (
	Regulus    CharacterIndex = "regulus"
	AKnight    CharacterIndex = "aknight"
	Lilya      CharacterIndex = "lilya"
	Eagle      CharacterIndex = "eagle"
	Jessica    CharacterIndex = "jessica"
	Charlie    CharacterIndex = "charlie"
	Bkornblume CharacterIndex = "bkornblume"
)

var Calculator = map[CharacterIndex]func(CalParams) []DamageResponse{
	Regulus:    RegulusDmgCalculate,
	AKnight:    AKnightDmgCalculate,
	Lilya:      LilyaDmgCalculate,
	Eagle:      EagleDmgCalculate,
	Jessica:    JessicaDmgCalculate,
	Charlie:    CharlieDmgCalculate,
	Bkornblume: BkornblumeDmgCalculate,
}
