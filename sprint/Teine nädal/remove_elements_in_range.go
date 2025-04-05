package sprint

//import(
 //   "fmt"
//)

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
    // Swap indices if from is greater th an to
    if from > to {   // from 0 > to 15
        from, to = to, from  // vahetab ümber 15, 0
      
    }
   // fmt.Println(from, to)
    // Handle out-of-bounds and empty ranges
    if from < 0 {  //-5 = 0
        from = 0
    }
    if to > len(arr) {
        to = len(arr)
    }
   // fmt.Println(len(arr))
   // fmt.Println(arr)
   // fmt.Println(from, to)
    
    if from > to {
        return arr // Empty range, return a copy of the original slice
    }
    //fmt.Println(arr)
    //fmt.Println(arr)
    // Remove elements by slicing (creates a new slice)
   // fmt.Println(from, to)
    return append(arr[:from], arr[to:]...)
}

//