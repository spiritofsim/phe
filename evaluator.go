package handeval

import (
	"encoding/binary"
	"errors"
)

// Eval evaluates poker hand
// Can evaluate 5,6 and 7 hands
// Rank is the weight of the hand. Bigger rank is best hand
// HandType is the name of combination
func Eval(ranks []byte, cards ...Card) (rank uint32, handType HandType, err error) {
	if len(cards) != 7 && len(cards) != 6 && len(cards) != 5 {
		err = errors.New("cards can be 7,6 or 5 length")
	}

	var p uint32 = 53
	for i := 0; i < len(cards); i++ {
		p = evalCard(p+uint32(cards[i]), ranks)
	}

	if len(cards) == 5 || len(cards) == 6 {
		p = evalCard(p, ranks)
	}

	return p & 0x00000fff, HandType(p >> 12), nil
}

func evalCard(card uint32, ranks []byte) uint32 {
	offset := uint32(card) * 4
	x := ranks[offset : offset+4]
	r := binary.LittleEndian.Uint32(x)
	return r
}
