package main

import (
	"testing"
)

func TestCardValues(t *testing.T) {
	testCases := []struct {
		rank     string
		expected int
	}{
		{"Ace", 11},
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
