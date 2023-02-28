package main

import (
	t "Poker_Symulator/src/types"
	"fmt"
)

func main() {
	p := t.NewPartie()

	s := t.Spieler{Name: "Niklas", Guthaben: 50.0, Status: true}
	p.SpielerHinzufügen(s)
	//k1, k2 := p.GetSpielerkarte(s.Name)

	//fmt.Println(k1)
	//fmt.Println(k2)
	fmt.Println(p)
}
