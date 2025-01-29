package main

type Vector2 struct {
	x int
	y int
}

func addVector2(a Vector2, b Vector2) Vector2 {
	x := a.x + b.x
	y := a.y + b.y
	return Vector2{x: x, y: y}
}
