package main

import (
	"bytes"
	"errors"
	"fmt"
)

type GameObject interface {
	draw()
}

func generateBG() [45 * 95 * 2]byte {
	const width = 95 * 2
	const height = 45

	var bg [width * height]byte
	for i := 0; i < len(bg); i += 2 {
		bg[i] = '#'
		bg[i+1] = ' '
	}
	fmt.Println(len(bg))
	emptySpace := bytes.Repeat([]byte{' ', ' '}, 40)
	offset := (width - len(emptySpace)) / 2
	for i := offset + (2 * width); i <= len(bg)-2*width; i += width {
		copy(bg[i:], emptySpace)
	}

	return bg
}

// var spaceBuffer = [windowWidth * 2 * windowHeight]byte(bytes.Repeat([]byte{' '}, windowWidth*2*windowHeight))
var spaceBuffer = generateBG()

type GameManager struct {
	CurrBuffer [windowWidth * 2 * windowHeight]byte
	prevBuffer [windowWidth * 2 * windowHeight]byte

	drawable map[int]*GameObject
	count    int
	console  map[string]string
	// TODO: Switch this to projectiles
	grenades    [100]Grenade // should be able to use only 100 grenades at once
	numGrenades uint8

	enemies    [100]Enemy
	numEnemies uint8

	ptrPlayer *Player
}

// Registers with Game manager and returns unique id.
func (g *GameManager) registerAsObject(obj GameObject) int {
	ptr := &obj
	g.drawable[g.count] = ptr
	g.count++
	return g.count - 1
}

func (g *GameManager) createNewGrenade(pos Vector2) error {
	//  100 is max for grenades
	if g.numGrenades != 100 {
		g.grenades[g.numGrenades] = Grenade{pos: pos, vel: Vector2{0, 0}, sprite: "O", trailSprite: "*", step: 0, amplitude: 1, creationID: g.numGrenades}
		g.grenades[g.numGrenades].id = g.registerAsObject(&g.grenades[g.numGrenades])
		g.numGrenades++
		return nil
	} else {
		return errors.New("too many projectiles made")
	}
}

func (g *GameManager) createNewEnemy(pos Vector2, ptrPlayer *Player) error {
	//  100 is max for enemies
	if g.numEnemies != 100 {
		g.enemies[g.numEnemies] = Enemy{
			pos: fVector2{x: float32(pos.x),
				y: float32(pos.y)},
			player:     ptrPlayer,
			sprite:     '?',
			vel:        Vector2{0, 0},
			damage:     0,
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
