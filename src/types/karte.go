package types

type Farbe string

const (
	Herz  Farbe = "Herz"
	Pik   Farbe = "Pik"
	Kreuz Farbe = "Kreuz"
	Karo  Farbe = "Karo"
)

type Wert int8

const (
	Zwei   Wert = 2
	Drei   Wert = 3
	Vier   Wert = 4
	Fünf   Wert = 5
	Sechs  Wert = 6
	Sieben Wert = 7
	Acht   Wert = 8
	Neun   Wert = 9
	Zehn   Wert = 10
	Bube   Wert = 11
	Dame   Wert = 12
	König  Wert = 13
	Ass    Wert = 14
)

type Karte struct {
	Färbung    Farbe
	Wertigkeit Wert
}

type Kartendeck struct {
}
