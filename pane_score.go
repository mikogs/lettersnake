package main

import (
	"context"
	"fmt"

	"gopkg.pl/mikogs/lettersnake/pkg/lettersnake"
	"gopkg.pl/mikogs/termui"
)

type scorePane struct {
	g *lettersnake.Game
}

func (w *scorePane) Render(pane *termui.Pane) {
}

func (w scorePane) Iterate(pane *termui.Pane) {
	pane.Write(1, 0, fmt.Sprintf("Correct: %d/%d", w.g.NumCorrectWords(), w.g.NumUsedWords()))
}

func (w scorePane) HasBackend() bool {
	return false
}

func (w *scorePane) Backend(ctx context.Context) {
}
