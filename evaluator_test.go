package phe

import (
	"testing"
	"time"
)

// Issue #1 test
func TestIssue1(t *testing.T) {
	fullHouse, err := Eval(Cardjs, Card8c, Card8s, Cardah, Cardtc, Cardas, Cardad)
	requireNoErr(t, err)

	pair, err := Eval(Cardjs, Card8c, Card8s, Cardah, Cardtc, Card6h, Card7h)
	requireNoErr(t, err)

	requireEquals(t, 1, fullHouse.Compare(pair))
}

func TestFullHouseBetterThanStreet(t *testing.T) {
	fullHouse, err := Eval(Cardah, Cardac, Card2c, Card2d, Card2s)
	requireNoErr(t, err)
	straight, err := Eval(Cardah, Cardks, Cardqh, Cardjd, Cardtd)
	requireNoErr(t, err)

	requireEquals(t, 1, fullHouse.Compare(straight))
}

func TestRoyalStreetFlashIsBetterThanStreetFlashWithKing(t *testing.T) {
	rsf, err := Eval(Cardah, Cardkh, Cardqh, Cardjh, Cardth)
	requireNoErr(t, err)
	sf, err := Eval(Cardkh, Cardqh, Cardjh, Cardth, Card9h)
	requireNoErr(t, err)

	requireEquals(t, 1, rsf.Compare(sf))
}

// TestEvalAllCombs will enumerate all 133784560 possible 7-card poker hands
func TestEval7CardCombs(t *testing.T) {
	var count = 0

	started := time.Now()
	defer func() {
		elapsed := time.Since(started)
		t.Logf("%v Processed in %v", count, elapsed)
	}()

	var handTypeSum = make([]int, len(handTypes)) // Stores number of combs
	for c0 := uint32(1); c0 < 47; c0++ {
		for c1 := c0 + 1; c1 < 48; c1++ {
			for c2 := c1 + 1; c2 < 49; c2++ {
				for c3 := c2 + 1; c3 < 50; c3++ {
					for c4 := c3 + 1; c4 < 51; c4++ {
						for c5 := c4 + 1; c5 < 52; c5++ {
							for c6 := c5 + 1; c6 < 53; c6++ {
								h, err := Eval(Card(c0), Card(c1), Card(c2), Card(c3), Card(c4), Card(c5), Card(c6))
								requireNoErr(t, err)
								handTypeSum[h.Type]++
								count++
							}
						}
					}
				}
			}
		}
	}

	requireEquals(t, 133784560, count)
	requireEquals(t, 23294460, handTypeSum[HandTypeHighCard])
	requireEquals(t, 58627800, handTypeSum[HandTypeOnePair])
	requireEquals(t, 31433400, handTypeSum[HandTypeTwoPairs])
	requireEquals(t, 6461620, handTypeSum[HandTypeThreeOfaKind])
	requireEquals(t, 6180020, handTypeSum[HandTypeStraight])
	requireEquals(t, 4047644, handTypeSum[HandTypeFlush])
	requireEquals(t, 3473184, handTypeSum[HandTypeFullHouse])
	requireEquals(t, 224848, handTypeSum[HandTypeFourOfaKind])
	requireEquals(t, 41584, handTypeSum[HandTypeStraightFlush])
}

// TestEvalAllCobs will enumerate all 2598960 possible 5-card poker hands
// Number of combinations taken from http://suffe.cool/poker/evaluator.html
func TestEval5CardCombs(t *testing.T) {
	var count = 0

	started := time.Now()
	defer func() {
		elapsed := time.Since(started)
		t.Logf("%v Processed in %v", count, elapsed)
	}()

	var handTypeSum = make([]int, len(handTypes)) // Stores number of combs
	for c0 := uint32(1); c0 < 49; c0++ {
		for c1 := c0 + 1; c1 < 50; c1++ {
			for c2 := c1 + 1; c2 < 51; c2++ {
				for c3 := c2 + 1; c3 < 52; c3++ {
					for c4 := c3 + 1; c4 < 53; c4++ {
						h, err := Eval(Card(c0), Card(c1), Card(c2), Card(c3), Card(c4))
						requireNoErr(t, err)
						handTypeSum[h.Type]++
						count++
					}
				}
			}
		}
	}

	requireEquals(t, 2598960, count)
	requireEquals(t, 1302540, handTypeSum[HandTypeHighCard])
	requireEquals(t, 1098240, handTypeSum[HandTypeOnePair])
	requireEquals(t, 123552, handTypeSum[HandTypeTwoPairs])
	requireEquals(t, 54912, handTypeSum[HandTypeThreeOfaKind])
	requireEquals(t, 10200, handTypeSum[HandTypeStraight])
	requireEquals(t, 5108, handTypeSum[HandTypeFlush])
	requireEquals(t, 3744, handTypeSum[HandTypeFullHouse])
	requireEquals(t, 624, handTypeSum[HandTypeFourOfaKind])
	requireEquals(t, 40, handTypeSum[HandTypeStraightFlush])
}

func TestBadCardCount(t *testing.T) {
	_, err := Eval(Card(1))
	requireErr(t, err)
}

func TestBadCardNumberEval(t *testing.T) {
	_, err := Eval(Card(100), Card(100), Card(100), Card(100), Card(100))
	requireErr(t, err)
}
