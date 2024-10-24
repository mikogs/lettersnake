package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const NOT_STARTED = 0
const CONTINUE_GAME = 1
const GAME_OVER = 2

type letter struct {
	x int
	y int
	l string
}

type game struct {
	words              []string
	started            bool
	nextWordIndex      int
	currentWord        string
	currentTranslation string
	wordsGiven         int
	wordsCorrect       int
	letters            []letter
	size               [2]int
	direction          int
	snake              []segment
	remove             *segment
	consumedLetters    string
}

type segment struct {
	x int
	y int
}

func newGame() *game {
	return &game{
		words:     []string{},
		started:   false,
		direction: DOWN,
		snake: []segment{
			{x: 3, y: 5},
			{x: 3, y: 4},
			{x: 3, y: 3},
			{x: 3, y: 2},
			{x: 3, y: 1},
		},
	}
}

func (g *game) readWordsFromFile(fp string) {
	fn := fp
	f, err := os.Open(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening wordlist file %s: %s", fn, err.Error())
		os.Exit(1)
	}
	defer f.Close()

	// TODO: Validation - for now, code assumes that the file contains correct data
	i := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		i++
		g.words = append(g.words, line)
	}
}

func (g *game) randomizeWords() {
	rand.Shuffle(len(g.words), func(i, j int) {
		g.words[i], g.words[j] = g.words[j], g.words[i]
	})
}

func (g *game) isStarted() bool {
	return g.started
}

func (g *game) stopGame() {
	g.started = false
}

func (g *game) startGame() {
	g.started = true
	g.nextWordIndex = 0
	g.currentWord = ""
	g.currentTranslation = ""
	g.wordsGiven = 0
}

func (g *game) getNumberOfUsedWords() int {
	return g.wordsGiven
}

func (g *game) getNumberOfAllWords() int {
	return len(g.words)
}

func (g *game) getCurrentWord() string {
	return g.currentWord
}

func (g *game) isCurrentWordEmpty() bool {
	return g.currentWord == ""
}

func (g *game) useNewWord() {
	curWordArr := strings.Split(g.words[g.nextWordIndex], ":")
	g.currentWord = curWordArr[1]
	g.currentTranslation = curWordArr[0]
	g.setLettersFromCurrentWord()
	g.nextWordIndex++
	g.wordsGiven++
	g.consumedLetters = ""
}

func (g *game) setSize(w int, h int) {
	g.size[0] = w
	g.size[1] = h
}

func (g *game) setLettersFromCurrentWord() {
	g.letters = make([]letter, 0)
	for i := 0; i < len(g.currentWord); i++ {
		g.letters = append(g.letters, letter{
			x: rand.Intn(g.size[0]-2) + 1,
			y: rand.Intn(g.size[1]-2) + 1,
			l: string(g.currentWord[i]),
		})
	}
}

func (g *game) iterate() int {
	// If there is no word then take the next one
	if g.isCurrentWordEmpty() {
		g.useNewWord()
	}

	// Check hitting edges and eating itself (going backwards)
	switch g.direction {
	case DOWN:
		if g.snake[0].x == g.snake[1].x && g.snake[0].y == g.snake[1].y {
			g.stopGame()
			return GAME_OVER
		}
		if g.snake[0].y == g.size[1]-1 {
			g.stopGame()
			return GAME_OVER
		}
	case UP:
		if g.snake[0].x == g.snake[1].x && g.snake[1].y == g.snake[0].y {
			g.stopGame()
			return GAME_OVER
		}
		if g.snake[0].y == 0 {
			g.stopGame()
			return GAME_OVER
		}
	case LEFT:
		if g.snake[0].y == g.snake[1].y && g.snake[0].x == g.snake[1].x {
			g.stopGame()
			return GAME_OVER
		}
		if g.snake[0].x == 0 {
			g.stopGame()
			return GAME_OVER
		}
	case RIGHT:
		if g.snake[0].y == g.snake[1].y && g.snake[1].x == g.snake[0].x {
			g.stopGame()
			return GAME_OVER
		}
		if g.snake[0].x == g.size[0]-1 {
			g.stopGame()
			return GAME_OVER
		}
	}

	consumedLetter := false
	newLetters := []letter{}
	var addTail *segment
	for _, l := range g.letters {
		if g.snake[0].x == l.x && g.snake[0].y == l.y {
			g.consumedLetters += l.l
			consumedLetter = true
		} else {
			newLetters = append(newLetters, l)
		}
	}
	g.letters = newLetters

	if !consumedLetter {
		g.remove = &segment{
			x: g.snake[len(g.snake)-1].x,
			y: g.snake[len(g.snake)-1].y,
		}
	} else {
		g.remove = nil
		addTail = &segment{
			x: g.snake[len(g.snake)-1].x,
			y: g.snake[len(g.snake)-1].y,
		}
	}

	for i := len(g.snake) - 1; i > 0; i-- {
		g.snake[i].x = g.snake[i-1].x
		g.snake[i].y = g.snake[i-1].y
	}
	if addTail != nil {
		g.snake = append(g.snake, *addTail)
	}

	switch g.direction {
	case DOWN:
		g.snake[0].y++
	case UP:
		g.snake[0].y--
	case LEFT:
		g.snake[0].x--
	case RIGHT:
		g.snake[0].x++
	}

	if len(g.letters) == 0 {
		if g.currentWord == g.consumedLetters {
			g.wordsCorrect++
		}
		if g.nextWordIndex == len(g.words) {
			g.stopGame()
			return GAME_OVER
		}
		g.useNewWord()
	}

	return CONTINUE_GAME
}
