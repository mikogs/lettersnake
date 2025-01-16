# Snakey Letters

[![Go Reference](https://pkg.go.dev/badge/github.com/cli-games/snakey-letters.svg)](https://pkg.go.dev/github.com/cli-games/snakey-letters) [![Go Report Card](https://goreportcard.com/badge/github.com/cli-games/snakey-letters)](https://goreportcard.com/report/github.com/cli-games/snakey-letters)

This project is a small, terminal-based snake game designed for my kids. The concept is that you get a word in one language and you need to collect letters to form its translation in another.

See screenshot below:

![Snakey-Letters](screenshot.png)

### Running

To play the game just run:

    go run *.go start -f words/words-pl-en-animals.txt

### Instructions
Use arrows to steer the snake.

### Words file
The words for the game are provided via the `-f` argument, and the file's format is straightforward.

Every line contains a word in one language and its translation in another (which needs to be guessed). Words are delimetered by a colon (`:`).
Space cannot be used, so an underscore (`_`) is preffered.

For example, a sample file might look like this:

    hol:hall
    garaż:garage
    jadalnia:dining_room

