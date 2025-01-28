package main

type Vector2 struct {
	x, y int
}

type fVector2 struct {
	x, y float32
}

func addVector2(a Vector2, b Vector2) Vector2 {
	return Vector2{x: (a.x + b.x), y: a.y + b.y}
}

func addfVector2(a fVector2, b fVector2) fVector2 {
	return fVector2{x: (a.x + b.x), y: (a.y + b.y)}
}

