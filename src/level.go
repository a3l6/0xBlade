package main

import "fmt"


type Level struct {
	sprite []string
	upperBound int
	lowerBound int
	leftBound int
	rightBound int
}

func (l *Level) draw() {
	fmt.Printf("\033[s")
	for idx, val := range l.sprite {
		fmt.Printf("\033[%d;%dH%s", idx, 0, val)
	}
	fmt.Printf("\033[u")}