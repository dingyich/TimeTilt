package buffalo

func (bm *BuffaloMachine) GetGeneralWinningRate() {
	// reelBaseCards := []model.Card{bm.CardNine, bm.CardTen, bm.CardJack, bm.CardQueen, bm.CardKing, bm.CardAce,
	// 	bm.CardEagle, bm.CardTiger, bm.CardWolf, bm.CardDeer, bm.CardBuffalo, bm.CardCoin}

	// r1config := Config.ReelOneConfig()
	// r2config := Config.ReelTwoConfig()
	// r3config := Config.ReelThreeConfig()
	// r4config := Config.ReelFourConfig()
	// r5config := Config.ReelFiveConfig()
	// sizes := []int{len(bm.ReelOne.Cards), len(bm.ReelTwo.Cards), len(bm.ReelThree.Cards), len(bm.ReelFour.Cards), len(bm.ReelFive.Cards)}

	// for i, c := range reelBaseCards {
	// 	payouts := GetPayRate(c.Name)
	// 	distributions := []int{r1config[i], r2config[i], r3config[i], r4config[i], r5config[i]}
	// }
}

func getSingleCardWinningRate() float64 {
	return 0
}
