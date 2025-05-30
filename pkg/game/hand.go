package game

// HandRank represents the rank of a poker hand
type HandRank int

const (
	HighCard HandRank = iota
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

// String returns the string representation of a hand rank
func (h HandRank) String() string {
	switch h {
	case HighCard:
		return "High Card"
	case Pair:
		return "Pair"
	case TwoPair:
		return "Two Pair"
	case ThreeOfAKind:
		return "Three of a Kind"
	case Straight:
		return "Straight"
	case Flush:
		return "Flush"
	case FullHouse:
		return "Full House"
	case FourOfAKind:
		return "Four of a Kind"
	case StraightFlush:
		return "Straight Flush"
	case RoyalFlush:
		return "Royal Flush"
	default:
		return "Unknown"
	}
}

// HandEvaluation represents the evaluation of a poker hand
type HandEvaluation struct {
	Rank  HandRank
	Cards []Card
	Value int // Used for comparing hands of the same rank
}

// EvaluateHand evaluates the best 5-card hand from the given cards
func EvaluateHand(cards []Card) HandEvaluation {
	if len(cards) < 5 {
		return HandEvaluation{Rank: HighCard, Cards: cards, Value: 0}
	}

	// Check for royal flush
	if royal := checkRoyalFlush(cards); royal.Rank == RoyalFlush {
		return royal
	}

	// Check for straight flush
	if straightFlush := checkStraightFlush(cards); straightFlush.Rank == StraightFlush {
		return straightFlush
	}

	// Check for four of a kind
	if fourOfAKind := checkFourOfAKind(cards); fourOfAKind.Rank == FourOfAKind {
		return fourOfAKind
	}

	// Check for full house
	if fullHouse := checkFullHouse(cards); fullHouse.Rank == FullHouse {
		return fullHouse
	}

	// Check for flush
	if flush := checkFlush(cards); flush.Rank == Flush {
		return flush
	}

	// Check for straight
	if straight := checkStraight(cards); straight.Rank == Straight {
		return straight
	}

	// Check for three of a kind
	if threeOfAKind := checkThreeOfAKind(cards); threeOfAKind.Rank == ThreeOfAKind {
		return threeOfAKind
	}

	// Check for two pair
	if twoPair := checkTwoPair(cards); twoPair.Rank == TwoPair {
		return twoPair
	}

	// Check for pair
	if pair := checkPair(cards); pair.Rank == Pair {
		return pair
	}

	// High card
	return checkHighCard(cards)
}

// Helper functions for hand evaluation
func checkRoyalFlush(cards []Card) HandEvaluation {
	// Implementation placeholder
	return HandEvaluation{Rank: HighCard, Cards: cards, Value: 0}
}

func checkStraightFlush(cards []Card) HandEvaluation {
	// Implementation placeholder
	return HandEvaluation{Rank: HighCard, Cards: cards, Value: 0}
}

func checkFourOfAKind(cards []Card) HandEvaluation {
	// Implementation placeholder
	return HandEvaluation{Rank: HighCard, Cards: cards, Value: 0}
}

func checkFullHouse(cards []Card) HandEvaluation {
	// Implementation placeholder
	return HandEvaluation{Rank: HighCard, Cards: cards, Value: 0}
}

func checkFlush(cards []Card) HandEvaluation {
	// Implementation placeholder
	return HandEvaluation{Rank: HighCard, Cards: cards, Value: 0}
}

func checkStraight(cards []Card) HandEvaluation {
	// Implementation placeholder
	return HandEvaluation{Rank: HighCard, Cards: cards, Value: 0}
}

func checkThreeOfAKind(cards []Card) HandEvaluation {
	// Implementation placeholder
	return HandEvaluation{Rank: HighCard, Cards: cards, Value: 0}
}

func checkTwoPair(cards []Card) HandEvaluation {
	// Implementation placeholder
	return HandEvaluation{Rank: HighCard, Cards: cards, Value: 0}
}

func checkPair(cards []Card) HandEvaluation {
	// Implementation placeholder
	return HandEvaluation{Rank: HighCard, Cards: cards, Value: 0}
}

func checkHighCard(cards []Card) HandEvaluation {
	// Implementation placeholder
	return HandEvaluation{Rank: HighCard, Cards: cards, Value: 0}
}

// CompareHands compares two hand evaluations and returns:
// 1 if hand1 is better, -1 if hand2 is better, 0 if they are equal
func CompareHands(hand1, hand2 HandEvaluation) int {
	if hand1.Rank > hand2.Rank {
		return 1
	}
	if hand1.Rank < hand2.Rank {
		return -1
	}
	if hand1.Value > hand2.Value {
		return 1
	}
	if hand1.Value < hand2.Value {
		return -1
	}
	return 0
}
