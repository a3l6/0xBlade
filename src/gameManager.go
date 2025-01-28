package main

import (
	"errors"
	"fmt"
	"sync"
)

type GameObject interface {
	draw()
}

type GameManager struct {
	drawable map[int]*GameObject
	count    int
	console  map[string]string
	// TODO: Switch this to projectiles
	grenades    [100]Grenade // should be able to use only 100 grenades at once
	numGrenades uint8

	enemies    [100]Enemy
	numEnemies uint8

	ptrPlayer *Player

	occupiedCoordinates map[Vector2]int
	muCoords            sync.Mutex
	muConsole           sync.Mutex
}

const PERMANENT int = -1
const NOT_FOUND int = 0
const PLAYER int = 1
const ENEMY int = 2
const GRENADE int = 3

func (g *GameManager) getCoordinate(pos Vector2) int {
	g.muCoords.Lock()
	defer g.muCoords.Unlock()
	return g.occupiedCoordinates[pos]
}

func (g *GameManager) setCoordinate(pos Vector2, master int) {
	g.muCoords.Lock()
	current := g.occupiedCoordinates[pos]
	if current == NOT_FOUND {
		g.occupiedCoordinates[pos] = master
	} else if master == NOT_FOUND {
		g.occupiedCoordinates[pos] = master
	} else if current == ENEMY && master == GRENADE {
		g.occupiedCoordinates[pos] = master
	} else if current == PLAYER && master == ENEMY {
		g.occupiedCoordinates[pos] = master
	} else {
	}

	g.muCoords.Unlock()
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
		g.enemies[g.numEnemies] = Enemy{pos: fVector2{x: float32(pos.x), y: float32(pos.y)}, player: ptrPlayer, sprite: "?", vel: Vector2{0, 0}, damage: 0, health: 100, creationId: int(g.numEnemies)}
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
	g.muConsole.Lock()
	g.console[key] = val
	g.muConsole.Unlock()
}

func (g *GameManager) read(key string) string {
	g.muConsole.Lock()
	x := g.console[key]
	g.muConsole.Unlock()
	return x
}

func (g *GameManager) drawScreen() {

	// See README.md #1 for explanation of why this over usual for loop
	// FUTURE ME: don't change to traditional for loop unless absolutely necessary
	for _, val := range g.drawable {
		(*val).draw()
	}

	fmt.Printf("\033[s")
	g.muConsole.Lock()
	var console string
	for key, val := range g.console {
		console += fmt.Sprintf("%s : %s", key, val)
	}
	g.muConsole.Unlock()
	fmt.Printf("\033[%d;%dH", 47, 0)
	fmt.Printf("\033[2K")
	fmt.Printf("%s", console)
	fmt.Printf("\033[u")
	//g.console = make(map[string]string)
}
