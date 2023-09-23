package main

import "testing"

func TestNewGame(t *testing.T) {
	game := NewGame()
	if game.Playing {
		t.Errorf("Expected game to not be playing, but got true")
	}
}

func TestStart(t *testing.T) {
	game := NewGame()
	game.Start()
	if !game.Playing {
		t.Errorf("Expected game to be playing, but got false")
	}
}

func TestHit(t *testing.T) {
	game := NewGame()
	game.player = Player{}

	game.Hit()

	if len(game.player.Hand) != 1 {
		t.Errorf("Expected player to have 1 card, but got %v", len(game.player.Hand))
	}

	for i := 0; i < 20; i++ {
		game.Hit()
	}

	if !game.player.IsBust() {
		t.Errorf("Expected player to be bust, but got false")
	}
}

func TestStand(t *testing.T) {
	game := NewGame()
	game.Stand()
	if game.Playing {
		t.Errorf("Expected game to not be playing, but got true")
	}
}

func TestExit(t *testing.T) {
	game := NewGame()
	game.Exit()
	if game.Playing {
		t.Errorf("Expected game to not be playing, but got true")
	}
}

func TestDealerTurn(t *testing.T) {
	game := NewGame()
	game.Stand()
	dealerScore := game.dealer.GetScore()
	if dealerScore < game.player.GetScore() {
		t.Errorf("Expected dealer score to be greater than or exqual to player score, but got %v", dealerScore)
	}
}
