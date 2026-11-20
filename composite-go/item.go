package main

import (
	"fmt"
	"strings"
)

// Item ist die gemeinsame Schnittstelle (Component) für Blätter und Composites.
type Item interface {
	Price() float64
	Display(level int)
}

// ---------- Blätter ----------

// Glass ist eine einzelne Glasscheibe (Leaf).
type Glass struct {
	Name string
	Cost float64
}

func (g *Glass) Price() float64 { return g.Cost }

func (g *Glass) Display(level int) {
	fmt.Printf("%sGlas: %s (%.2f €)\n", strings.Repeat("  ", level), g.Name, g.Cost)
}

// Frame ist ein Rahmen (Leaf).
type Frame struct {
	Name string
	Cost float64
}

func (f *Frame) Price() float64 { return f.Cost }

func (f *Frame) Display(level int) {
	fmt.Printf("%sRahmen: %s (%.2f €)\n", strings.Repeat("  ", level), f.Name, f.Cost)
}

// ---------- Composite ----------

// Window besteht aus 1..* Items (Glas, Rahmen oder weiteren Fenstern).
type Window struct {
	Name  string
	items []Item
}

func NewWindow(name string) *Window {
	return &Window{Name: name}
}

func (w *Window) Add(items ...Item) {
	w.items = append(w.items, items...)
}

// Price summiert rekursiv die Preise aller enthaltenen Items.
func (w *Window) Price() float64 {
	sum := 0.0
	for _, it := range w.items {
		sum += it.Price()
	}
	return sum
}

func (w *Window) Display(level int) {
	fmt.Printf("%sFenster: %s (gesamt %.2f €)\n", strings.Repeat("  ", level), w.Name, w.Price())
	for _, it := range w.items {
		it.Display(level + 1)
	}
}
