package sprint
//import "fmt"

func SortWordArr(a []string) []string {
    for i := 0; i < len(a); i++ {
        for j := i + 1; j < len(a); j++ {
            //fmt.Println(a[j])
            if a[i][0] > a[j][0] { // Compare ASCII values of first characters
                a[i], a[j] = a[j], a[i] // Swap elements if out of order
            }
        } 
    }  
    return a
}