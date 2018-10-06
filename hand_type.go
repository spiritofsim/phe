package handeval

// This array is for Stringer
// index is the HandType
var handTypes = []string{
	"Invalid",
	"High card",
	"One pair",
	"Two pairs",
	"Three of a kind",
	"Straight",
	"Flush",
	"Full house",
	"Four of a kind",
	"Straight flush",
}

// HandType is a best type of the hand
type HandType int

// TODO: Remove invalid, return err instead
const (
	HandTypeInvalid       HandType = 0
	HandTypeHighCard      HandType = 1
	HandTypeOnePair       HandType = 2
	HandTypeTwoPairs      HandType = 3
	HandTypeThreeOfaKind  HandType = 4
	HandTypeStraight      HandType = 5
	HandTypeFlush         HandType = 6
	HandTypeFullHouse     HandType = 7
	HandTypeFourOfaKind   HandType = 8
	HandTypeStraightFlush HandType = 9
)

// String returns string representation of hand type
func (ht HandType) String() string {
	return handTypes[ht]
}
