package types

type farbe string

const (
	Herz  farbe = "Herz"
	Pik   farbe = "Pik"
	Kreuz farbe = "Kreuz"
	Karo  farbe = "Karo"
)

type wert int8

const (
	Zwei   wert = 2
	Drei   wert = 3
	Vier   wert = 4
	Fünf   wert = 5
	Sechs  wert = 6
	Sieben wert = 7
	Acht   wert = 8
	Neun   wert = 9
	Zehn   wert = 10
	Bube   wert = 11
	Dame   wert = 12
	König  wert = 13
	Ass    wert = 14
)

const (
	Deck  string = "Deck"
	Tisch string = "Tisch"
)

type Karte struct {
	Färbung       farbe
	Wertigkeit    wert
	Zugehörigkeit string // Zugehörigkeit ist der Spielername. Wem gehört diese Karte gerade.
	// "Deck" ist Niemand und die Karte befindet sich im Deck.
	// "Tisch" ist der Tisch und die Karte befindet sich in den Gemeinschaftskarten.
	// Alles andere ist der Name des Spielers, der die Karte besitzt.
}
