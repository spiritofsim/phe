package phe

import "testing"

func TestEqualHands(t *testing.T) {
	h1 := NewHand(HandTypeStraight, 123)
	h2 := NewHand(HandTypeStraight, 123)
	requireEquals(t, 0, h1.Compare(h2))
	requireEquals(t, 0, h2.Compare(h1))
}

func TestBetterHandType(t *testing.T) {
	h1 := NewHand(HandTypeStraight, 123)
	h2 := NewHand(HandTypeOnePair, 1234)
	requireEquals(t, 1, h1.Compare(h2))
	requireEquals(t, -1, h2.Compare(h1))
}

func TestEqualHandType(t *testing.T) {
	h1 := NewHand(HandTypeStraight, 2)
	h2 := NewHand(HandTypeStraight, 1)
	requireEquals(t, 1, h1.Compare(h2))
	requireEquals(t, -1, h2.Compare(h1))
}
