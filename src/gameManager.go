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
	grenades    [100]Grenade // should be able to use only 200 grenades at once
	numGrenades uint8
}

// Registers with Game manager and returns id.
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

func (g *GameManager) deleteObject(id int, creationID uint8) {
	delete(g.drawable, id)
	// reset
	g.grenades[creationID] = Grenade{pos: Vector2{0, 0}, vel: Vector2{0, 0}, sprite: "O", trailSprite: "*", step: 0, amplitude: 1, creationID: g.numGrenades}
}

func (g *GameManager) drawScreen() {
	// See README.md #1 for explanation of why this over usual for loop
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
