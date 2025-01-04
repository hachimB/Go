package main

import "fmt"

type Book struct {
	id    int
	title string
}

func (b *Book) setTitle(newTitle string) {
	b.title = newTitle
}

func test(b *Book) {
	b.id++
}

func change(x *int) {
	*x = 3
}

func pointer() {
	a := 45
	change(&a)
	fmt.Println(a)

	b := &a
	fmt.Println(b, &b, *&b, *b)

	book := Book{1, "old"}
	fmt.Println(book)

	book.setTitle("New")
	test(&book)
	fmt.Println(book)
}
