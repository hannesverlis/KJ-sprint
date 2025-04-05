package sprint

type Point struct {
	X    float32 // X-coordinate
	Y    float32 // Y-coordinate
	Text string  // Text associated with the point
}

func MakePoint(x, y float32, text string) Point {
	return Point{
		X:    x,
		Y:    y,
		Text: text,
	}
}
