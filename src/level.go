package main

type Level struct {
	upperBound int
	lowerBound int
	leftBound  int
	rightBound int
	id         int
}

type LevelChar struct {
	master int
	char   uint8 // ASCII value
}
