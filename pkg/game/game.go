package game

// GamePhase represents the current phase of the game
type GamePhase int

const (
	PreFlop GamePhase = iota
	Flop
	Turn
	River
	Showdown
)

// String returns the string representation of a game phase
func (p GamePhase) String() string {
	switch p {
	case PreFlop:
		return "Pre-Flop"
	case Flop:
		return "Flop"
	case Turn:
		return "Turn"
	case River:
		return "River"
	case Showdown:
		return "Showdown"
	default:
		return "Unknown"
	}
}

// GameState represents the current state of the game
type GameState struct {
	Players       []*Player
	Deck          *Deck
	CommunityCards []Card
	CurrentPhase  GamePhase
	Pot           int
	CurrentBet    int
	SmallBlind    int
	BigBlind      int
	DealerPos     int
	CurrentPos    int
	LastRaisePos  int
	MinRaise      int
}

// NewGameState creates a new game state
func NewGameState(players []*Player, smallBlind, bigBlind int) *GameState {
	return &GameState{
		Players:       players,
		Deck:          NewDeck(),
		CommunityCards: make([]Card, 0, 5),
		CurrentPhase:  PreFlop,
		Pot:           0,
		CurrentBet:    0,
		SmallBlind:    smallBlind,
		BigBlind:      bigBlind,
		DealerPos:     0,
		CurrentPos:    0,
		LastRaisePos:  -1,
		MinRaise:      bigBlind,
	}
}

// StartNewHand starts a new hand
func (g *GameState) StartNewHand() {
	// Reset game state
	g.Deck = NewDeck()
	g.Deck.Shuffle()
	g.CommunityCards = make([]Card, 0, 5)
	g.CurrentPhase = PreFlop
	g.Pot = 0
	g.CurrentBet = 0
	g.LastRaisePos = -1
	g.MinRaise = g.BigBlind

	// Reset players for new hand
	for _, p := range g.Players {
		p.ResetForNewHand()
	}

	// Move dealer button
	g.DealerPos = (g.DealerPos + 1) % len(g.Players)
	
	// Find next active players for small blind, big blind, and first to act
	sbPos := g.findNextActivePosition(g.DealerPos)
	g.Players[sbPos].PlaceBet(g.SmallBlind)
	
	bbPos := g.findNextActivePosition(sbPos)
	g.Players[bbPos].PlaceBet(g.BigBlind)
	g.CurrentBet = g.BigBlind
	
	// Deal cards to players
	for i := 0; i < 2; i++ {
		for _, p := range g.Players {
			if p.IsActive() || p.Status == AllInStatus {
				card, ok := g.Deck.DrawOne()
				if ok {
					p.Cards = append(p.Cards, card)
				}
			}
		}
	}
	
	// Set current position to player after big blind
	g.CurrentPos = g.findNextActivePosition(bbPos)
}

// findNextActivePosition finds the next active player position
func (g *GameState) findNextActivePosition(pos int) int {
	count := 0
	nextPos := (pos + 1) % len(g.Players)
	
	// If we've checked all positions and found no active players, return the original position
	for count < len(g.Players) {
		if g.Players[nextPos].IsActive() {
			return nextPos
		}
		nextPos = (nextPos + 1) % len(g.Players)
		count++
	}
	
	return pos
}

// DealFlop deals the flop
func (g *GameState) DealFlop() {
	if g.CurrentPhase != PreFlop {
		return
	}
	
	// Burn a card
	_, _ = g.Deck.DrawOne()
	
	// Deal three cards for the flop
	for i := 0; i < 3; i++ {
		card, ok := g.Deck.DrawOne()
		if ok {
			g.CommunityCards = append(g.CommunityCards, card)
		}
	}
	
	g.CurrentPhase = Flop
	g.CurrentBet = 0
	g.LastRaisePos = -1
	g.CurrentPos = g.findNextActivePosition(g.DealerPos)
}

// DealTurn deals the turn
func (g *GameState) DealTurn() {
	if g.CurrentPhase != Flop {
		return
	}
	
	// Burn a card
	_, _ = g.Deck.DrawOne()
	
	// Deal the turn
	card, ok := g.Deck.DrawOne()
	if ok {
		g.CommunityCards = append(g.CommunityCards, card)
	}
	
	g.CurrentPhase = Turn
	g.CurrentBet = 0
	g.LastRaisePos = -1
	g.CurrentPos = g.findNextActivePosition(g.DealerPos)
}

// DealRiver deals the river
func (g *GameState) DealRiver() {
	if g.CurrentPhase != Turn {
		return
	}
	
	// Burn a card
	_, _ = g.Deck.DrawOne()
	
	// Deal the river
	card, ok := g.Deck.DrawOne()
	if ok {
		g.CommunityCards = append(g.CommunityCards, card)
	}
	
	g.CurrentPhase = River
	g.CurrentBet = 0
	g.LastRaisePos = -1
	g.CurrentPos = g.findNextActivePosition(g.DealerPos)
}

// ProcessAction processes a player action
func (g *GameState) ProcessAction(action PlayerAction, amount int) bool {
	player := g.Players[g.CurrentPos]
	
	switch action {
	case Fold:
		player.Fold()
	case Check:
		if g.CurrentBet > player.Bet {
			return false // Can't check if there's a bet
		}
	case Call:
		callAmount := g.CurrentBet - player.Bet
		if !player.PlaceBet(callAmount) {
			return false
		}
		g.Pot += callAmount
	case Bet:
		if g.CurrentBet > 0 {
			return false // Can't bet if there's already a bet
		}
		if amount < g.BigBlind {
			return false // Bet must be at least the big blind
		}
		if !player.PlaceBet(amount) {
			return false
		}
		g.Pot += amount
		g.CurrentBet = amount
		g.LastRaisePos = g.CurrentPos
		g.MinRaise = amount
	case Raise:
		raiseAmount := g.CurrentBet - player.Bet + amount
		if amount < g.MinRaise {
			return false // Raise must be at least the minimum raise
		}
		if !player.PlaceBet(raiseAmount) {
			return false
		}
		g.Pot += raiseAmount
		g.CurrentBet = player.Bet
		g.LastRaisePos = g.CurrentPos
		g.MinRaise = amount
	case AllIn:
		allInAmount := player.Chips
		player.PlaceBet(allInAmount)
		g.Pot += allInAmount
		if player.Bet > g.CurrentBet {
			g.CurrentBet = player.Bet
			g.LastRaisePos = g.CurrentPos
			g.MinRaise = g.CurrentBet - (g.CurrentBet - allInAmount)
		}
	}
	
	// Move to next player
	g.CurrentPos = g.findNextActivePosition(g.CurrentPos)
	
	// Check if betting round is over
	if g.CurrentPos == g.LastRaisePos || g.countActivePlayers() <= 1 {
		g.advancePhase()
	}
	
	return true
}

// countActivePlayers counts the number of active players
func (g *GameState) countActivePlayers() int {
	count := 0
	for _, p := range g.Players {
		if p.IsActive() {
			count++
		}
	}
	return count
}

// advancePhase advances the game to the next phase
func (g *GameState) advancePhase() {
	switch g.CurrentPhase {
	case PreFlop:
		g.DealFlop()
	case Flop:
		g.DealTurn()
	case Turn:
		g.DealRiver()
	case River:
		g.CurrentPhase = Showdown
		g.determineWinners()
	}
}

// determineWinners determines the winners of the hand
func (g *GameState) determineWinners() {
	// Implementation placeholder for determining winners
	// This would evaluate each player's hand and distribute the pot
}

// GetCurrentPlayer returns the current player
func (g *GameState) GetCurrentPlayer() *Player {
	return g.Players[g.CurrentPos]
}

// IsHandOver returns whether the hand is over
func (g *GameState) IsHandOver() bool {
	return g.CurrentPhase == Showdown || g.countActivePlayers() <= 1
}
