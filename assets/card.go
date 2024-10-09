package main

type Suit string
type Value string

type Card struct {
	Suit  Suit  `json:"suit"`
	Value Value `json:"value"`
}

func (c *Card) getSuit() Suit {
	return c.Suit
}

func (c *Card) getValue() Value {
	return c.Value
}
