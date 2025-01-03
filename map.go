package main

import "fmt"

func maps() {
	mp := map[string]int{"a": 23, "b": 25}
	mp["c"] = 90
	fmt.Printf("%T \n", mp)
	fmt.Println(mp)

	m := map[string][]int{"i": {2, 3, 4}}
	m["j"] = []int{45, 7}

	//Deletion
	delete(m, "i")

	//Checking if an element exists or not
	value, ok := m["o"]
	fmt.Println(value, ok)

	ma := map[string]string{"first": "ok", "second": "okk"}
	fmt.Println(ma)

	fmt.Println(m)
}
