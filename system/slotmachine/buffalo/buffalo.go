package buffalo

import (
	"TimeTilt/system/slotmachine/model"
	"fmt"
	"math/rand"
	"time"
)

type BuffaloMachine struct {
	CardNine    model.Card
	CardTen     model.Card
	CardJack    model.Card
	CardQueen   model.Card
	CardKing    model.Card
	CardAce     model.Card
	CardEagle   model.Card
	CardTiger   model.Card
	CardWolf    model.Card
	CardDeer    model.Card
	CardBuffalo model.Card
	CardCoin    model.Card
	CardSunset  model.Card
	CardSunset2 model.Card
	CardSunset3 model.Card

	ReelOne   model.Reel
	ReelTwo   model.Reel
	ReelThree model.Reel
	ReelFour  model.Reel
	ReelFive  model.Reel

	WindowSize int

	// in play data
	FreeGameOn      bool
	FreeGames       int
	GoldCollected   int
	FreeGameWinning float64

	CreditBalance  int64
	CreditBase     int
	CentBase       int
	SpinMultiplyer int
	DollarAmount   float64
	Bet            float64
}

func NewMachine(config model.StartConfig) BuffaloMachine {
	bm := initm()
	bm.WindowSize = config.WindowSize
	bm.CentBase = config.CentBase
	bm.CreditBase = config.BaseCredit
	bm.SpinMultiplyer = config.SpinMultiplyer
	bm.Bet = config.Bet
	return bm
}

func (bm *BuffaloMachine) Deposit(dollar float64) {
	bm.DollarAmount += dollar
	bm.CreditBalance = int64(bm.DollarAmount) * int64(100) / int64(bm.CentBase)

	fmt.Printf("$ %v accepted.\n", dollar)
}

func (bm *BuffaloMachine) CashOut() {
	fmt.Println("Thanks for playing.")
	fmt.Printf("Cashing out: $ %v.\n", bm.DollarAmount)
	fmt.Print("Printing ticket")

	for i := 0; i < 6; i++ {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println()
	fmt.Println("------")
}

func (bm *BuffaloMachine) Show() {
	fmt.Println("Reel One: ")
	for _, c := range bm.ReelOne.Cards {
		fmt.Printf("%v", c.Symbol)
	}
	fmt.Println()

	fmt.Println("Reel Two: ")
	for _, c := range bm.ReelTwo.Cards {
		fmt.Printf("%v", c.Symbol)
	}
	fmt.Println()

	fmt.Println("Reel Three: ")
	for _, c := range bm.ReelThree.Cards {
		fmt.Printf("%v", c.Symbol)
	}
	fmt.Println()

	fmt.Println("Reel Four: ")
	for _, c := range bm.ReelFour.Cards {
		fmt.Printf("%v", c.Symbol)
	}
	fmt.Println()

	fmt.Println("Reel Five: ")
	for _, c := range bm.ReelFive.Cards {
		fmt.Printf("%v", c.Symbol)
	}
	fmt.Println()
}

func (bm *BuffaloMachine) Shuffle() {
	for !bm.ValidateShuffle(bm.WindowSize - 1) {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(bm.ReelOne.Cards), func(i, j int) {
			bm.ReelOne.Cards[i], bm.ReelOne.Cards[j] = bm.ReelOne.Cards[j], bm.ReelOne.Cards[i]
		})

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(bm.ReelTwo.Cards), func(i, j int) {
			bm.ReelTwo.Cards[i], bm.ReelTwo.Cards[j] = bm.ReelTwo.Cards[j], bm.ReelTwo.Cards[i]
		})

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(bm.ReelThree.Cards), func(i, j int) {
			bm.ReelThree.Cards[i], bm.ReelThree.Cards[j] = bm.ReelThree.Cards[j], bm.ReelThree.Cards[i]
		})

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(bm.ReelFour.Cards), func(i, j int) {
			bm.ReelFour.Cards[i], bm.ReelFour.Cards[j] = bm.ReelFour.Cards[j], bm.ReelFour.Cards[i]
		})

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(bm.ReelFive.Cards), func(i, j int) {
			bm.ReelFive.Cards[i], bm.ReelFive.Cards[j] = bm.ReelFive.Cards[j], bm.ReelFive.Cards[i]
		})
	}
}

func (bm *BuffaloMachine) ValidateShuffle(distance int) bool {
	if !isValidDistance(bm.ReelOne.Cards, distance) {
		return false
	}

	if !isValidDistance(bm.ReelTwo.Cards, distance) {
		return false
	}

	if !isValidDistance(bm.ReelThree.Cards, distance) {
		return false
	}

	if !isValidDistance(bm.ReelFour.Cards, distance) {
		return false
	}

	if !isValidDistance(bm.ReelFive.Cards, distance) {
		return false
	}

	return true
}

func (bm *BuffaloMachine) EnoughBalance() bool {
	return bm.CreditBalance >= int64(bm.CreditBase)*int64(bm.SpinMultiplyer)
}

func isValidDistance(cards []model.Card, distance int) bool {
	for i := 0; i < len(cards); i++ {
		left := i
		count := 0
		for j := 0; j < distance+1; j++ {
			index := left + j
			if index >= len(cards) {
				index -= len(cards)
			}

			c := cards[index]
			if c.IsScatter || c.IsWild {
				count++
			}
		}

		if count > 1 {
			return false
		}
	}

	return true
}

func (bm *BuffaloMachine) Spin() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(bm.ReelOne.Cards))
	cs := shift(bm.ReelOne.Cards, n)
	bm.ReelOne.Cards = cs

	rand.Seed(time.Now().UnixNano())
	n = rand.Intn(len(bm.ReelTwo.Cards))
	cs = shift(bm.ReelTwo.Cards, n)
	bm.ReelTwo.Cards = cs

	rand.Seed(time.Now().UnixNano())
	n = rand.Intn(len(bm.ReelThree.Cards))
	cs = shift(bm.ReelThree.Cards, n)
	bm.ReelThree.Cards = cs

	rand.Seed(time.Now().UnixNano())
	n = rand.Intn(len(bm.ReelFour.Cards))
	cs = shift(bm.ReelFour.Cards, n)
	bm.ReelFour.Cards = cs

	rand.Seed(time.Now().UnixNano())
	n = rand.Intn(len(bm.ReelFive.Cards))
	cs = shift(bm.ReelFive.Cards, n)
	bm.ReelFive.Cards = cs
}

func shift(cards []model.Card, n int) []model.Card {
	res := make([]model.Card, 0)
	for i := range cards {
		index := i + n
		if index >= len(cards) {
			index -= len(cards)
		}
		res = append(res, cards[index])
	}

	return res
}

func (bm *BuffaloMachine) ShowWindow() {
	for i := 0; i < bm.WindowSize; i++ {
		fmt.Print(bm.ReelOne.Cards[i].Symbol)
		fmt.Print(bm.ReelTwo.Cards[i].Symbol)
		fmt.Print(bm.ReelThree.Cards[i].Symbol)
		fmt.Print(bm.ReelFour.Cards[i].Symbol)
		fmt.Print(bm.ReelFive.Cards[i].Symbol)
		fmt.Println()
	}
	fmt.Println()
}

func (bm *BuffaloMachine) Audit() model.AuditResult {
	return bm.basicAudit()
}

func (bm *BuffaloMachine) basicAudit() model.AuditResult {
	stats := make(map[string]model.HiveStats)
	reel1 := getReelStats(bm.ReelOne.Cards, bm.WindowSize, bm.FreeGameOn)
	reel2 := getReelStats(bm.ReelTwo.Cards, bm.WindowSize, bm.FreeGameOn)
	reel3 := getReelStats(bm.ReelThree.Cards, bm.WindowSize, bm.FreeGameOn)
	reel4 := getReelStats(bm.ReelFour.Cards, bm.WindowSize, bm.FreeGameOn)
	reel5 := getReelStats(bm.ReelFive.Cards, bm.WindowSize, bm.FreeGameOn)

	scatterCount := 0
	if reel1.ContainsScatter {
		scatterCount++
	}
	if reel2.ContainsScatter {
		scatterCount++
	}
	if reel3.ContainsScatter {
		scatterCount++
	}
	if reel4.ContainsScatter {
		scatterCount++
	}
	if reel5.ContainsScatter {
		scatterCount++
	}

	for n, hs := range reel1.Stats {
		hs2, found2 := reel2.Stats[n]
		hs3, found3 := reel3.Stats[n]
		hs4, found4 := reel4.Stats[n]
		hs5, found5 := reel5.Stats[n]

		var fhs model.HiveStats
		m1 := hs.Count
		var m2, m3, m4, m5 int

		if found2 || reel2.ContainsWild {
			m2 = hs2.Count
			if reel2.ContainsWild {
				m2 += 1
			}
			if found3 || reel3.ContainsWild {
				m3 = hs3.Count
				if reel3.ContainsWild {
					m3 += 1
				}
				if found4 || reel4.ContainsWild {
					m4 = hs4.Count
					if reel4.ContainsWild {
						m4 += 1
					}
					if found5 {
						m5 = hs5.Count
						fhs = model.HiveStats{
							Name:       n,
							PayLevel:   5,
							Multiplyer: m1 * m2 * m3 * m4 * m5,
						}
					} else {
						fhs = model.HiveStats{
							Name:       n,
							PayLevel:   4,
							Multiplyer: m1 * m2 * m3 * m4,
						}
					}
				} else {
					fhs = model.HiveStats{
						Name:       n,
						PayLevel:   3,
						Multiplyer: m1 * m2 * m3,
					}
				}
			} else {
				fhs = model.HiveStats{
					Name:       n,
					PayLevel:   2,
					Multiplyer: m1 * m2,
				}
			}
		} else {
			fhs = model.HiveStats{
				Name:       n,
				PayLevel:   1,
				Multiplyer: m1,
			}
		}

		stats[n] = fhs
	}

	scatterhs := model.HiveStats{
		Name:       CardCoinName,
		PayLevel:   scatterCount,
		Multiplyer: 1,
	}
	stats[CardCoinName] = scatterhs

	return model.AuditResult{
		Stats:        stats,
		Multiplyer:   reel2.Multiplyer * reel3.Multiplyer * reel4.Multiplyer,
		ScatterCount: scatterCount,
	}
}

func getReelStats(cards []model.Card, windowSize int, freeGameOn bool) model.ReelAuditStats {
	res := make(map[string]model.HiveStats)
	m := 1
	wild := false
	scatter := false
	for i := 0; i < windowSize; i++ {
		c := cards[i]
		if c.IsWild {
			m = c.Multiplyer
			wild = true
			continue
		}

		if c.IsScatter {
			scatter = true
			continue
		}

		if _, found := res[c.Name]; !found {
			res[c.Name] = model.HiveStats{
				Name:       c.Name,
				Count:      1,
				Multiplyer: 1,
			}
		} else {
			hs := res[c.Name]
			hs.Count++
			res[c.Name] = hs
		}
	}

	if freeGameOn {
		if m == 1 {
			m = 2
		}
	}

	return model.ReelAuditStats{
		Stats:           res,
		ContainsWild:    wild,
		Multiplyer:      m,
		ContainsScatter: scatter,
	}
}

func initm() BuffaloMachine {
	cardNine := model.Card{
		Name:      CardNineName,
		Symbol:    "|    9    |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardNinePayrate(),
	}

	cardTen := model.Card{
		Name:      CardTenName,
		Symbol:    "|   10    |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardTenPayrate(),
	}

	cardJack := model.Card{
		Name:      CardJackName,
		Symbol:    "|    J    |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardJackPayrate(),
	}

	cardQueen := model.Card{
		Name:      CardQueenName,
		Symbol:    "|    Q    |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardQueenPayrate(),
	}

	cardKing := model.Card{
		Name:      CardKingName,
		Symbol:    "|    K    |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardKingPayrate(),
	}

	cardAce := model.Card{
		Name:      CardAceName,
		Symbol:    "|    A    |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardAcePayrate(),
	}

	cardEagle := model.Card{
		Name:      CardEagleName,
		Symbol:    "|  eagle  |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardEaglePayrate(),
	}

	cardTiger := model.Card{
		Name:      CardTigerName,
		Symbol:    "|  tiger  |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardTigerPayrate(),
	}

	cardWolf := model.Card{
		Name:      CardWolfName,
		Symbol:    "|   wolf  |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardWolfPayrate(),
	}

	cardDeer := model.Card{
		Name:      CardDeerName,
		Symbol:    "|   deer  |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardDeerPayrate(),
	}

	cardBuffalo := model.Card{
		Name:      CardBuffaloName,
		Symbol:    "| BUFFALO |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardBuffaloPayrate(),
	}

	cardCoin := model.Card{
		Name:      CardCoinName,
		Symbol:    "|<<<OOO>>>|",
		IsScatter: true,
		IsWild:    false,
		Payrate:   Payrate.CardCoinPayrate(),
	}

	cardSunset := model.Card{
		Name:       CardSunsetName,
		Symbol:     "|  XXXXX  |",
		IsScatter:  false,
		IsWild:     true,
		Payrate:    Payrate.CardSunsetPayrate(),
		Multiplyer: 1,
	}

	cardSunset2 := model.Card{
		Name:       CardSunset2Name,
		Symbol:     "|   *2*   |",
		IsScatter:  false,
		IsWild:     true,
		Payrate:    Payrate.CardSunsetPayrate(),
		Multiplyer: 2,
	}

	cardSunset3 := model.Card{
		Name:       CardSunset3Name,
		Symbol:     "|   *3*   |",
		IsScatter:  false,
		IsWild:     true,
		Payrate:    Payrate.CardSunsetPayrate(),
		Multiplyer: 3,
	}

	reelBaseCards := []model.Card{cardNine, cardTen, cardJack, cardQueen, cardKing, cardAce, cardEagle, cardTiger, cardWolf, cardDeer, cardBuffalo, cardCoin, cardSunset}
	var reelOneCards, reelTwoCards, reelThreeCards, reelFourCards, reelFiveCards []model.Card

	reelOneConfig := Config.ReelOneConfig()
	for i := 0; i < len(reelOneConfig); i++ {
		count := reelOneConfig[i]
		c := reelBaseCards[i]
		for j := 0; j < count; j++ {
			reelOneCards = append(reelOneCards, c)
		}
	}

	reelTwoConfig := Config.ReelTwoConfig()
	for i := 0; i < len(reelTwoConfig); i++ {
		count := reelTwoConfig[i]
		c := reelBaseCards[i]
		for j := 0; j < count; j++ {
			reelTwoCards = append(reelTwoCards, c)
		}
	}

	reelThreeConfig := Config.ReelThreeConfig()
	for i := 0; i < len(reelThreeConfig); i++ {
		count := reelThreeConfig[i]
		c := reelBaseCards[i]
		for j := 0; j < count; j++ {
			reelThreeCards = append(reelThreeCards, c)
		}
	}

	reelFourConfig := Config.ReelFourConfig()
	for i := 0; i < len(reelFourConfig); i++ {
		count := reelFourConfig[i]
		c := reelBaseCards[i]
		for j := 0; j < count; j++ {
			reelFourCards = append(reelFourCards, c)
		}
	}

	reelFiveConfig := Config.ReelFiveConfig()
	for i := 0; i < len(reelFiveConfig); i++ {
		count := reelFiveConfig[i]
		c := reelBaseCards[i]
		for j := 0; j < count; j++ {
			reelFiveCards = append(reelFiveCards, c)
		}
	}

	return BuffaloMachine{
		CardNine:    cardNine,
		CardTen:     cardTen,
		CardJack:    cardJack,
		CardQueen:   cardQueen,
		CardKing:    cardKing,
		CardAce:     cardAce,
		CardEagle:   cardEagle,
		CardTiger:   cardTiger,
		CardWolf:    cardWolf,
		CardDeer:    cardDeer,
		CardBuffalo: cardBuffalo,
		CardCoin:    cardCoin,
		CardSunset:  cardSunset,
		CardSunset2: cardSunset2,
		CardSunset3: cardSunset3,

		ReelOne:   model.Reel{Cards: reelOneCards},
		ReelTwo:   model.Reel{Cards: reelTwoCards},
		ReelThree: model.Reel{Cards: reelThreeCards},
		ReelFour:  model.Reel{Cards: reelFourCards},
		ReelFive:  model.Reel{Cards: reelFiveCards},
	}
}

func MockMachine() BuffaloMachine {
	cardNine := model.Card{
		Name:      CardNineName,
		Symbol:    "|   9   |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardNinePayrate(),
	}

	cardTen := model.Card{
		Name:      CardTenName,
		Symbol:    "|  10   |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardTenPayrate(),
	}

	cardJack := model.Card{
		Name:      CardJackName,
		Symbol:    "|   J   |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardJackPayrate(),
	}

	cardQueen := model.Card{
		Name:      CardQueenName,
		Symbol:    "|   Q   |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardQueenPayrate(),
	}

	cardKing := model.Card{
		Name:      CardKingName,
		Symbol:    "|   K   |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardKingPayrate(),
	}

	cardAce := model.Card{
		Name:      CardAceName,
		Symbol:    "|   A   |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardAcePayrate(),
	}

	cardEagle := model.Card{
		Name:      CardEagleName,
		Symbol:    "| eagle |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardEaglePayrate(),
	}

	cardTiger := model.Card{
		Name:      CardTigerName,
		Symbol:    "| tiger |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardTigerPayrate(),
	}

	cardWolf := model.Card{
		Name:      CardWolfName,
		Symbol:    "|  wolf |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardWolfPayrate(),
	}

	cardDeer := model.Card{
		Name:      CardDeerName,
		Symbol:    "|  deer |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardDeerPayrate(),
	}

	cardBuffalo := model.Card{
		Name:      CardBuffaloName,
		Symbol:    "| buffa |",
		IsScatter: false,
		IsWild:    false,
		Payrate:   Payrate.CardBuffaloPayrate(),
	}

	cardCoin := model.Card{
		Name:      CardCoinName,
		Symbol:    "|   O   |",
		IsScatter: true,
		IsWild:    false,
		Payrate:   Payrate.CardCoinPayrate(),
	}

	cardSunset := model.Card{
		Name:       CardSunsetName,
		Symbol:     "|  ***  |",
		IsScatter:  false,
		IsWild:     true,
		Payrate:    Payrate.CardSunsetPayrate(),
		Multiplyer: 1,
	}

	// cardSunset2 := model.Card{
	// 	Name:       CardSunset2Name,
	// 	Symbol:     "|  *2*  |",
	// 	IsScatter:  false,
	// 	IsWild:     true,
	// 	Payrate:    Payrate.CardSunsetPayrate(),
	// 	Multiplyer: 2,
	// }

	cardSunset3 := model.Card{
		Name:       CardSunset3Name,
		Symbol:     "|  *3*  |",
		IsScatter:  false,
		IsWild:     true,
		Payrate:    Payrate.CardSunsetPayrate(),
		Multiplyer: 3,
	}

	var reelOneCards, reelTwoCards, reelThreeCards, reelFourCards, reelFiveCards []model.Card
	reelOneCards = []model.Card{cardJack, cardWolf, cardKing, cardDeer}
	reelTwoCards = []model.Card{cardWolf, cardSunset3, cardJack, cardBuffalo}
	reelThreeCards = []model.Card{cardSunset3, cardBuffalo, cardNine, cardKing}
	reelFourCards = []model.Card{cardAce, cardKing, cardKing, cardQueen}
	reelFiveCards = []model.Card{cardNine, cardDeer, cardCoin, cardKing}

	return BuffaloMachine{
		CardNine:    cardNine,
		CardTen:     cardTen,
		CardJack:    cardJack,
		CardQueen:   cardQueen,
		CardKing:    cardKing,
		CardAce:     cardAce,
		CardEagle:   cardEagle,
		CardTiger:   cardTiger,
		CardWolf:    cardWolf,
		CardDeer:    cardDeer,
		CardBuffalo: cardBuffalo,
		CardCoin:    cardCoin,
		CardSunset:  cardSunset,

		ReelOne:   model.Reel{Cards: reelOneCards},
		ReelTwo:   model.Reel{Cards: reelTwoCards},
		ReelThree: model.Reel{Cards: reelThreeCards},
		ReelFour:  model.Reel{Cards: reelFourCards},
		ReelFive:  model.Reel{Cards: reelFiveCards},

		WindowSize: 4,
	}
}
