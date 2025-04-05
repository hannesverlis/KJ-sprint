package sprint
//import "fmt"

func ToRoman(num int) string {
	if num < 1 || num > 3999 {
		return "Invalid input"
	}

	// Define Roman numeral representations for different values
	romanNumerals := []struct {
		Value int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""

	// Iterate through the Roman numeral representations
	for _, numeral := range romanNumerals {
      //  fmt.Println(numeral)
        //fmt.Println(romanNumerals)
		// Append the symbol as many times as the value fits into num
		for num >= numeral.Value {
			result += numeral.Symbol
			num -= numeral.Value
		}
	}

	return result
}
