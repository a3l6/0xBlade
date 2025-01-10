package main

import (
	"errors"
	"fmt"
)

type Drawable interface {
	draw()
}

type GameManager struct {
	drawable map[int]*Drawable
	count    int
	console  map[string]string
	// TODO: Switch this to projectiles
	grenades    [100]Grenade // should be able to use only 100 grenades at once
	numGrenades uint8

	enemies    [100]Enemy
	numEnemies uint8
}

// Registers with Game manager and returns unique id.
func (g *GameManager) registerAsObject(obj Drawable) int {
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
		g.enemies[g.numEnemies] = Enemy{pos: pos, player: ptrPlayer, sprite: "?", vel: Vector2{0, 0}, damage: 0, health: 100}
		g.enemies[g.numEnemies].id = g.registerAsObject(&g.enemies[g.numEnemies])
		g.numEnemies++
		return nil
	} else {
		return errors.New("too many enemies spawned")
	}
}

// Run all objects that have a step function.
// No array of stepable stuff, just manually add a new one.
// Make sure to call at a specific frame rate .
func (g *GameManager) StepAll() {
	for _, val := range g.enemies {
		val.Step()
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

func (g *GameManager) drawScreen() {
	// See README.md #1 for explanation of why this over usual for loop
	// FUTURE ME: don't change to traditional for loop unless absolutely necessary
	for _, val := range g.drawable {
		(*val).draw()
	}

	fmt.Printf("\033[s")
	var console string
	for key, val := range g.console {
		console += fmt.Sprintf("%s : %s", key, val)
	}
	fmt.Printf("\033[%d;%dH", 47, 0)
	fmt.Printf("\033[2K")
	fmt.Printf("%s", console)
	fmt.Printf("\033[u")
	//g.console = make(map[string]string)
}
