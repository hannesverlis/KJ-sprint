package sprint
//array 8 osa 1

func GenerateRange(min, max int) []int { // sisend (5, 10)
	if min > max { // 5 > 10, ei ole tõene, programm jätkab
		return nil 		//return []int{} // handle invalid input, empty slice
	}

	rangeSize := max - min //määrab elementide arve, 10-5=5
	result := []int{}      // Start with an empty slice literal, result tuleb luua ridade 14 ja 17 jaoks

	for i := 0; i < rangeSize; i++ { // i < 5
		result = append(result, min+i) // Append to the slice.  append lisab listile liikmeid koo väärtusega. Seda rida run jooksutab ni ikaua kui i < rangesize saab täis. result on mällu kirjutattav tulemus. iga ringiga kirjutab resultile ühe juurde
	}

	return result

}
