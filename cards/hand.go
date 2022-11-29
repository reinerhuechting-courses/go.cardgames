package cards

// Datentyp für die Hand eines Spielers.
// Eine Hand ist genau wie ein Deck aufgebaut,
// daher definieren wir den Typ Hand auch als Deck.
//
// Anmerkung: Anders als man von manchen anderen Sprachen erwarten würde,
// übernimmt der Datentyp Hand durch diese Definition *nicht* die Methoden von Deck.
// Dafür müsste man ihn als Mix-In oder als Typ-Alias (nicht empfohlen) definieren.
type Hand Deck

// Liefert eine menschenlesbare String-Repräsentation einer Hand.
// Verwendet dafür die String()-Methode von Deck.
func (hand Hand) String() string {
	// Wir können die String()-Methode von Deck benutzen, indem wir eine Hand
	// explizit in ein Deck konvertieren und dann darauf String() aufrufen.
	return Deck(hand).String()
}

// Liefert die Anzahl der Karten der Hand.
func (hand Hand) Length() int {
	return Deck(hand).Length()
}

// Getter für eine bestimmte Karte.
func (hand Hand) Get(n int) PlayingCard {
	return Deck(hand).Get(n)
}

// Fügt die gegebenen Spielkarten zum Deck hinzu (ans Ende)
func (hand *Hand) Add(cards ...PlayingCard) {
	(*Deck)(hand).Add(cards...)
}

// Entfernt die Karte mit der angegebenen Position aus der Hand.
func (hand *Hand) Remove(n int) {
	beforeN := hand.cards[:n]
	afterN := hand.cards[n+1:]
	hand.cards = append(beforeN, afterN...)
}

// Liefert eine Hand, die die angegebenen Spielkarten enthält.
func NewHand(cards ...PlayingCard) Hand {
	return Hand{cards}
}
