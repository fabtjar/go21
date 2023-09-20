package main

import (
	"fmt"
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

func (d *Deck) TakeRandomCard() Card {
	randIndex := rand.Intn(len(*d))
	card := (*d)[randIndex]
	(*d)[randIndex] = (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return card
}

func hit(card Card) {
	fmt.Println("... Hit!")
	fmt.Println("You got a", card.Rank, "of", card.Suit)
}

func stand() {
	fmt.Println("... Stand!")
}

func showHand(cards []Card) {
	fmt.Printf("... Hand: %v\n", cards)
}

func getScore(cards []Card) int {
	score := 0
	hasAce := false
	for _, card := range cards {
		score += card.GetValue()
		if card.Rank == "Ace" {
			hasAce = true
		}
	}
	if hasAce && score+10 <= 21 {
		score += 10
	}
	return score
}

func showScore(cards []Card) {
	fmt.Printf("... Score: %v\n", getScore(cards))
}

func invalid() {
	fmt.Println("... Invalid command!")
}

func main() {
	deck := NewDeck()
	cards := []Card{}

	var command string
	fmt.Println("Welcome to Go 21!")
	for {
		fmt.Print("> ")
		fmt.Scanf("%s", &command)

		switch command {
		case "hit":
			cards = append(cards, deck.TakeRandomCard())
			hit(cards[len(cards)-1])
		case "stand":
			stand()
		case "cards", "hand":
			showHand(cards)
		case "score":
			showScore(cards)
		case "exit", "quit":
			return
		default:
			invalid()
		}
	}

}
