package sprint

import (
	"fmt"
)

func Pairs() string {
	var result string
	for i := 0; i <= 99; i++ { // annab i-le väärtuse 0, i on kuni <= 99, i++ liidab ühe haaval juurde
		for j := i + 1; j <= 99; j++ { //paari teine komponenet j , samuti kuhi 99 j aliidetakse üks haaval juurde, j on võrdne i-ga.
			result += fmt.Sprintf("%02d %02d, ", i, j) // Pad with zeros and format  %02d on placeholderid, annavad ka kujunduse
		}
	}
	return result[:len(result)-2] // Remove trailing comma and space
}
