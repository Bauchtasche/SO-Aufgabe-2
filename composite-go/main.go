package main

import "fmt"

func main() {
	// kleines Fenster unten rechts: eigener Rahmen + 2x2 Glas
	klein := NewWindow("unten rechts")
	klein.Add(
		&Frame{Name: "kleiner Rahmen", Cost: 40},
		&Glass{Name: "oben links", Cost: 15},
		&Glass{Name: "oben rechts", Cost: 15},
		&Glass{Name: "unten links", Cost: 15},
		&Glass{Name: "unten rechts", Cost: 15},
	)

	// großes Fenster: Rahmen + 3 Scheiben + kleines Fenster
	fenster := NewWindow("Hauptfenster")
	fenster.Add(
		&Frame{Name: "Hauptrahmen", Cost: 120},
		&Glass{Name: "oben links", Cost: 50},
		&Glass{Name: "oben rechts", Cost: 50},
		&Glass{Name: "unten links", Cost: 50},
		klein,
	)

	fenster.Display(0)
	fmt.Printf("\nGesamtpreis: %.2f €\n", fenster.Price())
}
