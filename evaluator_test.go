package pha

import (
	"io/ioutil"
	"testing"
	"time"
)

// TestEvalAllCombs will enumerate all 133784560 possible 7-card poker hands
func TestEval7CardCombs(t *testing.T) {
	ranks, err := ioutil.ReadFile("ranks.dat")
	requireNoErr(t, err)

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
	ranks, err := ioutil.ReadFile("ranks.dat")
	requireNoErr(t, err)

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
						_, ht, err := Eval(ranks, Card(c0), Card(c1), Card(c2), Card(c3), Card(c4))
						requireNoErr(t, err)
						handTypeSum[ht]++
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
	ranks, err := ioutil.ReadFile("ranks.dat")
	requireNoErr(t, err)

	_, _, err = Eval(ranks, Card(1))
	requireErr(t, err)
}

func TestBadCardNumberEval(t *testing.T) {
	ranks, err := ioutil.ReadFile("ranks.dat")
	requireNoErr(t, err)
	_, _, err = Eval(ranks, Card(100), Card(100), Card(100), Card(100), Card(100))
	requireErr(t, err)
}

func requireNoErr(t testing.TB, err error) {
	if err != nil {
		t.Error(err)
	}
}

func requireErr(t testing.TB, err error) {
	if err == nil {
		t.Error("error expected")
	}
}

func requireEquals(t testing.TB, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("%v expected; %v provided", expected, actual)
	}
}

func requireTrue(t testing.TB, val bool) {
	if !val {
		t.Error("must be true")
	}
}
