package sprint

//8 Arrays / osa 3 Balance out // tehtud kl 22:43

func BalanceOut(arr []bool) []bool {
    trueCount := 0
    for _, val := range arr {
        if val {
            trueCount++
        }
    }

    falseCount := len(arr) - trueCount  // Calculate falseCount directly
    diff := trueCount - falseCount
    if diff < 0 {
        diff = -diff
    }

    result := make([]bool, len(arr)+diff)
    copy(result, arr)

    // Correctly determine which value to add
    addTrue := trueCount < falseCount 
    for i := len(arr); i < len(result); i++ {
        result[i] = addTrue
    }

    return result
}

//BalanceOut([]bool{true, false, false, false})
//>> [true false false false true true]
