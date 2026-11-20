# Composite Pattern in Go – Fenster-Beispiel

Umsetzung des Composite-Beispiels aus der Vorlesung (Item / Window / Glass / Frame) in Go.
Der verwendete Prompt steht in [PROMPT.md](PROMPT.md).

## Struktur

| Rolle      | Typ      |
|------------|----------|
| Component  | `Item` (Interface: `Price()`, `Display()`) |
| Leaf       | `Glass`, `Frame` |
| Composite  | `Window` (enthält 1..* `Item`) |

## Ausführen

```
go run .
```

## Ausgabe

```
Fenster: Hauptfenster (gesamt 370.00 €)
  Rahmen: Hauptrahmen (120.00 €)
  Glas: oben links (50.00 €)
  Glas: oben rechts (50.00 €)
  Glas: unten links (50.00 €)
  Fenster: unten rechts (gesamt 100.00 €)
    Rahmen: kleiner Rahmen (40.00 €)
    Glas: oben links (15.00 €)
    Glas: oben rechts (15.00 €)
    Glas: unten links (15.00 €)
    Glas: unten rechts (15.00 €)

Gesamtpreis: 370.00 €
```
