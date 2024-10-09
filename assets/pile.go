package main

type Pile struct {
	Cards []*Card
}

func NewPile() *Pile {
	cards := make([]*Card, 0)
	return &Pile{Cards: cards}
}

func (p *Pile) AddCard(c *Card) {
	p.Cards = append(p.Cards, c)
}

func (p *Pile) Reset() {
	p.Cards = make([]*Card, 0)
}
