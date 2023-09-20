package main

import (
	"testing"
)

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
