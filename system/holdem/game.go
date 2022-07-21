package holdem

import (
	"TimeTilt/system/model"
	"fmt"
)

func NewGame() {
	deck := model.NewDeck()
	deck.Show()
	fmt.Println("----")

	deck.Shuffle()
	deck.Show()
}
