package handeval

import (
	"io/ioutil"
	"testing"
)

// This tests just for examples

func TestRoyalStreetFlashIsBetterThanStreetFlashWithKing(t *testing.T) {
	ranks, err := ioutil.ReadFile("ranks.dat")
	requireNoErr(t, err)

	rsfRank, rsfHandType, err := Eval(ranks, ParseCard("ah"), ParseCard("kh"), ParseCard("qh"), ParseCard("jh"), ParseCard("th"))
	requireNoErr(t, err)
	ksfRank, ksfHandType, err := Eval(ranks, ParseCard("kh"), ParseCard("qh"), ParseCard("jh"), ParseCard("th"), ParseCard("9h"))
	requireNoErr(t, err)

	requireEquals(t, HandTypeStraightFlush, rsfHandType)
	requireEquals(t, HandTypeStraightFlush, ksfHandType)
	requireTrue(t, rsfRank > ksfRank)
}
