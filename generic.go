package main

import "fmt"

type Number interface {
	~int | ~uint | ~float64
}

func add[T Number](x, y T) T {
	return x + y
}

func minimum[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func indexOf[C comparable](slice []C, target C) int {
	for i, val := range slice {
		if val == target {
			return i
		}
	}
	return -1
}

func swap[A any](a, b A) (A, A) {
	return b, a
}

type Strct[F, S any] struct {
	name F
	age  S
}

func generic() {
	fmt.Println(add(93.5, 54))
	fmt.Println(indexOf([]int{1, 2, 3, 4, 5}, 0))
	fmt.Println(swap(123, 5))

	s := Strct[string, int]{"Hachim", 21}
	fmt.Println(s.name, s.age, s)

}
