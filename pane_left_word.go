package main

import (
	"context"
	"strings"

	"gopkg.pl/mikogs/lettersnake/pkg/lettersnake"
	"gopkg.pl/mikogs/termui"
)

type leftWordPane struct {
	g *lettersnake.Game
}

func (w *leftWordPane) Render(pane *termui.Pane) {
}

func (w leftWordPane) Iterate(pane *termui.Pane) {
	trim := 20 - len(w.g.CurrentTranslation())
	pane.Write(1, 0, w.g.CurrentTranslation()+strings.Repeat(" ", trim))
}

func (w leftWordPane) HasBackend() bool {
	return false
}

func (w *leftWordPane) Backend(ctx context.Context) {
}
