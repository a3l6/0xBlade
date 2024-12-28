package main

import "fmt"

type Drawable interface {
	draw()
}

type GameManager struct {
	drawable map[int]*Drawable
	count    int
	console  map[string]string
	grenades []Grenade
}

// Registers with Game manager and returns id.
func (g *GameManager) registerAsObject(obj Drawable) int {
	ptr := &obj
	g.drawable[g.count] = ptr
	g.count++
	return g.count - 1
}

func (g *GameManager) createNewGrenade(pos Vector2) {
	g.grenades = append(g.grenades, Grenade{pos: pos, vel: Vector2{0, 0}, sprite: "O", trailSprite: "*", step: 0, amplitude: 1})
	idx := len(g.grenades) - 1
	g.grenades[idx].id = g.registerAsObject(&g.grenades[idx])
}

func (g *GameManager) deleteObject(id int) {
	if _, ok := g.drawable[id]; ok {
		fmt.Printf("\0332J")
	} else {
		fmt.Print("Hello World")
	}
	delete(g.drawable, id)
	if len(g.grenades) > 1 {
		g.grenades = g.grenades[1:] // shouldn't be too slow because how many grenades are even gonna be on the page??

	} else {
		g.grenades = make([]Grenade, 0)
	}
}

func (g *GameManager) drawScreen() {
	for i := 0; i < len(g.drawable); i++ {
		fmt.Println(g.drawable[i])
		(*g.drawable[i]).draw()
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
