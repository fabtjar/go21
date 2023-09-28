package main

import (
	"errors"
	"math/rand"
)

type Rank int

const (
	ACE Rank = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

func (r Rank) String() string {
	ranks := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	if r < ACE || r > KING {
		return "Unknown"
	}
	return ranks[r]
}

type Suit int

const (
	HEARTS Suit = iota
	DIAMONDS
	CLUBS
	SPADES
)

func (s Suit) String() string {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	if s < HEARTS || s > SPADES {
		return "Unknown"
	}
	return suits[s]
}

type Card struct {
	Rank Rank
	Suit Suit
}

func (c *Card) String() string {
	return c.Rank.String() + " of " + c.Suit.String()
}

func (c *Card) GetValue() int {
	switch c.Rank {
	case ACE:
		return 1
	case JACK, QUEEN, KING:
		return 10
	default:
		return int(c.Rank) + 1
	}
}

type Deck []Card

func NewDeck() Deck {
	ranks := []Rank{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}
	suits := []Suit{HEARTS, DIAMONDS, CLUBS, SPADES}

	deck := make(Deck, 0, 52)

	for _, suit := range suits {
		for _, rank := range ranks {
			card := Card{Rank: rank, Suit: suit}
			deck = append(deck, card)
		}
	}

	return deck
}

func (d *Deck) TakeRandomCard() (Card, error) {
	if len(*d) == 0 {
		return Card{}, errors.New("deck is empty")
	}

	randIndex := rand.Intn(len(*d))
	card := (*d)[randIndex]
	(*d)[randIndex] = (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return card, nil
}
