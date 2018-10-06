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

// String returns string representation of hand type
func (ht HandType) String() string {
	return handTypes[ht]
}
