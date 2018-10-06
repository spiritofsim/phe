// TODO: Rename to pha - poker hand eval
package handeval

// This map used for Stringer
var cards = map[Card]string{
	1:  "2♣",
	2:  "2♦",
	3:  "2♥",
	4:  "2♠",
	5:  "3♣",
	6:  "3♦",
	7:  "3♥",
	8:  "3♠",
	9:  "4♣",
	10: "4♦",
	11: "4♥",
	12: "4♠",
	13: "5♣",
	14: "5♦",
	15: "5♥",
	16: "5♠",
	17: "6♣",
	18: "6♦",
	19: "6♥",
	20: "6♠",
	21: "7♣",
	22: "7♦",
	23: "7♥",
	24: "7♠",
	25: "8♣",
	26: "8♦",
	27: "8♥",
	28: "8♠",
	29: "9♣",
	30: "9♦",
	31: "9♥",
	32: "9♠",
	33: "t♣",
	34: "t♦",
	35: "t♥",
	36: "t♠",
	37: "j♣",
	38: "j♦",
	39: "j♥",
	40: "j♠",
	41: "q♣",
	42: "q♦",
	43: "q♥",
	44: "q♠",
	45: "k♣",
	46: "k♦",
	47: "k♥",
	48: "k♠",
	49: "a♣",
	50: "a♦",
	51: "a♥",
	52: "a♠",
}

// This map used for parsing
var cardsBack = map[string]Card{
	"2c": 1,
	"2d": 2,
	"2h": 3,
	"2s": 4,
	"3c": 5,
	"3d": 6,
	"3h": 7,
	"3s": 8,
	"4c": 9,
	"4d": 10,
	"4h": 11,
	"4s": 12,
	"5c": 13,
	"5d": 14,
	"5h": 15,
	"5s": 16,
	"6c": 17,
	"6d": 18,
	"6h": 19,
	"6s": 20,
	"7c": 21,
	"7d": 22,
	"7h": 23,
	"7s": 24,
	"8c": 25,
	"8d": 26,
	"8h": 27,
	"8s": 28,
	"9c": 29,
	"9d": 30,
	"9h": 31,
	"9s": 32,
	"tc": 33,
	"td": 34,
	"th": 35,
	"ts": 36,
	"jc": 37,
	"jd": 38,
	"jh": 39,
	"js": 40,
	"qc": 41,
	"qd": 42,
	"qh": 43,
	"qs": 44,
	"kc": 45,
	"kd": 46,
	"kh": 47,
	"ks": 48,
	"ac": 49,
	"ad": 50,
	"ah": 51,
	"as": 52,
}

// Card in a deck
type Card uint32

// ParseCard return card from it string representation
// Suit can be: s,h,d,c
func ParseCard(str string) Card {
	return cardsBack[str]
}

// String convert card to its string representaion
// Suits will be: ♠,♥,♦,♣
func (c Card) String() string {
	return cards[c]
}
