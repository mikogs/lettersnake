package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	tui "github.com/mikolajgs/terminal-ui"
)

const DOWN = 0
const UP = 1
const LEFT = 2
const RIGHT = 3

type gameInterface struct {
	t           *tui.TUI
	game        *tui.TUIPane
	top         *tui.TUIPane
	leftBottom  *tui.TUIPane
	rightBottom *tui.TUIPane
	g           *game
	sizeSet     bool
}

func newGameInterface(g *game) *gameInterface {
	gi := &gameInterface{}

	gi.g = g
	gi.t = tui.NewTUI()
	p := gi.t.GetPane()

	pTop, pGameBottom := p.SplitHorizontally(-3, tui.UNIT_CHAR)
	pGame, pBottom := pGameBottom.SplitHorizontally(3, tui.UNIT_CHAR)

	pBottomLeft, pBottomRight := pBottom.SplitVertically(50, tui.UNIT_PERCENT)

	gi.game = pGame
	gi.top = pTop
	gi.leftBottom = pBottomLeft
	gi.rightBottom = pBottomRight

	gi.initStyle()
	gi.initIteration()

	gi.initKeyboard()

	return gi
}

func (gi *gameInterface) initStyle() {
	s := tui.NewTUIPaneStyleFrame()
	gi.game.SetStyle(s)
	gi.top.SetStyle(s)
	gi.leftBottom.SetStyle(s)
	gi.rightBottom.SetStyle(s)
}

func (gi *gameInterface) initIteration() {
	f := func(p *tui.TUIPane) int {
		if !gi.sizeSet {
			gi.g.setSize(gi.game.GetWidth()-2, gi.game.GetHeight()-2)
			gi.sizeSet = true
		}

		if !gi.g.isStarted() {
			p.Write(2, 1, "Press the S key to start the game", false)
			return NOT_STARTED
		}

		r := gi.g.iterate()
		if r == GAME_OVER {
			p.Write(2, 0, "** Game over! **", false)
			return r
		}

		if r == CONTINUE_GAME {
			for i := 0; i < len(gi.g.letters); i++ {
				gi.game.Write(gi.g.letters[i].x, gi.g.letters[i].y, gi.wrapInRandomColour(gi.g.letters[i].l), false)
			}
			gi.drawSnake()
		}

		return r
	}

	gi.game.SetOnDraw(f)
	gi.game.SetOnIterate(f)

	gi.leftBottom.SetOnIterate(func(p *tui.TUIPane) int {
		trim := 20 - len(gi.g.currentTranslation)
		p.Write(1, 0, gi.g.currentTranslation+strings.Repeat(" ", trim), false)
		return 0
	})

	gi.rightBottom.SetOnIterate(func(p *tui.TUIPane) int {
		trim := 20 - len(gi.g.consumedLetters)
		p.Write(1, 0, gi.g.consumedLetters+strings.Repeat(" ", trim), false)
		return 0
	})

	gi.top.SetOnIterate(func(p *tui.TUIPane) int {
		p.Write(1, 0, fmt.Sprintf("Correct: %d/%d", gi.g.wordsCorrect, gi.g.wordsGiven), false)
		return 0
	})
}

func (gi *gameInterface) drawSnake() {
	for i := 0; i < len(gi.g.snake); i++ {
		gi.game.Write(gi.g.snake[i].x, gi.g.snake[i].y, gi.getSnakeSegment(i), false)
	}
	if gi.g.remove != nil {
		gi.game.Write(gi.g.remove.x, gi.g.remove.y, " ", false)
	}
}

func (gi *gameInterface) setSpeed(i int) {
	gi.t.SetLoopSleep(i)
}

func (gi *gameInterface) initKeyboard() {
	gi.t.SetOnKeyPress(func(t *tui.TUI, b []byte) {
		if string(b) == "x" {
			t.Exit(0)
		}
		if string(b) == "s" {
			if !gi.g.isStarted() {
				gi.clearPane(gi.game)

				gi.g.startGame()
			}
			return
		}
		// TODO: Keys should be handled differently, maybe in raw mode
		// left arrow pressed
		if string(b) == "D" {
			if gi.g.direction != RIGHT {
				gi.g.direction = LEFT
			}
			return
		}
		// right arrow pressed
		if string(b) == "C" {
			if gi.g.direction != LEFT {
				gi.g.direction = RIGHT
			}
			return
		}
		// down arrow pressed
		if string(b) == "B" {
			if gi.g.direction != UP {
				gi.g.direction = DOWN
			}
			return
		}
		// up arrow pressed
		if string(b) == "A" {
			if gi.g.direction != DOWN {
				gi.g.direction = UP
			}
			return
		}
	})
}

func (gi *gameInterface) clearPane(p *tui.TUIPane) {
	for y := 0; y < p.GetHeight()-2; y++ {
		gi.clearPaneLine(p, y)
	}
}

func (gi *gameInterface) clearPaneLine(p *tui.TUIPane, y int) {
	p.Write(0, y, strings.Repeat(" ", p.GetWidth()-2), false)
}

func (gi *gameInterface) run() {
	gi.t.Run(os.Stdout, os.Stderr)
}

func (gi *gameInterface) wrapInRandomColour(s string) string {
	colours := []string{"\033[1;93m", "\033[1;92m", "\033[1;95m", "\033[1;96m"}
	reset := "\033[0m"
	return colours[rand.Intn(len(colours))] + s + reset
}

func (gi *gameInterface) getSnakeSegment(i int) string {
	// 125-159
	c := 125 + i
	s := "▓"
	if i > 0 {
		s = "▒"
	}
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", c, s)
}
