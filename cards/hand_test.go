package cards

import "fmt"

func ExampleHand_String() {
	h1 := NewHand(
		PlayingCard{"Hearts", "Seven"},
		PlayingCard{"Spades", "Ace"},
		PlayingCard{"Spades", "King"},
	)

	fmt.Println(h1)

	// Output:
	// Seven of Hearts
	// Ace of Spades
	// King of Spades
}

func ExampleHand_Length() {
	h1 := NewHand(
		PlayingCard{"Spades", "Ace"},
		PlayingCard{"Spades", "Queen"},
		PlayingCard{"Hearts", "Queen"},
		PlayingCard{"Hearts", "Seven"},
	)

	fmt.Println(h1.Length())

	// Output:
	// 4
}

func ExampleHand_Get() {
	h1 := NewHand(
		PlayingCard{"Hearts", "Seven"},
		PlayingCard{"Spades", "Ace"},
		PlayingCard{"Spades", "King"},
	)

	fmt.Println(h1.Get(0))
	fmt.Println(h1.Get(1))
	fmt.Println(h1.Get(2))

	// Output:
	// Seven of Hearts
	// Ace of Spades
	// King of Spades
}

func ExampleHand_Add() {
	h1 := NewHand()
	h1.Add(
		PlayingCard{"Hearts", "Seven"},
		PlayingCard{"Spades", "Ace"},
		PlayingCard{"Spades", "King"},
	)

	fmt.Println(h1.Get(0))
	fmt.Println(h1.Get(1))
	fmt.Println(h1.Get(2))

	// Output:
	// Seven of Hearts
	// Ace of Spades
	// King of Spades
}

func ExampleHand_Remove() {
	h1 := NewHand(
		PlayingCard{"Hearts", "Seven"},
		PlayingCard{"Spades", "Ace"},
		PlayingCard{"Spades", "King"},
		PlayingCard{"Clubs", "Eight"},
		PlayingCard{"Diamonds", "Jack"},
	)

	h1.Remove(0)
	fmt.Println(h1)

	h1.Remove(2)
	fmt.Println(h1)

	h1.Remove(2)
	fmt.Println(h1)

	// Output:
	// Ace of Spades
	// King of Spades
	// Eight of Clubs
	// Jack of Diamonds
	//
	// Ace of Spades
	// King of Spades
	// Jack of Diamonds
	//
	// Ace of Spades
	// King of Spades
}
