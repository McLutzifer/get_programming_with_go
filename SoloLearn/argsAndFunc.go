package main

import "fmt"

func welcome(name string) {
	fmt.Println("Hello, " + name)
}

func convert(x int) {
	fmt.Println(x / 100)
}

func mult(x, y int) {
	fmt.Println(x * y)
}

func sum(x, y int) int {
	return x + y
}

func main() {
	welcome("David")
	welcome("Lukas")
	convert(300)
	mult(3, 4)
	result := sum(42, 8)
	fmt.Println(result)
}
