package main

import "fmt"

func array() {
	var arr [3]int

	arr2 := [...][2]int{{1, 2}, {12, 344}, {12, 34}}
	arr2[0] = [2]int{90, 98}

	arr3 := [2]string{"Hi", "everyone"}
	arr3[0] = "iuo"

	var arr4 [2]bool

	fmt.Println(arr, arr4, arr3)
	fmt.Println(arr2)
}
