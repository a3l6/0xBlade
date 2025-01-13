package main

import "fmt"

type Level struct {
	sprite []string
	width  uint8
	height uint8

	upperBound int
	lowerBound int
	leftBound  int
	rightBound int
	level      [][]LevelChar
}

const Level_ID uint8 = 0
const Player_ID uint8 = 1
const Grenade_ID uint8 = 2
const Enemy_ID uint8 = 3

type LevelChar struct {
	char uint8
	id   uint8
}

func (l *Level) print(char LevelChar, x uint8, y uint8) {
	l.level[y][x] = char
}

func (l *Level) println(chars string, x uint8, y uint8) {
	var output []LevelChar
	l.level[y] = output
}

func (l *Level) render() {

}

func (l *Level) renderln() {

}

func (l *Level) draw() {
	fmt.Printf("\033[s")
	for idx, val := range l.sprite {
		fmt.Printf("\033[%d;%dH%s", idx, 0, val)
	}
	fmt.Printf("\033[u")
}
