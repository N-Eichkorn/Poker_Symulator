package types

import "errors"

type Spieler struct {
	Name     string   // Name des Spielers
	Guthaben float64  // Kontostand des Spielers
	Karten   [2]Karte // Karten des Spielers
	Status   bool     // true = Spieler ist aktiver Spieler
	// false = Spieler ist ausgeschieden
}

func (s *Spieler) GuthabenVerringern(wert float64) error {
	if wert < s.Guthaben {
		return errors.New("nicht genug Guthaben!")
	}
	s.Guthaben = s.Guthaben - wert
	return nil
}

func (s *Spieler) GurhabenHinzufügen(wert float64) {
	s.Guthaben = s.Guthaben + wert
}

func (s *Spieler) KartenErhalten(k [2]Karte) {
	s.Karten = k
}
