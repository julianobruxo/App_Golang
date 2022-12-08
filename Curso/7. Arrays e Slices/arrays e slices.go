package main

import "fmt"

func main() {
	fmt.Println("Arrays e slices")

	var array1 [5]string
	array1[0] = "Posição 0"
	array1[1] = "Posição 1"
	fmt.Println(array1)

	array2 := [3]string{"Pos1", "pos2", "pos3"}
	fmt.Println(array2)

	array3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array3[0])

	slice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice1)

	slice1 = append(slice1, 12)
	fmt.Println(slice1)

	array2[0] = "Posição inicial"
	fmt.Println(array2)

	//Arrays internos

	slice2 := make([]int, 10, 12)
	fmt.Println(slice2)

	slice2 = append(slice2, 5, 6, 9, 10, 11, 12, 13, 14, 15)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

}
