package main

type Player struct {
	Hand []Card
}

func (p *Player) DrawCard(deck Deck) Card {
	card, err := deck.TakeRandomCard()
	if err != nil {
		panic(err)
	}
	p.Hand = append(p.Hand, card)
	return card
}

func (p *Player) GetScore() int {
	score := 0
	hasAce := false
	for _, card := range p.Hand {
		score += card.GetValue()
		if card.Rank == ACE {
			hasAce = true
		}
	}
	if hasAce && score+10 <= 21 {
		score += 10
	}
	return score
}

func (p *Player) IsBust() bool {
	return p.GetScore() > 21
}
