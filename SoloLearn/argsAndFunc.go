package main

import "fmt"

func welcome(name string) {
	fmt.Println("Hello, " + name)
}

func convert(x int) {
	fmt.Println(x / 100)
}

func multiply(x, y int) {
	fmt.Println(x * y)
}

func sum(x, y int) int {
	return x + y
}

func concat(a, b string) string {
	return a + b
}

func swap(x, y int) (int, int) {
	return x, y
}

func welcome2() {
	fmt.Println("Welcome")
}

func main() {
	defer welcome2()
	welcome("David")
	welcome("Lukas")
	convert(300)
	multiply(3, 4)
	result := sum(42, 8)
	fmt.Println(result)
	a, b := swap(42, 8)
	con := concat("hey", "ho")
	fmt.Println(con)
	fmt.Printf("a: %d and b: %v\n", a, b)
}
