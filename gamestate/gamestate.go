package gamestate

import "github.com/reinerhuechting-courses/go.cardgames/cards"

type GameState struct {
	Players       []cards.Hand
	Deck          cards.Deck
	UsedDeck      cards.Deck
	CurrentPlayer int
}
