package main

import (
	"log"
	"os"
	"syscall/js"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"go-wasm-poker/pkg/db"
	"go-wasm-poker/pkg/game"
	"go-wasm-poker/pkg/ui"
)

func main() {
	// Initialize mock database
	mockDB := db.NewMockSpaceTimeDB()
	err := mockDB.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer mockDB.Disconnect()

	// Register JavaScript callbacks for WASM
	js.Global().Set("startNewGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// This would be used to start a new game from JavaScript
		log.Println("Starting new game from JavaScript")
		return nil
	}))

	go func() {
		w := app.NewWindow(
			app.Title("Texas Hold'em Poker"),
			app.Size(800, 600),
		)
		if err := run(w, mockDB); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window, mockDB *db.MockSpaceTimeDB) error {
	// Create sample players
	players := []*game.Player{
		game.NewPlayer("1", "Player 1", 1000, 0),
		game.NewPlayer("2", "Player 2", 1000, 1),
		game.NewPlayer("3", "Player 3", 1000, 2),
		game.NewPlayer("4", "Player 4", 1000, 3),
	}

	// Create game state
	gameState := game.NewGameState(players, 5, 10)
	gameState.StartNewHand()

	// Save initial game state to mock database
	gameID := "game-1"
	err := mockDB.SaveGameState(gameID, gameState)
	if err != nil {
		log.Printf("Failed to save game state: %v", err)
	}

	// Create UI theme and game UI
	theme := ui.NewTheme()
	gameUI := ui.NewGameUI(theme, gameState)

	// Operations variable
	var ops op.Ops

	// Event loop
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			gameUI.SetWindowSize(e.Size)
			gameUI.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}
