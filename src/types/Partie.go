package types

import (
	"math/rand"
)

type Partie struct {
	SpielerListe map[string]Spieler
	Karten       [52]Karte
	BigBlind     float64
	SmalBlind    float64
}

func NewPartie() *Partie {
	p := new(Partie)
	p.Karten = [52]Karte{
		{Färbung: Herz, Wertigkeit: Zwei, Zugehörigkeit: "Deck"},
		{Färbung: Herz, Wertigkeit: Drei, Zugehörigkeit: "Deck"},
		{Färbung: Herz, Wertigkeit: Vier, Zugehörigkeit: "Deck"},
		{Färbung: Herz, Wertigkeit: Fünf, Zugehörigkeit: "Deck"},
		{Färbung: Herz, Wertigkeit: Sechs, Zugehörigkeit: "Deck"},
		{Färbung: Herz, Wertigkeit: Sieben, Zugehörigkeit: "Deck"},
		{Färbung: Herz, Wertigkeit: Acht, Zugehörigkeit: "Deck"},
		{Färbung: Herz, Wertigkeit: Neun, Zugehörigkeit: "Deck"},
		{Färbung: Herz, Wertigkeit: Zehn, Zugehörigkeit: "Deck"},
		{Färbung: Herz, Wertigkeit: Bube, Zugehörigkeit: "Deck"},
		{Färbung: Herz, Wertigkeit: Dame, Zugehörigkeit: "Deck"},
		{Färbung: Herz, Wertigkeit: König, Zugehörigkeit: "Deck"},
		{Färbung: Herz, Wertigkeit: Ass, Zugehörigkeit: "Deck"},
		{Färbung: Karo, Wertigkeit: Zwei, Zugehörigkeit: "Deck"},
		{Färbung: Karo, Wertigkeit: Drei, Zugehörigkeit: "Deck"},
		{Färbung: Karo, Wertigkeit: Vier, Zugehörigkeit: "Deck"},
		{Färbung: Karo, Wertigkeit: Fünf, Zugehörigkeit: "Deck"},
		{Färbung: Karo, Wertigkeit: Sechs, Zugehörigkeit: "Deck"},
		{Färbung: Karo, Wertigkeit: Sieben, Zugehörigkeit: "Deck"},
		{Färbung: Karo, Wertigkeit: Acht, Zugehörigkeit: "Deck"},
		{Färbung: Karo, Wertigkeit: Neun, Zugehörigkeit: "Deck"},
		{Färbung: Karo, Wertigkeit: Zehn, Zugehörigkeit: "Deck"},
		{Färbung: Karo, Wertigkeit: Bube, Zugehörigkeit: "Deck"},
		{Färbung: Karo, Wertigkeit: Dame, Zugehörigkeit: "Deck"},
		{Färbung: Karo, Wertigkeit: König, Zugehörigkeit: "Deck"},
		{Färbung: Karo, Wertigkeit: Ass, Zugehörigkeit: "Deck"},
		{Färbung: Kreuz, Wertigkeit: Zwei, Zugehörigkeit: "Deck"},
		{Färbung: Kreuz, Wertigkeit: Drei, Zugehörigkeit: "Deck"},
		{Färbung: Kreuz, Wertigkeit: Vier, Zugehörigkeit: "Deck"},
		{Färbung: Kreuz, Wertigkeit: Fünf, Zugehörigkeit: "Deck"},
		{Färbung: Kreuz, Wertigkeit: Sechs, Zugehörigkeit: "Deck"},
		{Färbung: Kreuz, Wertigkeit: Sieben, Zugehörigkeit: "Deck"},
		{Färbung: Kreuz, Wertigkeit: Acht, Zugehörigkeit: "Deck"},
		{Färbung: Kreuz, Wertigkeit: Neun, Zugehörigkeit: "Deck"},
		{Färbung: Kreuz, Wertigkeit: Zehn, Zugehörigkeit: "Deck"},
		{Färbung: Kreuz, Wertigkeit: Bube, Zugehörigkeit: "Deck"},
		{Färbung: Kreuz, Wertigkeit: Dame, Zugehörigkeit: "Deck"},
		{Färbung: Kreuz, Wertigkeit: König, Zugehörigkeit: "Deck"},
		{Färbung: Kreuz, Wertigkeit: Ass, Zugehörigkeit: "Deck"},
		{Färbung: Pik, Wertigkeit: Zwei, Zugehörigkeit: "Deck"},
		{Färbung: Pik, Wertigkeit: Drei, Zugehörigkeit: "Deck"},
		{Färbung: Pik, Wertigkeit: Vier, Zugehörigkeit: "Deck"},
		{Färbung: Pik, Wertigkeit: Fünf, Zugehörigkeit: "Deck"},
		{Färbung: Pik, Wertigkeit: Sechs, Zugehörigkeit: "Deck"},
		{Färbung: Pik, Wertigkeit: Sieben, Zugehörigkeit: "Deck"},
		{Färbung: Pik, Wertigkeit: Acht, Zugehörigkeit: "Deck"},
		{Färbung: Pik, Wertigkeit: Neun, Zugehörigkeit: "Deck"},
		{Färbung: Pik, Wertigkeit: Zehn, Zugehörigkeit: "Deck"},
		{Färbung: Pik, Wertigkeit: Bube, Zugehörigkeit: "Deck"},
		{Färbung: Pik, Wertigkeit: Dame, Zugehörigkeit: "Deck"},
		{Färbung: Pik, Wertigkeit: König, Zugehörigkeit: "Deck"},
		{Färbung: Pik, Wertigkeit: Ass, Zugehörigkeit: "Deck"},
	}
	p.SpielerListe = make(map[string]Spieler)
	return p
}

func (p *Partie) GetSpielerkarte(spielername string) (Karte, Karte) {

	var fehler bool = true
	var vergebeneKarten []int
	var random1 int
	var random2 int

	for fehler {
		fehler = false
		random1 = rand.Intn(len(p.Karten) - 1)
		random2 = rand.Intn(len(p.Karten) - 1)
		if random1 == random2 {
			fehler = true
		}
		for _, v := range vergebeneKarten {
			if random1 == v || random2 == v {
				fehler = true
			}
		}
	}
	karte1 := &p.Karten[random1]
	karte2 := &p.Karten[random2]

	karte1.Zugehörigkeit = p.SpielerListe[spielername].Name
	karte2.Zugehörigkeit = p.SpielerListe[spielername].Name

	k1 := *karte1
	k2 := *karte2

	return k1, k2
}

func (p *Partie) SpielerHinzufügen(s Spieler) *Spieler {
	name := s.Name
	p.SpielerListe[name] = s

	return &s
}

func (p *Partie) SpielerEntfernen(s Spieler) {
	delete(p.SpielerListe, s.Name)
}
