package phe

import (
	"math/rand"
	"testing"
)

func BenchmarkCardsEval(b *testing.B) {
	b.Run("7-cards", func(b *testing.B) { bench(b, 7) })
	b.Run("6-cards", func(b *testing.B) { bench(b, 6) })
	b.Run("5-cards", func(b *testing.B) { bench(b, 5) })
}

// TODO :Fix benchmark. If no hands generation it is must faster. ResetTimer not working
func bench(b *testing.B, cnt int) {
	// generate random hands here
	hands := make([][]Card, b.N)
	for i := 0; i < b.N; i++ {
		hands[i] = randomHand(cnt)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Eval(hands[i]...)
		requireNoErr(b, err)
	}
}

func randomHand(cnt int) []Card {
	cards := make([]Card, cnt)
	for i := 0; i < cnt; i++ {
		cards[i] = Card(rand.Int31n(52))
	}
	return cards
}
