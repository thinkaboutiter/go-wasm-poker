package db

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"go-wasm-poker/pkg/game"
)

// MockSpaceTimeDB is a mock implementation of SpaceTimeDB for the poker game
// This serves as a placeholder until an official Go client for SpaceTimeDB becomes available
type MockSpaceTimeDB struct {
	gameStates     map[string]*game.GameState
	playerProfiles map[string]*PlayerProfile
	gameHistory    map[string][]*GameHistoryEntry
	mu             sync.RWMutex
}

// PlayerProfile represents a player's profile in the database
type PlayerProfile struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	TotalChips    int       `json:"total_chips"`
	GamesPlayed   int       `json:"games_played"`
	GamesWon      int       `json:"games_won"`
	BiggestPot    int       `json:"biggest_pot"`
	LastLoginTime time.Time `json:"last_login_time"`
}

// GameHistoryEntry represents a single game history entry
type GameHistoryEntry struct {
	GameID      string    `json:"game_id"`
	Timestamp   time.Time `json:"timestamp"`
	Players     []string  `json:"players"`
	Winner      string    `json:"winner"`
	PotSize     int       `json:"pot_size"`
	HandSummary string    `json:"hand_summary"`
}

// NewMockSpaceTimeDB creates a new mock SpaceTimeDB
func NewMockSpaceTimeDB() *MockSpaceTimeDB {
	return &MockSpaceTimeDB{
		gameStates:     make(map[string]*game.GameState),
		playerProfiles: make(map[string]*PlayerProfile),
		gameHistory:    make(map[string][]*GameHistoryEntry),
	}
}

// SaveGameState saves the current game state
func (db *MockSpaceTimeDB) SaveGameState(gameID string, state *game.GameState) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	
	// In a real implementation, we would serialize the game state
	// and send it to SpaceTimeDB
	db.gameStates[gameID] = state
	
	log.Printf("Game state saved for game %s", gameID)
	return nil
}

// LoadGameState loads a game state
func (db *MockSpaceTimeDB) LoadGameState(gameID string) (*game.GameState, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	
	state, exists := db.gameStates[gameID]
	if !exists {
		return nil, errors.New("game state not found")
	}
	
	return state, nil
}

// SavePlayerProfile saves a player profile
func (db *MockSpaceTimeDB) SavePlayerProfile(profile *PlayerProfile) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	
	db.playerProfiles[profile.ID] = profile
	
	log.Printf("Player profile saved for player %s", profile.ID)
	return nil
}

// LoadPlayerProfile loads a player profile
func (db *MockSpaceTimeDB) LoadPlayerProfile(playerID string) (*PlayerProfile, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	
	profile, exists := db.playerProfiles[playerID]
	if !exists {
		return nil, errors.New("player profile not found")
	}
	
	return profile, nil
}

// AddGameHistoryEntry adds a game history entry
func (db *MockSpaceTimeDB) AddGameHistoryEntry(gameID string, entry *GameHistoryEntry) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	
	if _, exists := db.gameHistory[gameID]; !exists {
		db.gameHistory[gameID] = make([]*GameHistoryEntry, 0)
	}
	
	db.gameHistory[gameID] = append(db.gameHistory[gameID], entry)
	
	log.Printf("Game history entry added for game %s", gameID)
	return nil
}

// GetGameHistory gets the game history for a game
func (db *MockSpaceTimeDB) GetGameHistory(gameID string) ([]*GameHistoryEntry, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	
	history, exists := db.gameHistory[gameID]
	if !exists {
		return nil, errors.New("game history not found")
	}
	
	return history, nil
}

// SerializeGameState serializes a game state to JSON
func (db *MockSpaceTimeDB) SerializeGameState(state *game.GameState) (string, error) {
	data, err := json.Marshal(state)
	if err != nil {
		return "", err
	}
	
	return string(data), nil
}

// DeserializeGameState deserializes a game state from JSON
func (db *MockSpaceTimeDB) DeserializeGameState(data string) (*game.GameState, error) {
	var state game.GameState
	err := json.Unmarshal([]byte(data), &state)
	if err != nil {
		return nil, err
	}
	
	return &state, nil
}

// Connect simulates connecting to SpaceTimeDB
func (db *MockSpaceTimeDB) Connect() error {
	log.Println("Connecting to SpaceTimeDB (mock)...")
	// Simulate connection delay
	time.Sleep(500 * time.Millisecond)
	log.Println("Connected to SpaceTimeDB (mock)")
	return nil
}

// Disconnect simulates disconnecting from SpaceTimeDB
func (db *MockSpaceTimeDB) Disconnect() error {
	log.Println("Disconnecting from SpaceTimeDB (mock)...")
	// Simulate disconnection delay
	time.Sleep(200 * time.Millisecond)
	log.Println("Disconnected from SpaceTimeDB (mock)")
	return nil
}
