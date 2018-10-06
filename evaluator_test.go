package handeval

import (
	"io/ioutil"
	"testing"
)

// TestEvalAllCobs will enumerate all 133784560 possible 7-card poker hands
func TestEval7CardCombs(t *testing.T) {
	ranks, err := ioutil.ReadFile("ranks.dat")
	requireNoErr(t, err)

	var handTypeSum = make([]int, len(handTypes)) // Stores number of combs
	var count = 0
	for c0 := uint32(1); c0 < 47; c0++ {
		for c1 := c0 + 1; c1 < 48; c1++ {
			for c2 := c1 + 1; c2 < 49; c2++ {
				for c3 := c2 + 1; c3 < 50; c3++ {
					for c4 := c3 + 1; c4 < 51; c4++ {
						for c5 := c4 + 1; c5 < 52; c5++ {
							for c6 := c5 + 1; c6 < 53; c6++ {
								_, ht, err := Eval(ranks, Card(c0), Card(c1), Card(c2), Card(c3), Card(c4), Card(c5), Card(c6))
								requireNoErr(t, err)
								handTypeSum[ht]++
								count++
							}
						}
					}
				}
			}
		}
	}

	for i, hts := range handTypeSum {
		t.Logf("%s\t:\t%d", HandType(i), hts)
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
	requireEquals(t, 0, handTypeSum[HandTypeInvalid])
}

func requireNoErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func requireEquals(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("%v expected; %v provided", expected, actual)
	}
}
