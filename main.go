package main

import (
	"fmt"
)

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
			card, err := deck.TakeRandomCard()
			if err != nil {
				fmt.Println("... Deck is empty!")
				return
			}
			cards = append(cards, card)
			hit(card)
			score := getScore(cards)
			if score > 21 {
				fmt.Printf("... BUST! Score: %v\n", score)
				return
			}
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
