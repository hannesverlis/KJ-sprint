package sprint

func Countdown(n int) string {
    if n < 0 || n > 9 {
        return "Invalid input"
    }

    result := ""
    for i := n; i >= 0; i -= 2 {
        result += string('0' + i)
       // result += string(rune('0' + i)) // Convert integer to character
        if i > 0 {
            result += ", "
        }
    }

    if n%2 == 0 {
        result += "!"
    } else { 

          result += "0!"
    }
  
    return result
}