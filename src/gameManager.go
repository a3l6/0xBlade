package main

import (
	"errors"
	"fmt"
	random "math/rand"
	"strings"
)

type GameObject interface {
	draw()
}

func generateBG() [45 * 95 * 2]rune {
	const width = 95 * 2
	const height = 45

	var bg [width * height]rune
	for i := 0; i < len(bg); i += 2 {
		bg[i] = '#'
		bg[i+1] = ' '
		// This is not a space
		// Char above is U+2000, EN QUAD
		// This is here because I want to have a space between the # but I want to detect what is a wall
	}

	emptySpace := []rune(strings.Repeat(string([]rune{' ', ' '}), 40)) // These are spaces
	//	emptySpace := bytes.Repeat([]rune{' ', ' '}, 41)
	offset := (width - len(emptySpace)) / 2
	for i := offset + (2 * width); i <= len(bg)-2*width; i += width {
		copy(bg[i:], emptySpace)
	}

	return bg
}

// var spaceBuffer = [windowWidth * 2 * windowHeight]byte(bytes.Repeat([]byte{' '}, windowWidth*2*windowHeight))
var spaceBuffer = generateBG()

type GameManager struct {
	CurrBuffer [windowWidth * windowHeight]rune
	prevBuffer [windowWidth * windowHeight]rune

	drawable map[int]*GameObject
	count    int
	console  map[string]string
	// TODO: Switch this to projectiles
	grenades    [100]Grenade // should be able to use only 100 grenades at once
	numGrenades uint8

	enemies    [100]Enemy
	numEnemies uint8

	ptrPlayer *Player
	direction uint8

	difficultySeed uint8 // Ensure proper setting to use per time
}

// Registers with Game manager and returns unique id.
func (g *GameManager) registerAsObject(obj GameObject) int {
	ptr := &obj
	g.drawable[g.count] = ptr
	g.count++
	return g.count - 1
}

func (g *GameManager) createNewGrenade(pos Vector2, direction uint8) error {
	//  100 is max for grenades
	if g.numGrenades != 100 {
		g.grenades[g.numGrenades] = Grenade{pos: pos, vel: Vector2{0, 0}, sprite: "O", trailSprite: "*", step: 0, amplitude: 1, creationID: g.numGrenades, direction: direction}
		g.grenades[g.numGrenades].id = g.registerAsObject(&g.grenades[g.numGrenades])
		g.numGrenades++
		return nil
	} else {
		return errors.New("too many projectiles made")
	}
}

func (g *GameManager) tryToSpawnEnemy() {
	randNum := random.Intn(int(g.difficultySeed)) // difficultySeed adjusted to match per frame
	randomPos := random.Intn(len(g.prevBuffer))
	char := g.CurrBuffer[randomPos]

	//copy(g.CurrBuffer[10:], []rune(fmt.Sprintf("\"%s\"", )))
	if char == ' ' && randNum == 1 {
		x := randomPos % windowWidth
		y := randomPos / windowWidth

		g.createNewEnemy(Vector2{x: int(x), y: int(y)})
	}
}

func (g *GameManager) createNewEnemy(pos Vector2) error {
	//  100 is max for enemies
	if g.numEnemies != 100 {
		g.enemies[g.numEnemies] = Enemy{
			pos: fVector2{x: float32(pos.x),
				y: float32(pos.y)},
			sprite:     '?',
			vel:        Vector2{0, 0},
			health:     100,
			creationId: int(g.numEnemies)}
		g.enemies[g.numEnemies].id = g.registerAsObject(&g.enemies[g.numEnemies])
		g.numEnemies++
		return nil
	} else {
		return errors.New("too many enemies spawned")
	}
}

func (g *GameManager) killEnemy(id int, creationID int) {
	delete(g.drawable, id)
	g.enemies[creationID] = Enemy{}
}

// Run all objects that have a step function.
// No array of stepable stuff, just manually add a new one.
// Make sure to call at a specific frame rate .
func (g *GameManager) StepAll() {
	for idx := range g.enemies {
		// keep as this cause using val just edits the copy
		g.enemies[idx].Step()
	}

	for idx := range g.grenades {
		g.grenades[idx].Step()
	}
}

func (g *GameManager) deleteObject(id int, creationID uint8) {
	delete(g.drawable, id)
	// reset
	g.grenades[creationID] = Grenade{pos: Vector2{0, 0}, vel: Vector2{0, 0}, sprite: "O", trailSprite: "*", step: 0, amplitude: 1, creationID: g.numGrenades}
}

func (g *GameManager) writeToConsole(key string, val string) {
	g.console[key] = val
}

func (g *GameManager) read(key string) string {
	x := g.console[key]
	return x
}

// Calls Drawable.Draw methods.
// Compares buffer and only prints diff.
func (g *GameManager) drawScreen() {
	// See README.md #1 for explanation of why this is used over bare for loop
	// FUTURE ME: Don't change this to traditional loop

	copy(g.CurrBuffer[:], spaceBuffer[:])

	g.tryToSpawnEnemy()

	for _, val := range g.drawable {
		(*val).draw()
	}

	for i := range g.CurrBuffer {
		if g.CurrBuffer[i] != g.prevBuffer[i] {
			x, y := i%windowWidth, i/windowWidth
			fmt.Printf("\033[%d;%dH%c", y+1, x+1, g.CurrBuffer[i])
		}
	}
	// TODO: Make elegant handling of console
	// Another window would be really nice
	copy(g.prevBuffer[:], g.CurrBuffer[:])
}
