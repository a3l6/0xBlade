package main

import (
	"fmt"
)

type Level struct {
	upperBound int
	lowerBound int
	leftBound  int
	rightBound int
	id         int
	sprite     string
}

type LevelChar struct {
	master int
	char   uint8 // ASCII value
}

func (l *Level) draw() {
	fmt.Printf("\033[s")
	fmt.Print(l.sprite)
	fmt.Printf("\033[u")
}
