package sprint
//lõpetatud 6 juuni kl 11:28

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
}