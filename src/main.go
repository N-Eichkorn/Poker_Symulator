package main

import (
	t "Poker_Symulator/src/types"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("hallo welt")
	fmt.Println(t.Karte{Färbung: t.Herz, Wertigkeit: t.Zwei})

	//p := t.NewGame()

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
}
