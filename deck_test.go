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
	if card.Rank < ACE || card.Rank > KING {
		t.Errorf("Expected card to have a valid rank, but got %v", card.Rank)
	}
	if card.Suit < HEARTS || card.Suit > SPADES {
		t.Errorf("Expected card to have a valid suit, but got %v", card.Suit)
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
		Card{Rank: ACE, Suit: HEARTS},
		Card{Rank: JACK, Suit: DIAMONDS},
		Card{Rank: QUEEN, Suit: CLUBS},
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
		rank     Rank
		expected int
	}{
		{ACE, 1},
		{TWO, 2},
		{THREE, 3},
		{FOUR, 4},
		{FIVE, 5},
		{SIX, 6},
		{SEVEN, 7},
		{EIGHT, 8},
		{NINE, 9},
		{TEN, 10},
		{JACK, 10},
		{QUEEN, 10},
		{KING, 10},
	}
	for _, tc := range testCases {
		card := Card{Rank: tc.rank}
		if card.GetValue() != tc.expected {
			t.Errorf("%v should have a value of %v, but got %v", tc.rank, tc.expected, card.GetValue())
		}
	}
}
