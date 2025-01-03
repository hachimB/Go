package main

import "fmt"

/* Slices :
	pointer -> to the first element of the array
	length -> len(slice)
	capacity -> cap(slice)
**/

func slice() {
	arr := [5]int{12, 34, 45, 5, 6} // array notation
	fmt.Printf("Array's type: %T\n", arr)
	sl := arr[1:]
	fmt.Printf("Slice's type: %T\n", sl)
	sl[0] = 800 // will change the value of the second element of the array also
	fmt.Println(arr, sl, cap(sl), len(sl))

	slic := []int{1, 2, 4} // slice notation
	fmt.Printf("%T ", slic)
	fmt.Println(slic, cap(slic), len(slic))
	fmt.Println()

	s := []string{"Hello", "world!"}
	for x := 0; x < 10; x++ {
		s = append(s, "ok")
		fmt.Println(s, len(s), cap(s))
	}
	fmt.Println(s)

	for i, val := range s {
		fmt.Println(i, val)
	}

	// We can also create a slice like this:
	ms := make([]int, 3)
	fmt.Println(ms)
}
