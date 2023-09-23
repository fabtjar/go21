package main

import (
	"fmt"
)

type Game struct {
	deck    Deck
	player  Player
	dealer  Player
	Playing bool
}

func NewGame() Game {
	return Game{
		deck:    NewDeck(),
		player:  Player{},
		dealer:  Player{},
		Playing: false,
	}
}

func (g *Game) Start() {
	fmt.Println("Welcome to Go 21!")
	g.Playing = true
}

func (g *Game) Hit() {
	card := g.player.DrawCard(g.deck)
	fmt.Println("You draw a", card.Rank, "of", card.Suit)
	if g.player.IsBust() {
		fmt.Println("You BUST!")
		g.Playing = false
	}
}

func (g *Game) Stand() {
	fmt.Println("You stand!")
	g.dealerTurn()
	g.Playing = false
}

func (g *Game) Hand() {
	fmt.Println("Your hand:", g.player.Hand)
}

func (g *Game) Score() {
	fmt.Println("Your score:", g.player.GetScore())
}

func (g *Game) Exit() {
	fmt.Println("Goodbye!")
	g.Playing = false
}

func (g *Game) dealerTurn() {
	fmt.Println("Dealer's turn!")

	target := g.player.GetScore()

	for g.dealer.GetScore() < target {
		card := g.dealer.DrawCard(g.deck)
		fmt.Println("Dealer hits and gets a", card.Rank, "of", card.Suit)
	}

	score := g.dealer.GetScore()
	if score > 21 {
		fmt.Printf("Dealer BUST! Score: %v\n", score)
		fmt.Println("You win!")
	} else {
		fmt.Printf("Dealer stands with score: %v\n", score)
		fmt.Println("You lose!")
	}
}
