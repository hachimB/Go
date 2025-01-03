package main

import "fmt"

func Sum(x, y int) string {
	result := fmt.Sprintf("The sum of %d and %d is %d", x, y, x+y)
	return result
}
