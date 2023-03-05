package main

import (
	t "Poker_Symulator/src/types"
	"fmt"
)

func main() {
	p := t.NewPartie()

	s := t.Spieler{Name: "Niklas", Guthaben: 50.0, Status: true}
	p.SpielerHinzufügen(s)
	//fmt.Println(*s1)
	p.PottErhoehen(100)
	s = p.PottAuszahlen(s)
	fmt.Println(p)
	fmt.Println(s)
}
