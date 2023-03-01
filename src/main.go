package main

import (
	t "Poker_Symulator/src/types"
	"fmt"
)

func main() {
	p := t.NewPartie()

	s := t.Spieler{Name: "Niklas", Guthaben: 50.0, Status: true}
	s1 := p.SpielerHinzufügen(s)
	fmt.Println(*s1)
	fmt.Println(p.GetSpielerkarte(s1.Name))
	fmt.Println(p.GemeinschaftskarteAusteilen())
	fmt.Println(p)

}
