package main

import "errors"

type LevelChar struct {
	char   rune
	master int
}

func StringToLevelChar(master int, a string) []LevelChar {
	output := make([]LevelChar, len(a))
	for idx, val := range a {
		output[idx] = LevelChar{master: master, char: val}
	}
	return output
}

type Level struct {
	upperBound int
	lowerBound int
	leftBound  int
	rightBound int
	sprite     []string
	level      [windowHeight][windowWidth]LevelChar
}

func (l *Level) println(y int, chars []LevelChar) error {
	if len(chars) != int(windowWidth) {
		return errors.New("line exceeds max width")
	}
	if y < 0 || y >= int(windowHeight) {
		return errors.New("y index out of bounds")
	}

	copy(l.level[y][:], chars)
	return nil
}

func (l *Level) print(pos Vector2, char LevelChar) error {
	if pos.x >= int(windowWidth) || pos.y >= int(windowHeight) {
		return errors.New("position out of bounds")
	}
	if char.master != -1 {
		gameManager.sendCollision(char.master, l.level[pos.y][pos.x].master)
	}
	l.level[pos.y][pos.x] = char
	return nil
}

// Prints multiple chars. Assuming its all on ONE line!
func (l *Level) printM(pos Vector2, chars []LevelChar) error {
	for idx, val := range chars {
		newPos := Vector2{x: pos.x + idx, y: pos.y}
		err := l.print(newPos, val)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *Level) render() []string {
	output := make([]string, windowHeight)
	for idx, val := range l.level {
		for _, val := range val {
			output[idx] += string(val.char)
		}
	}
	return output
}
