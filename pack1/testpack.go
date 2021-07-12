package pack1

import "fmt"

func TestPack1() {
	fmt.Println("Test Pack 1")

	var arr [2]int

	arr[0] = 1
	arr[1] = 2

	fmt.Println("arr[1] = ", arr[1])

	arr2 := &arr

	fmt.Println("arr[1] = ", arr[1])
	fmt.Println("arr2[1] = ", arr2[1])

	arr2[1] = 3

	fmt.Println("arr[1] = ", arr[1])
	fmt.Println("arr2[1] = ", arr2[1])

}
