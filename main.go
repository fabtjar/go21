package main

import (
	"fmt"
)

func dealerTurn(deck Deck, target int) {
	fmt.Println("Dealer's turn!")
	dealer := Player{}

	for dealer.GetScore() < target {
		card := dealer.DrawCard(deck)
		fmt.Println("Dealer hits and gets a", card.Rank, "of", card.Suit)
	}

	score := dealer.GetScore()
	if score > 21 {
		fmt.Printf("Dealer BUST! Score: %v\n", score)
		fmt.Println("You win!")
	} else {
		fmt.Printf("Dealer stands with score: %v\n", score)
		fmt.Println("You lose!")
	}
}

func playerTurn(deck Deck) {
	fmt.Println("Your turn!")
	player := Player{}

	var command string
	for {
		fmt.Print("> ")
		fmt.Scanf("%s", &command)
		switch command {
		case "hit":
			card := player.DrawCard(deck)
			fmt.Println("You draw a", card.Rank, "of", card.Suit)
			if player.IsBust() {
				fmt.Println("You BUST!")
				return
			}
		case "stick", "stand":
			dealerTurn(deck, player.GetScore())
			return
		case "cards", "hand":
			fmt.Println("Your hand:", player.Hand)
		case "score":
			fmt.Println("Your score:", player.GetScore())
		case "exit", "quit":
			return
		default:
			fmt.Println("Invalid command!")
		}
	}
}

func main() {
	fmt.Println("Welcome to Go 21!")
	deck := NewDeck()
	playerTurn(deck)
}
