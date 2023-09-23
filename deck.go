package main

import (
	"errors"
	"math/rand"
	"strconv"
)

type Card struct {
	Rank string
	Suit string
}

func (c *Card) GetValue() int {
	switch c.Rank {
	case "Ace":
		return 1
	case "Jack", "Queen", "King":
		return 10
	default:
		value, _ := strconv.Atoi(c.Rank)
		return value
	}
}

func (c Card) String() string {
	return c.Rank + " of " + c.Suit
}

type Deck []Card

func NewDeck() Deck {
	ranks := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}

	deck := Deck{}

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
