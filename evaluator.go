package phe

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"io/ioutil"
)

var ranks []byte

func init() {
	data, _ := base64.StdEncoding.DecodeString(ranksData)
	reader, _ := gzip.NewReader(bytes.NewReader(data))
	defer reader.Close()
	ranks, _ = ioutil.ReadAll(reader)
}

// Eval evaluates poker hand
// Can evaluate 5,6 and 7 hands
// Rank is the weight of the hand. Bigger rank is best hand
// HandType is the name of combination
// You can load ranks from ranks.dat or embed it in your code
func Eval(cards ...Card) (uint32, HandType, error) {
	size := len(cards) // len is just for shorten code
	if size != 7 && size != 6 && size != 5 {
		return 0, 0, errors.New("cards can be 7,6 or 5 length")
	}

	var p uint32 = 53
	for i := 0; i < size; i++ {
		p = evalCard(p+uint32(cards[i]), ranks)
	}

	if size == 5 || size == 6 {
		p = evalCard(p, ranks)
	}

	tp := HandType(p >> 12)
	if tp == 0 {
		return 0, 0, errors.New("wrong cards")
	}

	rank := p & 0x00000fff
	return rank, tp, nil
}

func evalCard(card uint32, ranks []byte) uint32 {
	offset := uint32(card) * 4
	x := ranks[offset : offset+4]
	r := binary.LittleEndian.Uint32(x)
	return r
}
