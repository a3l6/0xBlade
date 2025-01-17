package main

import (
	"errors"
	"fmt"
)


type Level struct {
	sprite []string
	upperBound int
	lowerBound int
	leftBound int
	rightBound int
<<<<<<< HEAD
	level      [][]LevelChar
}

const Level_ID uint8 = 0
const Player_ID uint8 = 1
const Grenade_ID uint8 = 2
const Enemy_ID uint8 = 3

type LevelChar struct {
	char   uint8 // ASCII value
	master uint8
}

// Sets LevelChar at some coordinate. Does not render.
func (l *Level) print(master uint8, char uint8, x uint8, y uint8) error {
	// Bounds check
	if x > l.width || y > l.height {
		return errors.New("x or y value is out of bounds")
	}

	l.level[y][x] = LevelChar{char: char, master: master}
	return nil
}

// Sets line of LevelChars. Does not render.
func (l *Level) println(master uint8, chars string, y uint8) error {

	// Checking for max uint8
	if y > 255 {
		return errors.New("x or y value is too big")
	}
	if len(chars) != int(l.width) {
		return errors.New("string value needs to be the same as the level width")
	}

	var output []LevelChar
	for _, val := range chars {
		char := uint8(val)
		lc := LevelChar{char: char, master: master}
		output = append(output, lc)
	}
	l.level[y] = output
	return nil
}

// Converts ALL rows to strings and prints in parallel.
func (l *Level) render() {

}

// Converts ONE row to a string and returns
func (l *Level) renderln() string {
	return ""
=======
>>>>>>> parent of 896a7a0 (Started on new level)
}

func (l *Level) draw() {
	fmt.Printf("\033[s")
	for idx, val := range l.sprite {
		fmt.Printf("\033[%d;%dH%s", idx, 0, val)
	}
	fmt.Printf("\033[u")}