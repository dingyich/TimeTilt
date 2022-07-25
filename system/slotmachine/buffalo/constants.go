package buffalo

const (
	CardNineName  string = "nine"
	CardTenName   string = "ten"
	CardJackName  string = "jack"
	CardQueenName string = "queen"
	CardKingName  string = "king"
	CardAceName   string = "ace"

	CardEagleName   string = "eagle"
	CardTigerName   string = "tiger"
	CardWolfName    string = "wolf"
	CardDeerName    string = "deer"
	CardBuffaloName string = "buffalo"

	CardCoinName    string = "coin"
	CardSunsetName  string = "sunset"
	CardSunset2Name string = "sunset2"
	CardSunset3Name string = "sunset3"

	Payrate payrate = ""
	Config  config  = ""
)

type payrate string

func (p payrate) CardCoinPayrate() []int   { return []int{0, 0, 0, 80, 400, 800} }
func (p payrate) CardSunsetPayrate() []int { return []int{0, 0, 0, 0, 0} }

func (p payrate) CardNinePayrate() []int  { return []int{0, 2, 5, 10, 40} }
func (p payrate) CardTenPayrate() []int   { return []int{0, 0, 5, 20, 50} }
func (p payrate) CardJackPayrate() []int  { return []int{0, 0, 5, 30, 60} }
func (p payrate) CardQueenPayrate() []int { return []int{0, 0, 5, 40, 80} }
func (p payrate) CardKingPayrate() []int  { return []int{0, 0, 10, 50, 80} }
func (p payrate) CardAcePayrate() []int   { return []int{0, 0, 10, 60, 100} }

func (p payrate) CardEaglePayrate() []int   { return []int{0, 10, 40, 100, 150} }
func (p payrate) CardTigerPayrate() []int   { return []int{0, 10, 40, 100, 150} }
func (p payrate) CardWolfPayrate() []int    { return []int{0, 20, 80, 200, 250} }
func (p payrate) CardDeerPayrate() []int    { return []int{0, 20, 80, 200, 250} }
func (p payrate) CardBuffaloPayrate() []int { return []int{0, 50, 200, 250, 300} }

func GetPayRate(name string) []int {
	switch name {
	case CardNineName:
		return Payrate.CardNinePayrate()

	case CardTenName:
		return Payrate.CardTenPayrate()

	case CardJackName:
		return Payrate.CardJackPayrate()

	case CardQueenName:
		return Payrate.CardQueenPayrate()

	case CardKingName:
		return Payrate.CardKingPayrate()

	case CardAceName:
		return Payrate.CardAcePayrate()

	case CardEagleName:
		return Payrate.CardEaglePayrate()

	case CardTigerName:
		return Payrate.CardTigerPayrate()

	case CardWolfName:
		return Payrate.CardWolfPayrate()

	case CardDeerName:
		return Payrate.CardDeerPayrate()

	case CardBuffaloName:
		return Payrate.CardBuffaloPayrate()

	case CardCoinName:
		return Payrate.CardCoinPayrate()

	default:
		return nil
	}
}

type config string

func (c config) ReelBaseConfig() []int  { return []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} }
func (c config) ReelOneConfig() []int   { return []int{3, 3, 3, 3, 3, 3, 1, 1, 1, 1, 1, 1, 0} }
func (c config) ReelTwoConfig() []int   { return []int{3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 1, 1} }
func (c config) ReelThreeConfig() []int { return []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1} }
func (c config) ReelFourConfig() []int  { return []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1} }
func (c config) ReelFiveConfig() []int  { return []int{1, 1, 1, 1, 1, 1, 4, 4, 4, 4, 4, 1, 0} }
