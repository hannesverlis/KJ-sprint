package sprint

//tehtud 5 juuni kl 15:13
//youtube  https://www.youtube.com/watch?v=DYqpu3jF2_4

func IsLower(str string) bool {
    for _, r := range str {  //
        if r < 'a' || r > 'z' { // Check if character is outside 'a' to 'z' range
            return false
        }
    }
    return true
}