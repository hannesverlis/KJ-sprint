package sprint

import (
	"math"
)

func Casting(n float64) int {
	return int(math.Round(n)) // Round to nearest integer and cast to int
}

//func main() {
// Test cases
//  floats := []float64{3.14, 2.5, -1.7, 0.5}
//for _, f := range floats {
//  result := roundFloatToInt(f)
//fmt.Printf("Original: %.2f, Rounded: %d\n", f, result)
//}
//}
