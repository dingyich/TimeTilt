package model

import (
	"fmt"
	"math/rand"
)

type Card struct {
	Suit  string
	Value string
}

type Deck []Card

func NewDeck() Deck {
	deck := make([]Card, 0)
	suits := []string{"spade", "heart", "club", "diamond"}
	values := []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

	for _, s := range suits {
		for _, v := range values {
			c := Card{Suit: s, Value: v}
			deck = append(deck, c)
		}
	}

	return deck
}

func NewMultiDeck(decks int) Deck {
	deck := make([]Card, 0)
	for i := 0; i < decks; i++ {
		d := NewDeck()
		deck = append(deck, d...)
	}

	return deck
}

func (d *Deck) Shuffle() {
	dd := *d
	for i := 0; i < len(dd); i++ {
		j := rand.Intn(len(dd))
		dd[j], dd[i] = dd[i], dd[j]
		fmt.Printf("j: %v...", j)
		d.Show()
	}
}

func (d *Deck) Show() {
	dd := *d
	for i := 0; i < len(dd); i++ {
		fmt.Printf("%s%s ", dd[i].Value, dd[i].Suit[0:1])
	}
	fmt.Print("\n")
}
