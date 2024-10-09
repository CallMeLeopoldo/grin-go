package main

type Player struct {
	Name       string
	GameScore  int
	MatchScore int
	Hand       []*Card
}

func (p *Player) GetHands() []*Card {
	return p.Hand
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) GetMatchScore() int {
	return p.MatchScore
}

func (p *Player) GetGameScore() int {
	return p.GameScore
}
