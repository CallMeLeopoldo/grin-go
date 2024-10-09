package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type Deck struct {
	Cards []*Card
}

func NewDeck() (*Deck, error) {
	_, err := os.Stat("assets/deck.json")
	if err != nil {
		return &Deck{}, fmt.Errorf("could not get initial deck: %v", err)
	}
	data, err := os.ReadFile("assets/deck.json")
	if err != nil {
		return &Deck{}, fmt.Errorf("could read cards from data: %v", err)
	}
	deck := &Deck{}
	err = json.Unmarshal(data, deck)
	if err != nil {
		return &Deck{}, fmt.Errorf("could not unmarshall data into deck: %v", err)
	}

	return deck, err
}

func (d *Deck) Shuffle() *Deck {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
	return d
}

func (d *Deck) Draw() {
	d.Cards = d.Cards[:len(d.Cards)-1]
}

func (d *Deck) Reset(cards []*Card) {
	d.Cards = cards
}
