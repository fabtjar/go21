package main

import (
	"testing"
)

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
