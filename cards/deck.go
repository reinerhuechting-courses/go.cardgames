package cards

import "math/rand"

// Datentyp für einen Kartenstapel.
// Ein Kartenstapel ist nichts anderes als eine Liste von Karten.
// Wir definieren ihn als Struct, das eine Liste von Karten enthält.
// Alternativ könnte man das Deck direkt als Liste von Karten definieren,
// dann wären aber manche Methoden schwerer lesbar.
type Deck struct {
	cards []PlayingCard
}

// Liefert eine menschenlesbare String-Repräsentation eines Kartenstapels.
// Verwendet dafür die String()-Methode von PlayingCard.
func (deck Deck) String() string {
	result := ""
	for _, card := range deck.cards {
		result += card.String() + "\n"
	}
	return result
}

// Liefert die Anzahl der Karten des Stapels.
func (deck Deck) Length() int {
	return len(deck.cards)
}

// Getter für eine bestimmte Karte.
func (deck Deck) Get(n int) PlayingCard {
	return deck.cards[n]
}

// Fügt die gegebenen Spielkarten zum Deck hinzu (ans Ende)
func (deck *Deck) Add(cards ...PlayingCard) {
	deck.cards = append(deck.cards, cards...)
}

// Hilfsfunktion für's Ziehen von Karten:
// Liefert die letzte Karte auf dem Stapel.
// Diese Funktion ist nicht sicher, wenn der Stapel leer ist.
func (deck Deck) GetLast() PlayingCard {
	return deck.Get(deck.Length() - 1)
}

// Hilfsfunktion für's Ziehen von Karten:
// Entfernt die letzte Karte vom Stapel.
// Diese Funktion ist nicht sicher, wenn der Stapel leer ist.
//
// Anmerkung: Der Receiver dieser Methode ist ein Pointer.
// D.h. die Methode bekommt nur die Adresse eines Kartenstapels,
// nicht den Stapel selbst. Dies ist notwendig, da der Stapel
// von der Methode verändert wird.
func (deck *Deck) RemoveLast() {
	// Wir weisen dem Deck die Slice zu, die das letzte Element nicht mehr enthält.
	// Die Stern-Ausdrücke "*deck" dereferenzieren den Pointer.
	// D.h. hier sagen wir dem Compiler, dass er bitte das eigentliche Deck
	// verwenden soll, statt des Pointers.
	deck.cards = deck.cards[:deck.Length()-1]
}

// Zieht eine Karte vom Stapel.
// D.h. entfernt eine Karte und liefert diese.
// Diese Funktion ist nicht sicher, wenn der Stapel leer ist.
func (deck *Deck) DrawCard() PlayingCard {
	result := deck.GetLast()
	deck.RemoveLast()
	return result
}

// Mischt einen Kartenstapel.
func (deck Deck) Shuffle() {
	// Wir benutzen die Funktion `Shuffle()`` aus dem Package math/rand.
	// Diese kann beliebige Listen durcheinander bringen.
	// Dazu benötigt sie die Länge der Liste sowie eine Funktion, die zwei Elemente
	// erwartet und diese vertauscht. Wir müssen also eine solche Funktion definieren
	// und dann an `Shuffle()` übergeben.

	// Definiere eine Funktion, die zwei beliebige Karten aus dem Deck vertauscht.
	swapCards := func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	}

	// Verwende rand.Shuffle() mit unserer Tausch-Funktion, um den Stapel zu mischen.
	rand.Shuffle(deck.Length(), swapCards)
}

// Erwartet eine Hand und eine Zahl n.
// Zieht n Karten vom Deck in die Hand.
func (deck *Deck) Deal(hand *Hand, n int) {
	for i := 0; i < n; i++ {
		hand.Add(deck.DrawCard())
	}
}

// Liefert einen Kartenstapel, der die angegebenen Spielkarten enthält.
func NewDeck(cards ...PlayingCard) Deck {
	return Deck{cards}
}

// Liefert einen "frischen" sortierten (Skat-)Kartenstapel mit 32 Karten.
func NewDeck32() Deck {
	suits := []string{"Clubs", "Spades", "Hearts", "Diamonds"}
	ranks := []string{"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	deck := Deck{}
	for _, c := range suits {
		for _, v := range ranks {
			deck.Add(PlayingCard{c, v})
		}
	}
	return deck
}

/* TODO:
 * - Suchfunktion für bestimmte Karten auf der Hand des Spielers.
 *   - Nach Farbe, Wert oder beidem.
 * - Funktion zum Ausspielen einer Karte von der Hand.
 * - Erste einfache Spiellogik.
 */
