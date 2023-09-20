package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}
}

func TestTakeRandomCard(t *testing.T) {
	deck := NewDeck()
	card, _ := deck.TakeRandomCard()
	if len(deck) != 51 {
		t.Errorf("Expected deck length of 51, but got %v", len(deck))
	}
	if card.Rank == "" {
		t.Errorf("Expected card to have a rank, but got %v", card.Rank)
	}
	if card.Suit == "" {
		t.Errorf("Expected card to have a suit, but got %v", card.Suit)
	}
}

func TestTakeAllCards(t *testing.T) {
	deck := NewDeck()
	for i := 0; i < 52; i++ {
		_, err := deck.TakeRandomCard()
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
	}
	if len(deck) != 0 {
		t.Errorf("Expected deck length of 0, but got %v", len(deck))
	}
}

func TestTakeRandomCardWithEmptyDeck(t *testing.T) {
	deck := Deck{}
	card, err := deck.TakeRandomCard()
	if err == nil {
		t.Errorf("Expected error, but got %v", card)
	}
	if len(deck) != 0 {
		t.Errorf("Expected deck length of 0, but got %v", len(deck))
	}
}

func DoCardsMatch(card1 Card, card2 Card) bool {
	return card1.Rank == card2.Rank && card1.Suit == card2.Suit
}

func TestTakeRandomCardIsRandom(t *testing.T) {
	deck := Deck{
		Card{Rank: "Ace", Suit: "Hearts"},
		Card{Rank: "Jack", Suit: "Diamonds"},
		Card{Rank: "Queen", Suit: "Clubs"},
	}
	card1, _ := deck.TakeRandomCard()
	card2, _ := deck.TakeRandomCard()
	card3, _ := deck.TakeRandomCard()
	if DoCardsMatch(card1, card2) || DoCardsMatch(card1, card3) || DoCardsMatch(card2, card3) {
		t.Errorf("Expected cards to be different, but got %v, %v, %v", card1, card2, card3)
	}
}

func TestCardValues(t *testing.T) {
	testCases := []struct {
		rank     string
		expected int
	}{
		{"Ace", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
		{"10", 10},
		{"Jack", 10},
		{"Queen", 10},
		{"King", 10},
	}
	for _, tc := range testCases {
		card := Card{Rank: tc.rank}
		if card.GetValue() != tc.expected {
			t.Errorf("%v should have a value of %v, but got %v", tc.rank, tc.expected, card.GetValue())
		}
	}
}

func TestShowScore(t *testing.T) {
	testCases := []struct {
		ranks    []string
		expected int
	}{
		{[]string{}, 0},
		{[]string{"1"}, 1},
		{[]string{"2", "3"}, 5},
		{[]string{"4", "5", "6"}, 15},
		{[]string{"Ace", "Jack"}, 21},
		{[]string{"Ace", "Ace", "Ace", "Ace"}, 14},
		{[]string{"Ace", "Queen", "King"}, 21},
	}
	for _, tc := range testCases {
		cards := []Card{}
		for _, rank := range tc.ranks {
			cards = append(cards, Card{Rank: rank})
		}
		score := getScore(cards)
		if score != tc.expected {
			t.Errorf("%v should have a score of %v, but got %v", tc.ranks, tc.expected, score)
		}
	}
}
