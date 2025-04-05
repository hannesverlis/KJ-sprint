package sprint


type Point struct {
    X float32 // Use consistent field names: X and Y
    Y float32
    Text string // Optional text field
}

// PointDiff determines which Point has a greater X value (or greater Y if X values are equal)
func PointDiff(p1, p2 Point) Point {
    if p1.X > p2.X || (p1.X == p2.X && p1.Y >= p2.Y) {
        return p1
    }
    return p2
}
/*
// NewPoint creates a new Point with the given coordinates and an optional text label
func NewPoint(x, y float32, text ...string) Point {
    p := Point{X: x, Y: y}
    if len(text) > 0 {
        p.Text = text[0]
    }
    return p
}
*/
/*
type Point struct {
X float32
Y float32
Text string
}

func MakePoint(x, y float32, text string) Point {
	return Point {
		X: x,
		Y: y,
		Text: text,
	}
}*/