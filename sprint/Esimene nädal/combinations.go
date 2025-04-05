package sprint

import "fmt"

func Combinations() string {
	var triplets string

	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			for k := j + 1; k < 10; k++ {
				triplets += fmt.Sprintf("%d%d%d, ", i, j, k)
			}
		}
	}

	// Remove trailing comma and space.
	triplets = triplets[:len(triplets)-2]

	return triplets

}
