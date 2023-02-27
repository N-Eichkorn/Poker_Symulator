package types

import (
	"math/rand"
)

type Partie struct {
	SpielerListe map[string]Spieler
	Karten       []Karte
	BigBlind     float64
	SmalBlind    float64
}

func NewPartie() *Partie {
	p := new(Partie)
	p.Karten = []Karte{
		{Färbung: Herz, Wertigkeit: Zwei},
		{Färbung: Herz, Wertigkeit: Drei},
		{Färbung: Herz, Wertigkeit: Vier},
		{Färbung: Herz, Wertigkeit: Fünf},
		{Färbung: Herz, Wertigkeit: Sechs},
		{Färbung: Herz, Wertigkeit: Sieben},
		{Färbung: Herz, Wertigkeit: Acht},
		{Färbung: Herz, Wertigkeit: Neun},
		{Färbung: Herz, Wertigkeit: Zehn},
		{Färbung: Herz, Wertigkeit: Bube},
		{Färbung: Herz, Wertigkeit: Dame},
		{Färbung: Herz, Wertigkeit: König},
		{Färbung: Herz, Wertigkeit: Ass},
		{Färbung: Karo, Wertigkeit: Zwei},
		{Färbung: Karo, Wertigkeit: Drei},
		{Färbung: Karo, Wertigkeit: Vier},
		{Färbung: Karo, Wertigkeit: Fünf},
		{Färbung: Karo, Wertigkeit: Sechs},
		{Färbung: Karo, Wertigkeit: Sieben},
		{Färbung: Karo, Wertigkeit: Acht},
		{Färbung: Karo, Wertigkeit: Neun},
		{Färbung: Karo, Wertigkeit: Zehn},
		{Färbung: Karo, Wertigkeit: Bube},
		{Färbung: Karo, Wertigkeit: Dame},
		{Färbung: Karo, Wertigkeit: König},
		{Färbung: Karo, Wertigkeit: Ass},
		{Färbung: Kreuz, Wertigkeit: Zwei},
		{Färbung: Kreuz, Wertigkeit: Drei},
		{Färbung: Kreuz, Wertigkeit: Vier},
		{Färbung: Kreuz, Wertigkeit: Fünf},
		{Färbung: Kreuz, Wertigkeit: Sechs},
		{Färbung: Kreuz, Wertigkeit: Sieben},
		{Färbung: Kreuz, Wertigkeit: Acht},
		{Färbung: Kreuz, Wertigkeit: Neun},
		{Färbung: Kreuz, Wertigkeit: Zehn},
		{Färbung: Kreuz, Wertigkeit: Bube},
		{Färbung: Kreuz, Wertigkeit: Dame},
		{Färbung: Kreuz, Wertigkeit: König},
		{Färbung: Kreuz, Wertigkeit: Ass},
		{Färbung: Pik, Wertigkeit: Zwei},
		{Färbung: Pik, Wertigkeit: Drei},
		{Färbung: Pik, Wertigkeit: Vier},
		{Färbung: Pik, Wertigkeit: Fünf},
		{Färbung: Pik, Wertigkeit: Sechs},
		{Färbung: Pik, Wertigkeit: Sieben},
		{Färbung: Pik, Wertigkeit: Acht},
		{Färbung: Pik, Wertigkeit: Neun},
		{Färbung: Pik, Wertigkeit: Zehn},
		{Färbung: Pik, Wertigkeit: Bube},
		{Färbung: Pik, Wertigkeit: Dame},
		{Färbung: Pik, Wertigkeit: König},
		{Färbung: Pik, Wertigkeit: Ass},
	}
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

func (p *Partie) SpielerHinzufügen(name string, guthaben float64) {
	p.SpielerListe[name] = Spieler{Name: name, Guthaben: guthaben, Status: true}
}
