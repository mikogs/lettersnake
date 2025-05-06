package main

import (
	"context"
	"strings"

	"gopkg.pl/mikogs/lettersnake/pkg/lettersnake"
	"gopkg.pl/mikogs/termui"
)

type rightWordPane struct {
	g *lettersnake.Game
}

func (w *rightWordPane) Render(pane *termui.Pane) {
}

func (w rightWordPane) Iterate(pane *termui.Pane) {
	trim := 20 - len(w.g.ConsumedLetters())
	pane.Write(1, 0, w.g.ConsumedLetters()+strings.Repeat(" ", trim))
}

func (w rightWordPane) HasBackend() bool {
	return false
}

func (w *rightWordPane) Backend(ctx context.Context) {
}
