package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Level struct {
	sprite     []string
	upperBound int
	lowerBound int
	leftBound  int
	rightBound int
	id         int
	level      [windowHeight][windowWidth]LevelChar

	width  uint16
	height uint8
}

const Level_ID uint8 = 0
const Player_ID uint8 = 1
const Grenade_ID uint8 = 2
const Enemy_ID uint8 = 3

type LevelChar struct {
	master int
	char   uint8 // ASCII value
}

// Sets LevelChar at some coordinate. Does not render.
func (l *Level) print(master int, char uint8, x uint8, y uint8) error {
	// Bounds check
	//	if x > l.width || y > l.height {
	//		return errors.New("x or y value is out of bounds")
	//	}
	gameManager.NewCollision(master, l.level[y][x].master)
	l.level[y][x*2] = LevelChar{char: char, master: master}
	return nil
}

// Sets line of LevelChars. Does not render.
func (l *Level) println(master int, chars string, y uint8) error {

	if len(chars) != int(l.width) {
		return errors.New("string value needs to be the same as the level width")
	}

	var output [windowWidth]LevelChar
	for idx, val := range chars {
		gameManager.NewCollision(master, l.level[y][idx].master)
		char := uint8(val)
		lc := LevelChar{char: char, master: master}
		output[idx] = lc
	}
	l.level[y] = output
	return nil
}

// Converts ALL rows to strings and prints in parallel.
func (l *Level) render() []string {
	output := make([]string, l.height)
	for y := range l.level {
		for _, x := range l.level[y] {
			output[y] += string(x.char)
		}
	}
	fileName := "out1.txt"
	data := strings.Join(output, "\n")
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Cannot create file! ", err)
	}
	defer file.Close()

	_, err = file.WriteString(data) //fmt.Sprintf("LENGTH %d", len(l.level)))

	if err != nil {
		fmt.Println("Error writing to file, ", err)
	}
	return output
}

// Converts ONE row to a string and returns
func (l *Level) renderln() string {
	return ""
}

func (l *Level) draw() {
	for idx, val := range l.sprite {
		l.println(l.id, val, uint8(idx))
	}
	/*fmt.Printf("\033[s")
	for idx, val := range l.sprite {
		fmt.Printf("\033[%d;%dH%s", idx, 0, val)
	}
	fmt.Printf("\033[u")*/
}

func (l *Level) collision(x int) {}
