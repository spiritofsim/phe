package phe

import (
	"testing"
)

// This tests just for examples
func TestRoyalStreetFlashIsBetterThanStreetFlashWithKing(t *testing.T) {
	ranks, err := LoadRanks("ranks.dat.gz")
	requireNoErr(t, err)

	rsfRank, rsfHandType, err := Eval(ranks, Cardah, Cardkh, Cardqh, Cardjh, Cardth)
	requireNoErr(t, err)
	ksfRank, ksfHandType, err := Eval(ranks, Cardkh, Cardqh, Cardjh, Cardth, Card9h)
	requireNoErr(t, err)

	requireEquals(t, HandTypeStraightFlush, rsfHandType)
	requireEquals(t, HandTypeStraightFlush, ksfHandType)
	requireTrue(t, rsfRank > ksfRank)
}
