package main

import "log"

type Board struct {
	Player []*Player
	Pile   *Pile
	Deck   *Deck
}

func NewBoard(players []*Player) (*Board, error) {
	pile := NewPile()
	deck, err := NewDeck()
	if err != nil {
		log.Fatalf("error getting deck: %v", err)
	}
	return &Board{
		Player: players,
		Pile:   pile,
		Deck:   deck,
	}, err
}

func (b *Board) GetDeck() *Deck {
	return b.Deck
}

func (b *Board) GetPile() *Pile {
	return b.Pile
}
