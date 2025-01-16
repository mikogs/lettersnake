package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-phings/broccli"
)

func versionHandler(c *broccli.CLI) int {
	fmt.Fprintf(os.Stdout, VERSION+"\n")
	return 0
}

func main() {
	cli := broccli.NewCLI("snakey-letters", "Classic snake but with letters and words!", "")
	cmd := cli.AddCmd("start", "Starts the game", startHandler)
	cmd.AddFlag("words", "f", "", "Text file with wordlist", broccli.TypePathFile, broccli.IsExistent|broccli.IsRequired)
	cmd.AddFlag("speed", "s", "", "Snake speed", broccli.TypeInt, 0)
	_ = cli.AddCmd("version", "Shows version", versionHandler)
	if len(os.Args) == 2 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		os.Args = []string{"App", "version"}
	}
	os.Exit(cli.Run())
}

func startHandler(c *broccli.CLI) int {
	g := newGame()
	g.readWordsFromFile(c.Flag("words"))
	g.randomizeWords()

	speed := c.Flag("speed")
	if speed == "" {
		speed = "200"
	}

	speedInt, _ := strconv.Atoi(speed)

	gi := newGameInterface(g)
	gi.setSpeed(speedInt)

	gi.run()
	return 0
}
