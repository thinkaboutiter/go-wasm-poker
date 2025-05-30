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

// PlayerUI represents the UI for a player
type PlayerUI struct {
	player      *game.Player
	position    image.Point
	isActive    bool
	theme       *Theme
}

// NewPlayerUI creates a new player UI
func NewPlayerUI(player *game.Player, position image.Point, theme *Theme) *PlayerUI {
	return &PlayerUI{
		player:   player,
		position: position,
		theme:    theme,
	}
}

// SetActive sets whether the player is active
func (p *PlayerUI) SetActive(active bool) {
	p.isActive = active
}

// Layout lays out the player UI
func (p *PlayerUI) Layout(gtx layout.Context) layout.Dimensions {
	// Save operations
	defer op.Save(gtx.Ops).Load()
	
	// Translate to player position
	op.Offset(layout.FPt(p.position)).Add(gtx.Ops)
	
	// Player background color
	bgColor := p.theme.PlayerColor
	if p.isActive {
		bgColor = p.theme.ActivePlayer
	}
	
	// Player area size
	size := image.Point{X: 150, Y: 100}
	
	// Draw player background
	r := clip.Rect{Max: size}.Op()
	paint.FillShape(gtx.Ops, bgColor, r)
	
	// Layout player info
	layout.Inset{
		Top:    unit.Dp(5),
		Bottom: unit.Dp(5),
		Left:   unit.Dp(5),
		Right:  unit.Dp(5),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:    layout.Vertical,
			Spacing: layout.SpaceStart,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				label := material.Body1(p.theme.Theme, p.player.Name)
				label.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
				return label.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				label := material.Body2(p.theme.Theme, "Chips: "+string(rune(p.player.Chips)))
				label.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
				return label.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				label := material.Body2(p.theme.Theme, "Bet: "+string(rune(p.player.Bet)))
				label.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
				return label.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return p.layoutPlayerCards(gtx)
			}),
		)
	})
	
	return layout.Dimensions{Size: size}
}

// layoutPlayerCards lays out the player's cards
func (p *PlayerUI) layoutPlayerCards(gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceStart,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if len(p.player.Cards) > 0 {
				return drawSmallCard(gtx, p.player.Cards[0], p.theme)
			}
			return drawSmallEmptyCard(gtx, p.theme)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if len(p.player.Cards) > 1 {
				return drawSmallCard(gtx, p.player.Cards[1], p.theme)
			}
			return drawSmallEmptyCard(gtx, p.theme)
		}),
	)
}

// drawSmallCard draws a small card
func drawSmallCard(gtx layout.Context, card game.Card, theme *Theme) layout.Dimensions {
	size := image.Point{X: 40, Y: 60}
	r := clip.Rect{Max: size}.Op()
	
	// Draw card background
	paint.FillShape(gtx.Ops, theme.CardFront, r)
	
	// Draw card border
	borderColor := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	borderWidth := float32(1)
	drawBorder(gtx, size, borderColor, borderWidth)
	
	// Draw card text
	cardText := card.String()
	textColor := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	if card.Suit == game.Hearts || card.Suit == game.Diamonds {
		textColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	}
	
	op.Offset(image.Pt(5, 15)).Add(gtx.Ops)
	paint.ColorOp{Color: textColor}.Add(gtx.Ops)
	widget.Label{}.Layout(gtx, theme.Theme.Shaper, theme.Theme.Face, unit.Sp(12), cardText)
	
	return layout.Dimensions{Size: size}
}

// drawSmallEmptyCard draws a small empty card placeholder
func drawSmallEmptyCard(gtx layout.Context, theme *Theme) layout.Dimensions {
	size := image.Point{X: 40, Y: 60}
	r := clip.Rect{Max: size}.Op()
	
	// Draw card background with low opacity
	emptyColor := color.NRGBA{R: 200, G: 200, B: 200, A: 100}
	paint.FillShape(gtx.Ops, emptyColor, r)
	
	return layout.Dimensions{Size: size}
}
