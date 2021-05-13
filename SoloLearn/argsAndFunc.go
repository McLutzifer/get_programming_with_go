package main

import "fmt"

func welcome(name string) {
	fmt.Println("Hello, " + name)
}

func convert(x int) {
	fmt.Println(x / 100)
}

func main() {
	welcome("David")
	welcome("Lukas")
	convert(300)
}
