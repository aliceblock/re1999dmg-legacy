package damage_calculator

type CharacterIndex string

const (
	Regulus      CharacterIndex = "regulus"
	AKnight      CharacterIndex = "aknight"
	Lilya        CharacterIndex = "lilya"
	Eagle        CharacterIndex = "eagle"
	JessicaR10   CharacterIndex = "jessica_r10"
	JessicaR12   CharacterIndex = "jessica_r12"
	JessicaP2R10 CharacterIndex = "jessica_p2_r10"
	JessicaP2R12 CharacterIndex = "jessica_p2_r12"
	Charlie      CharacterIndex = "charlie"
	Bkornblume   CharacterIndex = "bkornblume"
	BkornblumeP2 CharacterIndex = "bkornblume_p2"
	BkornblumeP5 CharacterIndex = "bkornblume_p5"
	Centurion    CharacterIndex = "centurion"
	Eternity     CharacterIndex = "eternity"
)

var Calculator = map[CharacterIndex]func(CalParams) []DamageResponse{
	Regulus:      RegulusDmgCalculate,
	AKnight:      AKnightDmgCalculate,
	Lilya:        LilyaDmgCalculate,
	Eagle:        EagleDmgCalculate,
	JessicaR10:   JessicaR10DmgCalculate,
	JessicaR12:   JessicaR12DmgCalculate,
	JessicaP2R10: JessicaP2R10DmgCalculate,
	JessicaP2R12: JessicaP2R12DmgCalculate,
	Charlie:      CharlieDmgCalculate,
	Bkornblume:   BkornblumeDmgCalculate,
	BkornblumeP2: BkornblumeP2DmgCalculate,
	BkornblumeP5: BkornblumeP5DmgCalculate,
	Centurion:    CenturionDmgCalculate,
	Eternity:     EternityDmgCalculate,
}
