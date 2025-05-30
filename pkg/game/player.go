package game

// PlayerAction represents an action a player can take
type PlayerAction int

const (
	Fold PlayerAction = iota
	Check
	Call
	Bet
	Raise
	AllIn
)

// String returns the string representation of a player action
func (a PlayerAction) String() string {
	switch a {
	case Fold:
		return "Fold"
	case Check:
		return "Check"
	case Call:
		return "Call"
	case Bet:
		return "Bet"
	case Raise:
		return "Raise"
	case AllIn:
		return "All-In"
	default:
		return "Unknown"
	}
}

// PlayerStatus represents the status of a player in the game
type PlayerStatus int

const (
	Active PlayerStatus = iota
	Folded
	AllInStatus
	Out
)

// Player represents a player in the game
type Player struct {
	ID       string
	Name     string
	Chips    int
	Cards    []Card
	Bet      int
	Status   PlayerStatus
	Position int
}

// NewPlayer creates a new player
func NewPlayer(id, name string, chips int, position int) *Player {
	return &Player{
		ID:       id,
		Name:     name,
		Chips:    chips,
		Cards:    make([]Card, 0),
		Bet:      0,
		Status:   Active,
		Position: position,
	}
}

// PlaceBet places a bet for the player
func (p *Player) PlaceBet(amount int) bool {
	if amount > p.Chips {
		return false
	}
	p.Bet += amount
	p.Chips -= amount
	if p.Chips == 0 {
		p.Status = AllInStatus
	}
	return true
}

// CollectWinnings adds chips to the player's stack
func (p *Player) CollectWinnings(amount int) {
	p.Chips += amount
}

// Fold makes the player fold
func (p *Player) Fold() {
	p.Status = Folded
}

// ResetForNewHand resets the player for a new hand
func (p *Player) ResetForNewHand() {
	p.Cards = make([]Card, 0)
	p.Bet = 0
	if p.Chips > 0 {
		p.Status = Active
	} else {
		p.Status = Out
	}
}

// IsActive returns whether the player is active in the current hand
func (p *Player) IsActive() bool {
	return p.Status == Active
}

// CanAct returns whether the player can act
func (p *Player) CanAct() bool {
	return p.Status == Active || p.Status == AllInStatus
}
