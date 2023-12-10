package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_matchType(t *testing.T) {
	assert.Equal(t, matchType(HandBid{hand: "AAAAA"}), FiveOfAKind, "they should be equal")
	assert.Equal(t, matchType(HandBid{hand: "AAAAK"}), FourOfAKind, "they should be equal")
	assert.Equal(t, matchType(HandBid{hand: "AAAKK"}), FullHouse, "they should be equal")
	assert.Equal(t, matchType(HandBid{hand: "65432"}), HighCard, "they should be equal")
	assert.Equal(t, matchType(HandBid{hand: "65422"}), OnePair, "they should be equal")
	assert.Equal(t, matchType(HandBid{hand: "55422"}), TwoPair, "they should be equal")
	assert.Equal(t, matchType(HandBid{hand: "55532"}), ThreeOfAKind, "they should be equal")
}

func Test_sample1(t *testing.T) {
	assert.Equal(t, part1(readInput("sample.txt")), 6440, "they should be equal")
}

func TestMain_input1(t *testing.T) {
	assert.Equal(t, part1(readInput("input.txt")), 249483956, "they should be equal")
}
