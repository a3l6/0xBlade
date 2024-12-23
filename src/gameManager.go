package main

type Drawable interface {
	draw()
}


type GameManager struct {
	drawable map[int]Drawable
	objects map[int]uintptr
	count int
}


// Registers with Game manager and returns id.
func (g *GameManager) registerAsObject(obj uintptr) int{
	g.objects[g.count] = obj
	g.count++
	return g.count
}

func (g *GameManager) deleteObject(id int) {
	delete(g.objects, id)
}

func (g *GameManager) registerAsDrawable(obj uintptr) int{
	id := g.registerAsObject(obj)
	
	return id
}

func (g GameManager) drawScreen() {
	for i:=0; i < len(g.drawable); i++ {
		g.drawable[i].draw()
	}
}