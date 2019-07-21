package phe

// Hand contains information about hand type and rank
type Hand struct {
	Type HandType
	Rank uint32
}

func NewHand(ht HandType, rank uint32) Hand {
	return Hand{Type: ht, Rank: rank}
}

// Compare return 1 if this hand better than other; -1 if other better than this and 0 if hands are equal
func (h Hand) Compare(other Hand) int {
	if h.Rank == other.Rank && h.Type == other.Type {
		return 0
	}

	if h.Type > other.Type {
		return 1
	}

	if h.Type < other.Type {
		return -1
	}

	if h.Rank > other.Rank {
		return 1
	}
	return -1
}
