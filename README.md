# Lettersnake

[![Go Reference](https://pkg.go.dev/badge/gopkg.pl/mikogs/lettersnake.svg)](https://pkg.go.dev/gopkg.pl/mikogs/lettersnake) [![Go Report Card](https://goreportcard.com/badge/gopkg.pl/mikogs/lettersnake)](https://goreportcard.com/report/github.com/gopkg.pl/mikogs/lettersnake)

This project is a small, terminal-based snake game designed for my kids. The concept is that you get a word in one language and you need to collect letters to form its translation in another.

See screenshot below:

![Lettersnake](screenshot.png)

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

