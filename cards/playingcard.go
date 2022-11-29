package cards

import "fmt"

// Datentyp für Spielkarten.
// Enthält Felder für Farbe ("Suit") und Wert ("Rank").
// Beides sind Strings, weil Farben Begriffe wie "Kreuz", "Pik" etc. sind und
// weil Werte zwar Zahlen wie "7", "8", etc., aber auch Namen wie "Bube" sind.
type PlayingCard struct {
	Suit string
	Rank string
}

// Liefert eine String-Repräsentation einer Spielkarte.
// Durch die Existenz dieser Funktion wird fmt.Println() anstelle der generischen
// Repräsentation eines Structs mit geschweiften Klammern diese Funktion verwenden.
func (card PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", card.Rank, card.Suit)
}
