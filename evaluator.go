package handeval

import (
	"encoding/binary"
	"errors"
)

// Eval evaluates poker hand
// Can evaluate 5,6 and 7 hands
// Rank is the weight of the hand. Bigger rank is best hand
// HandType is the name of combination
// You can load ranks from ranks.dat or embed it in your code
func Eval(ranks []byte, cards ...Card) (rank uint32, handType HandType, err error) {
	len := len(cards) // len is just for shorten code
	if len != 7 && len != 6 && len != 5 {
		err = errors.New("cards can be 7,6 or 5 length")
		return
	}

	var p uint32 = 53
	for i := 0; i < len; i++ {
		p = evalCard(p+uint32(cards[i]), ranks)
	}

	if len == 5 || len == 6 {
		p = evalCard(p, ranks)
	}

	handType = HandType(p >> 12)
	rank = p & 0x00000fff
	return
}

func evalCard(card uint32, ranks []byte) uint32 {
	offset := uint32(card) * 4
	x := ranks[offset : offset+4]
	r := binary.LittleEndian.Uint32(x)
	return r
}
