package cards

import "fmt"

func ExamplePlayingCard_String() {
	c1 := PlayingCard{"Clubs", "Seven"}
	c2 := PlayingCard{"Spades", "Ace"}

	fmt.Println(c1.String())
	fmt.Println(c1)
	fmt.Println(c2.String())
	fmt.Println(c2)

	// Output:
	// Seven of Clubs
	// Seven of Clubs
	// Ace of Spades
	// Ace of Spades
}
