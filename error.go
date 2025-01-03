package main

import (
	"errors"
	"fmt"
)

/*
	As soon as a panic occur none of the code below that panic is going to execute.
	recover() allows to the programmer to deal with the runtime error himself
	we can only use a recover() inside a defer statement
**/

func deferFunc() {
	fmt.Println("Defer")
	r := recover() // r will contain the panic which will occur in the code
	if r == nil {
		fmt.Println("There is no error")
		return
	}
	fmt.Println(r) // The error will be printed
}

func division(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Error: Division by 0")
	}
	return a / b, nil
}

func errorHandlingWithPanic() {

	defer deferFunc()
	fmt.Println("Hello!") // This will be excuted before the defer
	panic("This will cause the crash of the program")
}

func errorHandler() {
	result, err := division(12, 2)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
