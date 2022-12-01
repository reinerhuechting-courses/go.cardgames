package maumau

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/reinerhuechting-courses/go.cardgames/cards"
	"github.com/reinerhuechting-courses/go.cardgames/gamestate"
)

// Spielvorbereitung:
// Erzeugt Spieler und Kartenstapel, mischt den Stapel und deckt die erste Karte auf.
func NewGame(numberOfPlayers int) gamestate.GameState {
	// Neuen GameState mit der gewünschten Spielerzahl,
	// einem 32-Karten-Deck und einem leeren Ablagestapel erzeugen.
	state := gamestate.GameState{
		Players:       make([]cards.Hand, numberOfPlayers),
		Deck:          cards.NewDeck32(),
		UsedDeck:      cards.NewDeck(),
		CurrentPlayer: 0,
	}

	// Deck mischen und Karten verteilen.
	state.Deck.Shuffle()
	for i := range state.Players {
		state.Deck.Deal(&state.Players[i], 6)
	}

	// Eine erste Karte ziehen und auf den Ablagestapel legen.
	state.UsedDeck.Add(state.Deck.DrawCard())
	return state
}

func RunGame() {
	// Zufallsgenerator initialisieren, damit nicht immer die gleichen Karten kommen.
	rand.Seed(int64(time.Now().Second()))

	// Ein Zwei-Spieler-Spiel vorbereiten.
	state := NewGame(2)

	// Oberste Karte auf dem Ablagestapel anzeigen.
	fmt.Println("Oberste Karte des Stapels:")
	fmt.Println(state.UsedDeck.GetLast())
	fmt.Println()

	// Aktuellem Spieler seine Karten zeigen.
	fmt.Printf("Spieler %d, dies sind deine Karten:\n", state.CurrentPlayer)
	fmt.Println(state.Players[state.CurrentPlayer])

	// Spieler nach seiner Wahl fragen.
	fmt.Printf("Spieler %d, welche Karte möchtest du spielen?\n", state.CurrentPlayer)
	var selectedCardNumber int
	fmt.Scanln(&selectedCardNumber)

	// Gewählte Karte ausspielen.
	card := state.Players[state.CurrentPlayer].Get(selectedCardNumber)
	state.Players[state.CurrentPlayer].Remove(selectedCardNumber)
	state.UsedDeck.Add(card)

	// Oberste Karte auf dem Ablagestapel anzeigen.
	fmt.Println("Oberste Karte des Stapels:")
	fmt.Println(state.UsedDeck.GetLast())
	fmt.Println()
}
