package main

import (
	"testing"
)

func TestDrawCard(t *testing.T) {
	deck := NewDeck()
	p := Player{}

	p.DrawCard(deck)
	if len(p.Hand) != 1 {
		t.Errorf("Expected hand length 1, got %v", len(p.Hand))
	}

	for i := 0; i < 10; i++ {
		p.DrawCard(deck)
	}
	if len(p.Hand) != 11 {
		t.Errorf("Expected hand length 11, got %v", len(p.Hand))
	}
}

func TestDrawCardError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got none")
		}
	}()

	deck := []Card{}
	p := Player{}

	p.DrawCard(deck)
}

func TestGetScore(t *testing.T) {
	testCases := []struct {
		ranks    []Rank
		expected int
	}{
		{[]Rank{}, 0},
		{[]Rank{ACE}, 11},
		{[]Rank{TWO, THREE}, 5},
		{[]Rank{FOUR, FIVE, SIX}, 15},
		{[]Rank{ACE, JACK}, 21},
		{[]Rank{ACE, ACE, ACE, ACE}, 14},
		{[]Rank{ACE, QUEEN, KING}, 21},
	}
	for _, tc := range testCases {
		p := Player{}
		for _, rank := range tc.ranks {
			p.Hand = append(p.Hand, Card{Rank: rank})
		}
		if p.GetScore() != tc.expected {
			t.Errorf("Expected score %v, got %v", tc.expected, p.GetScore())
		}
	}
}

func TestIsBust(t *testing.T) {
	testCases := []struct {
		ranks    []Rank
		expected bool
	}{
		{[]Rank{}, false},
		{[]Rank{TEN}, false},
		{[]Rank{ACE, JACK}, false},
		{[]Rank{JACK, QUEEN, KING}, true},
	}

	for _, tc := range testCases {
		p := Player{}
		for _, rank := range tc.ranks {
			p.Hand = append(p.Hand, Card{Rank: rank})
		}
		if p.IsBust() != tc.expected {
			t.Errorf("Expected bust %v, got %v", tc.expected, p.IsBust())
		}
	}
}
