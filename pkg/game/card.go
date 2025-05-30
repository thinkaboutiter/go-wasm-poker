package game

import (
	"fmt"
	"math/rand"
	"time"
)

// Suit represents a card suit
type Suit int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

// String returns the string representation of a suit
func (s Suit) String() string {
	switch s {
	case Spades:
		return "♠"
	case Hearts:
		return "♥"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		return "?"
	}
}

// Rank represents a card rank
type Rank int

const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

// String returns the string representation of a rank
func (r Rank) String() string {
	switch r {
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		return "?"
	}
}

// Card represents a playing card
type Card struct {
	Rank Rank
	Suit Suit
}

// String returns the string representation of a card
func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Rank.String(), c.Suit.String())
}

// Deck represents a deck of cards
type Deck struct {
	Cards []Card
}

// NewDeck creates a new deck of cards
func NewDeck() *Deck {
	cards := make([]Card, 0, 52)
	for suit := Spades; suit <= Clubs; suit++ {
		for rank := Two; rank <= Ace; rank++ {
			cards = append(cards, Card{Rank: rank, Suit: suit})
		}
	}
	return &Deck{Cards: cards}
}

// Shuffle shuffles the deck
func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// Draw draws n cards from the deck
func (d *Deck) Draw(n int) []Card {
	if n > len(d.Cards) {
		n = len(d.Cards)
	}
	cards := d.Cards[:n]
	d.Cards = d.Cards[n:]
	return cards
}

// DrawOne draws one card from the deck
func (d *Deck) DrawOne() (Card, bool) {
	if len(d.Cards) == 0 {
		return Card{}, false
	}
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card, true
}
