package pha

import (
	"io/ioutil"
	"math/rand"
	"testing"
)

func BenchmarkCardsEval(b *testing.B) {
	ranks, err := ioutil.ReadFile("ranks.dat")
	requireNoErr(b, err)

	b.Run("7-cards", func(b *testing.B) { bench(b, 7, ranks) })
	b.Run("6-cards", func(b *testing.B) { bench(b, 6, ranks) })
	b.Run("5-cards", func(b *testing.B) { bench(b, 5, ranks) })
}

// TODO :Fix benchmark. If no hands generation it is must faster. ResetTimer not working
func bench(b *testing.B, cnt int, ranks []byte) {
	// generate random hands here
	hands := make([][]Card, b.N)
	for i := 0; i < b.N; i++ {
		hands[i] = randomHand(cnt)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Eval(ranks, hands[i]...)
	}
}

func randomHand(cnt int) []Card {
	cards := make([]Card, cnt)
	for i := 0; i < cnt; i++ {
		cards[i] = Card(rand.Int31n(52))
	}
	return cards
}
