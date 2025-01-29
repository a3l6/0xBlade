package main

import "fmt"

type Object interface {
	collision(int) // I'm just gonna make all objects have a collsion
}

type GameManager struct {
	count   int
	objects map[int]*Object
}

// Takes in pointer to object and returns the id that it is held under
func (g *GameManager) registerObject(obj Object) int {
	g.objects[g.count] = &obj
	g.count++
	return g.count - 1
}

func (g *GameManager) sendCollision(sender int, reciever int) {
	if sender > len(g.objects) || reciever > len(g.objects) {
		return
	}
	(*g.objects[reciever]).collision(sender)
}

func (g *GameManager) drawScreen() {
	rendered := level.render()
	fmt.Printf("\033[2J\033[s")
	for i, val := range rendered {
		fmt.Printf("\033[%d;%dH%s", i, 0, val)
	}
	fmt.Printf("\033[u")

}
