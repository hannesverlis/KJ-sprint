package sprint

import "fmt"

// Point structure (assuming this is the same as before)
type Point struct {
    X float32
    Y float32
    Text string
}

// PointText function to format the text
func PointText(p Point) Point {
    return Point{
        X: p.X, 
        Y: p.Y,
        Text: fmt.Sprintf("Text at (%f, %f)", p.X, p.Y),
    }
}
/*
    p := Point{X: 3, Y: 5}
newPoint := PointText(p)
fmt.Println(newPoint) // Output: {3 5 Text at (3, 5)}
return (Point)
}*/
