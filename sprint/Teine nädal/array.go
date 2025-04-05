package main
// https://www.youtube.com/watch?v=e-oBn806Pzc

import "fmt"

func main() {
arr := [3]int{4, 55, 66} 
sum := 0

arr2D := [2][3]int{{1, 2, 3}, {2, 4, 6}} // jada jadas. esimene näitab suure jada komponentide arvu, teine nurksulg näitab väikse jada elementide arvu
fmt.Println(arr2D[1][0])
fmt.Println(len(arr))

for i := 0; i < len(arr); i++ {
	sum += arr[i]
}
fmt.Println(sum)

}