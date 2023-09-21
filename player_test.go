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
		ranks    []string
		expected bool
	}{
		{[]string{}, false},
		{[]string{"1"}, false},
		{[]string{"Ace", "Jack"}, false},
		{[]string{"Jack", "Queen", "King"}, true},
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
