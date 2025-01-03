package main

import "fmt"

type ShapeInterface interface {
	getPerimeter() uint
	getArea() float64
}

type Triangle struct {
	a, b, c, h, base uint
}

type Square struct {
	side uint
}

func (t Triangle) getPerimeter() uint {
	val := t.a + t.b + t.c
	return val
}

func (t Triangle) getArea() float64 {
	return 0.5 * float64(t.base) * float64(t.h)
}

func (s Square) getPerimeter() uint {
	return 4 * s.side
}

func (s Square) getArea() float64 {
	return float64(s.side * s.side)
}

func intrfce() {

	t := Triangle{a: 2, b: 4, c: 5, h: 3, base: 10}
	square := Square{side: 5}

	var s ShapeInterface = square
	var tri ShapeInterface = t

	fmt.Printf("The area of the square %.2f\n", s.getArea())
	fmt.Printf("The perimeter of the square is %d\n", s.getPerimeter())

	fmt.Printf("The area of the triangle %.2f\n", tri.getArea())
	fmt.Printf("The perimeter of the triangle is %d\n", tri.getPerimeter())
}
