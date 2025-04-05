package sprint

func ReverseAlphabet(step int) string {
	if step <= 0 {
		step = 1 // Default step if 0 or negative
	}

	var result []rune // Initialize empty rune slice

	// Iterate in reverse through alphabet
	for r := 'z'; r >= 'a'; r -= rune(step) { // rune määramine, '' ülakomad näitavad runet, rune piirid z kuni a, -=  lahutab maha stepi võrra
		result = append(result, r) // Add the current rune to the result.append onsisse ehitatud käsk, mis lisab runele väärtuse juurde
	}
	return string(result)
}
