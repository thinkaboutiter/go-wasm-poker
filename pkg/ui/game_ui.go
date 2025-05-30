package ui

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"go-wasm-poker/pkg/game"
)

// Theme holds the material design theme
type Theme struct {
	*material.Theme
	CardBack     color.NRGBA
	CardFront    color.NRGBA
	TableColor   color.NRGBA
	ButtonColor  color.NRGBA
	PlayerColor  color.NRGBA
	ActivePlayer color.NRGBA
}

// NewTheme creates a new theme
func NewTheme() *Theme {
	th := material.NewTheme()
	return &Theme{
		Theme:        th,
		CardBack:     color.NRGBA{R: 0x20, G: 0x20, B: 0x80, A: 0xFF},
		CardFront:    color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		TableColor:   color.NRGBA{R: 0x00, G: 0x80, B: 0x40, A: 0xFF},
		ButtonColor:  color.NRGBA{R: 0x30, G: 0x30, B: 0x30, A: 0xFF},
		PlayerColor:  color.NRGBA{R: 0x40, G: 0x40, B: 0x40, A: 0xFF},
		ActivePlayer: color.NRGBA{R: 0x60, G: 0x60, B: 0x00, A: 0xFF},
	}
}

// GameUI represents the UI for the poker game
type GameUI struct {
	theme        *Theme
	gameState    *game.GameState
	foldButton   widget.Clickable
	checkButton  widget.Clickable
	callButton   widget.Clickable
	betButton    widget.Clickable
	raiseButton  widget.Clickable
	allInButton  widget.Clickable
	betSlider    widget.Float
	betAmount    int
	windowSize   image.Point
	cardImages   map[string]image.Image // For future card images
}

// NewGameUI creates a new game UI
func NewGameUI(theme *Theme, gameState *game.GameState) *GameUI {
	return &GameUI{
		theme:      theme,
		gameState:  gameState,
		betAmount:  gameState.BigBlind,
		cardImages: make(map[string]image.Image),
	}
}

// SetWindowSize sets the window size
func (ui *GameUI) SetWindowSize(size image.Point) {
	ui.windowSize = size
}

// Layout lays out the UI
func (ui *GameUI) Layout(gtx layout.Context) layout.Dimensions {
	// Save operations
	ops := gtx.Ops
	
	// Draw poker table background
	paintRect(gtx, gtx.Constraints.Max, ui.theme.TableColor)
	
	// Layout the game elements
	layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceEvenly,
	}.Layout(gtx,
		layout.Flexed(0.7, func(gtx layout.Context) layout.Dimensions {
			return ui.layoutGameArea(gtx)
		}),
		layout.Flexed(0.3, func(gtx layout.Context) layout.Dimensions {
			return ui.layoutControls(gtx)
		}),
	)
	
	// Handle button clicks
	ui.handleButtonClicks()
	
	// Return dimensions
	return layout.Dimensions{Size: gtx.Constraints.Max}
}

// layoutGameArea lays out the game area
func (ui *GameUI) layoutGameArea(gtx layout.Context) layout.Dimensions {
	// Draw community cards
	communityCardsArea := layout.Inset{
		Top:    unit.Dp(20),
		Bottom: unit.Dp(20),
	}
	
	return communityCardsArea.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:    layout.Horizontal,
			Spacing: layout.SpaceEvenly,
			Alignment: layout.Middle,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return ui.layoutCommunityCards(gtx)
			}),
		)
	})
}

// layoutCommunityCards lays out the community cards
func (ui *GameUI) layoutCommunityCards(gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceAround,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if len(ui.gameState.CommunityCards) > 0 {
				return ui.drawCard(gtx, ui.gameState.CommunityCards[0])
			}
			return ui.drawEmptyCard(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if len(ui.gameState.CommunityCards) > 1 {
				return ui.drawCard(gtx, ui.gameState.CommunityCards[1])
			}
			return ui.drawEmptyCard(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if len(ui.gameState.CommunityCards) > 2 {
				return ui.drawCard(gtx, ui.gameState.CommunityCards[2])
			}
			return ui.drawEmptyCard(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if len(ui.gameState.CommunityCards) > 3 {
				return ui.drawCard(gtx, ui.gameState.CommunityCards[3])
			}
			return ui.drawEmptyCard(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if len(ui.gameState.CommunityCards) > 4 {
				return ui.drawCard(gtx, ui.gameState.CommunityCards[4])
			}
			return ui.drawEmptyCard(gtx)
		}),
	)
}

// layoutControls lays out the player controls
func (ui *GameUI) layoutControls(gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceEvenly,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ui.layoutPlayerInfo(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ui.layoutActionButtons(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ui.layoutBetSlider(gtx)
		}),
	)
}

// layoutPlayerInfo lays out the player information
func (ui *GameUI) layoutPlayerInfo(gtx layout.Context) layout.Dimensions {
	player := ui.gameState.GetCurrentPlayer()
	
	return layout.Inset{
		Top:    unit.Dp(10),
		Bottom: unit.Dp(10),
		Left:   unit.Dp(10),
		Right:  unit.Dp(10),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:    layout.Vertical,
			Spacing: layout.SpaceEvenly,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				label := material.H6(ui.theme.Theme, "Player: "+player.Name)
				return label.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				label := material.Body1(ui.theme.Theme, "Chips: "+string(rune(player.Chips)))
				return label.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				label := material.Body1(ui.theme.Theme, "Current Bet: "+string(rune(player.Bet)))
				return label.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				label := material.Body1(ui.theme.Theme, "Pot: "+string(rune(ui.gameState.Pot)))
				return label.Layout(gtx)
			}),
		)
	})
}

// layoutActionButtons lays out the action buttons
func (ui *GameUI) layoutActionButtons(gtx layout.Context) layout.Dimensions {
	return layout.Inset{
		Top:    unit.Dp(10),
		Bottom: unit.Dp(10),
		Left:   unit.Dp(10),
		Right:  unit.Dp(10),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:    layout.Horizontal,
			Spacing: layout.SpaceEvenly,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(ui.theme.Theme, &ui.foldButton, "Fold")
				return btn.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(ui.theme.Theme, &ui.checkButton, "Check")
				return btn.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(ui.theme.Theme, &ui.callButton, "Call")
				return btn.Layout(gtx)
			}),
		)
	})
}

// layoutBetSlider lays out the bet slider and bet/raise buttons
func (ui *GameUI) layoutBetSlider(gtx layout.Context) layout.Dimensions {
	return layout.Inset{
		Top:    unit.Dp(10),
		Bottom: unit.Dp(10),
		Left:   unit.Dp(10),
		Right:  unit.Dp(10),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:    layout.Vertical,
			Spacing: layout.SpaceEvenly,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				slider := material.Slider(ui.theme.Theme, &ui.betSlider, 0, 1)
				return slider.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:    layout.Horizontal,
					Spacing: layout.SpaceEvenly,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(ui.theme.Theme, &ui.betButton, "Bet")
						return btn.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(ui.theme.Theme, &ui.raiseButton, "Raise")
						return btn.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(ui.theme.Theme, &ui.allInButton, "All In")
						return btn.Layout(gtx)
					}),
				)
			}),
		)
	})
}

// drawCard draws a card
func (ui *GameUI) drawCard(gtx layout.Context, card game.Card) layout.Dimensions {
	size := image.Point{X: 80, Y: 120}
	r := clip.Rect{Max: size}.Op()
	
	// Draw card background
	paint.FillShape(gtx.Ops, ui.theme.CardFront, r)
	
	// Draw card border
	borderColor := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	borderWidth := float32(2)
	drawBorder(gtx, size, borderColor, borderWidth)
	
	// Draw card text
	cardText := card.String()
	textColor := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	if card.Suit == game.Hearts || card.Suit == game.Diamonds {
		textColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	}
	
	op.Offset(image.Pt(10, 30)).Add(gtx.Ops)
	paint.ColorOp{Color: textColor}.Add(gtx.Ops)
	widget.Label{}.Layout(gtx, ui.theme.Theme.Shaper, ui.theme.Theme.Face, unit.Sp(20), cardText)
	
	return layout.Dimensions{Size: size}
}

// drawEmptyCard draws an empty card placeholder
func (ui *GameUI) drawEmptyCard(gtx layout.Context) layout.Dimensions {
	size := image.Point{X: 80, Y: 120}
	r := clip.Rect{Max: size}.Op()
	
	// Draw card background with low opacity
	emptyColor := color.NRGBA{R: 200, G: 200, B: 200, A: 100}
	paint.FillShape(gtx.Ops, emptyColor, r)
	
	return layout.Dimensions{Size: size}
}

// handleButtonClicks handles button clicks
func (ui *GameUI) handleButtonClicks() {
	if ui.foldButton.Clicked() {
		ui.gameState.ProcessAction(game.Fold, 0)
	}
	
	if ui.checkButton.Clicked() {
		ui.gameState.ProcessAction(game.Check, 0)
	}
	
	if ui.callButton.Clicked() {
		ui.gameState.ProcessAction(game.Call, 0)
	}
	
	if ui.betButton.Clicked() {
		ui.gameState.ProcessAction(game.Bet, ui.betAmount)
	}
	
	if ui.raiseButton.Clicked() {
		ui.gameState.ProcessAction(game.Raise, ui.betAmount)
	}
	
	if ui.allInButton.Clicked() {
		ui.gameState.ProcessAction(game.AllIn, 0)
	}
	
	// Update bet amount based on slider
	player := ui.gameState.GetCurrentPlayer()
	ui.betAmount = int(ui.betSlider.Value * float32(player.Chips))
	if ui.betAmount < ui.gameState.BigBlind {
		ui.betAmount = ui.gameState.BigBlind
	}
}

// Helper functions

// paintRect paints a rectangle
func paintRect(gtx layout.Context, size image.Point, col color.NRGBA) {
	r := clip.Rect{Max: size}.Op()
	paint.FillShape(gtx.Ops, col, r)
}

// drawBorder draws a border around a rectangle
func drawBorder(gtx layout.Context, size image.Point, col color.NRGBA, width float32) {
	// Top border
	topRect := clip.Rect{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: size.X, Y: int(width)},
	}.Op()
	paint.FillShape(gtx.Ops, col, topRect)
	
	// Bottom border
	bottomRect := clip.Rect{
		Min: image.Point{X: 0, Y: size.Y - int(width)},
		Max: image.Point{X: size.X, Y: size.Y},
	}.Op()
	paint.FillShape(gtx.Ops, col, bottomRect)
	
	// Left border
	leftRect := clip.Rect{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: int(width), Y: size.Y},
	}.Op()
	paint.FillShape(gtx.Ops, col, leftRect)
	
	// Right border
	rightRect := clip.Rect{
		Min: image.Point{X: size.X - int(width), Y: 0},
		Max: image.Point{X: size.X, Y: size.Y},
	}.Op()
	paint.FillShape(gtx.Ops, col, rightRect)
}
